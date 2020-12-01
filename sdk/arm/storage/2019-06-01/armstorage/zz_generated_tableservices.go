// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorage

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// TableServicesClient contains the methods for the TableServices group.
// Don't use this type directly, use NewTableServicesClient() instead.
type TableServicesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewTableServicesClient creates a new instance of TableServicesClient with the specified values.
func NewTableServicesClient(con *armcore.Connection, subscriptionID string) TableServicesClient {
	return TableServicesClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client TableServicesClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// GetServiceProperties - Gets the properties of a storage account’s Table service, including properties for Storage Analytics and CORS (Cross-Origin Resource
// Sharing) rules.
func (client TableServicesClient) GetServiceProperties(ctx context.Context, resourceGroupName string, accountName string, options *TableServicesGetServicePropertiesOptions) (TableServicePropertiesResponse, error) {
	req, err := client.getServicePropertiesCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return TableServicePropertiesResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return TableServicePropertiesResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return TableServicePropertiesResponse{}, client.getServicePropertiesHandleError(resp)
	}
	result, err := client.getServicePropertiesHandleResponse(resp)
	if err != nil {
		return TableServicePropertiesResponse{}, err
	}
	return result, nil
}

// getServicePropertiesCreateRequest creates the GetServiceProperties request.
func (client TableServicesClient) getServicePropertiesCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *TableServicesGetServicePropertiesOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/tableServices/{tableServiceName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{tableServiceName}", url.PathEscape("default"))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getServicePropertiesHandleResponse handles the GetServiceProperties response.
func (client TableServicesClient) getServicePropertiesHandleResponse(resp *azcore.Response) (TableServicePropertiesResponse, error) {
	result := TableServicePropertiesResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.TableServiceProperties)
	return result, err
}

// getServicePropertiesHandleError handles the GetServiceProperties error response.
func (client TableServicesClient) getServicePropertiesHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - List all table services for the storage account.
func (client TableServicesClient) List(ctx context.Context, resourceGroupName string, accountName string, options *TableServicesListOptions) (ListTableServicesResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return ListTableServicesResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return ListTableServicesResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ListTableServicesResponse{}, client.listHandleError(resp)
	}
	result, err := client.listHandleResponse(resp)
	if err != nil {
		return ListTableServicesResponse{}, err
	}
	return result, nil
}

// listCreateRequest creates the List request.
func (client TableServicesClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *TableServicesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/tableServices"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client TableServicesClient) listHandleResponse(resp *azcore.Response) (ListTableServicesResponse, error) {
	result := ListTableServicesResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.ListTableServices)
	return result, err
}

// listHandleError handles the List error response.
func (client TableServicesClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// SetServiceProperties - Sets the properties of a storage account’s Table service, including properties for Storage Analytics and CORS (Cross-Origin Resource
// Sharing) rules.
func (client TableServicesClient) SetServiceProperties(ctx context.Context, resourceGroupName string, accountName string, parameters TableServiceProperties, options *TableServicesSetServicePropertiesOptions) (TableServicePropertiesResponse, error) {
	req, err := client.setServicePropertiesCreateRequest(ctx, resourceGroupName, accountName, parameters, options)
	if err != nil {
		return TableServicePropertiesResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return TableServicePropertiesResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return TableServicePropertiesResponse{}, client.setServicePropertiesHandleError(resp)
	}
	result, err := client.setServicePropertiesHandleResponse(resp)
	if err != nil {
		return TableServicePropertiesResponse{}, err
	}
	return result, nil
}

// setServicePropertiesCreateRequest creates the SetServiceProperties request.
func (client TableServicesClient) setServicePropertiesCreateRequest(ctx context.Context, resourceGroupName string, accountName string, parameters TableServiceProperties, options *TableServicesSetServicePropertiesOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/tableServices/{tableServiceName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{tableServiceName}", url.PathEscape("default"))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// setServicePropertiesHandleResponse handles the SetServiceProperties response.
func (client TableServicesClient) setServicePropertiesHandleResponse(resp *azcore.Response) (TableServicePropertiesResponse, error) {
	result := TableServicePropertiesResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.TableServiceProperties)
	return result, err
}

// setServicePropertiesHandleError handles the SetServiceProperties error response.
func (client TableServicesClient) setServicePropertiesHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}