package redisenterprise

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

// DatabasesClient is the REST API for managing Redis Enterprise resources in Azure.
type DatabasesClient struct {
	BaseClient
}

// NewDatabasesClient creates an instance of the DatabasesClient client.
func NewDatabasesClient(subscriptionID string) DatabasesClient {
	return NewDatabasesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDatabasesClientWithBaseURI creates an instance of the DatabasesClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewDatabasesClientWithBaseURI(baseURI string, subscriptionID string) DatabasesClient {
	return DatabasesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates a database
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterName - the name of the RedisEnterprise cluster.
// databaseName - the name of the database.
// parameters - parameters supplied to the create or update database operation.
func (client DatabasesClient) Create(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters Database) (result DatabasesCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DatabasesClient.Create")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, resourceGroupName, clusterName, databaseName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redisenterprise.DatabasesClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redisenterprise.DatabasesClient", "Create", nil, "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client DatabasesClient) CreatePreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters Database) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redisEnterprise/{clusterName}/databases/{databaseName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client DatabasesClient) CreateSender(req *http.Request) (future DatabasesCreateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client DatabasesClient) (d Database, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "redisenterprise.DatabasesCreateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("redisenterprise.DatabasesCreateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		d.Response.Response, err = future.GetResult(sender)
		if d.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "redisenterprise.DatabasesCreateFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && d.Response.Response.StatusCode != http.StatusNoContent {
			d, err = client.CreateResponder(d.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "redisenterprise.DatabasesCreateFuture", "Result", d.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client DatabasesClient) CreateResponder(resp *http.Response) (result Database, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a single database
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterName - the name of the RedisEnterprise cluster.
// databaseName - the name of the database.
func (client DatabasesClient) Delete(ctx context.Context, resourceGroupName string, clusterName string, databaseName string) (result DatabasesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DatabasesClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, clusterName, databaseName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redisenterprise.DatabasesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redisenterprise.DatabasesClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client DatabasesClient) DeletePreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redisEnterprise/{clusterName}/databases/{databaseName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client DatabasesClient) DeleteSender(req *http.Request) (future DatabasesDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client DatabasesClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "redisenterprise.DatabasesDeleteFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("redisenterprise.DatabasesDeleteFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client DatabasesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Export exports a database file from target database.
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterName - the name of the RedisEnterprise cluster.
// databaseName - the name of the database.
// parameters - storage information for exporting into the cluster
func (client DatabasesClient) Export(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters ExportClusterParameters) (result DatabasesExportFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DatabasesClient.Export")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.SasURI", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("redisenterprise.DatabasesClient", "Export", err.Error())
	}

	req, err := client.ExportPreparer(ctx, resourceGroupName, clusterName, databaseName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redisenterprise.DatabasesClient", "Export", nil, "Failure preparing request")
		return
	}

	result, err = client.ExportSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redisenterprise.DatabasesClient", "Export", nil, "Failure sending request")
		return
	}

	return
}

// ExportPreparer prepares the Export request.
func (client DatabasesClient) ExportPreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters ExportClusterParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redisEnterprise/{clusterName}/databases/{databaseName}/export", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ExportSender sends the Export request. The method will close the
// http.Response Body if it receives an error.
func (client DatabasesClient) ExportSender(req *http.Request) (future DatabasesExportFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client DatabasesClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "redisenterprise.DatabasesExportFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("redisenterprise.DatabasesExportFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// ExportResponder handles the response to the Export request. The method always
// closes the http.Response Body.
func (client DatabasesClient) ExportResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetMethod gets information about a database in a RedisEnterprise cluster.
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterName - the name of the RedisEnterprise cluster.
// databaseName - the name of the database.
func (client DatabasesClient) GetMethod(ctx context.Context, resourceGroupName string, clusterName string, databaseName string) (result Database, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DatabasesClient.GetMethod")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetMethodPreparer(ctx, resourceGroupName, clusterName, databaseName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redisenterprise.DatabasesClient", "GetMethod", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetMethodSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "redisenterprise.DatabasesClient", "GetMethod", resp, "Failure sending request")
		return
	}

	result, err = client.GetMethodResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redisenterprise.DatabasesClient", "GetMethod", resp, "Failure responding to request")
		return
	}

	return
}

// GetMethodPreparer prepares the GetMethod request.
func (client DatabasesClient) GetMethodPreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redisEnterprise/{clusterName}/databases/{databaseName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetMethodSender sends the GetMethod request. The method will close the
// http.Response Body if it receives an error.
func (client DatabasesClient) GetMethodSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetMethodResponder handles the response to the GetMethod request. The method always
// closes the http.Response Body.
func (client DatabasesClient) GetMethodResponder(resp *http.Response) (result Database, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Import imports a database file to target database.
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterName - the name of the RedisEnterprise cluster.
// databaseName - the name of the database.
// parameters - storage information for importing into the cluster
func (client DatabasesClient) Import(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters ImportClusterParameters) (result DatabasesImportFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DatabasesClient.Import")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.SasURI", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("redisenterprise.DatabasesClient", "Import", err.Error())
	}

	req, err := client.ImportPreparer(ctx, resourceGroupName, clusterName, databaseName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redisenterprise.DatabasesClient", "Import", nil, "Failure preparing request")
		return
	}

	result, err = client.ImportSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redisenterprise.DatabasesClient", "Import", nil, "Failure sending request")
		return
	}

	return
}

// ImportPreparer prepares the Import request.
func (client DatabasesClient) ImportPreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters ImportClusterParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redisEnterprise/{clusterName}/databases/{databaseName}/import", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ImportSender sends the Import request. The method will close the
// http.Response Body if it receives an error.
func (client DatabasesClient) ImportSender(req *http.Request) (future DatabasesImportFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client DatabasesClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "redisenterprise.DatabasesImportFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("redisenterprise.DatabasesImportFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// ImportResponder handles the response to the Import request. The method always
// closes the http.Response Body.
func (client DatabasesClient) ImportResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// ListByCluster gets all databases in the specified RedisEnterprise cluster.
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterName - the name of the RedisEnterprise cluster.
func (client DatabasesClient) ListByCluster(ctx context.Context, resourceGroupName string, clusterName string) (result DatabaseListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DatabasesClient.ListByCluster")
		defer func() {
			sc := -1
			if result.dl.Response.Response != nil {
				sc = result.dl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByClusterNextResults
	req, err := client.ListByClusterPreparer(ctx, resourceGroupName, clusterName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redisenterprise.DatabasesClient", "ListByCluster", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByClusterSender(req)
	if err != nil {
		result.dl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "redisenterprise.DatabasesClient", "ListByCluster", resp, "Failure sending request")
		return
	}

	result.dl, err = client.ListByClusterResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redisenterprise.DatabasesClient", "ListByCluster", resp, "Failure responding to request")
		return
	}
	if result.dl.hasNextLink() && result.dl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByClusterPreparer prepares the ListByCluster request.
func (client DatabasesClient) ListByClusterPreparer(ctx context.Context, resourceGroupName string, clusterName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redisEnterprise/{clusterName}/databases", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByClusterSender sends the ListByCluster request. The method will close the
// http.Response Body if it receives an error.
func (client DatabasesClient) ListByClusterSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByClusterResponder handles the response to the ListByCluster request. The method always
// closes the http.Response Body.
func (client DatabasesClient) ListByClusterResponder(resp *http.Response) (result DatabaseList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByClusterNextResults retrieves the next set of results, if any.
func (client DatabasesClient) listByClusterNextResults(ctx context.Context, lastResults DatabaseList) (result DatabaseList, err error) {
	req, err := lastResults.databaseListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "redisenterprise.DatabasesClient", "listByClusterNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByClusterSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "redisenterprise.DatabasesClient", "listByClusterNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByClusterResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redisenterprise.DatabasesClient", "listByClusterNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByClusterComplete enumerates all values, automatically crossing page boundaries as required.
func (client DatabasesClient) ListByClusterComplete(ctx context.Context, resourceGroupName string, clusterName string) (result DatabaseListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DatabasesClient.ListByCluster")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByCluster(ctx, resourceGroupName, clusterName)
	return
}

// ListKeys retrieves the access keys for the RedisEnterprise database.
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterName - the name of the RedisEnterprise cluster.
// databaseName - the name of the database.
func (client DatabasesClient) ListKeys(ctx context.Context, resourceGroupName string, clusterName string, databaseName string) (result AccessKeys, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DatabasesClient.ListKeys")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListKeysPreparer(ctx, resourceGroupName, clusterName, databaseName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redisenterprise.DatabasesClient", "ListKeys", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListKeysSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "redisenterprise.DatabasesClient", "ListKeys", resp, "Failure sending request")
		return
	}

	result, err = client.ListKeysResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redisenterprise.DatabasesClient", "ListKeys", resp, "Failure responding to request")
		return
	}

	return
}

// ListKeysPreparer prepares the ListKeys request.
func (client DatabasesClient) ListKeysPreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redisEnterprise/{clusterName}/databases/{databaseName}/listKeys", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListKeysSender sends the ListKeys request. The method will close the
// http.Response Body if it receives an error.
func (client DatabasesClient) ListKeysSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListKeysResponder handles the response to the ListKeys request. The method always
// closes the http.Response Body.
func (client DatabasesClient) ListKeysResponder(resp *http.Response) (result AccessKeys, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// RegenerateKey regenerates the RedisEnterprise database's access keys.
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterName - the name of the RedisEnterprise cluster.
// databaseName - the name of the database.
// parameters - specifies which key to regenerate.
func (client DatabasesClient) RegenerateKey(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters RegenerateKeyParameters) (result DatabasesRegenerateKeyFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DatabasesClient.RegenerateKey")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RegenerateKeyPreparer(ctx, resourceGroupName, clusterName, databaseName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redisenterprise.DatabasesClient", "RegenerateKey", nil, "Failure preparing request")
		return
	}

	result, err = client.RegenerateKeySender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redisenterprise.DatabasesClient", "RegenerateKey", nil, "Failure sending request")
		return
	}

	return
}

// RegenerateKeyPreparer prepares the RegenerateKey request.
func (client DatabasesClient) RegenerateKeyPreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters RegenerateKeyParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redisEnterprise/{clusterName}/databases/{databaseName}/regenerateKey", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RegenerateKeySender sends the RegenerateKey request. The method will close the
// http.Response Body if it receives an error.
func (client DatabasesClient) RegenerateKeySender(req *http.Request) (future DatabasesRegenerateKeyFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client DatabasesClient) (ak AccessKeys, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "redisenterprise.DatabasesRegenerateKeyFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("redisenterprise.DatabasesRegenerateKeyFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		ak.Response.Response, err = future.GetResult(sender)
		if ak.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "redisenterprise.DatabasesRegenerateKeyFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && ak.Response.Response.StatusCode != http.StatusNoContent {
			ak, err = client.RegenerateKeyResponder(ak.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "redisenterprise.DatabasesRegenerateKeyFuture", "Result", ak.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// RegenerateKeyResponder handles the response to the RegenerateKey request. The method always
// closes the http.Response Body.
func (client DatabasesClient) RegenerateKeyResponder(resp *http.Response) (result AccessKeys, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates a database
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterName - the name of the RedisEnterprise cluster.
// databaseName - the name of the database.
// parameters - parameters supplied to the create or update database operation.
func (client DatabasesClient) Update(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters DatabaseUpdate) (result DatabasesUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DatabasesClient.Update")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, clusterName, databaseName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redisenterprise.DatabasesClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redisenterprise.DatabasesClient", "Update", nil, "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client DatabasesClient) UpdatePreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters DatabaseUpdate) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redisEnterprise/{clusterName}/databases/{databaseName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client DatabasesClient) UpdateSender(req *http.Request) (future DatabasesUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client DatabasesClient) (d Database, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "redisenterprise.DatabasesUpdateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("redisenterprise.DatabasesUpdateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		d.Response.Response, err = future.GetResult(sender)
		if d.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "redisenterprise.DatabasesUpdateFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && d.Response.Response.StatusCode != http.StatusNoContent {
			d, err = client.UpdateResponder(d.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "redisenterprise.DatabasesUpdateFuture", "Result", d.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client DatabasesClient) UpdateResponder(resp *http.Response) (result Database, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
