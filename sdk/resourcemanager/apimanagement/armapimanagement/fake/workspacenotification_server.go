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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// WorkspaceNotificationServer is a fake server for instances of the armapimanagement.WorkspaceNotificationClient type.
type WorkspaceNotificationServer struct {
	// CreateOrUpdate is the fake for method WorkspaceNotificationClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, notificationName armapimanagement.NotificationName, options *armapimanagement.WorkspaceNotificationClientCreateOrUpdateOptions) (resp azfake.Responder[armapimanagement.WorkspaceNotificationClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method WorkspaceNotificationClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, notificationName armapimanagement.NotificationName, options *armapimanagement.WorkspaceNotificationClientGetOptions) (resp azfake.Responder[armapimanagement.WorkspaceNotificationClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByServicePager is the fake for method WorkspaceNotificationClient.NewListByServicePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByServicePager func(resourceGroupName string, serviceName string, workspaceID string, options *armapimanagement.WorkspaceNotificationClientListByServiceOptions) (resp azfake.PagerResponder[armapimanagement.WorkspaceNotificationClientListByServiceResponse])
}

// NewWorkspaceNotificationServerTransport creates a new instance of WorkspaceNotificationServerTransport with the provided implementation.
// The returned WorkspaceNotificationServerTransport instance is connected to an instance of armapimanagement.WorkspaceNotificationClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewWorkspaceNotificationServerTransport(srv *WorkspaceNotificationServer) *WorkspaceNotificationServerTransport {
	return &WorkspaceNotificationServerTransport{
		srv:                   srv,
		newListByServicePager: newTracker[azfake.PagerResponder[armapimanagement.WorkspaceNotificationClientListByServiceResponse]](),
	}
}

// WorkspaceNotificationServerTransport connects instances of armapimanagement.WorkspaceNotificationClient to instances of WorkspaceNotificationServer.
// Don't use this type directly, use NewWorkspaceNotificationServerTransport instead.
type WorkspaceNotificationServerTransport struct {
	srv                   *WorkspaceNotificationServer
	newListByServicePager *tracker[azfake.PagerResponder[armapimanagement.WorkspaceNotificationClientListByServiceResponse]]
}

// Do implements the policy.Transporter interface for WorkspaceNotificationServerTransport.
func (w *WorkspaceNotificationServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return w.dispatchToMethodFake(req, method)
}

func (w *WorkspaceNotificationServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if workspaceNotificationServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = workspaceNotificationServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "WorkspaceNotificationClient.CreateOrUpdate":
				res.resp, res.err = w.dispatchCreateOrUpdate(req)
			case "WorkspaceNotificationClient.Get":
				res.resp, res.err = w.dispatchGet(req)
			case "WorkspaceNotificationClient.NewListByServicePager":
				res.resp, res.err = w.dispatchNewListByServicePager(req)
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

func (w *WorkspaceNotificationServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if w.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/notifications/(?P<notificationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	notificationNameParam, err := parseWithCast(matches[regex.SubexpIndex("notificationName")], func(v string) (armapimanagement.NotificationName, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armapimanagement.NotificationName(p), nil
	})
	if err != nil {
		return nil, err
	}
	ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
	var options *armapimanagement.WorkspaceNotificationClientCreateOrUpdateOptions
	if ifMatchParam != nil {
		options = &armapimanagement.WorkspaceNotificationClientCreateOrUpdateOptions{
			IfMatch: ifMatchParam,
		}
	}
	respr, errRespr := w.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, serviceNameParam, workspaceIDParam, notificationNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NotificationContract, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WorkspaceNotificationServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if w.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/notifications/(?P<notificationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	notificationNameParam, err := parseWithCast(matches[regex.SubexpIndex("notificationName")], func(v string) (armapimanagement.NotificationName, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armapimanagement.NotificationName(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.Get(req.Context(), resourceGroupNameParam, serviceNameParam, workspaceIDParam, notificationNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NotificationContract, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WorkspaceNotificationServerTransport) dispatchNewListByServicePager(req *http.Request) (*http.Response, error) {
	if w.srv.NewListByServicePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByServicePager not implemented")}
	}
	newListByServicePager := w.newListByServicePager.get(req)
	if newListByServicePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/notifications`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
		if err != nil {
			return nil, err
		}
		workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
		if err != nil {
			return nil, err
		}
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		skipUnescaped, err := url.QueryUnescape(qp.Get("$skip"))
		if err != nil {
			return nil, err
		}
		skipParam, err := parseOptional(skipUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		var options *armapimanagement.WorkspaceNotificationClientListByServiceOptions
		if topParam != nil || skipParam != nil {
			options = &armapimanagement.WorkspaceNotificationClientListByServiceOptions{
				Top:  topParam,
				Skip: skipParam,
			}
		}
		resp := w.srv.NewListByServicePager(resourceGroupNameParam, serviceNameParam, workspaceIDParam, options)
		newListByServicePager = &resp
		w.newListByServicePager.add(req, newListByServicePager)
		server.PagerResponderInjectNextLinks(newListByServicePager, req, func(page *armapimanagement.WorkspaceNotificationClientListByServiceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByServicePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		w.newListByServicePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByServicePager) {
		w.newListByServicePager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to WorkspaceNotificationServerTransport
var workspaceNotificationServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
