package delegatednetwork

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

// DelegatedSubnetServiceClient is the DNC web api provides way to create, get and delete dnc controller
type DelegatedSubnetServiceClient struct {
	BaseClient
}

// NewDelegatedSubnetServiceClient creates an instance of the DelegatedSubnetServiceClient client.
func NewDelegatedSubnetServiceClient(subscriptionID string) DelegatedSubnetServiceClient {
	return NewDelegatedSubnetServiceClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDelegatedSubnetServiceClientWithBaseURI creates an instance of the DelegatedSubnetServiceClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewDelegatedSubnetServiceClientWithBaseURI(baseURI string, subscriptionID string) DelegatedSubnetServiceClient {
	return DelegatedSubnetServiceClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// DeleteDetails delete dnc DelegatedSubnet.
// Parameters:
// resourceGroupName - the name of the Azure Resource group of which a given DelegatedNetwork resource is part.
// This name must be at least 1 character in length, and no more than 90.
// resourceName - the name of the resource. It must be a minimum of 3 characters, and a maximum of 63.
func (client DelegatedSubnetServiceClient) DeleteDetails(ctx context.Context, resourceGroupName string, resourceName string) (result DelegatedSubnetServiceDeleteDetailsFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DelegatedSubnetServiceClient.DeleteDetails")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "resourceName", Name: validation.Pattern, Rule: `^[a-z][a-z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("delegatednetwork.DelegatedSubnetServiceClient", "DeleteDetails", err.Error())
	}

	req, err := client.DeleteDetailsPreparer(ctx, resourceGroupName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "delegatednetwork.DelegatedSubnetServiceClient", "DeleteDetails", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteDetailsSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "delegatednetwork.DelegatedSubnetServiceClient", "DeleteDetails", nil, "Failure sending request")
		return
	}

	return
}

// DeleteDetailsPreparer prepares the DeleteDetails request.
func (client DelegatedSubnetServiceClient) DeleteDetailsPreparer(ctx context.Context, resourceGroupName string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-08-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DelegatedNetwork/delegatedSubnets/{resourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteDetailsSender sends the DeleteDetails request. The method will close the
// http.Response Body if it receives an error.
func (client DelegatedSubnetServiceClient) DeleteDetailsSender(req *http.Request) (future DelegatedSubnetServiceDeleteDetailsFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client DelegatedSubnetServiceClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "delegatednetwork.DelegatedSubnetServiceDeleteDetailsFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("delegatednetwork.DelegatedSubnetServiceDeleteDetailsFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// DeleteDetailsResponder handles the response to the DeleteDetails request. The method always
// closes the http.Response Body.
func (client DelegatedSubnetServiceClient) DeleteDetailsResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetDetails gets details about the specified dnc DelegatedSubnet Link.
// Parameters:
// resourceGroupName - the name of the Azure Resource group of which a given DelegatedNetwork resource is part.
// This name must be at least 1 character in length, and no more than 90.
// resourceName - the name of the resource. It must be a minimum of 3 characters, and a maximum of 63.
func (client DelegatedSubnetServiceClient) GetDetails(ctx context.Context, resourceGroupName string, resourceName string) (result DelegatedSubnet, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DelegatedSubnetServiceClient.GetDetails")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "resourceName", Name: validation.Pattern, Rule: `^[a-z][a-z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("delegatednetwork.DelegatedSubnetServiceClient", "GetDetails", err.Error())
	}

	req, err := client.GetDetailsPreparer(ctx, resourceGroupName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "delegatednetwork.DelegatedSubnetServiceClient", "GetDetails", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetDetailsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "delegatednetwork.DelegatedSubnetServiceClient", "GetDetails", resp, "Failure sending request")
		return
	}

	result, err = client.GetDetailsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "delegatednetwork.DelegatedSubnetServiceClient", "GetDetails", resp, "Failure responding to request")
		return
	}

	return
}

// GetDetailsPreparer prepares the GetDetails request.
func (client DelegatedSubnetServiceClient) GetDetailsPreparer(ctx context.Context, resourceGroupName string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-08-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DelegatedNetwork/delegatedSubnets/{resourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetDetailsSender sends the GetDetails request. The method will close the
// http.Response Body if it receives an error.
func (client DelegatedSubnetServiceClient) GetDetailsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetDetailsResponder handles the response to the GetDetails request. The method always
// closes the http.Response Body.
func (client DelegatedSubnetServiceClient) GetDetailsResponder(resp *http.Response) (result DelegatedSubnet, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup get all the DelegatedSubnets resources in a resource group.
// Parameters:
// resourceGroupName - the name of the Azure Resource group of which a given DelegatedNetwork resource is part.
// This name must be at least 1 character in length, and no more than 90.
func (client DelegatedSubnetServiceClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result DelegatedSubnetsPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DelegatedSubnetServiceClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.ds.Response.Response != nil {
				sc = result.ds.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("delegatednetwork.DelegatedSubnetServiceClient", "ListByResourceGroup", err.Error())
	}

	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "delegatednetwork.DelegatedSubnetServiceClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.ds.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "delegatednetwork.DelegatedSubnetServiceClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.ds, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "delegatednetwork.DelegatedSubnetServiceClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.ds.hasNextLink() && result.ds.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client DelegatedSubnetServiceClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-08-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DelegatedNetwork/delegatedSubnets", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client DelegatedSubnetServiceClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client DelegatedSubnetServiceClient) ListByResourceGroupResponder(resp *http.Response) (result DelegatedSubnets, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client DelegatedSubnetServiceClient) listByResourceGroupNextResults(ctx context.Context, lastResults DelegatedSubnets) (result DelegatedSubnets, err error) {
	req, err := lastResults.delegatedSubnetsPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "delegatednetwork.DelegatedSubnetServiceClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "delegatednetwork.DelegatedSubnetServiceClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "delegatednetwork.DelegatedSubnetServiceClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client DelegatedSubnetServiceClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result DelegatedSubnetsIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DelegatedSubnetServiceClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}

// ListBySubscription get all the DelegatedSubnets resources in a subscription.
func (client DelegatedSubnetServiceClient) ListBySubscription(ctx context.Context) (result DelegatedSubnetsPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DelegatedSubnetServiceClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.ds.Response.Response != nil {
				sc = result.ds.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listBySubscriptionNextResults
	req, err := client.ListBySubscriptionPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "delegatednetwork.DelegatedSubnetServiceClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.ds.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "delegatednetwork.DelegatedSubnetServiceClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result.ds, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "delegatednetwork.DelegatedSubnetServiceClient", "ListBySubscription", resp, "Failure responding to request")
		return
	}
	if result.ds.hasNextLink() && result.ds.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client DelegatedSubnetServiceClient) ListBySubscriptionPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-08-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DelegatedNetwork/delegatedSubnets", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client DelegatedSubnetServiceClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client DelegatedSubnetServiceClient) ListBySubscriptionResponder(resp *http.Response) (result DelegatedSubnets, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySubscriptionNextResults retrieves the next set of results, if any.
func (client DelegatedSubnetServiceClient) listBySubscriptionNextResults(ctx context.Context, lastResults DelegatedSubnets) (result DelegatedSubnets, err error) {
	req, err := lastResults.delegatedSubnetsPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "delegatednetwork.DelegatedSubnetServiceClient", "listBySubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "delegatednetwork.DelegatedSubnetServiceClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "delegatednetwork.DelegatedSubnetServiceClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client DelegatedSubnetServiceClient) ListBySubscriptionComplete(ctx context.Context) (result DelegatedSubnetsIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DelegatedSubnetServiceClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySubscription(ctx)
	return
}

// PatchDetails patch delegated subnet resource
// Parameters:
// resourceGroupName - the name of the Azure Resource group of which a given DelegatedNetwork resource is part.
// This name must be at least 1 character in length, and no more than 90.
// resourceName - the name of the resource. It must be a minimum of 3 characters, and a maximum of 63.
// parameters - delegated subnet details.
func (client DelegatedSubnetServiceClient) PatchDetails(ctx context.Context, resourceGroupName string, resourceName string, parameters ResourceUpdateParameters) (result DelegatedSubnetServicePatchDetailsFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DelegatedSubnetServiceClient.PatchDetails")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "resourceName", Name: validation.Pattern, Rule: `^[a-z][a-z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("delegatednetwork.DelegatedSubnetServiceClient", "PatchDetails", err.Error())
	}

	req, err := client.PatchDetailsPreparer(ctx, resourceGroupName, resourceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "delegatednetwork.DelegatedSubnetServiceClient", "PatchDetails", nil, "Failure preparing request")
		return
	}

	result, err = client.PatchDetailsSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "delegatednetwork.DelegatedSubnetServiceClient", "PatchDetails", nil, "Failure sending request")
		return
	}

	return
}

// PatchDetailsPreparer prepares the PatchDetails request.
func (client DelegatedSubnetServiceClient) PatchDetailsPreparer(ctx context.Context, resourceGroupName string, resourceName string, parameters ResourceUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-08-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DelegatedNetwork/delegatedSubnets/{resourceName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PatchDetailsSender sends the PatchDetails request. The method will close the
// http.Response Body if it receives an error.
func (client DelegatedSubnetServiceClient) PatchDetailsSender(req *http.Request) (future DelegatedSubnetServicePatchDetailsFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client DelegatedSubnetServiceClient) (ds DelegatedSubnet, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "delegatednetwork.DelegatedSubnetServicePatchDetailsFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("delegatednetwork.DelegatedSubnetServicePatchDetailsFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		ds.Response.Response, err = future.GetResult(sender)
		if ds.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "delegatednetwork.DelegatedSubnetServicePatchDetailsFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && ds.Response.Response.StatusCode != http.StatusNoContent {
			ds, err = client.PatchDetailsResponder(ds.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "delegatednetwork.DelegatedSubnetServicePatchDetailsFuture", "Result", ds.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// PatchDetailsResponder handles the response to the PatchDetails request. The method always
// closes the http.Response Body.
func (client DelegatedSubnetServiceClient) PatchDetailsResponder(resp *http.Response) (result DelegatedSubnet, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// PutDetails put delegated subnet resource
// Parameters:
// resourceGroupName - the name of the Azure Resource group of which a given DelegatedNetwork resource is part.
// This name must be at least 1 character in length, and no more than 90.
// resourceName - the name of the resource. It must be a minimum of 3 characters, and a maximum of 63.
// parameters - delegated subnet details.
func (client DelegatedSubnetServiceClient) PutDetails(ctx context.Context, resourceGroupName string, resourceName string, parameters DelegatedSubnet) (result DelegatedSubnetServicePutDetailsFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DelegatedSubnetServiceClient.PutDetails")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "resourceName", Name: validation.Pattern, Rule: `^[a-z][a-z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("delegatednetwork.DelegatedSubnetServiceClient", "PutDetails", err.Error())
	}

	req, err := client.PutDetailsPreparer(ctx, resourceGroupName, resourceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "delegatednetwork.DelegatedSubnetServiceClient", "PutDetails", nil, "Failure preparing request")
		return
	}

	result, err = client.PutDetailsSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "delegatednetwork.DelegatedSubnetServiceClient", "PutDetails", nil, "Failure sending request")
		return
	}

	return
}

// PutDetailsPreparer prepares the PutDetails request.
func (client DelegatedSubnetServiceClient) PutDetailsPreparer(ctx context.Context, resourceGroupName string, resourceName string, parameters DelegatedSubnet) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-08-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.DelegatedSubnetProperties = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DelegatedNetwork/delegatedSubnets/{resourceName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutDetailsSender sends the PutDetails request. The method will close the
// http.Response Body if it receives an error.
func (client DelegatedSubnetServiceClient) PutDetailsSender(req *http.Request) (future DelegatedSubnetServicePutDetailsFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client DelegatedSubnetServiceClient) (ds DelegatedSubnet, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "delegatednetwork.DelegatedSubnetServicePutDetailsFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("delegatednetwork.DelegatedSubnetServicePutDetailsFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		ds.Response.Response, err = future.GetResult(sender)
		if ds.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "delegatednetwork.DelegatedSubnetServicePutDetailsFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && ds.Response.Response.StatusCode != http.StatusNoContent {
			ds, err = client.PutDetailsResponder(ds.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "delegatednetwork.DelegatedSubnetServicePutDetailsFuture", "Result", ds.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// PutDetailsResponder handles the response to the PutDetails request. The method always
// closes the http.Response Body.
func (client DelegatedSubnetServiceClient) PutDetailsResponder(resp *http.Response) (result DelegatedSubnet, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
