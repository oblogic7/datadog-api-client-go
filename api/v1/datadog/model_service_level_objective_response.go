/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// ServiceLevelObjectiveResponse TODO.
type ServiceLevelObjectiveResponse struct {
	Data ServiceLevelObjective `json:"data"`
	// An array of error messages. Each endpoint documents how/whether this field is used.
	Errors *[]string `json:"errors,omitempty"`
}

// NewServiceLevelObjectiveResponse instantiates a new ServiceLevelObjectiveResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceLevelObjectiveResponse(data ServiceLevelObjective) *ServiceLevelObjectiveResponse {
	this := ServiceLevelObjectiveResponse{}
	this.Data = data
	return &this
}

// NewServiceLevelObjectiveResponseWithDefaults instantiates a new ServiceLevelObjectiveResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceLevelObjectiveResponseWithDefaults() *ServiceLevelObjectiveResponse {
	this := ServiceLevelObjectiveResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ServiceLevelObjectiveResponse) GetData() ServiceLevelObjective {
	if o == nil {
		var ret ServiceLevelObjective
		return ret
	}

	return o.Data
}

// SetData sets field value
func (o *ServiceLevelObjectiveResponse) SetData(v ServiceLevelObjective) {
	o.Data = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *ServiceLevelObjectiveResponse) GetErrors() []string {
	if o == nil || o.Errors == nil {
		var ret []string
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLevelObjectiveResponse) GetErrorsOk() ([]string, bool) {
	if o == nil || o.Errors == nil {
		var ret []string
		return ret, false
	}
	return *o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ServiceLevelObjectiveResponse) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *ServiceLevelObjectiveResponse) SetErrors(v []string) {
	o.Errors = &v
}

func (o ServiceLevelObjectiveResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableServiceLevelObjectiveResponse struct {
	value *ServiceLevelObjectiveResponse
	isSet bool
}

func (v NullableServiceLevelObjectiveResponse) Get() *ServiceLevelObjectiveResponse {
	return v.value
}

func (v NullableServiceLevelObjectiveResponse) Set(val *ServiceLevelObjectiveResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceLevelObjectiveResponse) IsSet() bool {
	return v.isSet
}

func (v NullableServiceLevelObjectiveResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceLevelObjectiveResponse(val *ServiceLevelObjectiveResponse) *NullableServiceLevelObjectiveResponse {
	return &NullableServiceLevelObjectiveResponse{value: val, isSet: true}
}

func (v NullableServiceLevelObjectiveResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceLevelObjectiveResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
