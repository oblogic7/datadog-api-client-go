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
	"time"
)

// Linger please
var (
	_ _context.Context
)

// LogsApiService LogsApi service
type LogsApiService service

type apiAggregateLogsRequest struct {
	ctx        _context.Context
	ApiService *LogsApiService
	body       *LogsAggregateRequest
}

func (a *LogsApiService) buildAggregateLogsRequest(ctx _context.Context, body LogsAggregateRequest) (apiAggregateLogsRequest, error) {
	req := apiAggregateLogsRequest{
		ApiService: a,
		ctx:        ctx,
		body:       &body,
	}
	return req, nil
}

/*
 * AggregateLogs Aggregate events
 * The API endpoint to aggregate events into buckets and compute metrics and timeseries.
 */
func (a *LogsApiService) AggregateLogs(ctx _context.Context, body LogsAggregateRequest) (LogsAggregateResponse, *_nethttp.Response, error) {
	req, err := a.buildAggregateLogsRequest(ctx, body)
	if err != nil {
		var localVarReturnValue LogsAggregateResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.aggregateLogsExecute(req)
}

/*
 * Execute executes the request
 * @return LogsAggregateResponse
 */
func (a *LogsApiService) aggregateLogsExecute(r apiAggregateLogsRequest) (LogsAggregateResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue LogsAggregateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogsApiService.AggregateLogs")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/logs/analytics/aggregate"

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

type apiListLogsRequest struct {
	ctx        _context.Context
	ApiService *LogsApiService
	body       *LogsListRequest
}

type ListLogsOptionalParameters struct {
	Body *LogsListRequest
}

func NewListLogsOptionalParameters() *ListLogsOptionalParameters {
	this := ListLogsOptionalParameters{}
	return &this
}
func (r *ListLogsOptionalParameters) WithBody(body LogsListRequest) *ListLogsOptionalParameters {
	r.Body = &body
	return r
}

func (a *LogsApiService) buildListLogsRequest(ctx _context.Context, o ...ListLogsOptionalParameters) (apiListLogsRequest, error) {
	req := apiListLogsRequest{
		ApiService: a,
		ctx:        ctx,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type ListLogsOptionalParameters is allowed")
	}

	if o != nil {
		req.body = o[0].Body
	}
	return req, nil
}

/*
 * ListLogs Search logs
 * List endpoint returns logs that match a log search query.
 * [Results are paginated][1].
 *
 * Use this endpoint to build complex logs filtering and search.
 *
 * **If you are considering archiving logs for your organization,
 * consider use of the Datadog archive capabilities instead of the log list API.
 * See [Datadog Logs Archive documentation][2].**
 *
 * [1]: /logs/guide/collect-multiple-logs-with-pagination
 * [2]: https://docs.datadoghq.com/logs/archives
 */
func (a *LogsApiService) ListLogs(ctx _context.Context, o ...ListLogsOptionalParameters) (LogsListResponse, *_nethttp.Response, error) {
	req, err := a.buildListLogsRequest(ctx, o...)
	if err != nil {
		var localVarReturnValue LogsListResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.listLogsExecute(req)
}

/*
 * ListLogsWithPagination provides a paginated version of ListLogs returning a channel with all items.
 */
func (a *LogsApiService) ListLogsWithPagination(ctx _context.Context, o ...ListLogsOptionalParameters) (<-chan Log, func(), error) {
	ctx, cancel := _context.WithCancel(ctx)
	pageSize_ := int32(10)
	if len(o) == 0 {
		o = append(o, ListLogsOptionalParameters{})
	}
	if o[0].Body == nil {
		o[0].Body = NewLogsListRequest()
	}
	if o[0].Body.Page == nil {
		o[0].Body.Page = NewLogsListRequestPage()
	}
	if o[0].Body.Page.Limit != nil {
		pageSize_ = *o[0].Body.Page.Limit
	}
	o[0].Body.Page.Limit = &pageSize_

	items := make(chan Log, pageSize_)
	go func() {
		for {
			req, err := a.buildListLogsRequest(ctx, o...)
			if err != nil {
				break
			}

			resp, _, err := req.ApiService.listLogsExecute(req)
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

			o[0].Body.Page.Cursor = cursorMetaPageAfter
		}
		close(items)
	}()
	return items, cancel, nil
}

/*
 * Execute executes the request
 * @return LogsListResponse
 */
func (a *LogsApiService) listLogsExecute(r apiListLogsRequest) (LogsListResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue LogsListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogsApiService.ListLogs")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/logs/events/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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

type apiListLogsGetRequest struct {
	ctx         _context.Context
	ApiService  *LogsApiService
	filterQuery *string
	filterIndex *string
	filterFrom  *time.Time
	filterTo    *time.Time
	sort        *LogsSort
	pageCursor  *string
	pageLimit   *int32
}

type ListLogsGetOptionalParameters struct {
	FilterQuery *string
	FilterIndex *string
	FilterFrom  *time.Time
	FilterTo    *time.Time
	Sort        *LogsSort
	PageCursor  *string
	PageLimit   *int32
}

func NewListLogsGetOptionalParameters() *ListLogsGetOptionalParameters {
	this := ListLogsGetOptionalParameters{}
	return &this
}
func (r *ListLogsGetOptionalParameters) WithFilterQuery(filterQuery string) *ListLogsGetOptionalParameters {
	r.FilterQuery = &filterQuery
	return r
}
func (r *ListLogsGetOptionalParameters) WithFilterIndex(filterIndex string) *ListLogsGetOptionalParameters {
	r.FilterIndex = &filterIndex
	return r
}
func (r *ListLogsGetOptionalParameters) WithFilterFrom(filterFrom time.Time) *ListLogsGetOptionalParameters {
	r.FilterFrom = &filterFrom
	return r
}
func (r *ListLogsGetOptionalParameters) WithFilterTo(filterTo time.Time) *ListLogsGetOptionalParameters {
	r.FilterTo = &filterTo
	return r
}
func (r *ListLogsGetOptionalParameters) WithSort(sort LogsSort) *ListLogsGetOptionalParameters {
	r.Sort = &sort
	return r
}
func (r *ListLogsGetOptionalParameters) WithPageCursor(pageCursor string) *ListLogsGetOptionalParameters {
	r.PageCursor = &pageCursor
	return r
}
func (r *ListLogsGetOptionalParameters) WithPageLimit(pageLimit int32) *ListLogsGetOptionalParameters {
	r.PageLimit = &pageLimit
	return r
}

func (a *LogsApiService) buildListLogsGetRequest(ctx _context.Context, o ...ListLogsGetOptionalParameters) (apiListLogsGetRequest, error) {
	req := apiListLogsGetRequest{
		ApiService: a,
		ctx:        ctx,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type ListLogsGetOptionalParameters is allowed")
	}

	if o != nil {
		req.filterQuery = o[0].FilterQuery
		req.filterIndex = o[0].FilterIndex
		req.filterFrom = o[0].FilterFrom
		req.filterTo = o[0].FilterTo
		req.sort = o[0].Sort
		req.pageCursor = o[0].PageCursor
		req.pageLimit = o[0].PageLimit
	}
	return req, nil
}

/*
 * ListLogsGet Get a list of logs
 * List endpoint returns logs that match a log search query.
 * [Results are paginated][1].
 *
 * Use this endpoint to see your latest logs.
 *
 * **If you are considering archiving logs for your organization,
 * consider use of the Datadog archive capabilities instead of the log list API.
 * See [Datadog Logs Archive documentation][2].**
 *
 * [1]: /logs/guide/collect-multiple-logs-with-pagination
 * [2]: https://docs.datadoghq.com/logs/archives
 */
func (a *LogsApiService) ListLogsGet(ctx _context.Context, o ...ListLogsGetOptionalParameters) (LogsListResponse, *_nethttp.Response, error) {
	req, err := a.buildListLogsGetRequest(ctx, o...)
	if err != nil {
		var localVarReturnValue LogsListResponse
		return localVarReturnValue, nil, err
	}

	return req.ApiService.listLogsGetExecute(req)
}

/*
 * ListLogsGetWithPagination provides a paginated version of ListLogsGet returning a channel with all items.
 */
func (a *LogsApiService) ListLogsGetWithPagination(ctx _context.Context, o ...ListLogsGetOptionalParameters) (<-chan Log, func(), error) {
	ctx, cancel := _context.WithCancel(ctx)
	pageSize_ := int32(10)
	if len(o) == 0 {
		o = append(o, ListLogsGetOptionalParameters{})
	}
	if o[0].PageLimit != nil {
		pageSize_ = *o[0].PageLimit
	}
	o[0].PageLimit = &pageSize_

	items := make(chan Log, pageSize_)
	go func() {
		for {
			req, err := a.buildListLogsGetRequest(ctx, o...)
			if err != nil {
				break
			}

			resp, _, err := req.ApiService.listLogsGetExecute(req)
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
 * @return LogsListResponse
 */
func (a *LogsApiService) listLogsGetExecute(r apiListLogsGetRequest) (LogsListResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue LogsListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogsApiService.ListLogsGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/logs/events"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.filterQuery != nil {
		localVarQueryParams.Add("filter[query]", parameterToString(*r.filterQuery, ""))
	}
	if r.filterIndex != nil {
		localVarQueryParams.Add("filter[index]", parameterToString(*r.filterIndex, ""))
	}
	if r.filterFrom != nil {
		localVarQueryParams.Add("filter[from]", parameterToString(*r.filterFrom, ""))
	}
	if r.filterTo != nil {
		localVarQueryParams.Add("filter[to]", parameterToString(*r.filterTo, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
	if r.pageCursor != nil {
		localVarQueryParams.Add("page[cursor]", parameterToString(*r.pageCursor, ""))
	}
	if r.pageLimit != nil {
		localVarQueryParams.Add("page[limit]", parameterToString(*r.pageLimit, ""))
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

type apiSubmitLogRequest struct {
	ctx             _context.Context
	ApiService      *LogsApiService
	body            *[]HTTPLogItem
	contentEncoding *ContentEncoding
	ddtags          *string
}

type SubmitLogOptionalParameters struct {
	ContentEncoding *ContentEncoding
	Ddtags          *string
}

func NewSubmitLogOptionalParameters() *SubmitLogOptionalParameters {
	this := SubmitLogOptionalParameters{}
	return &this
}
func (r *SubmitLogOptionalParameters) WithContentEncoding(contentEncoding ContentEncoding) *SubmitLogOptionalParameters {
	r.ContentEncoding = &contentEncoding
	return r
}
func (r *SubmitLogOptionalParameters) WithDdtags(ddtags string) *SubmitLogOptionalParameters {
	r.Ddtags = &ddtags
	return r
}

func (a *LogsApiService) buildSubmitLogRequest(ctx _context.Context, body []HTTPLogItem, o ...SubmitLogOptionalParameters) (apiSubmitLogRequest, error) {
	req := apiSubmitLogRequest{
		ApiService: a,
		ctx:        ctx,
		body:       &body,
	}

	if len(o) > 1 {
		return req, reportError("only one argument of type SubmitLogOptionalParameters is allowed")
	}

	if o != nil {
		req.contentEncoding = o[0].ContentEncoding
		req.ddtags = o[0].Ddtags
	}
	return req, nil
}

/*
 * SubmitLog Send logs
 * Send your logs to your Datadog platform over HTTP. Limits per HTTP request are:
 *
 * - Maximum content size per payload (uncompressed): 5MB
 * - Maximum size for a single log: 1MB
 * - Maximum array size if sending multiple logs in an array: 1000 entries
 *
 * Any log exceeding 1MB is accepted and truncated by Datadog:
 * - For a single log request, the API truncates the log at 1MB and returns a 2xx.
 * - For a multi-logs request, the API processes all logs, truncates only logs larger than 1MB, and returns a 2xx.
 *
 * Datadog recommends sending your logs compressed.
 * Add the `Content-Encoding: gzip` header to the request when sending compressed logs.
 *
 * The status codes answered by the HTTP API are:
 * - 202: Accepted: the request has been accepted for processing
 * - 400: Bad request (likely an issue in the payload formatting)
 * - 401: Unauthorized (likely a missing API Key)
 * - 403: Permission issue (likely using an invalid API Key)
 * - 408: Request Timeout, request should be retried after some time
 * - 413: Payload too large (batch is above 5MB uncompressed)
 * - 429: Too Many Requests, request should be retried after some time
 * - 500: Internal Server Error, the server encountered an unexpected condition that prevented it from fulfilling the request, request should be retried after some time
 * - 503: Service Unavailable, the server is not ready to handle the request probably because it is overloaded, request should be retried after some time
 */
func (a *LogsApiService) SubmitLog(ctx _context.Context, body []HTTPLogItem, o ...SubmitLogOptionalParameters) (interface{}, *_nethttp.Response, error) {
	req, err := a.buildSubmitLogRequest(ctx, body, o...)
	if err != nil {
		var localVarReturnValue interface{}
		return localVarReturnValue, nil, err
	}

	return req.ApiService.submitLogExecute(req)
}

/*
 * Execute executes the request
 * @return interface{}
 */
func (a *LogsApiService) submitLogExecute(r apiSubmitLogRequest) (interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogsApiService.SubmitLog")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/logs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}
	if r.ddtags != nil {
		localVarQueryParams.Add("ddtags", parameterToString(*r.ddtags, ""))
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/logplex-1", "text/plain"}

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

	if r.contentEncoding != nil {
		localVarHeaderParams["Content-Encoding"] = parameterToString(*r.contentEncoding, "")
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
			var v HTTPLogErrors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v HTTPLogErrors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v HTTPLogErrors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 408 {
			var v HTTPLogErrors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 413 {
			var v HTTPLogErrors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v HTTPLogErrors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v HTTPLogErrors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v HTTPLogErrors
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
