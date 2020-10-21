// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"strings"
)

// VMInsightsOperations contains the methods for the VMInsights group.
type VMInsightsOperations interface {
	// GetOnboardingStatus - Retrieves the VM Insights onboarding status for the specified resource or resource scope.
	GetOnboardingStatus(ctx context.Context, resourceUri string, options *VMInsightsGetOnboardingStatusOptions) (*VMInsightsOnboardingStatusResponse, error)
}

// VMInsightsClient implements the VMInsightsOperations interface.
// Don't use this type directly, use NewVMInsightsClient() instead.
type VMInsightsClient struct {
	*Client
}

// NewVMInsightsClient creates a new instance of VMInsightsClient with the specified values.
func NewVMInsightsClient(c *Client) VMInsightsOperations {
	return &VMInsightsClient{Client: c}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *VMInsightsClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// GetOnboardingStatus - Retrieves the VM Insights onboarding status for the specified resource or resource scope.
func (client *VMInsightsClient) GetOnboardingStatus(ctx context.Context, resourceUri string, options *VMInsightsGetOnboardingStatusOptions) (*VMInsightsOnboardingStatusResponse, error) {
	req, err := client.GetOnboardingStatusCreateRequest(ctx, resourceUri, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetOnboardingStatusHandleError(resp)
	}
	result, err := client.GetOnboardingStatusHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetOnboardingStatusCreateRequest creates the GetOnboardingStatus request.
func (client *VMInsightsClient) GetOnboardingStatusCreateRequest(ctx context.Context, resourceUri string, options *VMInsightsGetOnboardingStatusOptions) (*azcore.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.Insights/vmInsightsOnboardingStatuses/default"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceUri)
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2018-11-27-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetOnboardingStatusHandleResponse handles the GetOnboardingStatus response.
func (client *VMInsightsClient) GetOnboardingStatusHandleResponse(resp *azcore.Response) (*VMInsightsOnboardingStatusResponse, error) {
	result := VMInsightsOnboardingStatusResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VMInsightsOnboardingStatus)
}

// GetOnboardingStatusHandleError handles the GetOnboardingStatus error response.
func (client *VMInsightsClient) GetOnboardingStatusHandleError(resp *azcore.Response) error {
	var err ResponseWithError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}