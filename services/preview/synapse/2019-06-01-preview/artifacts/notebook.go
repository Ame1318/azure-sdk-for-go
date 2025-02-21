package artifacts

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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// NotebookClient is the client for the Notebook methods of the Artifacts service.
type NotebookClient struct {
	BaseClient
}

// NewNotebookClient creates an instance of the NotebookClient client.
func NewNotebookClient(endpoint string) NotebookClient {
	return NotebookClient{New(endpoint)}
}

// CreateOrUpdateNotebook creates or updates a Note Book.
// Parameters:
// notebookName - the notebook name.
// notebook - note book resource definition.
// ifMatch - eTag of the Note book entity.  Should only be specified for update, for which it should match
// existing entity or can be * for unconditional update.
func (client NotebookClient) CreateOrUpdateNotebook(ctx context.Context, notebookName string, notebook NotebookResource, ifMatch string) (result NotebookCreateOrUpdateNotebookFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NotebookClient.CreateOrUpdateNotebook")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: notebook,
			Constraints: []validation.Constraint{{Target: "notebook.Name", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "notebook.Properties", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "notebook.Properties.BigDataPool", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "notebook.Properties.BigDataPool.Type", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "notebook.Properties.BigDataPool.ReferenceName", Name: validation.Null, Rule: true, Chain: nil},
						}},
						{Target: "notebook.Properties.SessionProperties", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "notebook.Properties.SessionProperties.DriverMemory", Name: validation.Null, Rule: true, Chain: nil},
								{Target: "notebook.Properties.SessionProperties.DriverCores", Name: validation.Null, Rule: true, Chain: nil},
								{Target: "notebook.Properties.SessionProperties.ExecutorMemory", Name: validation.Null, Rule: true, Chain: nil},
								{Target: "notebook.Properties.SessionProperties.ExecutorCores", Name: validation.Null, Rule: true, Chain: nil},
								{Target: "notebook.Properties.SessionProperties.NumExecutors", Name: validation.Null, Rule: true, Chain: nil},
							}},
						{Target: "notebook.Properties.Metadata", Name: validation.Null, Rule: true,
							Chain: []validation.Constraint{{Target: "notebook.Properties.Metadata.Kernelspec", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "notebook.Properties.Metadata.Kernelspec.Name", Name: validation.Null, Rule: true, Chain: nil},
									{Target: "notebook.Properties.Metadata.Kernelspec.DisplayName", Name: validation.Null, Rule: true, Chain: nil},
								}},
								{Target: "notebook.Properties.Metadata.LanguageInfo", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "notebook.Properties.Metadata.LanguageInfo.Name", Name: validation.Null, Rule: true, Chain: nil}}},
							}},
						{Target: "notebook.Properties.Nbformat", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "notebook.Properties.NbformatMinor", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "notebook.Properties.Cells", Name: validation.Null, Rule: true, Chain: nil},
					}}}}}); err != nil {
		return result, validation.NewError("artifacts.NotebookClient", "CreateOrUpdateNotebook", err.Error())
	}

	req, err := client.CreateOrUpdateNotebookPreparer(ctx, notebookName, notebook, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.NotebookClient", "CreateOrUpdateNotebook", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateNotebookSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.NotebookClient", "CreateOrUpdateNotebook", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdateNotebookPreparer prepares the CreateOrUpdateNotebook request.
func (client NotebookClient) CreateOrUpdateNotebookPreparer(ctx context.Context, notebookName string, notebook NotebookResource, ifMatch string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"notebookName": autorest.Encode("path", notebookName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	notebook.ID = nil
	notebook.Type = nil
	notebook.Etag = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithCustomBaseURL("{endpoint}", urlParameters),
		autorest.WithPathParameters("/notebooks/{notebookName}", pathParameters),
		autorest.WithJSON(notebook),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateNotebookSender sends the CreateOrUpdateNotebook request. The method will close the
// http.Response Body if it receives an error.
func (client NotebookClient) CreateOrUpdateNotebookSender(req *http.Request) (future NotebookCreateOrUpdateNotebookFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client NotebookClient) (nr NotebookResource, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "artifacts.NotebookCreateOrUpdateNotebookFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("artifacts.NotebookCreateOrUpdateNotebookFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		nr.Response.Response, err = future.GetResult(sender)
		if nr.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "artifacts.NotebookCreateOrUpdateNotebookFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && nr.Response.Response.StatusCode != http.StatusNoContent {
			nr, err = client.CreateOrUpdateNotebookResponder(nr.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "artifacts.NotebookCreateOrUpdateNotebookFuture", "Result", nr.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateOrUpdateNotebookResponder handles the response to the CreateOrUpdateNotebook request. The method always
// closes the http.Response Body.
func (client NotebookClient) CreateOrUpdateNotebookResponder(resp *http.Response) (result NotebookResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteNotebook deletes a Note book.
// Parameters:
// notebookName - the notebook name.
func (client NotebookClient) DeleteNotebook(ctx context.Context, notebookName string) (result NotebookDeleteNotebookFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NotebookClient.DeleteNotebook")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeleteNotebookPreparer(ctx, notebookName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.NotebookClient", "DeleteNotebook", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteNotebookSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.NotebookClient", "DeleteNotebook", nil, "Failure sending request")
		return
	}

	return
}

// DeleteNotebookPreparer prepares the DeleteNotebook request.
func (client NotebookClient) DeleteNotebookPreparer(ctx context.Context, notebookName string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"notebookName": autorest.Encode("path", notebookName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("{endpoint}", urlParameters),
		autorest.WithPathParameters("/notebooks/{notebookName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteNotebookSender sends the DeleteNotebook request. The method will close the
// http.Response Body if it receives an error.
func (client NotebookClient) DeleteNotebookSender(req *http.Request) (future NotebookDeleteNotebookFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client NotebookClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "artifacts.NotebookDeleteNotebookFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("artifacts.NotebookDeleteNotebookFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// DeleteNotebookResponder handles the response to the DeleteNotebook request. The method always
// closes the http.Response Body.
func (client NotebookClient) DeleteNotebookResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetNotebook gets a Note Book.
// Parameters:
// notebookName - the notebook name.
// ifNoneMatch - eTag of the Notebook entity. Should only be specified for get. If the ETag matches the
// existing entity tag, or if * was provided, then no content will be returned.
func (client NotebookClient) GetNotebook(ctx context.Context, notebookName string, ifNoneMatch string) (result NotebookResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NotebookClient.GetNotebook")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetNotebookPreparer(ctx, notebookName, ifNoneMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.NotebookClient", "GetNotebook", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetNotebookSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "artifacts.NotebookClient", "GetNotebook", resp, "Failure sending request")
		return
	}

	result, err = client.GetNotebookResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.NotebookClient", "GetNotebook", resp, "Failure responding to request")
		return
	}

	return
}

// GetNotebookPreparer prepares the GetNotebook request.
func (client NotebookClient) GetNotebookPreparer(ctx context.Context, notebookName string, ifNoneMatch string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"notebookName": autorest.Encode("path", notebookName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{endpoint}", urlParameters),
		autorest.WithPathParameters("/notebooks/{notebookName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(ifNoneMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-None-Match", autorest.String(ifNoneMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetNotebookSender sends the GetNotebook request. The method will close the
// http.Response Body if it receives an error.
func (client NotebookClient) GetNotebookSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetNotebookResponder handles the response to the GetNotebook request. The method always
// closes the http.Response Body.
func (client NotebookClient) GetNotebookResponder(resp *http.Response) (result NotebookResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNotModified),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetNotebooksByWorkspace lists Notebooks.
func (client NotebookClient) GetNotebooksByWorkspace(ctx context.Context) (result NotebookListResponsePage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NotebookClient.GetNotebooksByWorkspace")
		defer func() {
			sc := -1
			if result.nlr.Response.Response != nil {
				sc = result.nlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.getNotebooksByWorkspaceNextResults
	req, err := client.GetNotebooksByWorkspacePreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.NotebookClient", "GetNotebooksByWorkspace", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetNotebooksByWorkspaceSender(req)
	if err != nil {
		result.nlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "artifacts.NotebookClient", "GetNotebooksByWorkspace", resp, "Failure sending request")
		return
	}

	result.nlr, err = client.GetNotebooksByWorkspaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.NotebookClient", "GetNotebooksByWorkspace", resp, "Failure responding to request")
		return
	}
	if result.nlr.hasNextLink() && result.nlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// GetNotebooksByWorkspacePreparer prepares the GetNotebooksByWorkspace request.
func (client NotebookClient) GetNotebooksByWorkspacePreparer(ctx context.Context) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"endpoint": client.Endpoint,
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{endpoint}", urlParameters),
		autorest.WithPath("/notebooks"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetNotebooksByWorkspaceSender sends the GetNotebooksByWorkspace request. The method will close the
// http.Response Body if it receives an error.
func (client NotebookClient) GetNotebooksByWorkspaceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetNotebooksByWorkspaceResponder handles the response to the GetNotebooksByWorkspace request. The method always
// closes the http.Response Body.
func (client NotebookClient) GetNotebooksByWorkspaceResponder(resp *http.Response) (result NotebookListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// getNotebooksByWorkspaceNextResults retrieves the next set of results, if any.
func (client NotebookClient) getNotebooksByWorkspaceNextResults(ctx context.Context, lastResults NotebookListResponse) (result NotebookListResponse, err error) {
	req, err := lastResults.notebookListResponsePreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "artifacts.NotebookClient", "getNotebooksByWorkspaceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.GetNotebooksByWorkspaceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "artifacts.NotebookClient", "getNotebooksByWorkspaceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.GetNotebooksByWorkspaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.NotebookClient", "getNotebooksByWorkspaceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// GetNotebooksByWorkspaceComplete enumerates all values, automatically crossing page boundaries as required.
func (client NotebookClient) GetNotebooksByWorkspaceComplete(ctx context.Context) (result NotebookListResponseIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NotebookClient.GetNotebooksByWorkspace")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.GetNotebooksByWorkspace(ctx)
	return
}

// GetNotebookSummaryByWorkSpace lists a summary of Notebooks.
func (client NotebookClient) GetNotebookSummaryByWorkSpace(ctx context.Context) (result NotebookListResponsePage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NotebookClient.GetNotebookSummaryByWorkSpace")
		defer func() {
			sc := -1
			if result.nlr.Response.Response != nil {
				sc = result.nlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.getNotebookSummaryByWorkSpaceNextResults
	req, err := client.GetNotebookSummaryByWorkSpacePreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.NotebookClient", "GetNotebookSummaryByWorkSpace", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetNotebookSummaryByWorkSpaceSender(req)
	if err != nil {
		result.nlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "artifacts.NotebookClient", "GetNotebookSummaryByWorkSpace", resp, "Failure sending request")
		return
	}

	result.nlr, err = client.GetNotebookSummaryByWorkSpaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.NotebookClient", "GetNotebookSummaryByWorkSpace", resp, "Failure responding to request")
		return
	}
	if result.nlr.hasNextLink() && result.nlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// GetNotebookSummaryByWorkSpacePreparer prepares the GetNotebookSummaryByWorkSpace request.
func (client NotebookClient) GetNotebookSummaryByWorkSpacePreparer(ctx context.Context) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"endpoint": client.Endpoint,
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{endpoint}", urlParameters),
		autorest.WithPath("/notebooks/summary"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetNotebookSummaryByWorkSpaceSender sends the GetNotebookSummaryByWorkSpace request. The method will close the
// http.Response Body if it receives an error.
func (client NotebookClient) GetNotebookSummaryByWorkSpaceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetNotebookSummaryByWorkSpaceResponder handles the response to the GetNotebookSummaryByWorkSpace request. The method always
// closes the http.Response Body.
func (client NotebookClient) GetNotebookSummaryByWorkSpaceResponder(resp *http.Response) (result NotebookListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// getNotebookSummaryByWorkSpaceNextResults retrieves the next set of results, if any.
func (client NotebookClient) getNotebookSummaryByWorkSpaceNextResults(ctx context.Context, lastResults NotebookListResponse) (result NotebookListResponse, err error) {
	req, err := lastResults.notebookListResponsePreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "artifacts.NotebookClient", "getNotebookSummaryByWorkSpaceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.GetNotebookSummaryByWorkSpaceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "artifacts.NotebookClient", "getNotebookSummaryByWorkSpaceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.GetNotebookSummaryByWorkSpaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.NotebookClient", "getNotebookSummaryByWorkSpaceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// GetNotebookSummaryByWorkSpaceComplete enumerates all values, automatically crossing page boundaries as required.
func (client NotebookClient) GetNotebookSummaryByWorkSpaceComplete(ctx context.Context) (result NotebookListResponseIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NotebookClient.GetNotebookSummaryByWorkSpace")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.GetNotebookSummaryByWorkSpace(ctx)
	return
}

// RenameNotebook renames a notebook.
// Parameters:
// notebookName - the notebook name.
// request - proposed new name.
func (client NotebookClient) RenameNotebook(ctx context.Context, notebookName string, request RenameRequest) (result NotebookRenameNotebookFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NotebookClient.RenameNotebook")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: request,
			Constraints: []validation.Constraint{{Target: "request.NewName", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "request.NewName", Name: validation.MaxLength, Rule: 260, Chain: nil},
					{Target: "request.NewName", Name: validation.MinLength, Rule: 1, Chain: nil},
					{Target: "request.NewName", Name: validation.Pattern, Rule: `^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("artifacts.NotebookClient", "RenameNotebook", err.Error())
	}

	req, err := client.RenameNotebookPreparer(ctx, notebookName, request)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.NotebookClient", "RenameNotebook", nil, "Failure preparing request")
		return
	}

	result, err = client.RenameNotebookSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.NotebookClient", "RenameNotebook", nil, "Failure sending request")
		return
	}

	return
}

// RenameNotebookPreparer prepares the RenameNotebook request.
func (client NotebookClient) RenameNotebookPreparer(ctx context.Context, notebookName string, request RenameRequest) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"notebookName": autorest.Encode("path", notebookName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{endpoint}", urlParameters),
		autorest.WithPathParameters("/notebooks/{notebookName}/rename", pathParameters),
		autorest.WithJSON(request),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RenameNotebookSender sends the RenameNotebook request. The method will close the
// http.Response Body if it receives an error.
func (client NotebookClient) RenameNotebookSender(req *http.Request) (future NotebookRenameNotebookFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client NotebookClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "artifacts.NotebookRenameNotebookFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("artifacts.NotebookRenameNotebookFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// RenameNotebookResponder handles the response to the RenameNotebook request. The method always
// closes the http.Response Body.
func (client NotebookClient) RenameNotebookResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}
