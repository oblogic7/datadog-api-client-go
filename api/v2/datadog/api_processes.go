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
)

// Linger please
var (
	_ _context.Context
)

// ProcessesApiService ProcessesApi service
type ProcessesApiService service

type apiListProcessesRequest struct {
	ctx        _context.Context
	ApiService *ProcessesApiService
	search     *string
	tags       *string
	from       *int64
	to         *int64
	pageLimit  *int32
	pageCursor *string
}

type ListProcessesOptionalParameters struct {
	Search     *string
	Tags       *string
	From       *int64
	To         *int64
	PageLimit  *int32
	PageCursor *string
}

func NewListProcessesOptionalParameters() *ListProcessesOptionalParameters {
	this := ListProcessesOptionalParameters{}
	return &this
}
func (r *ListProcessesOptionalParameters) WithSearch(search string) *ListProcessesOptionalParameters {
	r.Search = &search
	return r
}
func (r *ListProcessesOptionalParameters) WithTags(tags string) *ListProcessesOptionalParameters {
	r.Tags = &tags
	return r
}
func (r *ListProcessesOptionalParameters) WithFrom(from int64) *ListProcessesOptionalParameters {
	r.From = &from
	return r
}
func (r *ListProcessesOptionalParameters) WithTo(to int64) *ListProcessesOptionalParameters {
	r.To = &to
	return r
}
func (r *ListProcessesOptionalParameters) WithPageLimit(pageLimit int32) *ListProcessesOptionalParameters {
	r.PageLimit = &pageLimit
	return r
}
func (r *ListProcessesOptionalParameters) WithPageCursor(pageCursor string) *ListProcessesOptionalParameters {
	r.PageCursor = &pageCursor
	return r
}

func (a *ProcessesApiService) buildListProcessesRequest(ctx _context.Context, o ...ListProcessesOptionalParameters) (apiListProcessesRequest, error) {
	req := apiListProcessesRequest{
		ApiService: a,
		ctx:        ctx,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type ListProcessesOptionalParameters is allowed")
	}

	if o != nil {
		req.search = o[0].Search
		req.tags = o[0].Tags
		req.from = o[0].From
		req.to = o[0].To
		req.pageLimit = o[0].PageLimit
		req.pageCursor = o[0].PageCursor
	}
	return req, nil
}

/*
 * ListProcesses Get all processes
 * Get all processes for your organization.
 */
func (a *ProcessesApiService) ListProcesses(ctx _context.Context, o ...ListProcessesOptionalParameters) (ProcessSummariesResponse, *_nethttp.Response, error) {
	req, err := a.buildListProcessesRequest(ctx, o...)
	if err != nil {
		var localVarReturnValue ProcessSummariesResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.listProcessesExecute(req)
}

/*
 * ListProcessesWithPagination provides a paginated version of ListProcesses returning a channel with all items.
 */
func (a *ProcessesApiService) ListProcessesWithPagination(ctx _context.Context, o ...ListProcessesOptionalParameters) (<-chan ProcessSummary, func(), error) {
	ctx, cancel := _context.WithCancel(ctx)
	pageSize_ := int32(1000)
	if len(o) == 0 {
		o = append(o, ListProcessesOptionalParameters{})
	}
	if o[0].PageLimit != nil {
		pageSize_ = *o[0].PageLimit
	}
	o[0].PageLimit = &pageSize_

	items := make(chan ProcessSummary, pageSize_)
	go func() {
		for {
			req, err := a.buildListProcessesRequest(ctx, o...)
			if err != nil {
				break
			}

			resp, _, err := req.ApiService.listProcessesExecute(req)
			if err != nil {
				break
			}
			respData, ok := resp.GetDataOk()
			if !ok {
				break
			}
			results := *respData

			for _, item := range results {
				select {
				case items <- item:
				case <-ctx.Done():
					close(items)
					return
				}
			}
			if len(results) < int(pageSize_) {
				break
			}
			cursorMeta, ok := resp.GetMetaOk()
			if !ok {
				break
			}
			cursorMetaPage, ok := cursorMeta.GetPageOk()
			if !ok {
				break
			}
			cursorMetaPageAfter, ok := cursorMetaPage.GetAfterOk()
			if !ok {
				break
			}

			o[0].PageCursor = cursorMetaPageAfter
		}
		close(items)
	}()
	return items, cancel, nil
}

/*
 * Execute executes the request
 * @return ProcessSummariesResponse
 */
func (a *ProcessesApiService) listProcessesExecute(r apiListProcessesRequest) (ProcessSummariesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ProcessSummariesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProcessesApiService.ListProcesses")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/processes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.search != nil {
		localVarQueryParams.Add("search", parameterToString(*r.search, ""))
	}
	if r.tags != nil {
		localVarQueryParams.Add("tags", parameterToString(*r.tags, ""))
	}
	if r.from != nil {
		localVarQueryParams.Add("from", parameterToString(*r.from, ""))
	}
	if r.to != nil {
		localVarQueryParams.Add("to", parameterToString(*r.to, ""))
	}
	if r.pageLimit != nil {
		localVarQueryParams.Add("page[limit]", parameterToString(*r.pageLimit, ""))
	}
	if r.pageCursor != nil {
		localVarQueryParams.Add("page[cursor]", parameterToString(*r.pageCursor, ""))
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
