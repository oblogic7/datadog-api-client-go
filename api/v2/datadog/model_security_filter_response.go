/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// SecurityFilterResponse Response object which includes a single security filter.
type SecurityFilterResponse struct {
	// The security filter's properties.
	Data *SecurityFilter `json:"data,omitempty"`
	// Optional metadata associated to the response.
	Meta *SecurityFilterMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewSecurityFilterResponse instantiates a new SecurityFilterResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityFilterResponse() *SecurityFilterResponse {
	this := SecurityFilterResponse{}
	return &this
}

// NewSecurityFilterResponseWithDefaults instantiates a new SecurityFilterResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityFilterResponseWithDefaults() *SecurityFilterResponse {
	this := SecurityFilterResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SecurityFilterResponse) GetData() SecurityFilter {
	if o == nil || o.Data == nil {
		var ret SecurityFilter
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityFilterResponse) GetDataOk() (*SecurityFilter, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SecurityFilterResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given SecurityFilter and assigns it to the Data field.
func (o *SecurityFilterResponse) SetData(v SecurityFilter) {
	o.Data = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SecurityFilterResponse) GetMeta() SecurityFilterMeta {
	if o == nil || o.Meta == nil {
		var ret SecurityFilterMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityFilterResponse) GetMetaOk() (*SecurityFilterMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SecurityFilterResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given SecurityFilterMeta and assigns it to the Meta field.
func (o *SecurityFilterResponse) SetMeta(v SecurityFilterMeta) {
	o.Meta = &v
}

func (o SecurityFilterResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

func (o *SecurityFilterResponse) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Data *SecurityFilter     `json:"data,omitempty"`
		Meta *SecurityFilterMeta `json:"meta,omitempty"`
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
	if all.Data != nil && all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Data = all.Data
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Meta = all.Meta
	return nil
}
