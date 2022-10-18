//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecuritydevops

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// AzureDevOpsConnectorClient contains the methods for the AzureDevOpsConnector group.
// Don't use this type directly, use NewAzureDevOpsConnectorClient() instead.
type AzureDevOpsConnectorClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAzureDevOpsConnectorClient creates a new instance of AzureDevOpsConnectorClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAzureDevOpsConnectorClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AzureDevOpsConnectorClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &AzureDevOpsConnectorClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates an Azure DevOps Connector.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// azureDevOpsConnectorName - Name of the AzureDevOps Connector.
// azureDevOpsConnector - Connector resource payload.
// options - AzureDevOpsConnectorClientBeginCreateOrUpdateOptions contains the optional parameters for the AzureDevOpsConnectorClient.BeginCreateOrUpdate
// method.
func (client *AzureDevOpsConnectorClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, azureDevOpsConnectorName string, azureDevOpsConnector AzureDevOpsConnector, options *AzureDevOpsConnectorClientBeginCreateOrUpdateOptions) (*runtime.Poller[AzureDevOpsConnectorClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, azureDevOpsConnectorName, azureDevOpsConnector, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[AzureDevOpsConnectorClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[AzureDevOpsConnectorClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates an Azure DevOps Connector.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
func (client *AzureDevOpsConnectorClient) createOrUpdate(ctx context.Context, resourceGroupName string, azureDevOpsConnectorName string, azureDevOpsConnector AzureDevOpsConnector, options *AzureDevOpsConnectorClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, azureDevOpsConnectorName, azureDevOpsConnector, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AzureDevOpsConnectorClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, azureDevOpsConnectorName string, azureDevOpsConnector AzureDevOpsConnector, options *AzureDevOpsConnectorClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityDevOps/azureDevOpsConnectors/{azureDevOpsConnectorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if azureDevOpsConnectorName == "" {
		return nil, errors.New("parameter azureDevOpsConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureDevOpsConnectorName}", url.PathEscape(azureDevOpsConnectorName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, azureDevOpsConnector)
}

// BeginDelete - Delete monitored AzureDevOps Connector details.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// azureDevOpsConnectorName - Name of the AzureDevOps Connector.
// options - AzureDevOpsConnectorClientBeginDeleteOptions contains the optional parameters for the AzureDevOpsConnectorClient.BeginDelete
// method.
func (client *AzureDevOpsConnectorClient) BeginDelete(ctx context.Context, resourceGroupName string, azureDevOpsConnectorName string, options *AzureDevOpsConnectorClientBeginDeleteOptions) (*runtime.Poller[AzureDevOpsConnectorClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, azureDevOpsConnectorName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[AzureDevOpsConnectorClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[AzureDevOpsConnectorClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete monitored AzureDevOps Connector details.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
func (client *AzureDevOpsConnectorClient) deleteOperation(ctx context.Context, resourceGroupName string, azureDevOpsConnectorName string, options *AzureDevOpsConnectorClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, azureDevOpsConnectorName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AzureDevOpsConnectorClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, azureDevOpsConnectorName string, options *AzureDevOpsConnectorClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityDevOps/azureDevOpsConnectors/{azureDevOpsConnectorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if azureDevOpsConnectorName == "" {
		return nil, errors.New("parameter azureDevOpsConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureDevOpsConnectorName}", url.PathEscape(azureDevOpsConnectorName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Returns a monitored AzureDevOps Connector resource for a given ID.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// azureDevOpsConnectorName - Name of the AzureDevOps Connector.
// options - AzureDevOpsConnectorClientGetOptions contains the optional parameters for the AzureDevOpsConnectorClient.Get
// method.
func (client *AzureDevOpsConnectorClient) Get(ctx context.Context, resourceGroupName string, azureDevOpsConnectorName string, options *AzureDevOpsConnectorClientGetOptions) (AzureDevOpsConnectorClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, azureDevOpsConnectorName, options)
	if err != nil {
		return AzureDevOpsConnectorClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AzureDevOpsConnectorClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AzureDevOpsConnectorClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AzureDevOpsConnectorClient) getCreateRequest(ctx context.Context, resourceGroupName string, azureDevOpsConnectorName string, options *AzureDevOpsConnectorClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityDevOps/azureDevOpsConnectors/{azureDevOpsConnectorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if azureDevOpsConnectorName == "" {
		return nil, errors.New("parameter azureDevOpsConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureDevOpsConnectorName}", url.PathEscape(azureDevOpsConnectorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AzureDevOpsConnectorClient) getHandleResponse(resp *http.Response) (AzureDevOpsConnectorClientGetResponse, error) {
	result := AzureDevOpsConnectorClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureDevOpsConnector); err != nil {
		return AzureDevOpsConnectorClientGetResponse{}, err
	}
	return result, nil
}

// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - AzureDevOpsConnectorClientListByResourceGroupOptions contains the optional parameters for the AzureDevOpsConnectorClient.ListByResourceGroup
// method.
func (client *AzureDevOpsConnectorClient) NewListByResourceGroupPager(resourceGroupName string, options *AzureDevOpsConnectorClientListByResourceGroupOptions) *runtime.Pager[AzureDevOpsConnectorClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[AzureDevOpsConnectorClientListByResourceGroupResponse]{
		More: func(page AzureDevOpsConnectorClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AzureDevOpsConnectorClientListByResourceGroupResponse) (AzureDevOpsConnectorClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AzureDevOpsConnectorClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AzureDevOpsConnectorClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AzureDevOpsConnectorClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *AzureDevOpsConnectorClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *AzureDevOpsConnectorClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityDevOps/azureDevOpsConnectors"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *AzureDevOpsConnectorClient) listByResourceGroupHandleResponse(resp *http.Response) (AzureDevOpsConnectorClientListByResourceGroupResponse, error) {
	result := AzureDevOpsConnectorClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureDevOpsConnectorListResponse); err != nil {
		return AzureDevOpsConnectorClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Returns a list of monitored AzureDevOps Connectors.
// Generated from API version 2022-09-01-preview
// options - AzureDevOpsConnectorClientListBySubscriptionOptions contains the optional parameters for the AzureDevOpsConnectorClient.ListBySubscription
// method.
func (client *AzureDevOpsConnectorClient) NewListBySubscriptionPager(options *AzureDevOpsConnectorClientListBySubscriptionOptions) *runtime.Pager[AzureDevOpsConnectorClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[AzureDevOpsConnectorClientListBySubscriptionResponse]{
		More: func(page AzureDevOpsConnectorClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AzureDevOpsConnectorClientListBySubscriptionResponse) (AzureDevOpsConnectorClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AzureDevOpsConnectorClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AzureDevOpsConnectorClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AzureDevOpsConnectorClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *AzureDevOpsConnectorClient) listBySubscriptionCreateRequest(ctx context.Context, options *AzureDevOpsConnectorClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.SecurityDevOps/azureDevOpsConnectors"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *AzureDevOpsConnectorClient) listBySubscriptionHandleResponse(resp *http.Response) (AzureDevOpsConnectorClientListBySubscriptionResponse, error) {
	result := AzureDevOpsConnectorClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureDevOpsConnectorListResponse); err != nil {
		return AzureDevOpsConnectorClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update monitored AzureDevOps Connector details.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// azureDevOpsConnectorName - Name of the AzureDevOps Connector.
// azureDevOpsConnector - Connector resource payload.
// options - AzureDevOpsConnectorClientBeginUpdateOptions contains the optional parameters for the AzureDevOpsConnectorClient.BeginUpdate
// method.
func (client *AzureDevOpsConnectorClient) BeginUpdate(ctx context.Context, resourceGroupName string, azureDevOpsConnectorName string, azureDevOpsConnector AzureDevOpsConnector, options *AzureDevOpsConnectorClientBeginUpdateOptions) (*runtime.Poller[AzureDevOpsConnectorClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, azureDevOpsConnectorName, azureDevOpsConnector, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[AzureDevOpsConnectorClientUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[AzureDevOpsConnectorClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Update monitored AzureDevOps Connector details.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
func (client *AzureDevOpsConnectorClient) update(ctx context.Context, resourceGroupName string, azureDevOpsConnectorName string, azureDevOpsConnector AzureDevOpsConnector, options *AzureDevOpsConnectorClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, azureDevOpsConnectorName, azureDevOpsConnector, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *AzureDevOpsConnectorClient) updateCreateRequest(ctx context.Context, resourceGroupName string, azureDevOpsConnectorName string, azureDevOpsConnector AzureDevOpsConnector, options *AzureDevOpsConnectorClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityDevOps/azureDevOpsConnectors/{azureDevOpsConnectorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if azureDevOpsConnectorName == "" {
		return nil, errors.New("parameter azureDevOpsConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureDevOpsConnectorName}", url.PathEscape(azureDevOpsConnectorName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, azureDevOpsConnector)
}