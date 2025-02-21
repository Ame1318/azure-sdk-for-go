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

// VirtualApplianceSitesClient is the network Client
type VirtualApplianceSitesClient struct {
	BaseClient
}

// NewVirtualApplianceSitesClient creates an instance of the VirtualApplianceSitesClient client.
func NewVirtualApplianceSitesClient(subscriptionID string) VirtualApplianceSitesClient {
	return NewVirtualApplianceSitesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVirtualApplianceSitesClientWithBaseURI creates an instance of the VirtualApplianceSitesClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewVirtualApplianceSitesClientWithBaseURI(baseURI string, subscriptionID string) VirtualApplianceSitesClient {
	return VirtualApplianceSitesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates the specified Network Virtual Appliance Site.
// Parameters:
// resourceGroupName - the name of the resource group.
// networkVirtualApplianceName - the name of the Network Virtual Appliance.
// siteName - the name of the site.
// parameters - parameters supplied to the create or update Network Virtual Appliance Site operation.
func (client VirtualApplianceSitesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, siteName string, parameters VirtualApplianceSite) (result VirtualApplianceSitesCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualApplianceSitesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, networkVirtualApplianceName, siteName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualApplianceSitesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualApplianceSitesClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client VirtualApplianceSitesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, siteName string, parameters VirtualApplianceSite) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkVirtualApplianceName": autorest.Encode("path", networkVirtualApplianceName),
		"resourceGroupName":           autorest.Encode("path", resourceGroupName),
		"siteName":                    autorest.Encode("path", siteName),
		"subscriptionId":              autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.Etag = nil
	parameters.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}/virtualApplianceSites/{siteName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualApplianceSitesClient) CreateOrUpdateSender(req *http.Request) (future VirtualApplianceSitesCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client VirtualApplianceSitesClient) (vas VirtualApplianceSite, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.VirtualApplianceSitesCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("network.VirtualApplianceSitesCreateOrUpdateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		vas.Response.Response, err = future.GetResult(sender)
		if vas.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "network.VirtualApplianceSitesCreateOrUpdateFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && vas.Response.Response.StatusCode != http.StatusNoContent {
			vas, err = client.CreateOrUpdateResponder(vas.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "network.VirtualApplianceSitesCreateOrUpdateFuture", "Result", vas.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client VirtualApplianceSitesClient) CreateOrUpdateResponder(resp *http.Response) (result VirtualApplianceSite, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified site from a Virtual Appliance.
// Parameters:
// resourceGroupName - the name of the resource group.
// networkVirtualApplianceName - the name of the Network Virtual Appliance.
// siteName - the name of the site.
func (client VirtualApplianceSitesClient) Delete(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, siteName string) (result VirtualApplianceSitesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualApplianceSitesClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, networkVirtualApplianceName, siteName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualApplianceSitesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualApplianceSitesClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client VirtualApplianceSitesClient) DeletePreparer(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, siteName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkVirtualApplianceName": autorest.Encode("path", networkVirtualApplianceName),
		"resourceGroupName":           autorest.Encode("path", resourceGroupName),
		"siteName":                    autorest.Encode("path", siteName),
		"subscriptionId":              autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}/virtualApplianceSites/{siteName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualApplianceSitesClient) DeleteSender(req *http.Request) (future VirtualApplianceSitesDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client VirtualApplianceSitesClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.VirtualApplianceSitesDeleteFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("network.VirtualApplianceSitesDeleteFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client VirtualApplianceSitesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the specified Virtual Appliance Site.
// Parameters:
// resourceGroupName - the name of the resource group.
// networkVirtualApplianceName - the name of the Network Virtual Appliance.
// siteName - the name of the site.
func (client VirtualApplianceSitesClient) Get(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, siteName string) (result VirtualApplianceSite, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualApplianceSitesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, networkVirtualApplianceName, siteName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualApplianceSitesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VirtualApplianceSitesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualApplianceSitesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client VirtualApplianceSitesClient) GetPreparer(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, siteName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkVirtualApplianceName": autorest.Encode("path", networkVirtualApplianceName),
		"resourceGroupName":           autorest.Encode("path", resourceGroupName),
		"siteName":                    autorest.Encode("path", siteName),
		"subscriptionId":              autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}/virtualApplianceSites/{siteName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualApplianceSitesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VirtualApplianceSitesClient) GetResponder(resp *http.Response) (result VirtualApplianceSite, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all Network Virtual Appliance Sites in a Network Virtual Appliance resource.
// Parameters:
// resourceGroupName - the name of the resource group.
// networkVirtualApplianceName - the name of the Network Virtual Appliance.
func (client VirtualApplianceSitesClient) List(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string) (result VirtualApplianceSiteListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualApplianceSitesClient.List")
		defer func() {
			sc := -1
			if result.vaslr.Response.Response != nil {
				sc = result.vaslr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, networkVirtualApplianceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualApplianceSitesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.vaslr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VirtualApplianceSitesClient", "List", resp, "Failure sending request")
		return
	}

	result.vaslr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualApplianceSitesClient", "List", resp, "Failure responding to request")
		return
	}
	if result.vaslr.hasNextLink() && result.vaslr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client VirtualApplianceSitesClient) ListPreparer(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkVirtualApplianceName": autorest.Encode("path", networkVirtualApplianceName),
		"resourceGroupName":           autorest.Encode("path", resourceGroupName),
		"subscriptionId":              autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}/virtualApplianceSites", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualApplianceSitesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client VirtualApplianceSitesClient) ListResponder(resp *http.Response) (result VirtualApplianceSiteListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client VirtualApplianceSitesClient) listNextResults(ctx context.Context, lastResults VirtualApplianceSiteListResult) (result VirtualApplianceSiteListResult, err error) {
	req, err := lastResults.virtualApplianceSiteListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.VirtualApplianceSitesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.VirtualApplianceSitesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualApplianceSitesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client VirtualApplianceSitesClient) ListComplete(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string) (result VirtualApplianceSiteListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualApplianceSitesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, networkVirtualApplianceName)
	return
}
