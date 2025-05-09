// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
package azeventhubs

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/log"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs/v2/internal"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs/v2/internal/amqpwrap"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs/v2/internal/exported"
	"github.com/Azure/go-amqp"
	"github.com/stretchr/testify/require"
)

func TestUnit_Processor_loadBalancing(t *testing.T) {
	cps := newCheckpointStoreForTest()
	firstProcessor := newProcessorForTest(t, "first-processor", cps, nil)
	newTestOwnership := func(base Ownership) Ownership {
		base.ConsumerGroup = "consumer-group"
		base.EventHubName = "event-hub"
		base.FullyQualifiedNamespace = "fqdn"
		return base
	}

	require.Equal(t, ProcessorStrategyBalanced, firstProcessor.lb.strategy)

	allPartitionIDs := []string{"1", "100", "1001"}
	lbinfo, err := firstProcessor.lb.getAvailablePartitions(context.Background(), allPartitionIDs)
	require.NoError(t, err)

	// this is a completely empty checkpoint store so nobody owns any partitions yet
	// which means that we get to claim them all
	require.Empty(t, lbinfo.aboveMax)
	require.Empty(t, lbinfo.current)
	require.True(t, lbinfo.claimMorePartitions)
	require.Equal(t, 3, lbinfo.maxAllowed, "only 1 possible owner (us), so we're allowed all the available partitions")

	expectedOwnerships := []Ownership{
		newTestOwnership(Ownership{
			PartitionID: "1",
			OwnerID:     "first-processor",
		}),
		newTestOwnership(Ownership{
			PartitionID: "100",
			OwnerID:     "first-processor",
		}),
		newTestOwnership(Ownership{
			PartitionID: "1001",
			OwnerID:     "first-processor",
		}),
	}

	require.Equal(t, expectedOwnerships, lbinfo.unownedOrExpired)

	// getAvailablePartitions doesn't mutate the checkpoint store.
	lbinfo, err = firstProcessor.lb.getAvailablePartitions(context.Background(), allPartitionIDs)
	require.NoError(t, err)
	require.Equal(t, expectedOwnerships, lbinfo.unownedOrExpired)

	// the balanced strategy claims one new partition per round, until balanced.
	// we'll do more in-depth testing in other tests, but this is just a basic
	// run through.
	firstProcessorOwnerships, err := firstProcessor.lb.LoadBalance(context.Background(), allPartitionIDs)
	require.NoError(t, err)

	expectedLoadBalancingOwnership := updateDynamicData(t, firstProcessorOwnerships[0], newTestOwnership(Ownership{
		PartitionID: "1001",
		OwnerID:     "first-processor",
	}), allPartitionIDs)
	require.Equal(t, []Ownership{expectedLoadBalancingOwnership}, firstProcessorOwnerships)

	// at this point this is our state:
	// 3 total partitions ("1", "100", "1001")
	// 1 of those partitions is owned by our client ("first-processor")
	// 2 are still available.

	secondProcessor := newProcessorForTest(t, "second-processor", cps, nil)

	// when we ask for available partitions we take into account the owners that are
	// present in the checkpoint store and ourselves, since we're about to try to claim
	// some partitions. So now it has to divide 3 partitions amongst two Processors.
	lbinfo, err = secondProcessor.lb.getAvailablePartitions(context.Background(), allPartitionIDs)
	require.NoError(t, err)
	require.Empty(t, lbinfo.aboveMax)
	require.Empty(t, lbinfo.current)
	require.True(t, lbinfo.claimMorePartitions)
	require.Equal(t, 1, lbinfo.maxAllowed, "the max is now 1 (instead of 2) because _our_ processor doesn't own enough")

	// there are two available partition ownerships - we should be getting one of them.
	newProcessorOwnerships, err := secondProcessor.lb.LoadBalance(context.Background(), allPartitionIDs)
	require.NoError(t, err)

	newExpectedLoadBalancingOwnership := updateDynamicData(t, newProcessorOwnerships[0], newTestOwnership(Ownership{
		PartitionID: "1001",
		OwnerID:     "second-processor",
	}), allPartitionIDs)

	require.Equal(t, []Ownership{newExpectedLoadBalancingOwnership}, newProcessorOwnerships)
	require.NotEqual(t, newExpectedLoadBalancingOwnership.PartitionID, expectedLoadBalancingOwnership.PartitionID, "partitions should not be assigned twice")

	//
	// now let's assign out the last partition - we'll pick a winner here and just use the second processor, but either one can technically claim it (or even attempt to at the same time!)
	//

	secondProcessorOwnershipsForLastPartition, err := secondProcessor.lb.LoadBalance(context.Background(), allPartitionIDs)
	require.NoError(t, err)

	require.Equal(t, 2, len(secondProcessorOwnershipsForLastPartition))

	// no overlap in partition assignments
	for _, o := range secondProcessorOwnershipsForLastPartition {
		require.NotEqual(t, firstProcessorOwnerships[0].PartitionID, o.PartitionID)
	}

	// and if we try to claim now with the first, it won't get anything new since
	// a. we're in a balanced state (all extra partitions assigned)
	// b. it has the minimum.
	time.Sleep(100 * time.Millisecond) // give a little gap so our last modified time is definitely greater
	afterBalanceOwnerships, err := firstProcessor.lb.LoadBalance(context.Background(), allPartitionIDs)
	require.NoError(t, err)
	require.Equal(t, 1, len(afterBalanceOwnerships))
	require.Equal(t, firstProcessorOwnerships[0].PartitionID, afterBalanceOwnerships[0].PartitionID)
	require.NotEqual(t, firstProcessorOwnerships[0].ETag, afterBalanceOwnerships[0].ETag, "ownership (etag) also gets updated each time we load balance")
	require.Greater(t, afterBalanceOwnerships[0].LastModifiedTime, firstProcessorOwnerships[0].LastModifiedTime, "ownership (last modified time) also gets updated each time we load balance")
}

func TestUnit_Processor_Run(t *testing.T) {
	cps := newCheckpointStoreForTest()

	processor, err := newProcessorImpl(simpleFakeConsumerClient(), cps, &ProcessorOptions{
		PartitionExpirationDuration: time.Hour,
	})

	require.NoError(t, err)

	procCtx, procCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer procCancel()

	partClientValue := atomic.Value{}

	go func() {
		partitionClient := processor.NextPartitionClient(context.Background())
		partClientValue.Store(partitionClient)
		procCancel()
	}()

	time.Sleep(time.Second)
	require.NoError(t, processor.Run(procCtx))

	partitionClient := partClientValue.Load().(*ProcessorPartitionClient)
	require.NotNil(t, partitionClient)
	require.Equal(t, "a", partitionClient.partitionID)
}

func TestUnit_Processor_Run_singleConsumerPerPartition(t *testing.T) {
	cps := newCheckpointStoreForTest()
	ehProps := EventHubProperties{
		PartitionIDs: []string{"a"},
	}

	partitionClientsCreated := 0

	cc := &fakeConsumerClient{
		details: consumerClientDetails{
			ConsumerGroup:           "consumer-group",
			EventHubName:            "event-hub",
			FullyQualifiedNamespace: "fqdn",
			ClientID:                "my-client-id",
		},
		getEventHubPropertiesResult: ehProps,
		newPartitionClientFn: func(partitionID string, options *PartitionClientOptions) (*PartitionClient, error) {
			partitionClientsCreated++
			return newFakePartitionClient(partitionID, ""), nil
		},
	}

	processor, err := newProcessorImpl(cc, cps, &ProcessorOptions{
		PartitionExpirationDuration: time.Hour,
	})
	require.NoError(t, err)

	consumersSyncMap := &sync.Map{}

	// to make the test easier (and less dependent on timing) we're calling through to the
	// pieces of the runImpl function
	_, err = processor.initNextClientsCh(context.Background())
	require.NoError(t, err)
	require.Empty(t, processor.nextClients)
	require.Equal(t, len(ehProps.PartitionIDs), cap(processor.nextClients))

	// the first dispatch - we have a single partition available ("a") and it gets assigned
	err = processor.dispatch(context.Background(), ehProps, consumersSyncMap)
	require.NoError(t, err)
	require.Equal(t, 1, len(processor.nextClients), "the client we created is ready to get picked up by NextPartitionClient()")

	consumers := syncMapToNormalMap(consumersSyncMap)
	origPartClient := consumers["a"]
	require.Equal(t, "a", origPartClient.partitionID)
	require.Equal(t, 1, partitionClientsCreated)

	// pull the client from the channel - it should be for the "a" partition
	procClient := processor.NextPartitionClient(context.Background())
	require.Equal(t, "a", procClient.partitionID)
	require.Empty(t, processor.nextClients)

	// the second dispatch - we reaffirm our ownership of "a" _but_ since we're already processing it no new
	// client is returned.
	err = processor.dispatch(context.Background(), ehProps, consumersSyncMap)
	require.NoError(t, err)

	// make sure we didn't create any new clients since we're already actively subscribed.
	consumers = syncMapToNormalMap(consumersSyncMap)
	afterSecondDispatchPartClient := consumers["a"]
	require.Equal(t, "a", afterSecondDispatchPartClient.partitionID)
	require.Same(t, origPartClient, afterSecondDispatchPartClient, "the client in our map is still the active one from before")
}

func TestUnit_Processor_Run_startPosition(t *testing.T) {
	cps := newCheckpointStoreForTest()

	err := cps.SetCheckpoint(context.Background(), Checkpoint{
		ConsumerGroup:           "consumer-group",
		EventHubName:            "event-hub",
		FullyQualifiedNamespace: "fqdn",
		PartitionID:             "a",
		SequenceNumber:          to.Ptr[int64](202),
	}, nil)
	require.NoError(t, err)

	fakeConsumerClient := simpleFakeConsumerClient()

	fakeConsumerClient.newPartitionClientFn = func(partitionID string, options *PartitionClientOptions) (*PartitionClient, error) {
		offsetExpr, err := getStartExpression(options.StartPosition)
		require.NoError(t, err)

		return newFakePartitionClient(partitionID, offsetExpr), nil
	}

	processor, err := newProcessorImpl(fakeConsumerClient, cps, &ProcessorOptions{
		PartitionExpirationDuration: time.Hour,
	})
	require.NoError(t, err)

	ehProps, err := processor.initNextClientsCh(context.Background())
	require.NoError(t, err)

	consumers := sync.Map{}
	err = processor.dispatch(context.Background(), ehProps, &consumers)
	require.NoError(t, err)

	checkpoints, err := cps.ListCheckpoints(context.Background(),
		processor.consumerClientDetails.FullyQualifiedNamespace,
		processor.consumerClientDetails.EventHubName,
		processor.consumerClientDetails.ConsumerGroup, nil)
	require.NoError(t, err)
	require.Equal(t, int64(202), *checkpoints[0].SequenceNumber)

	partClient := processor.NextPartitionClient(context.Background())
	require.Equal(t, "amqp.annotation.x-opt-sequence-number > '202'", partClient.innerClient.offsetExpression)

	err = partClient.UpdateCheckpoint(context.Background(), &ReceivedEventData{
		SequenceNumber: 405,
	}, nil)
	require.NoError(t, err)
	checkpoints, err = cps.ListCheckpoints(context.Background(),
		processor.consumerClientDetails.FullyQualifiedNamespace,
		processor.consumerClientDetails.EventHubName,
		processor.consumerClientDetails.ConsumerGroup, nil)
	require.NoError(t, err)
	require.Equal(t, 1, len(checkpoints))
	require.Equal(t, int64(405), *checkpoints[0].SequenceNumber)
}

func TestUnit_Processor_RunCancelledQuickly(t *testing.T) {
	cps := newCheckpointStoreForTest()
	processor := newProcessorForTest(t, "processor", cps, nil)

	precancelledContext, cancel := context.WithCancel(context.Background())
	cancel()

	require.NoError(t, processor.Run(precancelledContext))
	require.Nil(t, processor.NextPartitionClient(precancelledContext))
}

func syncMapToNormalMap(src *sync.Map) map[string]*ProcessorPartitionClient {
	dest := map[string]*ProcessorPartitionClient{}

	src.Range(func(key, value any) bool {
		dest[key.(string)] = value.(*ProcessorPartitionClient)
		return true
	})

	return dest
}

func TestUnit_Processor_Run_cancellation(t *testing.T) {
	cps := newCheckpointStoreForTest()

	processor, err := newProcessorImpl(&fakeConsumerClient{
		details: consumerClientDetails{
			ConsumerGroup:           "consumer-group",
			EventHubName:            "event-hub",
			FullyQualifiedNamespace: "fqdn",
			ClientID:                "my-client-id",
		},
	}, cps, &ProcessorOptions{
		PartitionExpirationDuration: time.Hour,
	})

	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// note that the cancellation here doesn't cause an error.
	err = processor.Run(ctx)
	require.NoError(t, err)
}

func TestUnit_Processor_getStartPosition_checkpointStore(t *testing.T) {
	ignoredSP := StartPosition{Offset: to.Ptr("IGNORED: checkpoint wins")}

	checkpointTests := []struct {
		Name           string
		Checkpoint     Checkpoint
		StartPositions StartPositions
		Actual         StartPosition
	}{
		{
			Name:           "Offset wins over Sequence, when present",
			Checkpoint:     Checkpoint{SequenceNumber: to.Ptr[int64](123), Offset: to.Ptr("offset from checkpoint")},
			StartPositions: StartPositions{Default: ignoredSP},
			Actual:         StartPosition{Offset: to.Ptr("offset from checkpoint")},
		},
		{
			Name:           "Use SequenceNumber, if no offset",
			Checkpoint:     Checkpoint{SequenceNumber: to.Ptr[int64](123)},
			StartPositions: StartPositions{Default: ignoredSP},
			Actual:         StartPosition{SequenceNumber: to.Ptr[int64](123)},
		},
		{
			Name:           "Overrides Default",
			Checkpoint:     Checkpoint{Offset: to.Ptr("offset from checkpoint")},
			StartPositions: StartPositions{Default: ignoredSP},
			Actual:         StartPosition{Offset: to.Ptr("offset from checkpoint")},
		},
		{
			Name:       "Overrides Partition Default",
			Checkpoint: Checkpoint{Offset: to.Ptr("offset from checkpoint")},
			StartPositions: StartPositions{
				PerPartition: map[string]StartPosition{"a": {Offset: to.Ptr("IGNORED: checkpoint wins")}},
			},
			Actual: StartPosition{Offset: to.Ptr("offset from checkpoint")},
		},
	}

	for _, td := range checkpointTests {
		cps := newCheckpointStoreForTest()
		td.Checkpoint.ConsumerGroup = "cg"
		td.Checkpoint.EventHubName = "eh"
		td.Checkpoint.FullyQualifiedNamespace = "ns"
		td.Checkpoint.PartitionID = "a"

		err := cps.SetCheckpoint(context.Background(), td.Checkpoint, nil)
		require.NoError(t, err)

		processor := newProcessorForTest(t, "client-id", cps, &ProcessorOptions{
			StartPositions: td.StartPositions,
		})

		called := false

		actualStartPosition, err := processor.getStartPosition(func() (map[string]Checkpoint, error) {
			called = true
			return processor.getCheckpointsMap(context.Background())
		}, Ownership{PartitionID: "a"})
		require.NoError(t, err)

		require.True(t, called)
		require.Equal(t, td.Actual, actualStartPosition)
	}

	t.Run("InvalidCheckpoint", func(t *testing.T) {
		cps := newCheckpointStoreForTest()

		err := cps.SetCheckpoint(context.Background(), Checkpoint{
			ConsumerGroup: "cg", EventHubName: "eh", FullyQualifiedNamespace: "ns", PartitionID: "a",
			// no offset or sequence number set
		}, nil)
		require.NoError(t, err)

		processor := newProcessorForTest(t, "client-id", cps, &ProcessorOptions{
			StartPositions: StartPositions{Default: ignoredSP},
		})

		called := false

		actualStartPosition, err := processor.getStartPosition(func() (map[string]Checkpoint, error) {
			called = true
			return processor.getCheckpointsMap(context.Background())
		}, Ownership{PartitionID: "a"})
		require.Empty(t, actualStartPosition)
		require.Error(t, err)
		require.True(t, called)
	})
}

func TestUnit_Processor_getStartPosition_noCheckpoint(t *testing.T) {
	t.Run("PartitionPreferredOverDefault", func(t *testing.T) {
		cps := newCheckpointStoreForTest()
		called := false

		processor := newProcessorForTest(t, "client-id", cps, &ProcessorOptions{
			StartPositions: StartPositions{
				Default: StartPosition{
					Offset: to.Ptr("default start position"),
				},
				PerPartition: map[string]StartPosition{
					"a": {
						Offset: to.Ptr("a partition start position"),
					},
				},
			},
		})

		actualStartPosition, err := processor.getStartPosition(func() (map[string]Checkpoint, error) {
			called = true
			return processor.getCheckpointsMap(context.Background())
		}, Ownership{PartitionID: "a"})
		require.NoError(t, err)
		require.True(t, called)
		require.Equal(t, StartPosition{Offset: to.Ptr("a partition start position")}, actualStartPosition)
	})

	t.Run("DefaultUsedOtherwise_HasPartitionMap", func(t *testing.T) {
		cps := newCheckpointStoreForTest()
		called := false

		processor := newProcessorForTest(t, "client-id", cps, &ProcessorOptions{
			StartPositions: StartPositions{
				Default: StartPosition{
					Offset: to.Ptr("default start position"),
				},
				PerPartition: map[string]StartPosition{
					"not-our-partition": {
						Offset: to.Ptr("ignored, doesn't match partition"),
					},
				},
			},
		})

		actualStartPosition, err := processor.getStartPosition(func() (map[string]Checkpoint, error) {
			called = true
			return processor.getCheckpointsMap(context.Background())
		}, Ownership{PartitionID: "a"})
		require.NoError(t, err)
		require.True(t, called)
		require.Equal(t, StartPosition{Offset: to.Ptr("default start position")}, actualStartPosition)
	})

	t.Run("DefaultUsedOtherwise_NoPartitionMap", func(t *testing.T) {
		cps := newCheckpointStoreForTest()
		called := false

		processor := newProcessorForTest(t, "client-id", cps, &ProcessorOptions{
			StartPositions: StartPositions{
				Default: StartPosition{
					Offset: to.Ptr("default start position"),
				},
			},
		})

		actualStartPosition, err := processor.getStartPosition(func() (map[string]Checkpoint, error) {
			called = true
			return processor.getCheckpointsMap(context.Background())
		}, Ownership{PartitionID: "a"})
		require.NoError(t, err)
		require.True(t, called)
		require.Equal(t, StartPosition{Offset: to.Ptr("default start position")}, actualStartPosition)
	})
}

func TestUnit_Processor_addPartitionClient(t *testing.T) {
	cps := newCheckpointStoreForTest()

	err := cps.SetCheckpoint(context.Background(), Checkpoint{
		ConsumerGroup: "cg", EventHubName: "eh", FullyQualifiedNamespace: "fqdn", PartitionID: "0",
		Offset: to.Ptr("100"), // ie, a legacy offset
	}, nil)
	require.NoError(t, err)

	processor := newProcessorForTest(t, "client-id", cps, &ProcessorOptions{})

	getCheckpoints := func() (map[string]Checkpoint, error) {
		return processor.getCheckpointsMap(context.Background())
	}

	geodrErr := &amqp.Error{
		Condition:   amqp.ErrCond("com.microsoft:georeplication:invalid-offset"),
		Description: "The supplied offset '1' is invalid for geo replication enabled namespace (current epoch:'1') (lots of other error text)",
	}

	t.Run("GeoDRFallback", func(t *testing.T) {
		processor.nextClients = make(chan *ProcessorPartitionClient, 1)

		gotGeodrFailure := false
		called := 0
		consumers := &sync.Map{}

		openPartitionClient := func(ctx context.Context, partitionID string, startPosition StartPosition) (partitionClient *PartitionClient, err error) {
			called++

			if startPosition.Offset != nil {
				gotGeodrFailure = true
				return nil, geodrErr
			}

			return &PartitionClient{}, nil
		}

		err = processor.addPartitionClient(context.Background(), Ownership{PartitionID: "0"}, getCheckpoints, openPartitionClient, consumers)
		require.NoError(t, err)
		require.True(t, gotGeodrFailure)
		require.Equal(t, 2, called)

		pc, _ := consumers.Load("0")
		require.NotNil(t, pc)
	})

	t.Run("Normal", func(t *testing.T) {
		processor.nextClients = make(chan *ProcessorPartitionClient, 1)

		called := 0
		consumers := &sync.Map{}

		openPartitionClient := func(ctx context.Context, partitionID string, startPosition StartPosition) (partitionClient *PartitionClient, err error) {
			called++
			return &PartitionClient{}, nil
		}

		err = processor.addPartitionClient(context.Background(), Ownership{PartitionID: "0"}, getCheckpoints, openPartitionClient, consumers)
		require.NoError(t, err)
		require.Equal(t, 1, called)

		pc, _ := consumers.Load("0")
		require.NotNil(t, pc)
	})

	t.Run("Failure", func(t *testing.T) {
		processor.nextClients = make(chan *ProcessorPartitionClient, 1)

		called := 0
		consumers := &sync.Map{}

		openPartitionClient := func(ctx context.Context, partitionID string, startPosition StartPosition) (partitionClient *PartitionClient, err error) {
			called++
			return nil, errors.New("any other error aborts immediately")
		}

		err = processor.addPartitionClient(context.Background(), Ownership{PartitionID: "0"}, getCheckpoints, openPartitionClient, consumers)
		require.EqualError(t, err, "any other error aborts immediately")
		require.Equal(t, 1, called, "We don't fall back if the error is not the specific GeoDR failure")

		pc, _ := consumers.Load("0")
		require.Nil(t, pc)
	})

	t.Run("GeoDR failure", func(t *testing.T) {
		processor.nextClients = make(chan *ProcessorPartitionClient, 1)

		called := 0
		consumers := &sync.Map{}
		gotGeodrFailure := false

		openPartitionClient := func(ctx context.Context, partitionID string, startPosition StartPosition) (partitionClient *PartitionClient, err error) {
			called++

			if startPosition.Offset != nil {
				gotGeodrFailure = true
				return nil, geodrErr
			}

			return nil, errors.New("error trying to open geodr partition client")
		}

		err = processor.addPartitionClient(context.Background(), Ownership{PartitionID: "0"}, getCheckpoints, openPartitionClient, consumers)
		require.EqualError(t, err, "error trying to open geodr partition client")
		require.Equal(t, 2, called)
		require.True(t, gotGeodrFailure)

		pc, _ := consumers.Load("0")
		require.Nil(t, pc)
	})
}

// updateDynamicData updates the passed in `expected` Ownership with any fields that are
// dynamically or randomly chosen. It returns the updated value.
func updateDynamicData(t *testing.T, src Ownership, expected Ownership, allPartitionIDs []string) Ownership {
	// these fields are dynamic (current time, randomly generated etag and randomly chosen partition ID) so we'll just copy them over so we can easily compare, after we validate they're
	// not bogus.
	require.NotEmpty(t, src.ETag)
	expected.ETag = src.ETag

	require.NotEqual(t, time.Time{}, src.LastModifiedTime)
	expected.LastModifiedTime = src.LastModifiedTime

	require.Contains(t, allPartitionIDs, src.PartitionID, "partition ID is randomly chosen but in our domain of partitions")
	expected.PartitionID = src.PartitionID

	return expected
}

func newProcessorForTest(t *testing.T, clientID string, cps CheckpointStore, userOptions *ProcessorOptions) *Processor {
	var options ProcessorOptions

	if userOptions != nil {
		options = *userOptions
	}

	options.PartitionExpirationDuration = time.Hour

	processor, err := newProcessorImpl(&fakeConsumerClient{
		details: consumerClientDetails{
			ConsumerGroup:           "consumer-group",
			EventHubName:            "event-hub",
			FullyQualifiedNamespace: "fqdn",
			ClientID:                clientID,
		},
	}, cps, &options)
	require.NoError(t, err)
	return processor
}

type fakeConsumerClient struct {
	details consumerClientDetails

	getEventHubPropertiesResult EventHubProperties
	getEventHubPropertiesErr    error

	partitionClients     map[string]newMockPartitionClientResult
	newPartitionClientFn func(partitionID string, options *PartitionClientOptions) (*PartitionClient, error)
}

type newMockPartitionClientResult struct {
	client *PartitionClient
	err    error
}

func (cc *fakeConsumerClient) GetEventHubProperties(ctx context.Context, options *GetEventHubPropertiesOptions) (EventHubProperties, error) {
	select {
	case <-ctx.Done():
		return EventHubProperties{}, ctx.Err()
	default:
		return cc.getEventHubPropertiesResult, cc.getEventHubPropertiesErr
	}
}

func (cc *fakeConsumerClient) NewPartitionClient(partitionID string, options *PartitionClientOptions) (*PartitionClient, error) {
	if *options.OwnerLevel != 0 {
		panic(fmt.Sprintf("Invalid owner level passed for the Processor: got %d instead of 0", *options.OwnerLevel))
	}

	if cc.newPartitionClientFn != nil {
		return cc.newPartitionClientFn(partitionID, options)
	}

	if cc.partitionClients == nil {
		panic("bad test, no partition clients defined")
	}

	value, exists := cc.partitionClients[partitionID]

	if !exists {
		panic(fmt.Sprintf("bad test, partition client needed for partition %s but didn't exist in test map", partitionID))
	}

	return value.client, value.err
}

func (cc *fakeConsumerClient) getDetails() consumerClientDetails {
	return cc.details
}

func simpleFakeConsumerClient() *fakeConsumerClient {
	return &fakeConsumerClient{
		details: consumerClientDetails{
			ConsumerGroup:           "consumer-group",
			EventHubName:            "event-hub",
			FullyQualifiedNamespace: "fqdn",
			ClientID:                "my-client-id",
		},
		getEventHubPropertiesResult: EventHubProperties{
			PartitionIDs: []string{"a"},
		},
		partitionClients: map[string]newMockPartitionClientResult{
			"a": {
				client: newFakePartitionClient("a", ""),
				err:    nil,
			},
		},
	}
}

type fakeLinksForPartitionClient struct {
	internal.LinksForPartitionClient[amqpwrap.AMQPReceiverCloser]
}

func (fc *fakeLinksForPartitionClient) Retry(ctx context.Context, eventName log.Event, operation string, partitionID string, retryOptions exported.RetryOptions, fn func(ctx context.Context, lwid internal.LinkWithID[amqpwrap.AMQPReceiverCloser]) error) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}

func (fc *fakeLinksForPartitionClient) Close(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}

func newFakePartitionClient(partitionID string, offsetExpr string) *PartitionClient {
	return &PartitionClient{
		partitionID:      partitionID,
		offsetExpression: offsetExpr,
		links:            &fakeLinksForPartitionClient{},
	}
}
