//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmanagedapplications_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armmanagedapplications"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/resources/resource-manager/Microsoft.Solutions/stable/2018-06-01/examples/getApplication.json
func ExampleApplicationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagedapplications.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewApplicationsClient().Get(ctx, "rg", "myManagedApplication", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Application = armmanagedapplications.Application{
	// 	Name: to.Ptr("myManagedApplication"),
	// 	Type: to.Ptr("Microsoft.Solutions/applications"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applications/myManagedApplication"),
	// 	Location: to.Ptr("East US 2"),
	// 	Kind: to.Ptr("ServiceCatalog"),
	// 	Properties: &armmanagedapplications.ApplicationProperties{
	// 		ManagedResourceGroupID: to.Ptr("/subscriptions/subid/resourceGroups/myManagedRG"),
	// 		ProvisioningState: to.Ptr(armmanagedapplications.ProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/resources/resource-manager/Microsoft.Solutions/stable/2018-06-01/examples/deleteApplication.json
func ExampleApplicationsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagedapplications.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewApplicationsClient().BeginDelete(ctx, "rg", "myManagedApplication", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/resources/resource-manager/Microsoft.Solutions/stable/2018-06-01/examples/createOrUpdateApplication.json
func ExampleApplicationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagedapplications.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewApplicationsClient().BeginCreateOrUpdate(ctx, "rg", "myManagedApplication", armmanagedapplications.Application{
		Location: to.Ptr("East US 2"),
		Kind:     to.Ptr("ServiceCatalog"),
		Properties: &armmanagedapplications.ApplicationProperties{
			ManagedResourceGroupID: to.Ptr("/subscriptions/subid/resourceGroups/myManagedRG"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Application = armmanagedapplications.Application{
	// 	Name: to.Ptr("myManagedApplication"),
	// 	Type: to.Ptr("Microsoft.Solutions/applications"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applications/myManagedApplication"),
	// 	Location: to.Ptr("East US 2"),
	// 	Kind: to.Ptr("ServiceCatalog"),
	// 	Properties: &armmanagedapplications.ApplicationProperties{
	// 		ManagedResourceGroupID: to.Ptr("/subscriptions/subid/resourceGroups/myManagedRG"),
	// 		ProvisioningState: to.Ptr(armmanagedapplications.ProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/resources/resource-manager/Microsoft.Solutions/stable/2018-06-01/examples/updateApplication.json
func ExampleApplicationsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagedapplications.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewApplicationsClient().Update(ctx, "rg", "myManagedApplication", &armmanagedapplications.ApplicationsClientUpdateOptions{Parameters: &armmanagedapplications.ApplicationPatchable{
		Kind: to.Ptr("ServiceCatalog"),
		Properties: &armmanagedapplications.ApplicationPropertiesPatchable{
			ApplicationDefinitionID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applicationDefinitions/myAppDef"),
			ManagedResourceGroupID:  to.Ptr("/subscriptions/subid/resourceGroups/myManagedRG"),
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Application = armmanagedapplications.Application{
	// 	Name: to.Ptr("myManagedApplication"),
	// 	Type: to.Ptr("Microsoft.Solutions/applications"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applications/myManagedApplication"),
	// 	Kind: to.Ptr("ServiceCatalog"),
	// 	Properties: &armmanagedapplications.ApplicationProperties{
	// 		ApplicationDefinitionID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applicationDefinitions/myAppDef"),
	// 		ManagedResourceGroupID: to.Ptr("/subscriptions/subid/resourceGroups/myManagedRG"),
	// 		ProvisioningState: to.Ptr(armmanagedapplications.ProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/resources/resource-manager/Microsoft.Solutions/stable/2018-06-01/examples/listApplicationsByResourceGroup.json
func ExampleApplicationsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagedapplications.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewApplicationsClient().NewListByResourceGroupPager("rg", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ApplicationListResult = armmanagedapplications.ApplicationListResult{
		// 	Value: []*armmanagedapplications.Application{
		// 		{
		// 			Name: to.Ptr("myManagedApplication"),
		// 			Type: to.Ptr("Microsoft.Solutions/applications"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applications/myManagedApplication"),
		// 			Location: to.Ptr("East US 2"),
		// 			Kind: to.Ptr("ServiceCatalog"),
		// 			Properties: &armmanagedapplications.ApplicationProperties{
		// 				ManagedResourceGroupID: to.Ptr("/subscriptions/subid/resourceGroups/myManagedRG"),
		// 				ProvisioningState: to.Ptr(armmanagedapplications.ProvisioningStateSucceeded),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myManagedApplication2"),
		// 			Type: to.Ptr("Microsoft.Solutions/applications"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applications/myManagedApplication2"),
		// 			Location: to.Ptr("West US"),
		// 			Kind: to.Ptr("ServiceCatalog"),
		// 			Properties: &armmanagedapplications.ApplicationProperties{
		// 				ManagedResourceGroupID: to.Ptr("/subscriptions/subid/resourceGroups/myManagedRG"),
		// 				ProvisioningState: to.Ptr(armmanagedapplications.ProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/resources/resource-manager/Microsoft.Solutions/stable/2018-06-01/examples/listApplicationsBySubscription.json
func ExampleApplicationsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagedapplications.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewApplicationsClient().NewListBySubscriptionPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ApplicationListResult = armmanagedapplications.ApplicationListResult{
		// 	Value: []*armmanagedapplications.Application{
		// 		{
		// 			Name: to.Ptr("myManagedApplication"),
		// 			Type: to.Ptr("Microsoft.Solutions/applications"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applications/myManagedApplication"),
		// 			Location: to.Ptr("East US 2"),
		// 			Kind: to.Ptr("ServiceCatalog"),
		// 			Properties: &armmanagedapplications.ApplicationProperties{
		// 				ManagedResourceGroupID: to.Ptr("/subscriptions/subid/resourceGroups/myManagedRG"),
		// 				ProvisioningState: to.Ptr(armmanagedapplications.ProvisioningStateSucceeded),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myManagedApplication2"),
		// 			Type: to.Ptr("Microsoft.Solutions/applications"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applications/myManagedApplication2"),
		// 			Location: to.Ptr("West US"),
		// 			Kind: to.Ptr("ServiceCatalog"),
		// 			Properties: &armmanagedapplications.ApplicationProperties{
		// 				ManagedResourceGroupID: to.Ptr("/subscriptions/subid/resourceGroups/myManagedRG"),
		// 				ProvisioningState: to.Ptr(armmanagedapplications.ProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/resources/resource-manager/Microsoft.Solutions/stable/2018-06-01/examples/getApplicationById.json
func ExampleApplicationsClient_GetByID() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagedapplications.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewApplicationsClient().GetByID(ctx, "myApplicationId", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Application = armmanagedapplications.Application{
	// 	Name: to.Ptr("myManagedApplication"),
	// 	Type: to.Ptr("Microsoft.Solutions/applications"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applications/myManagedApplication"),
	// 	Location: to.Ptr("East US 2"),
	// 	Kind: to.Ptr("ServiceCatalog"),
	// 	Properties: &armmanagedapplications.ApplicationProperties{
	// 		ManagedResourceGroupID: to.Ptr("/subscriptions/subid/resourceGroups/myManagedRG"),
	// 		ProvisioningState: to.Ptr(armmanagedapplications.ProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/resources/resource-manager/Microsoft.Solutions/stable/2018-06-01/examples/deleteApplicationById.json
func ExampleApplicationsClient_BeginDeleteByID() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagedapplications.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewApplicationsClient().BeginDeleteByID(ctx, "myApplicationId", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/resources/resource-manager/Microsoft.Solutions/stable/2018-06-01/examples/createOrUpdateApplicationById.json
func ExampleApplicationsClient_BeginCreateOrUpdateByID() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagedapplications.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewApplicationsClient().BeginCreateOrUpdateByID(ctx, "myApplicationId", armmanagedapplications.Application{
		Location: to.Ptr("East US 2"),
		Kind:     to.Ptr("ServiceCatalog"),
		Properties: &armmanagedapplications.ApplicationProperties{
			ManagedResourceGroupID: to.Ptr("/subscriptions/subid/resourceGroups/myManagedRG"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Application = armmanagedapplications.Application{
	// 	Name: to.Ptr("myManagedApplication"),
	// 	Type: to.Ptr("Microsoft.Solutions/applications"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applications/myManagedApplication"),
	// 	Location: to.Ptr("East US 2"),
	// 	Kind: to.Ptr("ServiceCatalog"),
	// 	Properties: &armmanagedapplications.ApplicationProperties{
	// 		ManagedResourceGroupID: to.Ptr("/subscriptions/subid/resourceGroups/myManagedRG"),
	// 		ProvisioningState: to.Ptr(armmanagedapplications.ProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/resources/resource-manager/Microsoft.Solutions/stable/2018-06-01/examples/updateApplicationById.json
func ExampleApplicationsClient_UpdateByID() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagedapplications.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewApplicationsClient().UpdateByID(ctx, "myApplicationId", &armmanagedapplications.ApplicationsClientUpdateByIDOptions{Parameters: &armmanagedapplications.Application{
		Kind: to.Ptr("ServiceCatalog"),
		Properties: &armmanagedapplications.ApplicationProperties{
			ApplicationDefinitionID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applicationDefinitions/myAppDef"),
			ManagedResourceGroupID:  to.Ptr("/subscriptions/subid/resourceGroups/myManagedRG"),
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Application = armmanagedapplications.Application{
	// 	Name: to.Ptr("myManagedApplication"),
	// 	Type: to.Ptr("Microsoft.Solutions/applications"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applications/myManagedApplication"),
	// 	Kind: to.Ptr("ServiceCatalog"),
	// 	Properties: &armmanagedapplications.ApplicationProperties{
	// 		ApplicationDefinitionID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applicationDefinitions/myAppDef"),
	// 		ManagedResourceGroupID: to.Ptr("/subscriptions/subid/resourceGroups/myManagedRG"),
	// 		ProvisioningState: to.Ptr(armmanagedapplications.ProvisioningStateSucceeded),
	// 	},
	// }
}