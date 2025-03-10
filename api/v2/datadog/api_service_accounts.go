/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// ServiceAccountsApiService ServiceAccountsApi service
type ServiceAccountsApiService service

type apiCreateServiceAccountApplicationKeyRequest struct {
	ctx              _context.Context
	ApiService       *ServiceAccountsApiService
	serviceAccountId string
	body             *ApplicationKeyCreateRequest
}

func (a *ServiceAccountsApiService) buildCreateServiceAccountApplicationKeyRequest(ctx _context.Context, serviceAccountId string, body ApplicationKeyCreateRequest) (apiCreateServiceAccountApplicationKeyRequest, error) {
	req := apiCreateServiceAccountApplicationKeyRequest{
		ApiService:       a,
		ctx:              ctx,
		serviceAccountId: serviceAccountId,
		body:             &body,
	}
	return req, nil
}

/*
 * CreateServiceAccountApplicationKey Create an application key for this service account
 * Create an application key for this service account.
 */
func (a *ServiceAccountsApiService) CreateServiceAccountApplicationKey(ctx _context.Context, serviceAccountId string, body ApplicationKeyCreateRequest) (ApplicationKeyResponse, *_nethttp.Response, error) {
	req, err := a.buildCreateServiceAccountApplicationKeyRequest(ctx, serviceAccountId, body)
	if err != nil {
		var localVarReturnValue ApplicationKeyResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.createServiceAccountApplicationKeyExecute(req)
}

/*
 * Execute executes the request
 * @return ApplicationKeyResponse
 */
func (a *ServiceAccountsApiService) createServiceAccountApplicationKeyExecute(r apiCreateServiceAccountApplicationKeyRequest) (ApplicationKeyResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue ApplicationKeyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAccountsApiService.CreateServiceAccountApplicationKey")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/service_accounts/{service_account_id}/application_keys"
	localVarPath = strings.Replace(localVarPath, "{"+"service_account_id"+"}", _neturl.PathEscape(parameterToString(r.serviceAccountId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["DD-API-KEY"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["appKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["DD-APPLICATION-KEY"] = key
			}
		}
	}
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type apiDeleteServiceAccountApplicationKeyRequest struct {
	ctx              _context.Context
	ApiService       *ServiceAccountsApiService
	serviceAccountId string
	appKeyId         string
}

func (a *ServiceAccountsApiService) buildDeleteServiceAccountApplicationKeyRequest(ctx _context.Context, serviceAccountId string, appKeyId string) (apiDeleteServiceAccountApplicationKeyRequest, error) {
	req := apiDeleteServiceAccountApplicationKeyRequest{
		ApiService:       a,
		ctx:              ctx,
		serviceAccountId: serviceAccountId,
		appKeyId:         appKeyId,
	}
	return req, nil
}

/*
 * DeleteServiceAccountApplicationKey Delete an application key for this service account
 * Delete an application key owned by this service account.
 */
func (a *ServiceAccountsApiService) DeleteServiceAccountApplicationKey(ctx _context.Context, serviceAccountId string, appKeyId string) (*_nethttp.Response, error) {
	req, err := a.buildDeleteServiceAccountApplicationKeyRequest(ctx, serviceAccountId, appKeyId)
	if err != nil {
		return nil, err
	}

	return req.ApiService.deleteServiceAccountApplicationKeyExecute(req)
}

/*
 * Execute executes the request
 */
func (a *ServiceAccountsApiService) deleteServiceAccountApplicationKeyExecute(r apiDeleteServiceAccountApplicationKeyRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAccountsApiService.DeleteServiceAccountApplicationKey")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/service_accounts/{service_account_id}/application_keys/{app_key_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"service_account_id"+"}", _neturl.PathEscape(parameterToString(r.serviceAccountId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"app_key_id"+"}", _neturl.PathEscape(parameterToString(r.appKeyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["DD-API-KEY"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["appKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["DD-APPLICATION-KEY"] = key
			}
		}
	}
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type apiGetServiceAccountApplicationKeyRequest struct {
	ctx              _context.Context
	ApiService       *ServiceAccountsApiService
	serviceAccountId string
	appKeyId         string
}

func (a *ServiceAccountsApiService) buildGetServiceAccountApplicationKeyRequest(ctx _context.Context, serviceAccountId string, appKeyId string) (apiGetServiceAccountApplicationKeyRequest, error) {
	req := apiGetServiceAccountApplicationKeyRequest{
		ApiService:       a,
		ctx:              ctx,
		serviceAccountId: serviceAccountId,
		appKeyId:         appKeyId,
	}
	return req, nil
}

/*
 * GetServiceAccountApplicationKey Get one application key for this service account
 * Get an application key owned by this service account.
 */
func (a *ServiceAccountsApiService) GetServiceAccountApplicationKey(ctx _context.Context, serviceAccountId string, appKeyId string) (PartialApplicationKeyResponse, *_nethttp.Response, error) {
	req, err := a.buildGetServiceAccountApplicationKeyRequest(ctx, serviceAccountId, appKeyId)
	if err != nil {
		var localVarReturnValue PartialApplicationKeyResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.getServiceAccountApplicationKeyExecute(req)
}

/*
 * Execute executes the request
 * @return PartialApplicationKeyResponse
 */
func (a *ServiceAccountsApiService) getServiceAccountApplicationKeyExecute(r apiGetServiceAccountApplicationKeyRequest) (PartialApplicationKeyResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue PartialApplicationKeyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAccountsApiService.GetServiceAccountApplicationKey")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/service_accounts/{service_account_id}/application_keys/{app_key_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"service_account_id"+"}", _neturl.PathEscape(parameterToString(r.serviceAccountId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"app_key_id"+"}", _neturl.PathEscape(parameterToString(r.appKeyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["DD-API-KEY"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["appKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["DD-APPLICATION-KEY"] = key
			}
		}
	}
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type apiListServiceAccountApplicationKeysRequest struct {
	ctx                  _context.Context
	ApiService           *ServiceAccountsApiService
	serviceAccountId     string
	pageSize             *int64
	pageNumber           *int64
	sort                 *ApplicationKeysSort
	filter               *string
	filterCreatedAtStart *string
	filterCreatedAtEnd   *string
}

type ListServiceAccountApplicationKeysOptionalParameters struct {
	PageSize             *int64
	PageNumber           *int64
	Sort                 *ApplicationKeysSort
	Filter               *string
	FilterCreatedAtStart *string
	FilterCreatedAtEnd   *string
}

func NewListServiceAccountApplicationKeysOptionalParameters() *ListServiceAccountApplicationKeysOptionalParameters {
	this := ListServiceAccountApplicationKeysOptionalParameters{}
	return &this
}
func (r *ListServiceAccountApplicationKeysOptionalParameters) WithPageSize(pageSize int64) *ListServiceAccountApplicationKeysOptionalParameters {
	r.PageSize = &pageSize
	return r
}
func (r *ListServiceAccountApplicationKeysOptionalParameters) WithPageNumber(pageNumber int64) *ListServiceAccountApplicationKeysOptionalParameters {
	r.PageNumber = &pageNumber
	return r
}
func (r *ListServiceAccountApplicationKeysOptionalParameters) WithSort(sort ApplicationKeysSort) *ListServiceAccountApplicationKeysOptionalParameters {
	r.Sort = &sort
	return r
}
func (r *ListServiceAccountApplicationKeysOptionalParameters) WithFilter(filter string) *ListServiceAccountApplicationKeysOptionalParameters {
	r.Filter = &filter
	return r
}
func (r *ListServiceAccountApplicationKeysOptionalParameters) WithFilterCreatedAtStart(filterCreatedAtStart string) *ListServiceAccountApplicationKeysOptionalParameters {
	r.FilterCreatedAtStart = &filterCreatedAtStart
	return r
}
func (r *ListServiceAccountApplicationKeysOptionalParameters) WithFilterCreatedAtEnd(filterCreatedAtEnd string) *ListServiceAccountApplicationKeysOptionalParameters {
	r.FilterCreatedAtEnd = &filterCreatedAtEnd
	return r
}

func (a *ServiceAccountsApiService) buildListServiceAccountApplicationKeysRequest(ctx _context.Context, serviceAccountId string, o ...ListServiceAccountApplicationKeysOptionalParameters) (apiListServiceAccountApplicationKeysRequest, error) {
	req := apiListServiceAccountApplicationKeysRequest{
		ApiService:       a,
		ctx:              ctx,
		serviceAccountId: serviceAccountId,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type ListServiceAccountApplicationKeysOptionalParameters is allowed")
	}

	if o != nil {
		req.pageSize = o[0].PageSize
		req.pageNumber = o[0].PageNumber
		req.sort = o[0].Sort
		req.filter = o[0].Filter
		req.filterCreatedAtStart = o[0].FilterCreatedAtStart
		req.filterCreatedAtEnd = o[0].FilterCreatedAtEnd
	}
	return req, nil
}

/*
 * ListServiceAccountApplicationKeys List application keys for this service account
 * List all application keys available for this service account.
 */
func (a *ServiceAccountsApiService) ListServiceAccountApplicationKeys(ctx _context.Context, serviceAccountId string, o ...ListServiceAccountApplicationKeysOptionalParameters) (ListApplicationKeysResponse, *_nethttp.Response, error) {
	req, err := a.buildListServiceAccountApplicationKeysRequest(ctx, serviceAccountId, o...)
	if err != nil {
		var localVarReturnValue ListApplicationKeysResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.listServiceAccountApplicationKeysExecute(req)
}

/*
 * Execute executes the request
 * @return ListApplicationKeysResponse
 */
func (a *ServiceAccountsApiService) listServiceAccountApplicationKeysExecute(r apiListServiceAccountApplicationKeysRequest) (ListApplicationKeysResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ListApplicationKeysResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAccountsApiService.ListServiceAccountApplicationKeys")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/service_accounts/{service_account_id}/application_keys"
	localVarPath = strings.Replace(localVarPath, "{"+"service_account_id"+"}", _neturl.PathEscape(parameterToString(r.serviceAccountId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.pageSize != nil {
		localVarQueryParams.Add("page[size]", parameterToString(*r.pageSize, ""))
	}
	if r.pageNumber != nil {
		localVarQueryParams.Add("page[number]", parameterToString(*r.pageNumber, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
	if r.filter != nil {
		localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
	}
	if r.filterCreatedAtStart != nil {
		localVarQueryParams.Add("filter[created_at][start]", parameterToString(*r.filterCreatedAtStart, ""))
	}
	if r.filterCreatedAtEnd != nil {
		localVarQueryParams.Add("filter[created_at][end]", parameterToString(*r.filterCreatedAtEnd, ""))
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["DD-API-KEY"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["appKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["DD-APPLICATION-KEY"] = key
			}
		}
	}
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type apiUpdateServiceAccountApplicationKeyRequest struct {
	ctx              _context.Context
	ApiService       *ServiceAccountsApiService
	serviceAccountId string
	appKeyId         string
	body             *ApplicationKeyUpdateRequest
}

func (a *ServiceAccountsApiService) buildUpdateServiceAccountApplicationKeyRequest(ctx _context.Context, serviceAccountId string, appKeyId string, body ApplicationKeyUpdateRequest) (apiUpdateServiceAccountApplicationKeyRequest, error) {
	req := apiUpdateServiceAccountApplicationKeyRequest{
		ApiService:       a,
		ctx:              ctx,
		serviceAccountId: serviceAccountId,
		appKeyId:         appKeyId,
		body:             &body,
	}
	return req, nil
}

/*
 * UpdateServiceAccountApplicationKey Edit an application key for this service account
 * Edit an application key owned by this service account.
 */
func (a *ServiceAccountsApiService) UpdateServiceAccountApplicationKey(ctx _context.Context, serviceAccountId string, appKeyId string, body ApplicationKeyUpdateRequest) (PartialApplicationKeyResponse, *_nethttp.Response, error) {
	req, err := a.buildUpdateServiceAccountApplicationKeyRequest(ctx, serviceAccountId, appKeyId, body)
	if err != nil {
		var localVarReturnValue PartialApplicationKeyResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.updateServiceAccountApplicationKeyExecute(req)
}

/*
 * Execute executes the request
 * @return PartialApplicationKeyResponse
 */
func (a *ServiceAccountsApiService) updateServiceAccountApplicationKeyExecute(r apiUpdateServiceAccountApplicationKeyRequest) (PartialApplicationKeyResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPatch
		localVarPostBody    interface{}
		localVarReturnValue PartialApplicationKeyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceAccountsApiService.UpdateServiceAccountApplicationKey")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/service_accounts/{service_account_id}/application_keys/{app_key_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"service_account_id"+"}", _neturl.PathEscape(parameterToString(r.serviceAccountId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"app_key_id"+"}", _neturl.PathEscape(parameterToString(r.appKeyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["DD-API-KEY"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["appKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["DD-APPLICATION-KEY"] = key
			}
		}
	}
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
