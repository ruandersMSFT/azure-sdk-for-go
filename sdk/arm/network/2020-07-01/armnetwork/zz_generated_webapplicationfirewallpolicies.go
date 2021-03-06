// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// WebApplicationFirewallPoliciesClient contains the methods for the WebApplicationFirewallPolicies group.
// Don't use this type directly, use NewWebApplicationFirewallPoliciesClient() instead.
type WebApplicationFirewallPoliciesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewWebApplicationFirewallPoliciesClient creates a new instance of WebApplicationFirewallPoliciesClient with the specified values.
func NewWebApplicationFirewallPoliciesClient(con *armcore.Connection, subscriptionID string) WebApplicationFirewallPoliciesClient {
	return WebApplicationFirewallPoliciesClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client WebApplicationFirewallPoliciesClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// CreateOrUpdate - Creates or update policy with specified rule set name within a resource group.
func (client WebApplicationFirewallPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, policyName string, parameters WebApplicationFirewallPolicy, options *WebApplicationFirewallPoliciesCreateOrUpdateOptions) (WebApplicationFirewallPolicyResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, policyName, parameters, options)
	if err != nil {
		return WebApplicationFirewallPolicyResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return WebApplicationFirewallPolicyResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return WebApplicationFirewallPolicyResponse{}, client.createOrUpdateHandleError(resp)
	}
	result, err := client.createOrUpdateHandleResponse(resp)
	if err != nil {
		return WebApplicationFirewallPolicyResponse{}, err
	}
	return result, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client WebApplicationFirewallPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, policyName string, parameters WebApplicationFirewallPolicy, options *WebApplicationFirewallPoliciesCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies/{policyName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{policyName}", url.PathEscape(policyName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client WebApplicationFirewallPoliciesClient) createOrUpdateHandleResponse(resp *azcore.Response) (WebApplicationFirewallPolicyResponse, error) {
	result := WebApplicationFirewallPolicyResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.WebApplicationFirewallPolicy)
	return result, err
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client WebApplicationFirewallPoliciesClient) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDelete - Deletes Policy.
func (client WebApplicationFirewallPoliciesClient) BeginDelete(ctx context.Context, resourceGroupName string, policyName string, options *WebApplicationFirewallPoliciesBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.delete(ctx, resourceGroupName, policyName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("WebApplicationFirewallPoliciesClient.Delete", "location", resp, client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client WebApplicationFirewallPoliciesClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("WebApplicationFirewallPoliciesClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes Policy.
func (client WebApplicationFirewallPoliciesClient) delete(ctx context.Context, resourceGroupName string, policyName string, options *WebApplicationFirewallPoliciesBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, policyName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client WebApplicationFirewallPoliciesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, policyName string, options *WebApplicationFirewallPoliciesBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies/{policyName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{policyName}", url.PathEscape(policyName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client WebApplicationFirewallPoliciesClient) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Retrieve protection policy with specified name within a resource group.
func (client WebApplicationFirewallPoliciesClient) Get(ctx context.Context, resourceGroupName string, policyName string, options *WebApplicationFirewallPoliciesGetOptions) (WebApplicationFirewallPolicyResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, policyName, options)
	if err != nil {
		return WebApplicationFirewallPolicyResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return WebApplicationFirewallPolicyResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return WebApplicationFirewallPolicyResponse{}, client.getHandleError(resp)
	}
	result, err := client.getHandleResponse(resp)
	if err != nil {
		return WebApplicationFirewallPolicyResponse{}, err
	}
	return result, nil
}

// getCreateRequest creates the Get request.
func (client WebApplicationFirewallPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, policyName string, options *WebApplicationFirewallPoliciesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies/{policyName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{policyName}", url.PathEscape(policyName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client WebApplicationFirewallPoliciesClient) getHandleResponse(resp *azcore.Response) (WebApplicationFirewallPolicyResponse, error) {
	result := WebApplicationFirewallPolicyResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.WebApplicationFirewallPolicy)
	return result, err
}

// getHandleError handles the Get error response.
func (client WebApplicationFirewallPoliciesClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Lists all of the protection policies within a resource group.
func (client WebApplicationFirewallPoliciesClient) List(resourceGroupName string, options *WebApplicationFirewallPoliciesListOptions) WebApplicationFirewallPolicyListResultPager {
	return &webApplicationFirewallPolicyListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp WebApplicationFirewallPolicyListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.WebApplicationFirewallPolicyListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client WebApplicationFirewallPoliciesClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *WebApplicationFirewallPoliciesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client WebApplicationFirewallPoliciesClient) listHandleResponse(resp *azcore.Response) (WebApplicationFirewallPolicyListResultResponse, error) {
	result := WebApplicationFirewallPolicyListResultResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.WebApplicationFirewallPolicyListResult)
	return result, err
}

// listHandleError handles the List error response.
func (client WebApplicationFirewallPoliciesClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListAll - Gets all the WAF policies in a subscription.
func (client WebApplicationFirewallPoliciesClient) ListAll(options *WebApplicationFirewallPoliciesListAllOptions) WebApplicationFirewallPolicyListResultPager {
	return &webApplicationFirewallPolicyListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listAllCreateRequest(ctx, options)
		},
		responder: client.listAllHandleResponse,
		errorer:   client.listAllHandleError,
		advancer: func(ctx context.Context, resp WebApplicationFirewallPolicyListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.WebApplicationFirewallPolicyListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listAllCreateRequest creates the ListAll request.
func (client WebApplicationFirewallPoliciesClient) listAllCreateRequest(ctx context.Context, options *WebApplicationFirewallPoliciesListAllOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client WebApplicationFirewallPoliciesClient) listAllHandleResponse(resp *azcore.Response) (WebApplicationFirewallPolicyListResultResponse, error) {
	result := WebApplicationFirewallPolicyListResultResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.WebApplicationFirewallPolicyListResult)
	return result, err
}

// listAllHandleError handles the ListAll error response.
func (client WebApplicationFirewallPoliciesClient) listAllHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
