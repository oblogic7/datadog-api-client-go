/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// UsageSpecifiedCustomReportsPage The object containing page total count for specified ID.
type UsageSpecifiedCustomReportsPage struct {
	// Total page count.
	TotalCount *int64 `json:"total_count,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewUsageSpecifiedCustomReportsPage instantiates a new UsageSpecifiedCustomReportsPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageSpecifiedCustomReportsPage() *UsageSpecifiedCustomReportsPage {
	this := UsageSpecifiedCustomReportsPage{}
	return &this
}

// NewUsageSpecifiedCustomReportsPageWithDefaults instantiates a new UsageSpecifiedCustomReportsPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageSpecifiedCustomReportsPageWithDefaults() *UsageSpecifiedCustomReportsPage {
	this := UsageSpecifiedCustomReportsPage{}
	return &this
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *UsageSpecifiedCustomReportsPage) GetTotalCount() int64 {
	if o == nil || o.TotalCount == nil {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSpecifiedCustomReportsPage) GetTotalCountOk() (*int64, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *UsageSpecifiedCustomReportsPage) HasTotalCount() bool {
	if o != nil && o.TotalCount != nil {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *UsageSpecifiedCustomReportsPage) SetTotalCount(v int64) {
	o.TotalCount = &v
}

func (o UsageSpecifiedCustomReportsPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.TotalCount != nil {
		toSerialize["total_count"] = o.TotalCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

func (o *UsageSpecifiedCustomReportsPage) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		TotalCount *int64 `json:"total_count,omitempty"`
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
	o.TotalCount = all.TotalCount
	return nil
}
