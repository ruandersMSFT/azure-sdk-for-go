// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// AlertRulesOperations contains the methods for the AlertRules group.
type AlertRulesOperations interface {
	// CreateOrUpdate - Creates or updates a classic metric alert rule.
	CreateOrUpdate(ctx context.Context, resourceGroupName string, ruleName string, parameters AlertRuleResource, options *AlertRulesCreateOrUpdateOptions) (*AlertRuleResourceResponse, error)
	// Delete - Deletes a classic metric alert rule
	Delete(ctx context.Context, resourceGroupName string, ruleName string, options *AlertRulesDeleteOptions) (*http.Response, error)
	// Get - Gets a classic metric alert rule
	Get(ctx context.Context, resourceGroupName string, ruleName string, options *AlertRulesGetOptions) (*AlertRuleResourceResponse, error)
	// ListByResourceGroup - List the classic metric alert rules within a resource group.
	ListByResourceGroup(ctx context.Context, resourceGroupName string, options *AlertRulesListByResourceGroupOptions) (*AlertRuleResourceCollectionResponse, error)
	// ListBySubscription - List the classic metric alert rules within a subscription.
	ListBySubscription(ctx context.Context, options *AlertRulesListBySubscriptionOptions) (*AlertRuleResourceCollectionResponse, error)
	// Update - Updates an existing classic metric AlertRuleResource. To update other fields use the CreateOrUpdate method.
	Update(ctx context.Context, resourceGroupName string, ruleName string, alertRulesResource AlertRuleResourcePatch, options *AlertRulesUpdateOptions) (*AlertRuleResourceResponse, error)
}

// AlertRulesClient implements the AlertRulesOperations interface.
// Don't use this type directly, use NewAlertRulesClient() instead.
type AlertRulesClient struct {
	*Client
	subscriptionID string
}

// NewAlertRulesClient creates a new instance of AlertRulesClient with the specified values.
func NewAlertRulesClient(c *Client, subscriptionID string) AlertRulesOperations {
	return &AlertRulesClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *AlertRulesClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// CreateOrUpdate - Creates or updates a classic metric alert rule.
func (client *AlertRulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, ruleName string, parameters AlertRuleResource, options *AlertRulesCreateOrUpdateOptions) (*AlertRuleResourceResponse, error) {
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
func (client *AlertRulesClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, ruleName string, parameters AlertRuleResource, options *AlertRulesCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.insights/alertrules/{ruleName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2016-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// CreateOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *AlertRulesClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*AlertRuleResourceResponse, error) {
	result := AlertRuleResourceResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.AlertRuleResource)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *AlertRulesClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Delete - Deletes a classic metric alert rule
func (client *AlertRulesClient) Delete(ctx context.Context, resourceGroupName string, ruleName string, options *AlertRulesDeleteOptions) (*http.Response, error) {
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
func (client *AlertRulesClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, ruleName string, options *AlertRulesDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.insights/alertrules/{ruleName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2016-03-01")
	req.URL.RawQuery = query.Encode()
	return req, nil
}

// DeleteHandleError handles the Delete error response.
func (client *AlertRulesClient) DeleteHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// Get - Gets a classic metric alert rule
func (client *AlertRulesClient) Get(ctx context.Context, resourceGroupName string, ruleName string, options *AlertRulesGetOptions) (*AlertRuleResourceResponse, error) {
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
func (client *AlertRulesClient) GetCreateRequest(ctx context.Context, resourceGroupName string, ruleName string, options *AlertRulesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.insights/alertrules/{ruleName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2016-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetHandleResponse handles the Get response.
func (client *AlertRulesClient) GetHandleResponse(resp *azcore.Response) (*AlertRuleResourceResponse, error) {
	result := AlertRuleResourceResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.AlertRuleResource)
}

// GetHandleError handles the Get error response.
func (client *AlertRulesClient) GetHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ListByResourceGroup - List the classic metric alert rules within a resource group.
func (client *AlertRulesClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, options *AlertRulesListByResourceGroupOptions) (*AlertRuleResourceCollectionResponse, error) {
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
func (client *AlertRulesClient) ListByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *AlertRulesListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.insights/alertrules"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2016-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *AlertRulesClient) ListByResourceGroupHandleResponse(resp *azcore.Response) (*AlertRuleResourceCollectionResponse, error) {
	result := AlertRuleResourceCollectionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.AlertRuleResourceCollection)
}

// ListByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *AlertRulesClient) ListByResourceGroupHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ListBySubscription - List the classic metric alert rules within a subscription.
func (client *AlertRulesClient) ListBySubscription(ctx context.Context, options *AlertRulesListBySubscriptionOptions) (*AlertRuleResourceCollectionResponse, error) {
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
func (client *AlertRulesClient) ListBySubscriptionCreateRequest(ctx context.Context, options *AlertRulesListBySubscriptionOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/microsoft.insights/alertrules"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2016-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *AlertRulesClient) ListBySubscriptionHandleResponse(resp *azcore.Response) (*AlertRuleResourceCollectionResponse, error) {
	result := AlertRuleResourceCollectionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.AlertRuleResourceCollection)
}

// ListBySubscriptionHandleError handles the ListBySubscription error response.
func (client *AlertRulesClient) ListBySubscriptionHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// Update - Updates an existing classic metric AlertRuleResource. To update other fields use the CreateOrUpdate method.
func (client *AlertRulesClient) Update(ctx context.Context, resourceGroupName string, ruleName string, alertRulesResource AlertRuleResourcePatch, options *AlertRulesUpdateOptions) (*AlertRuleResourceResponse, error) {
	req, err := client.UpdateCreateRequest(ctx, resourceGroupName, ruleName, alertRulesResource, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.UpdateHandleError(resp)
	}
	result, err := client.UpdateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateCreateRequest creates the Update request.
func (client *AlertRulesClient) UpdateCreateRequest(ctx context.Context, resourceGroupName string, ruleName string, alertRulesResource AlertRuleResourcePatch, options *AlertRulesUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.insights/alertrules/{ruleName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2016-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(alertRulesResource)
}

// UpdateHandleResponse handles the Update response.
func (client *AlertRulesClient) UpdateHandleResponse(resp *azcore.Response) (*AlertRuleResourceResponse, error) {
	result := AlertRuleResourceResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.AlertRuleResource)
}

// UpdateHandleError handles the Update error response.
func (client *AlertRulesClient) UpdateHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}