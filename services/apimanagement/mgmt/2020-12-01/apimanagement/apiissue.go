package apimanagement

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

// APIIssueClient is the apiManagement Client
type APIIssueClient struct {
	BaseClient
}

// NewAPIIssueClient creates an instance of the APIIssueClient client.
func NewAPIIssueClient(subscriptionID string) APIIssueClient {
	return NewAPIIssueClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAPIIssueClientWithBaseURI creates an instance of the APIIssueClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewAPIIssueClientWithBaseURI(baseURI string, subscriptionID string) APIIssueClient {
	return APIIssueClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a new Issue for an API or updates an existing one.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// apiid - API identifier. Must be unique in the current API Management service instance.
// issueID - issue identifier. Must be unique in the current API Management service instance.
// parameters - create parameters.
// ifMatch - eTag of the Entity. Not required when creating an entity, but required when updating an entity.
func (client APIIssueClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, apiid string, issueID string, parameters IssueContract, ifMatch string) (result IssueContract, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/APIIssueClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: apiid,
			Constraints: []validation.Constraint{{Target: "apiid", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "apiid", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: issueID,
			Constraints: []validation.Constraint{{Target: "issueID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "issueID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "issueID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.IssueContractProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.IssueContractProperties.Title", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.IssueContractProperties.Description", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.IssueContractProperties.UserID", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("apimanagement.APIIssueClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serviceName, apiid, issueID, parameters, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client APIIssueClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serviceName string, apiid string, issueID string, parameters IssueContract, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"apiId":             autorest.Encode("path", apiid),
		"issueId":           autorest.Encode("path", issueID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client APIIssueClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client APIIssueClient) CreateOrUpdateResponder(resp *http.Response) (result IssueContract, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified Issue from an API.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// apiid - API identifier. Must be unique in the current API Management service instance.
// issueID - issue identifier. Must be unique in the current API Management service instance.
// ifMatch - eTag of the Entity. ETag should match the current entity state from the header response of the GET
// request or it should be * for unconditional update.
func (client APIIssueClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, apiid string, issueID string, ifMatch string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/APIIssueClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: apiid,
			Constraints: []validation.Constraint{{Target: "apiid", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "apiid", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: issueID,
			Constraints: []validation.Constraint{{Target: "issueID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "issueID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "issueID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.APIIssueClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, serviceName, apiid, issueID, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client APIIssueClient) DeletePreparer(ctx context.Context, resourceGroupName string, serviceName string, apiid string, issueID string, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"apiId":             autorest.Encode("path", apiid),
		"issueId":           autorest.Encode("path", issueID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}", pathParameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client APIIssueClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client APIIssueClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the details of the Issue for an API specified by its identifier.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// apiid - API identifier. Must be unique in the current API Management service instance.
// issueID - issue identifier. Must be unique in the current API Management service instance.
// expandCommentsAttachments - expand the comment attachments.
func (client APIIssueClient) Get(ctx context.Context, resourceGroupName string, serviceName string, apiid string, issueID string, expandCommentsAttachments *bool) (result IssueContract, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/APIIssueClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: apiid,
			Constraints: []validation.Constraint{{Target: "apiid", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "apiid", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: issueID,
			Constraints: []validation.Constraint{{Target: "issueID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "issueID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "issueID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.APIIssueClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, serviceName, apiid, issueID, expandCommentsAttachments)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client APIIssueClient) GetPreparer(ctx context.Context, resourceGroupName string, serviceName string, apiid string, issueID string, expandCommentsAttachments *bool) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"apiId":             autorest.Encode("path", apiid),
		"issueId":           autorest.Encode("path", issueID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if expandCommentsAttachments != nil {
		queryParameters["expandCommentsAttachments"] = autorest.Encode("query", *expandCommentsAttachments)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client APIIssueClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client APIIssueClient) GetResponder(resp *http.Response) (result IssueContract, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetEntityTag gets the entity state (Etag) version of the Issue for an API specified by its identifier.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// apiid - API identifier. Must be unique in the current API Management service instance.
// issueID - issue identifier. Must be unique in the current API Management service instance.
func (client APIIssueClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, apiid string, issueID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/APIIssueClient.GetEntityTag")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: apiid,
			Constraints: []validation.Constraint{{Target: "apiid", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "apiid", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: issueID,
			Constraints: []validation.Constraint{{Target: "issueID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "issueID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "issueID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.APIIssueClient", "GetEntityTag", err.Error())
	}

	req, err := client.GetEntityTagPreparer(ctx, resourceGroupName, serviceName, apiid, issueID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueClient", "GetEntityTag", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetEntityTagSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueClient", "GetEntityTag", resp, "Failure sending request")
		return
	}

	result, err = client.GetEntityTagResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueClient", "GetEntityTag", resp, "Failure responding to request")
		return
	}

	return
}

// GetEntityTagPreparer prepares the GetEntityTag request.
func (client APIIssueClient) GetEntityTagPreparer(ctx context.Context, resourceGroupName string, serviceName string, apiid string, issueID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"apiId":             autorest.Encode("path", apiid),
		"issueId":           autorest.Encode("path", issueID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsHead(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetEntityTagSender sends the GetEntityTag request. The method will close the
// http.Response Body if it receives an error.
func (client APIIssueClient) GetEntityTagSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetEntityTagResponder handles the response to the GetEntityTag request. The method always
// closes the http.Response Body.
func (client APIIssueClient) GetEntityTagResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// ListByService lists all issues associated with the specified API.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// apiid - API identifier. Must be unique in the current API Management service instance.
// filter - |     Field     |     Usage     |     Supported operators     |     Supported functions
// |</br>|-------------|-------------|-------------|-------------|</br>| name | filter | ge, le, eq, ne, gt, lt
// | substringof, contains, startswith, endswith |</br>| userId | filter | ge, le, eq, ne, gt, lt |
// substringof, contains, startswith, endswith |</br>| state | filter | eq |     |</br>
// expandCommentsAttachments - expand the comment attachments.
// top - number of records to return.
// skip - number of records to skip.
func (client APIIssueClient) ListByService(ctx context.Context, resourceGroupName string, serviceName string, apiid string, filter string, expandCommentsAttachments *bool, top *int32, skip *int32) (result IssueCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/APIIssueClient.ListByService")
		defer func() {
			sc := -1
			if result.ic.Response.Response != nil {
				sc = result.ic.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: apiid,
			Constraints: []validation.Constraint{{Target: "apiid", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "apiid", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil}}}}},
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("apimanagement.APIIssueClient", "ListByService", err.Error())
	}

	result.fn = client.listByServiceNextResults
	req, err := client.ListByServicePreparer(ctx, resourceGroupName, serviceName, apiid, filter, expandCommentsAttachments, top, skip)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueClient", "ListByService", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByServiceSender(req)
	if err != nil {
		result.ic.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueClient", "ListByService", resp, "Failure sending request")
		return
	}

	result.ic, err = client.ListByServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueClient", "ListByService", resp, "Failure responding to request")
		return
	}
	if result.ic.hasNextLink() && result.ic.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByServicePreparer prepares the ListByService request.
func (client APIIssueClient) ListByServicePreparer(ctx context.Context, resourceGroupName string, serviceName string, apiid string, filter string, expandCommentsAttachments *bool, top *int32, skip *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"apiId":             autorest.Encode("path", apiid),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if expandCommentsAttachments != nil {
		queryParameters["expandCommentsAttachments"] = autorest.Encode("query", *expandCommentsAttachments)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByServiceSender sends the ListByService request. The method will close the
// http.Response Body if it receives an error.
func (client APIIssueClient) ListByServiceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByServiceResponder handles the response to the ListByService request. The method always
// closes the http.Response Body.
func (client APIIssueClient) ListByServiceResponder(resp *http.Response) (result IssueCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByServiceNextResults retrieves the next set of results, if any.
func (client APIIssueClient) listByServiceNextResults(ctx context.Context, lastResults IssueCollection) (result IssueCollection, err error) {
	req, err := lastResults.issueCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "apimanagement.APIIssueClient", "listByServiceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByServiceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "apimanagement.APIIssueClient", "listByServiceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueClient", "listByServiceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByServiceComplete enumerates all values, automatically crossing page boundaries as required.
func (client APIIssueClient) ListByServiceComplete(ctx context.Context, resourceGroupName string, serviceName string, apiid string, filter string, expandCommentsAttachments *bool, top *int32, skip *int32) (result IssueCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/APIIssueClient.ListByService")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByService(ctx, resourceGroupName, serviceName, apiid, filter, expandCommentsAttachments, top, skip)
	return
}

// Update updates an existing issue for an API.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// apiid - API identifier. Must be unique in the current API Management service instance.
// issueID - issue identifier. Must be unique in the current API Management service instance.
// parameters - update parameters.
// ifMatch - eTag of the Entity. ETag should match the current entity state from the header response of the GET
// request or it should be * for unconditional update.
func (client APIIssueClient) Update(ctx context.Context, resourceGroupName string, serviceName string, apiid string, issueID string, parameters IssueUpdateContract, ifMatch string) (result IssueContract, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/APIIssueClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: apiid,
			Constraints: []validation.Constraint{{Target: "apiid", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "apiid", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: issueID,
			Constraints: []validation.Constraint{{Target: "issueID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "issueID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "issueID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.APIIssueClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, serviceName, apiid, issueID, parameters, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client APIIssueClient) UpdatePreparer(ctx context.Context, resourceGroupName string, serviceName string, apiid string, issueID string, parameters IssueUpdateContract, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"apiId":             autorest.Encode("path", apiid),
		"issueId":           autorest.Encode("path", issueID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client APIIssueClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client APIIssueClient) UpdateResponder(resp *http.Response) (result IssueContract, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
