/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// SyntheticsGetBrowserTestLatestResultsResponse Object with the latest Synthetic browser test run.
type SyntheticsGetBrowserTestLatestResultsResponse struct {
	// Timestamp of the latest browser test run.
	LastTimestampFetched *int64 `json:"last_timestamp_fetched,omitempty"`
	// Result of the latest browser test run.
	Results []SyntheticsBrowserTestResultShort `json:"results,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewSyntheticsGetBrowserTestLatestResultsResponse instantiates a new SyntheticsGetBrowserTestLatestResultsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticsGetBrowserTestLatestResultsResponse() *SyntheticsGetBrowserTestLatestResultsResponse {
	this := SyntheticsGetBrowserTestLatestResultsResponse{}
	return &this
}

// NewSyntheticsGetBrowserTestLatestResultsResponseWithDefaults instantiates a new SyntheticsGetBrowserTestLatestResultsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticsGetBrowserTestLatestResultsResponseWithDefaults() *SyntheticsGetBrowserTestLatestResultsResponse {
	this := SyntheticsGetBrowserTestLatestResultsResponse{}
	return &this
}

// GetLastTimestampFetched returns the LastTimestampFetched field value if set, zero value otherwise.
func (o *SyntheticsGetBrowserTestLatestResultsResponse) GetLastTimestampFetched() int64 {
	if o == nil || o.LastTimestampFetched == nil {
		var ret int64
		return ret
	}
	return *o.LastTimestampFetched
}

// GetLastTimestampFetchedOk returns a tuple with the LastTimestampFetched field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsGetBrowserTestLatestResultsResponse) GetLastTimestampFetchedOk() (*int64, bool) {
	if o == nil || o.LastTimestampFetched == nil {
		return nil, false
	}
	return o.LastTimestampFetched, true
}

// HasLastTimestampFetched returns a boolean if a field has been set.
func (o *SyntheticsGetBrowserTestLatestResultsResponse) HasLastTimestampFetched() bool {
	if o != nil && o.LastTimestampFetched != nil {
		return true
	}

	return false
}

// SetLastTimestampFetched gets a reference to the given int64 and assigns it to the LastTimestampFetched field.
func (o *SyntheticsGetBrowserTestLatestResultsResponse) SetLastTimestampFetched(v int64) {
	o.LastTimestampFetched = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *SyntheticsGetBrowserTestLatestResultsResponse) GetResults() []SyntheticsBrowserTestResultShort {
	if o == nil || o.Results == nil {
		var ret []SyntheticsBrowserTestResultShort
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsGetBrowserTestLatestResultsResponse) GetResultsOk() (*[]SyntheticsBrowserTestResultShort, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return &o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *SyntheticsGetBrowserTestLatestResultsResponse) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []SyntheticsBrowserTestResultShort and assigns it to the Results field.
func (o *SyntheticsGetBrowserTestLatestResultsResponse) SetResults(v []SyntheticsBrowserTestResultShort) {
	o.Results = v
}

func (o SyntheticsGetBrowserTestLatestResultsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.LastTimestampFetched != nil {
		toSerialize["last_timestamp_fetched"] = o.LastTimestampFetched
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

func (o *SyntheticsGetBrowserTestLatestResultsResponse) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		LastTimestampFetched *int64                             `json:"last_timestamp_fetched,omitempty"`
		Results              []SyntheticsBrowserTestResultShort `json:"results,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.LastTimestampFetched = all.LastTimestampFetched
	o.Results = all.Results
	return nil
}
