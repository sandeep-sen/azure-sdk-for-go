// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
)

// JobsServer is a fake server for instances of the armappcontainers.JobsClient type.
type JobsServer struct {
	// BeginCreateOrUpdate is the fake for method JobsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, jobName string, jobEnvelope armappcontainers.Job, options *armappcontainers.JobsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armappcontainers.JobsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method JobsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, jobName string, options *armappcontainers.JobsClientBeginDeleteOptions) (resp azfake.PollerResponder[armappcontainers.JobsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method JobsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, jobName string, options *armappcontainers.JobsClientGetOptions) (resp azfake.Responder[armappcontainers.JobsClientGetResponse], errResp azfake.ErrorResponder)

	// GetDetector is the fake for method JobsClient.GetDetector
	// HTTP status codes to indicate success: http.StatusOK
	GetDetector func(ctx context.Context, resourceGroupName string, jobName string, detectorName string, options *armappcontainers.JobsClientGetDetectorOptions) (resp azfake.Responder[armappcontainers.JobsClientGetDetectorResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method JobsClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armappcontainers.JobsClientListByResourceGroupOptions) (resp azfake.PagerResponder[armappcontainers.JobsClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method JobsClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armappcontainers.JobsClientListBySubscriptionOptions) (resp azfake.PagerResponder[armappcontainers.JobsClientListBySubscriptionResponse])

	// NewListDetectorsPager is the fake for method JobsClient.NewListDetectorsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListDetectorsPager func(resourceGroupName string, jobName string, options *armappcontainers.JobsClientListDetectorsOptions) (resp azfake.PagerResponder[armappcontainers.JobsClientListDetectorsResponse])

	// ListSecrets is the fake for method JobsClient.ListSecrets
	// HTTP status codes to indicate success: http.StatusOK
	ListSecrets func(ctx context.Context, resourceGroupName string, jobName string, options *armappcontainers.JobsClientListSecretsOptions) (resp azfake.Responder[armappcontainers.JobsClientListSecretsResponse], errResp azfake.ErrorResponder)

	// ProxyGet is the fake for method JobsClient.ProxyGet
	// HTTP status codes to indicate success: http.StatusOK
	ProxyGet func(ctx context.Context, resourceGroupName string, jobName string, apiName string, options *armappcontainers.JobsClientProxyGetOptions) (resp azfake.Responder[armappcontainers.JobsClientProxyGetResponse], errResp azfake.ErrorResponder)

	// BeginStart is the fake for method JobsClient.BeginStart
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginStart func(ctx context.Context, resourceGroupName string, jobName string, options *armappcontainers.JobsClientBeginStartOptions) (resp azfake.PollerResponder[armappcontainers.JobsClientStartResponse], errResp azfake.ErrorResponder)

	// BeginStopExecution is the fake for method JobsClient.BeginStopExecution
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginStopExecution func(ctx context.Context, resourceGroupName string, jobName string, jobExecutionName string, options *armappcontainers.JobsClientBeginStopExecutionOptions) (resp azfake.PollerResponder[armappcontainers.JobsClientStopExecutionResponse], errResp azfake.ErrorResponder)

	// BeginStopMultipleExecutions is the fake for method JobsClient.BeginStopMultipleExecutions
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginStopMultipleExecutions func(ctx context.Context, resourceGroupName string, jobName string, options *armappcontainers.JobsClientBeginStopMultipleExecutionsOptions) (resp azfake.PollerResponder[armappcontainers.JobsClientStopMultipleExecutionsResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method JobsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, jobName string, jobEnvelope armappcontainers.JobPatchProperties, options *armappcontainers.JobsClientBeginUpdateOptions) (resp azfake.PollerResponder[armappcontainers.JobsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewJobsServerTransport creates a new instance of JobsServerTransport with the provided implementation.
// The returned JobsServerTransport instance is connected to an instance of armappcontainers.JobsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewJobsServerTransport(srv *JobsServer) *JobsServerTransport {
	return &JobsServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armappcontainers.JobsClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armappcontainers.JobsClientDeleteResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armappcontainers.JobsClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armappcontainers.JobsClientListBySubscriptionResponse]](),
		newListDetectorsPager:       newTracker[azfake.PagerResponder[armappcontainers.JobsClientListDetectorsResponse]](),
		beginStart:                  newTracker[azfake.PollerResponder[armappcontainers.JobsClientStartResponse]](),
		beginStopExecution:          newTracker[azfake.PollerResponder[armappcontainers.JobsClientStopExecutionResponse]](),
		beginStopMultipleExecutions: newTracker[azfake.PollerResponder[armappcontainers.JobsClientStopMultipleExecutionsResponse]](),
		beginUpdate:                 newTracker[azfake.PollerResponder[armappcontainers.JobsClientUpdateResponse]](),
	}
}

// JobsServerTransport connects instances of armappcontainers.JobsClient to instances of JobsServer.
// Don't use this type directly, use NewJobsServerTransport instead.
type JobsServerTransport struct {
	srv                         *JobsServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armappcontainers.JobsClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armappcontainers.JobsClientDeleteResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armappcontainers.JobsClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armappcontainers.JobsClientListBySubscriptionResponse]]
	newListDetectorsPager       *tracker[azfake.PagerResponder[armappcontainers.JobsClientListDetectorsResponse]]
	beginStart                  *tracker[azfake.PollerResponder[armappcontainers.JobsClientStartResponse]]
	beginStopExecution          *tracker[azfake.PollerResponder[armappcontainers.JobsClientStopExecutionResponse]]
	beginStopMultipleExecutions *tracker[azfake.PollerResponder[armappcontainers.JobsClientStopMultipleExecutionsResponse]]
	beginUpdate                 *tracker[azfake.PollerResponder[armappcontainers.JobsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for JobsServerTransport.
func (j *JobsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return j.dispatchToMethodFake(req, method)
}

func (j *JobsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if jobsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = jobsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "JobsClient.BeginCreateOrUpdate":
				res.resp, res.err = j.dispatchBeginCreateOrUpdate(req)
			case "JobsClient.BeginDelete":
				res.resp, res.err = j.dispatchBeginDelete(req)
			case "JobsClient.Get":
				res.resp, res.err = j.dispatchGet(req)
			case "JobsClient.GetDetector":
				res.resp, res.err = j.dispatchGetDetector(req)
			case "JobsClient.NewListByResourceGroupPager":
				res.resp, res.err = j.dispatchNewListByResourceGroupPager(req)
			case "JobsClient.NewListBySubscriptionPager":
				res.resp, res.err = j.dispatchNewListBySubscriptionPager(req)
			case "JobsClient.NewListDetectorsPager":
				res.resp, res.err = j.dispatchNewListDetectorsPager(req)
			case "JobsClient.ListSecrets":
				res.resp, res.err = j.dispatchListSecrets(req)
			case "JobsClient.ProxyGet":
				res.resp, res.err = j.dispatchProxyGet(req)
			case "JobsClient.BeginStart":
				res.resp, res.err = j.dispatchBeginStart(req)
			case "JobsClient.BeginStopExecution":
				res.resp, res.err = j.dispatchBeginStopExecution(req)
			case "JobsClient.BeginStopMultipleExecutions":
				res.resp, res.err = j.dispatchBeginStopMultipleExecutions(req)
			case "JobsClient.BeginUpdate":
				res.resp, res.err = j.dispatchBeginUpdate(req)
			default:
				res.err = fmt.Errorf("unhandled API %s", method)
			}

		}
		select {
		case resultChan <- res:
		case <-req.Context().Done():
		}
	}()

	select {
	case <-req.Context().Done():
		return nil, req.Context().Err()
	case res := <-resultChan:
		return res.resp, res.err
	}
}

func (j *JobsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if j.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := j.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/jobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armappcontainers.Job](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := j.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, jobNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		j.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		j.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		j.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (j *JobsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if j.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := j.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/jobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := j.srv.BeginDelete(req.Context(), resourceGroupNameParam, jobNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		j.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		j.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		j.beginDelete.remove(req)
	}

	return resp, nil
}

func (j *JobsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if j.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/jobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := j.srv.Get(req.Context(), resourceGroupNameParam, jobNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Job, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (j *JobsServerTransport) dispatchGetDetector(req *http.Request) (*http.Response, error) {
	if j.srv.GetDetector == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetDetector not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/jobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/detectors/(?P<detectorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
	if err != nil {
		return nil, err
	}
	detectorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("detectorName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := j.srv.GetDetector(req.Context(), resourceGroupNameParam, jobNameParam, detectorNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Diagnostics, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (j *JobsServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if j.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := j.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/jobs`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := j.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		j.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armappcontainers.JobsClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		j.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		j.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (j *JobsServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if j.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := j.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/jobs`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := j.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		j.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armappcontainers.JobsClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		j.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		j.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (j *JobsServerTransport) dispatchNewListDetectorsPager(req *http.Request) (*http.Response, error) {
	if j.srv.NewListDetectorsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListDetectorsPager not implemented")}
	}
	newListDetectorsPager := j.newListDetectorsPager.get(req)
	if newListDetectorsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/jobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/detectors`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
		if err != nil {
			return nil, err
		}
		resp := j.srv.NewListDetectorsPager(resourceGroupNameParam, jobNameParam, nil)
		newListDetectorsPager = &resp
		j.newListDetectorsPager.add(req, newListDetectorsPager)
		server.PagerResponderInjectNextLinks(newListDetectorsPager, req, func(page *armappcontainers.JobsClientListDetectorsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListDetectorsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		j.newListDetectorsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListDetectorsPager) {
		j.newListDetectorsPager.remove(req)
	}
	return resp, nil
}

func (j *JobsServerTransport) dispatchListSecrets(req *http.Request) (*http.Response, error) {
	if j.srv.ListSecrets == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListSecrets not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/jobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listSecrets`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := j.srv.ListSecrets(req.Context(), resourceGroupNameParam, jobNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).JobSecretsCollection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (j *JobsServerTransport) dispatchProxyGet(req *http.Request) (*http.Response, error) {
	if j.srv.ProxyGet == nil {
		return nil, &nonRetriableError{errors.New("fake for method ProxyGet not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/jobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/detectorProperties/(?P<apiName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
	if err != nil {
		return nil, err
	}
	apiNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("apiName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := j.srv.ProxyGet(req.Context(), resourceGroupNameParam, jobNameParam, apiNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Job, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (j *JobsServerTransport) dispatchBeginStart(req *http.Request) (*http.Response, error) {
	if j.srv.BeginStart == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginStart not implemented")}
	}
	beginStart := j.beginStart.get(req)
	if beginStart == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/jobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/start`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armappcontainers.JobExecutionTemplate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
		if err != nil {
			return nil, err
		}
		var options *armappcontainers.JobsClientBeginStartOptions
		if !reflect.ValueOf(body).IsZero() {
			options = &armappcontainers.JobsClientBeginStartOptions{
				Template: &body,
			}
		}
		respr, errRespr := j.srv.BeginStart(req.Context(), resourceGroupNameParam, jobNameParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginStart = &respr
		j.beginStart.add(req, beginStart)
	}

	resp, err := server.PollerResponderNext(beginStart, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		j.beginStart.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginStart) {
		j.beginStart.remove(req)
	}

	return resp, nil
}

func (j *JobsServerTransport) dispatchBeginStopExecution(req *http.Request) (*http.Response, error) {
	if j.srv.BeginStopExecution == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginStopExecution not implemented")}
	}
	beginStopExecution := j.beginStopExecution.get(req)
	if beginStopExecution == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/jobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/executions/(?P<jobExecutionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/stop`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
		if err != nil {
			return nil, err
		}
		jobExecutionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobExecutionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := j.srv.BeginStopExecution(req.Context(), resourceGroupNameParam, jobNameParam, jobExecutionNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginStopExecution = &respr
		j.beginStopExecution.add(req, beginStopExecution)
	}

	resp, err := server.PollerResponderNext(beginStopExecution, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		j.beginStopExecution.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginStopExecution) {
		j.beginStopExecution.remove(req)
	}

	return resp, nil
}

func (j *JobsServerTransport) dispatchBeginStopMultipleExecutions(req *http.Request) (*http.Response, error) {
	if j.srv.BeginStopMultipleExecutions == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginStopMultipleExecutions not implemented")}
	}
	beginStopMultipleExecutions := j.beginStopMultipleExecutions.get(req)
	if beginStopMultipleExecutions == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/jobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/stop`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := j.srv.BeginStopMultipleExecutions(req.Context(), resourceGroupNameParam, jobNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginStopMultipleExecutions = &respr
		j.beginStopMultipleExecutions.add(req, beginStopMultipleExecutions)
	}

	resp, err := server.PollerResponderNext(beginStopMultipleExecutions, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		j.beginStopMultipleExecutions.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginStopMultipleExecutions) {
		j.beginStopMultipleExecutions.remove(req)
	}

	return resp, nil
}

func (j *JobsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if j.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := j.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/jobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armappcontainers.JobPatchProperties](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := j.srv.BeginUpdate(req.Context(), resourceGroupNameParam, jobNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		j.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		j.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		j.beginUpdate.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to JobsServerTransport
var jobsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
