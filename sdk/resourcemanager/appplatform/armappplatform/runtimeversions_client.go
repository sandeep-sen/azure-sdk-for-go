//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappplatform

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// RuntimeVersionsClient contains the methods for the RuntimeVersions group.
// Don't use this type directly, use NewRuntimeVersionsClient() instead.
type RuntimeVersionsClient struct {
	internal *arm.Client
}

// NewRuntimeVersionsClient creates a new instance of RuntimeVersionsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRuntimeVersionsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*RuntimeVersionsClient, error) {
	cl, err := arm.NewClient(moduleName+".RuntimeVersionsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RuntimeVersionsClient{
		internal: cl,
	}
	return client, nil
}

// ListRuntimeVersions - Lists all of the available runtime versions supported by Microsoft.AppPlatform provider.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - options - RuntimeVersionsClientListRuntimeVersionsOptions contains the optional parameters for the RuntimeVersionsClient.ListRuntimeVersions
//     method.
func (client *RuntimeVersionsClient) ListRuntimeVersions(ctx context.Context, options *RuntimeVersionsClientListRuntimeVersionsOptions) (RuntimeVersionsClientListRuntimeVersionsResponse, error) {
	req, err := client.listRuntimeVersionsCreateRequest(ctx, options)
	if err != nil {
		return RuntimeVersionsClientListRuntimeVersionsResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RuntimeVersionsClientListRuntimeVersionsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RuntimeVersionsClientListRuntimeVersionsResponse{}, runtime.NewResponseError(resp)
	}
	return client.listRuntimeVersionsHandleResponse(resp)
}

// listRuntimeVersionsCreateRequest creates the ListRuntimeVersions request.
func (client *RuntimeVersionsClient) listRuntimeVersionsCreateRequest(ctx context.Context, options *RuntimeVersionsClientListRuntimeVersionsOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.AppPlatform/runtimeVersions"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listRuntimeVersionsHandleResponse handles the ListRuntimeVersions response.
func (client *RuntimeVersionsClient) listRuntimeVersionsHandleResponse(resp *http.Response) (RuntimeVersionsClientListRuntimeVersionsResponse, error) {
	result := RuntimeVersionsClientListRuntimeVersionsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailableRuntimeVersions); err != nil {
		return RuntimeVersionsClientListRuntimeVersionsResponse{}, err
	}
	return result, nil
}
