// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// ScheduledQueryRulesOperations contains the methods for the ScheduledQueryRules group.
type ScheduledQueryRulesOperations interface {
	// CreateOrUpdate - Creates or updates an log search rule.
	CreateOrUpdate(ctx context.Context, resourceGroupName string, ruleName string, parameters LogSearchRuleResource, options *ScheduledQueryRulesCreateOrUpdateOptions) (*LogSearchRuleResourceResponse, error)
	// Delete - Deletes a Log Search rule
	Delete(ctx context.Context, resourceGroupName string, ruleName string, options *ScheduledQueryRulesDeleteOptions) (*http.Response, error)
	// Get - Gets an Log Search rule
	Get(ctx context.Context, resourceGroupName string, ruleName string, options *ScheduledQueryRulesGetOptions) (*LogSearchRuleResourceResponse, error)
	// ListByResourceGroup - List the Log Search rules within a resource group.
	ListByResourceGroup(ctx context.Context, resourceGroupName string, options *ScheduledQueryRulesListByResourceGroupOptions) (*LogSearchRuleResourceCollectionResponse, error)
	// ListBySubscription - List the Log Search rules within a subscription group.
	ListBySubscription(ctx context.Context, options *ScheduledQueryRulesListBySubscriptionOptions) (*LogSearchRuleResourceCollectionResponse, error)
	// Update - Update log search Rule.
	Update(ctx context.Context, resourceGroupName string, ruleName string, parameters LogSearchRuleResourcePatch, options *ScheduledQueryRulesUpdateOptions) (*LogSearchRuleResourceResponse, error)
}

// ScheduledQueryRulesClient implements the ScheduledQueryRulesOperations interface.
// Don't use this type directly, use NewScheduledQueryRulesClient() instead.
type ScheduledQueryRulesClient struct {
	*Client
	subscriptionID string
}

// NewScheduledQueryRulesClient creates a new instance of ScheduledQueryRulesClient with the specified values.
func NewScheduledQueryRulesClient(c *Client, subscriptionID string) ScheduledQueryRulesOperations {
	return &ScheduledQueryRulesClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *ScheduledQueryRulesClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// CreateOrUpdate - Creates or updates an log search rule.
func (client *ScheduledQueryRulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, ruleName string, parameters LogSearchRuleResource, options *ScheduledQueryRulesCreateOrUpdateOptions) (*LogSearchRuleResourceResponse, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, ruleName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.CreateOrUpdateHandleError(resp)
	}
	result, err := client.CreateOrUpdateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ScheduledQueryRulesClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, ruleName string, parameters LogSearchRuleResource, options *ScheduledQueryRulesCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.insights/scheduledQueryRules/{ruleName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2018-04-16")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// CreateOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ScheduledQueryRulesClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*LogSearchRuleResourceResponse, error) {
	result := LogSearchRuleResourceResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.LogSearchRuleResource)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ScheduledQueryRulesClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Delete - Deletes a Log Search rule
func (client *ScheduledQueryRulesClient) Delete(ctx context.Context, resourceGroupName string, ruleName string, options *ScheduledQueryRulesDeleteOptions) (*http.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, ruleName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNoContent) {
		return nil, client.DeleteHandleError(resp)
	}
	return resp.Response, nil
}

// DeleteCreateRequest creates the Delete request.
func (client *ScheduledQueryRulesClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, ruleName string, options *ScheduledQueryRulesDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.insights/scheduledQueryRules/{ruleName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2018-04-16")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// DeleteHandleError handles the Delete error response.
func (client *ScheduledQueryRulesClient) DeleteHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets an Log Search rule
func (client *ScheduledQueryRulesClient) Get(ctx context.Context, resourceGroupName string, ruleName string, options *ScheduledQueryRulesGetOptions) (*LogSearchRuleResourceResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, ruleName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetHandleError(resp)
	}
	result, err := client.GetHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCreateRequest creates the Get request.
func (client *ScheduledQueryRulesClient) GetCreateRequest(ctx context.Context, resourceGroupName string, ruleName string, options *ScheduledQueryRulesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.insights/scheduledQueryRules/{ruleName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2018-04-16")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetHandleResponse handles the Get response.
func (client *ScheduledQueryRulesClient) GetHandleResponse(resp *azcore.Response) (*LogSearchRuleResourceResponse, error) {
	result := LogSearchRuleResourceResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.LogSearchRuleResource)
}

// GetHandleError handles the Get error response.
func (client *ScheduledQueryRulesClient) GetHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListByResourceGroup - List the Log Search rules within a resource group.
func (client *ScheduledQueryRulesClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, options *ScheduledQueryRulesListByResourceGroupOptions) (*LogSearchRuleResourceCollectionResponse, error) {
	req, err := client.ListByResourceGroupCreateRequest(ctx, resourceGroupName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ListByResourceGroupHandleError(resp)
	}
	result, err := client.ListByResourceGroupHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ScheduledQueryRulesClient) ListByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ScheduledQueryRulesListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.insights/scheduledQueryRules"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2018-04-16")
	if options != nil && options.Filter != nil {
		query.Set("$filter", *options.Filter)
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ScheduledQueryRulesClient) ListByResourceGroupHandleResponse(resp *azcore.Response) (*LogSearchRuleResourceCollectionResponse, error) {
	result := LogSearchRuleResourceCollectionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.LogSearchRuleResourceCollection)
}

// ListByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *ScheduledQueryRulesClient) ListByResourceGroupHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListBySubscription - List the Log Search rules within a subscription group.
func (client *ScheduledQueryRulesClient) ListBySubscription(ctx context.Context, options *ScheduledQueryRulesListBySubscriptionOptions) (*LogSearchRuleResourceCollectionResponse, error) {
	req, err := client.ListBySubscriptionCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ListBySubscriptionHandleError(resp)
	}
	result, err := client.ListBySubscriptionHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ScheduledQueryRulesClient) ListBySubscriptionCreateRequest(ctx context.Context, options *ScheduledQueryRulesListBySubscriptionOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/microsoft.insights/scheduledQueryRules"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2018-04-16")
	if options != nil && options.Filter != nil {
		query.Set("$filter", *options.Filter)
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ScheduledQueryRulesClient) ListBySubscriptionHandleResponse(resp *azcore.Response) (*LogSearchRuleResourceCollectionResponse, error) {
	result := LogSearchRuleResourceCollectionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.LogSearchRuleResourceCollection)
}

// ListBySubscriptionHandleError handles the ListBySubscription error response.
func (client *ScheduledQueryRulesClient) ListBySubscriptionHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Update - Update log search Rule.
func (client *ScheduledQueryRulesClient) Update(ctx context.Context, resourceGroupName string, ruleName string, parameters LogSearchRuleResourcePatch, options *ScheduledQueryRulesUpdateOptions) (*LogSearchRuleResourceResponse, error) {
	req, err := client.UpdateCreateRequest(ctx, resourceGroupName, ruleName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.UpdateHandleError(resp)
	}
	result, err := client.UpdateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateCreateRequest creates the Update request.
func (client *ScheduledQueryRulesClient) UpdateCreateRequest(ctx context.Context, resourceGroupName string, ruleName string, parameters LogSearchRuleResourcePatch, options *ScheduledQueryRulesUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.insights/scheduledQueryRules/{ruleName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2018-04-16")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// UpdateHandleResponse handles the Update response.
func (client *ScheduledQueryRulesClient) UpdateHandleResponse(resp *azcore.Response) (*LogSearchRuleResourceResponse, error) {
	result := LogSearchRuleResourceResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.LogSearchRuleResource)
}

// UpdateHandleError handles the Update error response.
func (client *ScheduledQueryRulesClient) UpdateHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}