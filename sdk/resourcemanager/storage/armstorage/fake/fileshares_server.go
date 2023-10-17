//go:build go1.18
// +build go1.18

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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
)

// FileSharesServer is a fake server for instances of the armstorage.FileSharesClient type.
type FileSharesServer struct {
	// Create is the fake for method FileSharesClient.Create
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	Create func(ctx context.Context, resourceGroupName string, accountName string, shareName string, fileShare armstorage.FileShare, options *armstorage.FileSharesClientCreateOptions) (resp azfake.Responder[armstorage.FileSharesClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method FileSharesClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, accountName string, shareName string, options *armstorage.FileSharesClientDeleteOptions) (resp azfake.Responder[armstorage.FileSharesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method FileSharesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, accountName string, shareName string, options *armstorage.FileSharesClientGetOptions) (resp azfake.Responder[armstorage.FileSharesClientGetResponse], errResp azfake.ErrorResponder)

	// Lease is the fake for method FileSharesClient.Lease
	// HTTP status codes to indicate success: http.StatusOK
	Lease func(ctx context.Context, resourceGroupName string, accountName string, shareName string, options *armstorage.FileSharesClientLeaseOptions) (resp azfake.Responder[armstorage.FileSharesClientLeaseResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method FileSharesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, accountName string, options *armstorage.FileSharesClientListOptions) (resp azfake.PagerResponder[armstorage.FileSharesClientListResponse])

	// Restore is the fake for method FileSharesClient.Restore
	// HTTP status codes to indicate success: http.StatusOK
	Restore func(ctx context.Context, resourceGroupName string, accountName string, shareName string, deletedShare armstorage.DeletedShare, options *armstorage.FileSharesClientRestoreOptions) (resp azfake.Responder[armstorage.FileSharesClientRestoreResponse], errResp azfake.ErrorResponder)

	// Update is the fake for method FileSharesClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, accountName string, shareName string, fileShare armstorage.FileShare, options *armstorage.FileSharesClientUpdateOptions) (resp azfake.Responder[armstorage.FileSharesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewFileSharesServerTransport creates a new instance of FileSharesServerTransport with the provided implementation.
// The returned FileSharesServerTransport instance is connected to an instance of armstorage.FileSharesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewFileSharesServerTransport(srv *FileSharesServer) *FileSharesServerTransport {
	return &FileSharesServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armstorage.FileSharesClientListResponse]](),
	}
}

// FileSharesServerTransport connects instances of armstorage.FileSharesClient to instances of FileSharesServer.
// Don't use this type directly, use NewFileSharesServerTransport instead.
type FileSharesServerTransport struct {
	srv          *FileSharesServer
	newListPager *tracker[azfake.PagerResponder[armstorage.FileSharesClientListResponse]]
}

// Do implements the policy.Transporter interface for FileSharesServerTransport.
func (f *FileSharesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "FileSharesClient.Create":
		resp, err = f.dispatchCreate(req)
	case "FileSharesClient.Delete":
		resp, err = f.dispatchDelete(req)
	case "FileSharesClient.Get":
		resp, err = f.dispatchGet(req)
	case "FileSharesClient.Lease":
		resp, err = f.dispatchLease(req)
	case "FileSharesClient.NewListPager":
		resp, err = f.dispatchNewListPager(req)
	case "FileSharesClient.Restore":
		resp, err = f.dispatchRestore(req)
	case "FileSharesClient.Update":
		resp, err = f.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (f *FileSharesServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if f.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Storage/storageAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/fileServices/default/shares/(?P<shareName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	body, err := server.UnmarshalRequestAsJSON[armstorage.FileShare](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	shareNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("shareName")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(expandUnescaped)
	var options *armstorage.FileSharesClientCreateOptions
	if expandParam != nil {
		options = &armstorage.FileSharesClientCreateOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := f.srv.Create(req.Context(), resourceGroupNameUnescaped, accountNameUnescaped, shareNameUnescaped, body, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).FileShare, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (f *FileSharesServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if f.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Storage/storageAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/fileServices/default/shares/(?P<shareName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	shareNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("shareName")])
	if err != nil {
		return nil, err
	}
	xMSSnapshotParam := getOptional(getHeaderValue(req.Header, "x-ms-snapshot"))
	includeUnescaped, err := url.QueryUnescape(qp.Get("$include"))
	if err != nil {
		return nil, err
	}
	includeParam := getOptional(includeUnescaped)
	var options *armstorage.FileSharesClientDeleteOptions
	if xMSSnapshotParam != nil || includeParam != nil {
		options = &armstorage.FileSharesClientDeleteOptions{
			XMSSnapshot: xMSSnapshotParam,
			Include:     includeParam,
		}
	}
	respr, errRespr := f.srv.Delete(req.Context(), resourceGroupNameUnescaped, accountNameUnescaped, shareNameUnescaped, options)
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

func (f *FileSharesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if f.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Storage/storageAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/fileServices/default/shares/(?P<shareName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	shareNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("shareName")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(expandUnescaped)
	xMSSnapshotParam := getOptional(getHeaderValue(req.Header, "x-ms-snapshot"))
	var options *armstorage.FileSharesClientGetOptions
	if expandParam != nil || xMSSnapshotParam != nil {
		options = &armstorage.FileSharesClientGetOptions{
			Expand:      expandParam,
			XMSSnapshot: xMSSnapshotParam,
		}
	}
	respr, errRespr := f.srv.Get(req.Context(), resourceGroupNameUnescaped, accountNameUnescaped, shareNameUnescaped, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).FileShare, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (f *FileSharesServerTransport) dispatchLease(req *http.Request) (*http.Response, error) {
	if f.srv.Lease == nil {
		return nil, &nonRetriableError{errors.New("fake for method Lease not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Storage/storageAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/fileServices/default/shares/(?P<shareName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/lease`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armstorage.LeaseShareRequest](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	shareNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("shareName")])
	if err != nil {
		return nil, err
	}
	xMSSnapshotParam := getOptional(getHeaderValue(req.Header, "x-ms-snapshot"))
	var options *armstorage.FileSharesClientLeaseOptions
	if xMSSnapshotParam != nil || !reflect.ValueOf(body).IsZero() {
		options = &armstorage.FileSharesClientLeaseOptions{
			XMSSnapshot: xMSSnapshotParam,
			Parameters:  &body,
		}
	}
	respr, errRespr := f.srv.Lease(req.Context(), resourceGroupNameUnescaped, accountNameUnescaped, shareNameUnescaped, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).LeaseShareResponse, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}

func (f *FileSharesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if f.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := f.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Storage/storageAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/fileServices/default/shares`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		accountNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		maxpagesizeUnescaped, err := url.QueryUnescape(qp.Get("$maxpagesize"))
		if err != nil {
			return nil, err
		}
		maxpagesizeParam := getOptional(maxpagesizeUnescaped)
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
		if err != nil {
			return nil, err
		}
		expandParam := getOptional(expandUnescaped)
		var options *armstorage.FileSharesClientListOptions
		if maxpagesizeParam != nil || filterParam != nil || expandParam != nil {
			options = &armstorage.FileSharesClientListOptions{
				Maxpagesize: maxpagesizeParam,
				Filter:      filterParam,
				Expand:      expandParam,
			}
		}
		resp := f.srv.NewListPager(resourceGroupNameUnescaped, accountNameUnescaped, options)
		newListPager = &resp
		f.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armstorage.FileSharesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		f.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		f.newListPager.remove(req)
	}
	return resp, nil
}

func (f *FileSharesServerTransport) dispatchRestore(req *http.Request) (*http.Response, error) {
	if f.srv.Restore == nil {
		return nil, &nonRetriableError{errors.New("fake for method Restore not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Storage/storageAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/fileServices/default/shares/(?P<shareName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/restore`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armstorage.DeletedShare](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	shareNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("shareName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.Restore(req.Context(), resourceGroupNameUnescaped, accountNameUnescaped, shareNameUnescaped, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (f *FileSharesServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if f.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Storage/storageAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/fileServices/default/shares/(?P<shareName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armstorage.FileShare](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	shareNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("shareName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.Update(req.Context(), resourceGroupNameUnescaped, accountNameUnescaped, shareNameUnescaped, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).FileShare, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}