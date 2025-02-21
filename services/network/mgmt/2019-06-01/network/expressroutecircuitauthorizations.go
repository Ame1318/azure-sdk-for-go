package network

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

// ExpressRouteCircuitAuthorizationsClient is the network Client
type ExpressRouteCircuitAuthorizationsClient struct {
	BaseClient
}

// NewExpressRouteCircuitAuthorizationsClient creates an instance of the ExpressRouteCircuitAuthorizationsClient
// client.
func NewExpressRouteCircuitAuthorizationsClient(subscriptionID string) ExpressRouteCircuitAuthorizationsClient {
	return NewExpressRouteCircuitAuthorizationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewExpressRouteCircuitAuthorizationsClientWithBaseURI creates an instance of the
// ExpressRouteCircuitAuthorizationsClient client using a custom endpoint.  Use this when interacting with an Azure
// cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewExpressRouteCircuitAuthorizationsClientWithBaseURI(baseURI string, subscriptionID string) ExpressRouteCircuitAuthorizationsClient {
	return ExpressRouteCircuitAuthorizationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates an authorization in the specified express route circuit.
// Parameters:
// resourceGroupName - the name of the resource group.
// circuitName - the name of the express route circuit.
// authorizationName - the name of the authorization.
// authorizationParameters - parameters supplied to the create or update express route circuit authorization
// operation.
func (client ExpressRouteCircuitAuthorizationsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string, authorizationParameters ExpressRouteCircuitAuthorization) (result ExpressRouteCircuitAuthorizationsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRouteCircuitAuthorizationsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, circuitName, authorizationName, authorizationParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitAuthorizationsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitAuthorizationsClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ExpressRouteCircuitAuthorizationsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string, authorizationParameters ExpressRouteCircuitAuthorization) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"authorizationName": autorest.Encode("path", authorizationName),
		"circuitName":       autorest.Encode("path", circuitName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	authorizationParameters.Etag = nil
	authorizationParameters.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/authorizations/{authorizationName}", pathParameters),
		autorest.WithJSON(authorizationParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitAuthorizationsClient) CreateOrUpdateSender(req *http.Request) (future ExpressRouteCircuitAuthorizationsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client ExpressRouteCircuitAuthorizationsClient) (erca ExpressRouteCircuitAuthorization, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitAuthorizationsCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("network.ExpressRouteCircuitAuthorizationsCreateOrUpdateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		erca.Response.Response, err = future.GetResult(sender)
		if erca.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitAuthorizationsCreateOrUpdateFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && erca.Response.Response.StatusCode != http.StatusNoContent {
			erca, err = client.CreateOrUpdateResponder(erca.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitAuthorizationsCreateOrUpdateFuture", "Result", erca.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitAuthorizationsClient) CreateOrUpdateResponder(resp *http.Response) (result ExpressRouteCircuitAuthorization, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified authorization from the specified express route circuit.
// Parameters:
// resourceGroupName - the name of the resource group.
// circuitName - the name of the express route circuit.
// authorizationName - the name of the authorization.
func (client ExpressRouteCircuitAuthorizationsClient) Delete(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string) (result ExpressRouteCircuitAuthorizationsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRouteCircuitAuthorizationsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, circuitName, authorizationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitAuthorizationsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitAuthorizationsClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ExpressRouteCircuitAuthorizationsClient) DeletePreparer(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"authorizationName": autorest.Encode("path", authorizationName),
		"circuitName":       autorest.Encode("path", circuitName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/authorizations/{authorizationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitAuthorizationsClient) DeleteSender(req *http.Request) (future ExpressRouteCircuitAuthorizationsDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client ExpressRouteCircuitAuthorizationsClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitAuthorizationsDeleteFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("network.ExpressRouteCircuitAuthorizationsDeleteFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitAuthorizationsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the specified authorization from the specified express route circuit.
// Parameters:
// resourceGroupName - the name of the resource group.
// circuitName - the name of the express route circuit.
// authorizationName - the name of the authorization.
func (client ExpressRouteCircuitAuthorizationsClient) Get(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string) (result ExpressRouteCircuitAuthorization, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRouteCircuitAuthorizationsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, circuitName, authorizationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitAuthorizationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitAuthorizationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitAuthorizationsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ExpressRouteCircuitAuthorizationsClient) GetPreparer(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"authorizationName": autorest.Encode("path", authorizationName),
		"circuitName":       autorest.Encode("path", circuitName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/authorizations/{authorizationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitAuthorizationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitAuthorizationsClient) GetResponder(resp *http.Response) (result ExpressRouteCircuitAuthorization, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all authorizations in an express route circuit.
// Parameters:
// resourceGroupName - the name of the resource group.
// circuitName - the name of the circuit.
func (client ExpressRouteCircuitAuthorizationsClient) List(ctx context.Context, resourceGroupName string, circuitName string) (result AuthorizationListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRouteCircuitAuthorizationsClient.List")
		defer func() {
			sc := -1
			if result.alr.Response.Response != nil {
				sc = result.alr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, circuitName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitAuthorizationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.alr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitAuthorizationsClient", "List", resp, "Failure sending request")
		return
	}

	result.alr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitAuthorizationsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.alr.hasNextLink() && result.alr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ExpressRouteCircuitAuthorizationsClient) ListPreparer(ctx context.Context, resourceGroupName string, circuitName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"circuitName":       autorest.Encode("path", circuitName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/authorizations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitAuthorizationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitAuthorizationsClient) ListResponder(resp *http.Response) (result AuthorizationListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ExpressRouteCircuitAuthorizationsClient) listNextResults(ctx context.Context, lastResults AuthorizationListResult) (result AuthorizationListResult, err error) {
	req, err := lastResults.authorizationListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.ExpressRouteCircuitAuthorizationsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.ExpressRouteCircuitAuthorizationsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteCircuitAuthorizationsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ExpressRouteCircuitAuthorizationsClient) ListComplete(ctx context.Context, resourceGroupName string, circuitName string) (result AuthorizationListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRouteCircuitAuthorizationsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, circuitName)
	return
}
