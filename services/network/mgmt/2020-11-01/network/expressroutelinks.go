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

// ExpressRouteLinksClient is the network Client
type ExpressRouteLinksClient struct {
	BaseClient
}

// NewExpressRouteLinksClient creates an instance of the ExpressRouteLinksClient client.
func NewExpressRouteLinksClient(subscriptionID string) ExpressRouteLinksClient {
	return NewExpressRouteLinksClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewExpressRouteLinksClientWithBaseURI creates an instance of the ExpressRouteLinksClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewExpressRouteLinksClientWithBaseURI(baseURI string, subscriptionID string) ExpressRouteLinksClient {
	return ExpressRouteLinksClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get retrieves the specified ExpressRouteLink resource.
// Parameters:
// resourceGroupName - the name of the resource group.
// expressRoutePortName - the name of the ExpressRoutePort resource.
// linkName - the name of the ExpressRouteLink resource.
func (client ExpressRouteLinksClient) Get(ctx context.Context, resourceGroupName string, expressRoutePortName string, linkName string) (result ExpressRouteLink, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRouteLinksClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, expressRoutePortName, linkName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteLinksClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ExpressRouteLinksClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteLinksClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ExpressRouteLinksClient) GetPreparer(ctx context.Context, resourceGroupName string, expressRoutePortName string, linkName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"expressRoutePortName": autorest.Encode("path", expressRoutePortName),
		"linkName":             autorest.Encode("path", linkName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}/links/{linkName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteLinksClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ExpressRouteLinksClient) GetResponder(resp *http.Response) (result ExpressRouteLink, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List retrieve the ExpressRouteLink sub-resources of the specified ExpressRoutePort resource.
// Parameters:
// resourceGroupName - the name of the resource group.
// expressRoutePortName - the name of the ExpressRoutePort resource.
func (client ExpressRouteLinksClient) List(ctx context.Context, resourceGroupName string, expressRoutePortName string) (result ExpressRouteLinkListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRouteLinksClient.List")
		defer func() {
			sc := -1
			if result.erllr.Response.Response != nil {
				sc = result.erllr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, expressRoutePortName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteLinksClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.erllr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ExpressRouteLinksClient", "List", resp, "Failure sending request")
		return
	}

	result.erllr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteLinksClient", "List", resp, "Failure responding to request")
		return
	}
	if result.erllr.hasNextLink() && result.erllr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ExpressRouteLinksClient) ListPreparer(ctx context.Context, resourceGroupName string, expressRoutePortName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"expressRoutePortName": autorest.Encode("path", expressRoutePortName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}/links", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteLinksClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ExpressRouteLinksClient) ListResponder(resp *http.Response) (result ExpressRouteLinkListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ExpressRouteLinksClient) listNextResults(ctx context.Context, lastResults ExpressRouteLinkListResult) (result ExpressRouteLinkListResult, err error) {
	req, err := lastResults.expressRouteLinkListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.ExpressRouteLinksClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.ExpressRouteLinksClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteLinksClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ExpressRouteLinksClient) ListComplete(ctx context.Context, resourceGroupName string, expressRoutePortName string) (result ExpressRouteLinkListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRouteLinksClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, expressRoutePortName)
	return
}
