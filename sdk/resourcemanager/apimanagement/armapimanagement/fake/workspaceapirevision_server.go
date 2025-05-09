// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
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

// WorkspaceAPIRevisionServer is a fake server for instances of the armapimanagement.WorkspaceAPIRevisionClient type.
type WorkspaceAPIRevisionServer struct {
	// NewListByServicePager is the fake for method WorkspaceAPIRevisionClient.NewListByServicePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByServicePager func(resourceGroupName string, serviceName string, workspaceID string, apiID string, options *armapimanagement.WorkspaceAPIRevisionClientListByServiceOptions) (resp azfake.PagerResponder[armapimanagement.WorkspaceAPIRevisionClientListByServiceResponse])
}

// NewWorkspaceAPIRevisionServerTransport creates a new instance of WorkspaceAPIRevisionServerTransport with the provided implementation.
// The returned WorkspaceAPIRevisionServerTransport instance is connected to an instance of armapimanagement.WorkspaceAPIRevisionClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewWorkspaceAPIRevisionServerTransport(srv *WorkspaceAPIRevisionServer) *WorkspaceAPIRevisionServerTransport {
	return &WorkspaceAPIRevisionServerTransport{
		srv:                   srv,
		newListByServicePager: newTracker[azfake.PagerResponder[armapimanagement.WorkspaceAPIRevisionClientListByServiceResponse]](),
	}
}

// WorkspaceAPIRevisionServerTransport connects instances of armapimanagement.WorkspaceAPIRevisionClient to instances of WorkspaceAPIRevisionServer.
// Don't use this type directly, use NewWorkspaceAPIRevisionServerTransport instead.
type WorkspaceAPIRevisionServerTransport struct {
	srv                   *WorkspaceAPIRevisionServer
	newListByServicePager *tracker[azfake.PagerResponder[armapimanagement.WorkspaceAPIRevisionClientListByServiceResponse]]
}

// Do implements the policy.Transporter interface for WorkspaceAPIRevisionServerTransport.
func (w *WorkspaceAPIRevisionServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return w.dispatchToMethodFake(req, method)
}

func (w *WorkspaceAPIRevisionServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if workspaceApiRevisionServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = workspaceApiRevisionServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "WorkspaceAPIRevisionClient.NewListByServicePager":
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

func (w *WorkspaceAPIRevisionServerTransport) dispatchNewListByServicePager(req *http.Request) (*http.Response, error) {
	if w.srv.NewListByServicePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByServicePager not implemented")}
	}
	newListByServicePager := w.newListByServicePager.get(req)
	if newListByServicePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/apis/(?P<apiId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/revisions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
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
		apiIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("apiId")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
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
		var options *armapimanagement.WorkspaceAPIRevisionClientListByServiceOptions
		if filterParam != nil || topParam != nil || skipParam != nil {
			options = &armapimanagement.WorkspaceAPIRevisionClientListByServiceOptions{
				Filter: filterParam,
				Top:    topParam,
				Skip:   skipParam,
			}
		}
		resp := w.srv.NewListByServicePager(resourceGroupNameParam, serviceNameParam, workspaceIDParam, apiIDParam, options)
		newListByServicePager = &resp
		w.newListByServicePager.add(req, newListByServicePager)
		server.PagerResponderInjectNextLinks(newListByServicePager, req, func(page *armapimanagement.WorkspaceAPIRevisionClientListByServiceResponse, createLink func() string) {
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

// set this to conditionally intercept incoming requests to WorkspaceAPIRevisionServerTransport
var workspaceApiRevisionServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
