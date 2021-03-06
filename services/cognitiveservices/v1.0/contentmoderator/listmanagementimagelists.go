package contentmoderator

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ListManagementImageListsClient is the you use the API to scan your content as it is generated. Content Moderator
// then processes your content and sends the results along with relevant information either back to your systems or to
// the built-in review tool. You can use this information to take decisions e.g. take it down, send to human judge,
// etc.
//
// When using the API, images need to have a minimum of 128 pixels and a maximum file size of 4MB.
// Text can be at most 1024 characters long.
// If the content passed to the text API or the image API exceeds the size limits, the API will return an error code
// that informs about the issue.
type ListManagementImageListsClient struct {
	BaseClient
}

// NewListManagementImageListsClient creates an instance of the ListManagementImageListsClient client.
func NewListManagementImageListsClient(endpoint string) ListManagementImageListsClient {
	return ListManagementImageListsClient{New(endpoint)}
}

// Create creates an image list.
// Parameters:
// contentType - the content type.
// body - schema of the body.
func (client ListManagementImageListsClient) Create(ctx context.Context, contentType string, body Body) (result ImageList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ListManagementImageListsClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, contentType, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentmoderator.ListManagementImageListsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "contentmoderator.ListManagementImageListsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentmoderator.ListManagementImageListsClient", "Create", resp, "Failure responding to request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ListManagementImageListsClient) CreatePreparer(ctx context.Context, contentType string, body Body) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}", urlParameters),
		autorest.WithPath("/contentmoderator/lists/v1.0/imagelists"),
		autorest.WithJSON(body),
		autorest.WithHeader("Content-Type", autorest.String(contentType)))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ListManagementImageListsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ListManagementImageListsClient) CreateResponder(resp *http.Response) (result ImageList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes image list with the list Id equal to list Id passed.
// Parameters:
// listID - list Id of the image list.
func (client ListManagementImageListsClient) Delete(ctx context.Context, listID string) (result String, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ListManagementImageListsClient.Delete")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, listID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentmoderator.ListManagementImageListsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "contentmoderator.ListManagementImageListsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentmoderator.ListManagementImageListsClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ListManagementImageListsClient) DeletePreparer(ctx context.Context, listID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"listId": autorest.Encode("path", listID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("{Endpoint}", urlParameters),
		autorest.WithPathParameters("/contentmoderator/lists/v1.0/imagelists/{listId}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ListManagementImageListsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ListManagementImageListsClient) DeleteResponder(resp *http.Response) (result String, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetAllImageLists gets all the Image Lists.
func (client ListManagementImageListsClient) GetAllImageLists(ctx context.Context) (result ListImageList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ListManagementImageListsClient.GetAllImageLists")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetAllImageListsPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentmoderator.ListManagementImageListsClient", "GetAllImageLists", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetAllImageListsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "contentmoderator.ListManagementImageListsClient", "GetAllImageLists", resp, "Failure sending request")
		return
	}

	result, err = client.GetAllImageListsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentmoderator.ListManagementImageListsClient", "GetAllImageLists", resp, "Failure responding to request")
		return
	}

	return
}

// GetAllImageListsPreparer prepares the GetAllImageLists request.
func (client ListManagementImageListsClient) GetAllImageListsPreparer(ctx context.Context) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{Endpoint}", urlParameters),
		autorest.WithPath("/contentmoderator/lists/v1.0/imagelists"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetAllImageListsSender sends the GetAllImageLists request. The method will close the
// http.Response Body if it receives an error.
func (client ListManagementImageListsClient) GetAllImageListsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetAllImageListsResponder handles the response to the GetAllImageLists request. The method always
// closes the http.Response Body.
func (client ListManagementImageListsClient) GetAllImageListsResponder(resp *http.Response) (result ListImageList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetDetails returns the details of the image list with list Id equal to list Id passed.
// Parameters:
// listID - list Id of the image list.
func (client ListManagementImageListsClient) GetDetails(ctx context.Context, listID string) (result ImageList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ListManagementImageListsClient.GetDetails")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetDetailsPreparer(ctx, listID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentmoderator.ListManagementImageListsClient", "GetDetails", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetDetailsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "contentmoderator.ListManagementImageListsClient", "GetDetails", resp, "Failure sending request")
		return
	}

	result, err = client.GetDetailsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentmoderator.ListManagementImageListsClient", "GetDetails", resp, "Failure responding to request")
		return
	}

	return
}

// GetDetailsPreparer prepares the GetDetails request.
func (client ListManagementImageListsClient) GetDetailsPreparer(ctx context.Context, listID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"listId": autorest.Encode("path", listID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{Endpoint}", urlParameters),
		autorest.WithPathParameters("/contentmoderator/lists/v1.0/imagelists/{listId}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetDetailsSender sends the GetDetails request. The method will close the
// http.Response Body if it receives an error.
func (client ListManagementImageListsClient) GetDetailsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetDetailsResponder handles the response to the GetDetails request. The method always
// closes the http.Response Body.
func (client ListManagementImageListsClient) GetDetailsResponder(resp *http.Response) (result ImageList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// RefreshIndexMethod refreshes the index of the list with list Id equal to list Id passed.
// Parameters:
// listID - list Id of the image list.
func (client ListManagementImageListsClient) RefreshIndexMethod(ctx context.Context, listID string) (result RefreshIndex, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ListManagementImageListsClient.RefreshIndexMethod")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RefreshIndexMethodPreparer(ctx, listID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentmoderator.ListManagementImageListsClient", "RefreshIndexMethod", nil, "Failure preparing request")
		return
	}

	resp, err := client.RefreshIndexMethodSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "contentmoderator.ListManagementImageListsClient", "RefreshIndexMethod", resp, "Failure sending request")
		return
	}

	result, err = client.RefreshIndexMethodResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentmoderator.ListManagementImageListsClient", "RefreshIndexMethod", resp, "Failure responding to request")
		return
	}

	return
}

// RefreshIndexMethodPreparer prepares the RefreshIndexMethod request.
func (client ListManagementImageListsClient) RefreshIndexMethodPreparer(ctx context.Context, listID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"listId": autorest.Encode("path", listID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}", urlParameters),
		autorest.WithPathParameters("/contentmoderator/lists/v1.0/imagelists/{listId}/RefreshIndex", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RefreshIndexMethodSender sends the RefreshIndexMethod request. The method will close the
// http.Response Body if it receives an error.
func (client ListManagementImageListsClient) RefreshIndexMethodSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// RefreshIndexMethodResponder handles the response to the RefreshIndexMethod request. The method always
// closes the http.Response Body.
func (client ListManagementImageListsClient) RefreshIndexMethodResponder(resp *http.Response) (result RefreshIndex, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates an image list with list Id equal to list Id passed.
// Parameters:
// listID - list Id of the image list.
// contentType - the content type.
// body - schema of the body.
func (client ListManagementImageListsClient) Update(ctx context.Context, listID string, contentType string, body Body) (result ImageList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ListManagementImageListsClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, listID, contentType, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentmoderator.ListManagementImageListsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "contentmoderator.ListManagementImageListsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentmoderator.ListManagementImageListsClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ListManagementImageListsClient) UpdatePreparer(ctx context.Context, listID string, contentType string, body Body) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"listId": autorest.Encode("path", listID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithCustomBaseURL("{Endpoint}", urlParameters),
		autorest.WithPathParameters("/contentmoderator/lists/v1.0/imagelists/{listId}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithHeader("Content-Type", autorest.String(contentType)))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ListManagementImageListsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ListManagementImageListsClient) UpdateResponder(resp *http.Response) (result ImageList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
