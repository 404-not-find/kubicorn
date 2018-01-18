package policy

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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// SetDefinitionsClient is the to manage and control access to your resources, you can define customized policies and
// assign them at a scope.
type SetDefinitionsClient struct {
	ManagementClient
}

// NewSetDefinitionsClient creates an instance of the SetDefinitionsClient client.
func NewSetDefinitionsClient(subscriptionID string) SetDefinitionsClient {
	return NewSetDefinitionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSetDefinitionsClientWithBaseURI creates an instance of the SetDefinitionsClient client.
func NewSetDefinitionsClientWithBaseURI(baseURI string, subscriptionID string) SetDefinitionsClient {
	return SetDefinitionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a policy set definition.
//
// policySetDefinitionName is the name of the policy set definition to create. parameters is the policy set definition
// properties.
func (client SetDefinitionsClient) CreateOrUpdate(policySetDefinitionName string, parameters SetDefinition) (result SetDefinition, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.SetDefinitionProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.SetDefinitionProperties.PolicyDefinitions", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "policy.SetDefinitionsClient", "CreateOrUpdate")
	}

	req, err := client.CreateOrUpdatePreparer(policySetDefinitionName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SetDefinitionsClient) CreateOrUpdatePreparer(policySetDefinitionName string, parameters SetDefinition) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policySetDefinitionName": autorest.Encode("path", policySetDefinitionName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policySetDefinitions/{policySetDefinitionName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client SetDefinitionsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client SetDefinitionsClient) CreateOrUpdateResponder(resp *http.Response) (result SetDefinition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateOrUpdateAtManagementGroup creates or updates a policy set definition at management group level.
//
// policySetDefinitionName is the name of the policy set definition to create. parameters is the policy set definition
// properties. managementGroupID is the ID of the management group.
func (client SetDefinitionsClient) CreateOrUpdateAtManagementGroup(policySetDefinitionName string, parameters SetDefinition, managementGroupID string) (result SetDefinition, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.SetDefinitionProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.SetDefinitionProperties.PolicyDefinitions", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "policy.SetDefinitionsClient", "CreateOrUpdateAtManagementGroup")
	}

	req, err := client.CreateOrUpdateAtManagementGroupPreparer(policySetDefinitionName, parameters, managementGroupID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "CreateOrUpdateAtManagementGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateAtManagementGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "CreateOrUpdateAtManagementGroup", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateAtManagementGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "CreateOrUpdateAtManagementGroup", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdateAtManagementGroupPreparer prepares the CreateOrUpdateAtManagementGroup request.
func (client SetDefinitionsClient) CreateOrUpdateAtManagementGroupPreparer(policySetDefinitionName string, parameters SetDefinition, managementGroupID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managementGroupId":       autorest.Encode("path", managementGroupID),
		"policySetDefinitionName": autorest.Encode("path", policySetDefinitionName),
	}

	const APIVersion = "2017-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policySetDefinitions/{policySetDefinitionName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateAtManagementGroupSender sends the CreateOrUpdateAtManagementGroup request. The method will close the
// http.Response Body if it receives an error.
func (client SetDefinitionsClient) CreateOrUpdateAtManagementGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateOrUpdateAtManagementGroupResponder handles the response to the CreateOrUpdateAtManagementGroup request. The method always
// closes the http.Response Body.
func (client SetDefinitionsClient) CreateOrUpdateAtManagementGroupResponder(resp *http.Response) (result SetDefinition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a policy set definition.
//
// policySetDefinitionName is the name of the policy set definition to delete.
func (client SetDefinitionsClient) Delete(policySetDefinitionName string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(policySetDefinitionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SetDefinitionsClient) DeletePreparer(policySetDefinitionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policySetDefinitionName": autorest.Encode("path", policySetDefinitionName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policySetDefinitions/{policySetDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SetDefinitionsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SetDefinitionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DeleteAtManagementGroup deletes a policy set definition at management group level.
//
// policySetDefinitionName is the name of the policy set definition to delete. managementGroupID is the ID of the
// management group.
func (client SetDefinitionsClient) DeleteAtManagementGroup(policySetDefinitionName string, managementGroupID string) (result autorest.Response, err error) {
	req, err := client.DeleteAtManagementGroupPreparer(policySetDefinitionName, managementGroupID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "DeleteAtManagementGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteAtManagementGroupSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "DeleteAtManagementGroup", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteAtManagementGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "DeleteAtManagementGroup", resp, "Failure responding to request")
	}

	return
}

// DeleteAtManagementGroupPreparer prepares the DeleteAtManagementGroup request.
func (client SetDefinitionsClient) DeleteAtManagementGroupPreparer(policySetDefinitionName string, managementGroupID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managementGroupId":       autorest.Encode("path", managementGroupID),
		"policySetDefinitionName": autorest.Encode("path", policySetDefinitionName),
	}

	const APIVersion = "2017-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policySetDefinitions/{policySetDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteAtManagementGroupSender sends the DeleteAtManagementGroup request. The method will close the
// http.Response Body if it receives an error.
func (client SetDefinitionsClient) DeleteAtManagementGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteAtManagementGroupResponder handles the response to the DeleteAtManagementGroup request. The method always
// closes the http.Response Body.
func (client SetDefinitionsClient) DeleteAtManagementGroupResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the policy set definition.
//
// policySetDefinitionName is the name of the policy set definition to get.
func (client SetDefinitionsClient) Get(policySetDefinitionName string) (result SetDefinition, err error) {
	req, err := client.GetPreparer(policySetDefinitionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client SetDefinitionsClient) GetPreparer(policySetDefinitionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policySetDefinitionName": autorest.Encode("path", policySetDefinitionName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policySetDefinitions/{policySetDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SetDefinitionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SetDefinitionsClient) GetResponder(resp *http.Response) (result SetDefinition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetAtManagementGroup gets the policy set definition at management group level.
//
// policySetDefinitionName is the name of the policy set definition to get. managementGroupID is the ID of the
// management group.
func (client SetDefinitionsClient) GetAtManagementGroup(policySetDefinitionName string, managementGroupID string) (result SetDefinition, err error) {
	req, err := client.GetAtManagementGroupPreparer(policySetDefinitionName, managementGroupID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "GetAtManagementGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetAtManagementGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "GetAtManagementGroup", resp, "Failure sending request")
		return
	}

	result, err = client.GetAtManagementGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "GetAtManagementGroup", resp, "Failure responding to request")
	}

	return
}

// GetAtManagementGroupPreparer prepares the GetAtManagementGroup request.
func (client SetDefinitionsClient) GetAtManagementGroupPreparer(policySetDefinitionName string, managementGroupID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managementGroupId":       autorest.Encode("path", managementGroupID),
		"policySetDefinitionName": autorest.Encode("path", policySetDefinitionName),
	}

	const APIVersion = "2017-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policySetDefinitions/{policySetDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetAtManagementGroupSender sends the GetAtManagementGroup request. The method will close the
// http.Response Body if it receives an error.
func (client SetDefinitionsClient) GetAtManagementGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetAtManagementGroupResponder handles the response to the GetAtManagementGroup request. The method always
// closes the http.Response Body.
func (client SetDefinitionsClient) GetAtManagementGroupResponder(resp *http.Response) (result SetDefinition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetBuiltIn gets the built in policy set definition.
//
// policySetDefinitionName is the name of the policy set definition to get.
func (client SetDefinitionsClient) GetBuiltIn(policySetDefinitionName string) (result SetDefinition, err error) {
	req, err := client.GetBuiltInPreparer(policySetDefinitionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "GetBuiltIn", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetBuiltInSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "GetBuiltIn", resp, "Failure sending request")
		return
	}

	result, err = client.GetBuiltInResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "GetBuiltIn", resp, "Failure responding to request")
	}

	return
}

// GetBuiltInPreparer prepares the GetBuiltIn request.
func (client SetDefinitionsClient) GetBuiltInPreparer(policySetDefinitionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policySetDefinitionName": autorest.Encode("path", policySetDefinitionName),
	}

	const APIVersion = "2017-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Authorization/policySetDefinitions/{policySetDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetBuiltInSender sends the GetBuiltIn request. The method will close the
// http.Response Body if it receives an error.
func (client SetDefinitionsClient) GetBuiltInSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetBuiltInResponder handles the response to the GetBuiltIn request. The method always
// closes the http.Response Body.
func (client SetDefinitionsClient) GetBuiltInResponder(resp *http.Response) (result SetDefinition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all the policy set definitions for a subscription.
func (client SetDefinitionsClient) List() (result SetDefinitionListResult, err error) {
	req, err := client.ListPreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client SetDefinitionsClient) ListPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policySetDefinitions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client SetDefinitionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client SetDefinitionsClient) ListResponder(resp *http.Response) (result SetDefinitionListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client SetDefinitionsClient) ListNextResults(lastResults SetDefinitionListResult) (result SetDefinitionListResult, err error) {
	req, err := lastResults.SetDefinitionListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// ListComplete gets all elements from the list without paging.
func (client SetDefinitionsClient) ListComplete(cancel <-chan struct{}) (<-chan SetDefinition, <-chan error) {
	resultChan := make(chan SetDefinition)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.List()
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}

// ListBuiltIn gets all the built in policy set definitions.
func (client SetDefinitionsClient) ListBuiltIn() (result SetDefinitionListResult, err error) {
	req, err := client.ListBuiltInPreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "ListBuiltIn", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBuiltInSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "ListBuiltIn", resp, "Failure sending request")
		return
	}

	result, err = client.ListBuiltInResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "ListBuiltIn", resp, "Failure responding to request")
	}

	return
}

// ListBuiltInPreparer prepares the ListBuiltIn request.
func (client SetDefinitionsClient) ListBuiltInPreparer() (*http.Request, error) {
	const APIVersion = "2017-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Authorization/policySetDefinitions"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListBuiltInSender sends the ListBuiltIn request. The method will close the
// http.Response Body if it receives an error.
func (client SetDefinitionsClient) ListBuiltInSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListBuiltInResponder handles the response to the ListBuiltIn request. The method always
// closes the http.Response Body.
func (client SetDefinitionsClient) ListBuiltInResponder(resp *http.Response) (result SetDefinitionListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListBuiltInNextResults retrieves the next set of results, if any.
func (client SetDefinitionsClient) ListBuiltInNextResults(lastResults SetDefinitionListResult) (result SetDefinitionListResult, err error) {
	req, err := lastResults.SetDefinitionListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "ListBuiltIn", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListBuiltInSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "ListBuiltIn", resp, "Failure sending next results request")
	}

	result, err = client.ListBuiltInResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "ListBuiltIn", resp, "Failure responding to next results request")
	}

	return
}

// ListBuiltInComplete gets all elements from the list without paging.
func (client SetDefinitionsClient) ListBuiltInComplete(cancel <-chan struct{}) (<-chan SetDefinition, <-chan error) {
	resultChan := make(chan SetDefinition)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListBuiltIn()
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListBuiltInNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}

// ListByManagementGroup gets all the policy set definitions for a subscription at management group.
//
// managementGroupID is the ID of the management group.
func (client SetDefinitionsClient) ListByManagementGroup(managementGroupID string) (result SetDefinitionListResult, err error) {
	req, err := client.ListByManagementGroupPreparer(managementGroupID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "ListByManagementGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByManagementGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "ListByManagementGroup", resp, "Failure sending request")
		return
	}

	result, err = client.ListByManagementGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "ListByManagementGroup", resp, "Failure responding to request")
	}

	return
}

// ListByManagementGroupPreparer prepares the ListByManagementGroup request.
func (client SetDefinitionsClient) ListByManagementGroupPreparer(managementGroupID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managementGroupId": autorest.Encode("path", managementGroupID),
	}

	const APIVersion = "2017-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policySetDefinitions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByManagementGroupSender sends the ListByManagementGroup request. The method will close the
// http.Response Body if it receives an error.
func (client SetDefinitionsClient) ListByManagementGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByManagementGroupResponder handles the response to the ListByManagementGroup request. The method always
// closes the http.Response Body.
func (client SetDefinitionsClient) ListByManagementGroupResponder(resp *http.Response) (result SetDefinitionListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByManagementGroupNextResults retrieves the next set of results, if any.
func (client SetDefinitionsClient) ListByManagementGroupNextResults(lastResults SetDefinitionListResult) (result SetDefinitionListResult, err error) {
	req, err := lastResults.SetDefinitionListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "ListByManagementGroup", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByManagementGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "ListByManagementGroup", resp, "Failure sending next results request")
	}

	result, err = client.ListByManagementGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "ListByManagementGroup", resp, "Failure responding to next results request")
	}

	return
}

// ListByManagementGroupComplete gets all elements from the list without paging.
func (client SetDefinitionsClient) ListByManagementGroupComplete(managementGroupID string, cancel <-chan struct{}) (<-chan SetDefinition, <-chan error) {
	resultChan := make(chan SetDefinition)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListByManagementGroup(managementGroupID)
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListByManagementGroupNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}