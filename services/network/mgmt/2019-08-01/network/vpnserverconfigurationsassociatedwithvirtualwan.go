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

// VpnServerConfigurationsAssociatedWithVirtualWanClient is the network Client
type VpnServerConfigurationsAssociatedWithVirtualWanClient struct {
	BaseClient
}

// NewVpnServerConfigurationsAssociatedWithVirtualWanClient creates an instance of the
// VpnServerConfigurationsAssociatedWithVirtualWanClient client.
func NewVpnServerConfigurationsAssociatedWithVirtualWanClient(subscriptionID string) VpnServerConfigurationsAssociatedWithVirtualWanClient {
	return NewVpnServerConfigurationsAssociatedWithVirtualWanClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVpnServerConfigurationsAssociatedWithVirtualWanClientWithBaseURI creates an instance of the
// VpnServerConfigurationsAssociatedWithVirtualWanClient client using a custom endpoint.  Use this when interacting
// with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewVpnServerConfigurationsAssociatedWithVirtualWanClientWithBaseURI(baseURI string, subscriptionID string) VpnServerConfigurationsAssociatedWithVirtualWanClient {
	return VpnServerConfigurationsAssociatedWithVirtualWanClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List gives the list of VpnServerConfigurations associated with Virtual Wan in a resource group.
// Parameters:
// resourceGroupName - the resource group name.
// virtualWANName - the name of the VirtualWAN whose associated VpnServerConfigurations is needed.
func (client VpnServerConfigurationsAssociatedWithVirtualWanClient) List(ctx context.Context, resourceGroupName string, virtualWANName string) (result VpnServerConfigurationsAssociatedWithVirtualWanListFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VpnServerConfigurationsAssociatedWithVirtualWanClient.List")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx, resourceGroupName, virtualWANName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnServerConfigurationsAssociatedWithVirtualWanClient", "List", nil, "Failure preparing request")
		return
	}

	result, err = client.ListSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnServerConfigurationsAssociatedWithVirtualWanClient", "List", nil, "Failure sending request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client VpnServerConfigurationsAssociatedWithVirtualWanClient) ListPreparer(ctx context.Context, resourceGroupName string, virtualWANName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"virtualWANName":    autorest.Encode("path", virtualWANName),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{virtualWANName}/vpnServerConfigurations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client VpnServerConfigurationsAssociatedWithVirtualWanClient) ListSender(req *http.Request) (future VpnServerConfigurationsAssociatedWithVirtualWanListFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client VpnServerConfigurationsAssociatedWithVirtualWanClient) (vscr VpnServerConfigurationsResponse, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.VpnServerConfigurationsAssociatedWithVirtualWanListFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("network.VpnServerConfigurationsAssociatedWithVirtualWanListFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		vscr.Response.Response, err = future.GetResult(sender)
		if vscr.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "network.VpnServerConfigurationsAssociatedWithVirtualWanListFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && vscr.Response.Response.StatusCode != http.StatusNoContent {
			vscr, err = client.ListResponder(vscr.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "network.VpnServerConfigurationsAssociatedWithVirtualWanListFuture", "Result", vscr.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client VpnServerConfigurationsAssociatedWithVirtualWanClient) ListResponder(resp *http.Response) (result VpnServerConfigurationsResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
