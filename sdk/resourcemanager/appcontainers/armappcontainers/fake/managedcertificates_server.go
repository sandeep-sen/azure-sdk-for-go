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

// ManagedCertificatesServer is a fake server for instances of the armappcontainers.ManagedCertificatesClient type.
type ManagedCertificatesServer struct {
	// BeginCreateOrUpdate is the fake for method ManagedCertificatesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, environmentName string, managedCertificateName string, options *armappcontainers.ManagedCertificatesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armappcontainers.ManagedCertificatesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method ManagedCertificatesClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, environmentName string, managedCertificateName string, options *armappcontainers.ManagedCertificatesClientDeleteOptions) (resp azfake.Responder[armappcontainers.ManagedCertificatesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ManagedCertificatesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, environmentName string, managedCertificateName string, options *armappcontainers.ManagedCertificatesClientGetOptions) (resp azfake.Responder[armappcontainers.ManagedCertificatesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ManagedCertificatesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, environmentName string, options *armappcontainers.ManagedCertificatesClientListOptions) (resp azfake.PagerResponder[armappcontainers.ManagedCertificatesClientListResponse])

	// Update is the fake for method ManagedCertificatesClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, environmentName string, managedCertificateName string, managedCertificateEnvelope armappcontainers.ManagedCertificatePatch, options *armappcontainers.ManagedCertificatesClientUpdateOptions) (resp azfake.Responder[armappcontainers.ManagedCertificatesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewManagedCertificatesServerTransport creates a new instance of ManagedCertificatesServerTransport with the provided implementation.
// The returned ManagedCertificatesServerTransport instance is connected to an instance of armappcontainers.ManagedCertificatesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewManagedCertificatesServerTransport(srv *ManagedCertificatesServer) *ManagedCertificatesServerTransport {
	return &ManagedCertificatesServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armappcontainers.ManagedCertificatesClientCreateOrUpdateResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armappcontainers.ManagedCertificatesClientListResponse]](),
	}
}

// ManagedCertificatesServerTransport connects instances of armappcontainers.ManagedCertificatesClient to instances of ManagedCertificatesServer.
// Don't use this type directly, use NewManagedCertificatesServerTransport instead.
type ManagedCertificatesServerTransport struct {
	srv                 *ManagedCertificatesServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armappcontainers.ManagedCertificatesClientCreateOrUpdateResponse]]
	newListPager        *tracker[azfake.PagerResponder[armappcontainers.ManagedCertificatesClientListResponse]]
}

// Do implements the policy.Transporter interface for ManagedCertificatesServerTransport.
func (m *ManagedCertificatesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return m.dispatchToMethodFake(req, method)
}

func (m *ManagedCertificatesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if managedCertificatesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = managedCertificatesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ManagedCertificatesClient.BeginCreateOrUpdate":
				res.resp, res.err = m.dispatchBeginCreateOrUpdate(req)
			case "ManagedCertificatesClient.Delete":
				res.resp, res.err = m.dispatchDelete(req)
			case "ManagedCertificatesClient.Get":
				res.resp, res.err = m.dispatchGet(req)
			case "ManagedCertificatesClient.NewListPager":
				res.resp, res.err = m.dispatchNewListPager(req)
			case "ManagedCertificatesClient.Update":
				res.resp, res.err = m.dispatchUpdate(req)
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

func (m *ManagedCertificatesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if m.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := m.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/managedEnvironments/(?P<environmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/managedCertificates/(?P<managedCertificateName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armappcontainers.ManagedCertificate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		environmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("environmentName")])
		if err != nil {
			return nil, err
		}
		managedCertificateNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedCertificateName")])
		if err != nil {
			return nil, err
		}
		var options *armappcontainers.ManagedCertificatesClientBeginCreateOrUpdateOptions
		if !reflect.ValueOf(body).IsZero() {
			options = &armappcontainers.ManagedCertificatesClientBeginCreateOrUpdateOptions{
				ManagedCertificateEnvelope: &body,
			}
		}
		respr, errRespr := m.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, environmentNameParam, managedCertificateNameParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		m.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		m.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		m.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (m *ManagedCertificatesServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if m.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/managedEnvironments/(?P<environmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/managedCertificates/(?P<managedCertificateName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	environmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("environmentName")])
	if err != nil {
		return nil, err
	}
	managedCertificateNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedCertificateName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.Delete(req.Context(), resourceGroupNameParam, environmentNameParam, managedCertificateNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ManagedCertificatesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if m.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/managedEnvironments/(?P<environmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/managedCertificates/(?P<managedCertificateName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	environmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("environmentName")])
	if err != nil {
		return nil, err
	}
	managedCertificateNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedCertificateName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.Get(req.Context(), resourceGroupNameParam, environmentNameParam, managedCertificateNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ManagedCertificate, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ManagedCertificatesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := m.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/managedEnvironments/(?P<environmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/managedCertificates`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		environmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("environmentName")])
		if err != nil {
			return nil, err
		}
		resp := m.srv.NewListPager(resourceGroupNameParam, environmentNameParam, nil)
		newListPager = &resp
		m.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armappcontainers.ManagedCertificatesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		m.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		m.newListPager.remove(req)
	}
	return resp, nil
}

func (m *ManagedCertificatesServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if m.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/managedEnvironments/(?P<environmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/managedCertificates/(?P<managedCertificateName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armappcontainers.ManagedCertificatePatch](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	environmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("environmentName")])
	if err != nil {
		return nil, err
	}
	managedCertificateNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedCertificateName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.Update(req.Context(), resourceGroupNameParam, environmentNameParam, managedCertificateNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ManagedCertificate, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ManagedCertificatesServerTransport
var managedCertificatesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
