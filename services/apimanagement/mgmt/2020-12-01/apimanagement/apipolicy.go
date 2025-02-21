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

// APIPolicyClient is the apiManagement Client
type APIPolicyClient struct {
	BaseClient
}

// NewAPIPolicyClient creates an instance of the APIPolicyClient client.
func NewAPIPolicyClient(subscriptionID string) APIPolicyClient {
	return NewAPIPolicyClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAPIPolicyClientWithBaseURI creates an instance of the APIPolicyClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewAPIPolicyClientWithBaseURI(baseURI string, subscriptionID string) APIPolicyClient {
	return APIPolicyClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates policy configuration for the API.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// apiid - API revision identifier. Must be unique in the current API Management service instance. Non-current
// revision has ;rev=n as a suffix where n is the revision number.
// parameters - the policy contents to apply.
// ifMatch - eTag of the Entity. Not required when creating an entity, but required when updating an entity.
func (client APIPolicyClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, apiid string, parameters PolicyContract, ifMatch string) (result PolicyContract, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/APIPolicyClient.CreateOrUpdate")
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
			Constraints: []validation.Constraint{{Target: "apiid", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "apiid", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "apiid", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.PolicyContractProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.PolicyContractProperties.Value", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("apimanagement.APIPolicyClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serviceName, apiid, parameters, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIPolicyClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.APIPolicyClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIPolicyClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client APIPolicyClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serviceName string, apiid string, parameters PolicyContract, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"apiId":             autorest.Encode("path", apiid),
		"policyId":          autorest.Encode("path", "policy"),
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/policies/{policyId}", pathParameters),
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
func (client APIPolicyClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client APIPolicyClient) CreateOrUpdateResponder(resp *http.Response) (result PolicyContract, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the policy configuration at the Api.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// apiid - API revision identifier. Must be unique in the current API Management service instance. Non-current
// revision has ;rev=n as a suffix where n is the revision number.
// ifMatch - eTag of the Entity. ETag should match the current entity state from the header response of the GET
// request or it should be * for unconditional update.
func (client APIPolicyClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, apiid string, ifMatch string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/APIPolicyClient.Delete")
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
			Constraints: []validation.Constraint{{Target: "apiid", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "apiid", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "apiid", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.APIPolicyClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, serviceName, apiid, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIPolicyClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.APIPolicyClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIPolicyClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client APIPolicyClient) DeletePreparer(ctx context.Context, resourceGroupName string, serviceName string, apiid string, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"apiId":             autorest.Encode("path", apiid),
		"policyId":          autorest.Encode("path", "policy"),
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/policies/{policyId}", pathParameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client APIPolicyClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client APIPolicyClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get the policy configuration at the API level.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// apiid - API revision identifier. Must be unique in the current API Management service instance. Non-current
// revision has ;rev=n as a suffix where n is the revision number.
// formatParameter - policy Export Format.
func (client APIPolicyClient) Get(ctx context.Context, resourceGroupName string, serviceName string, apiid string, formatParameter PolicyExportFormat) (result PolicyContract, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/APIPolicyClient.Get")
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
			Constraints: []validation.Constraint{{Target: "apiid", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "apiid", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "apiid", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.APIPolicyClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, serviceName, apiid, formatParameter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIPolicyClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.APIPolicyClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIPolicyClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client APIPolicyClient) GetPreparer(ctx context.Context, resourceGroupName string, serviceName string, apiid string, formatParameter PolicyExportFormat) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"apiId":             autorest.Encode("path", apiid),
		"policyId":          autorest.Encode("path", "policy"),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(string(formatParameter)) > 0 {
		queryParameters["format"] = autorest.Encode("query", formatParameter)
	} else {
		queryParameters["format"] = autorest.Encode("query", "xml")
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/policies/{policyId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client APIPolicyClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client APIPolicyClient) GetResponder(resp *http.Response) (result PolicyContract, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetEntityTag gets the entity state (Etag) version of the API policy specified by its identifier.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// apiid - API revision identifier. Must be unique in the current API Management service instance. Non-current
// revision has ;rev=n as a suffix where n is the revision number.
func (client APIPolicyClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, apiid string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/APIPolicyClient.GetEntityTag")
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
			Constraints: []validation.Constraint{{Target: "apiid", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "apiid", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "apiid", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.APIPolicyClient", "GetEntityTag", err.Error())
	}

	req, err := client.GetEntityTagPreparer(ctx, resourceGroupName, serviceName, apiid)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIPolicyClient", "GetEntityTag", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetEntityTagSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.APIPolicyClient", "GetEntityTag", resp, "Failure sending request")
		return
	}

	result, err = client.GetEntityTagResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIPolicyClient", "GetEntityTag", resp, "Failure responding to request")
		return
	}

	return
}

// GetEntityTagPreparer prepares the GetEntityTag request.
func (client APIPolicyClient) GetEntityTagPreparer(ctx context.Context, resourceGroupName string, serviceName string, apiid string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"apiId":             autorest.Encode("path", apiid),
		"policyId":          autorest.Encode("path", "policy"),
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/policies/{policyId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetEntityTagSender sends the GetEntityTag request. The method will close the
// http.Response Body if it receives an error.
func (client APIPolicyClient) GetEntityTagSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetEntityTagResponder handles the response to the GetEntityTag request. The method always
// closes the http.Response Body.
func (client APIPolicyClient) GetEntityTagResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// ListByAPI get the policy configuration at the API level.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// apiid - API revision identifier. Must be unique in the current API Management service instance. Non-current
// revision has ;rev=n as a suffix where n is the revision number.
func (client APIPolicyClient) ListByAPI(ctx context.Context, resourceGroupName string, serviceName string, apiid string) (result PolicyCollection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/APIPolicyClient.ListByAPI")
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
			Constraints: []validation.Constraint{{Target: "apiid", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "apiid", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "apiid", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.APIPolicyClient", "ListByAPI", err.Error())
	}

	req, err := client.ListByAPIPreparer(ctx, resourceGroupName, serviceName, apiid)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIPolicyClient", "ListByAPI", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByAPISender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.APIPolicyClient", "ListByAPI", resp, "Failure sending request")
		return
	}

	result, err = client.ListByAPIResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIPolicyClient", "ListByAPI", resp, "Failure responding to request")
		return
	}

	return
}

// ListByAPIPreparer prepares the ListByAPI request.
func (client APIPolicyClient) ListByAPIPreparer(ctx context.Context, resourceGroupName string, serviceName string, apiid string) (*http.Request, error) {
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

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/policies", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByAPISender sends the ListByAPI request. The method will close the
// http.Response Body if it receives an error.
func (client APIPolicyClient) ListByAPISender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByAPIResponder handles the response to the ListByAPI request. The method always
// closes the http.Response Body.
func (client APIPolicyClient) ListByAPIResponder(resp *http.Response) (result PolicyCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
