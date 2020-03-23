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

// HostMetrics TODO.
type HostMetrics struct {
	// TODO.
	Cpu *float64 `json:"cpu,omitempty"`
	// TODO.
	Iowait *float64 `json:"iowait,omitempty"`
	// TODO.
	Load *float64 `json:"load,omitempty"`
}

// NewHostMetrics instantiates a new HostMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostMetrics() *HostMetrics {
	this := HostMetrics{}
	return &this
}

// NewHostMetricsWithDefaults instantiates a new HostMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostMetricsWithDefaults() *HostMetrics {
	this := HostMetrics{}
	return &this
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *HostMetrics) GetCpu() float64 {
	if o == nil || o.Cpu == nil {
		var ret float64
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *HostMetrics) GetCpuOk() (float64, bool) {
	if o == nil || o.Cpu == nil {
		var ret float64
		return ret, false
	}
	return *o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *HostMetrics) HasCpu() bool {
	if o != nil && o.Cpu != nil {
		return true
	}

	return false
}

// SetCpu gets a reference to the given float64 and assigns it to the Cpu field.
func (o *HostMetrics) SetCpu(v float64) {
	o.Cpu = &v
}

// GetIowait returns the Iowait field value if set, zero value otherwise.
func (o *HostMetrics) GetIowait() float64 {
	if o == nil || o.Iowait == nil {
		var ret float64
		return ret
	}
	return *o.Iowait
}

// GetIowaitOk returns a tuple with the Iowait field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *HostMetrics) GetIowaitOk() (float64, bool) {
	if o == nil || o.Iowait == nil {
		var ret float64
		return ret, false
	}
	return *o.Iowait, true
}

// HasIowait returns a boolean if a field has been set.
func (o *HostMetrics) HasIowait() bool {
	if o != nil && o.Iowait != nil {
		return true
	}

	return false
}

// SetIowait gets a reference to the given float64 and assigns it to the Iowait field.
func (o *HostMetrics) SetIowait(v float64) {
	o.Iowait = &v
}

// GetLoad returns the Load field value if set, zero value otherwise.
func (o *HostMetrics) GetLoad() float64 {
	if o == nil || o.Load == nil {
		var ret float64
		return ret
	}
	return *o.Load
}

// GetLoadOk returns a tuple with the Load field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *HostMetrics) GetLoadOk() (float64, bool) {
	if o == nil || o.Load == nil {
		var ret float64
		return ret, false
	}
	return *o.Load, true
}

// HasLoad returns a boolean if a field has been set.
func (o *HostMetrics) HasLoad() bool {
	if o != nil && o.Load != nil {
		return true
	}

	return false
}

// SetLoad gets a reference to the given float64 and assigns it to the Load field.
func (o *HostMetrics) SetLoad(v float64) {
	o.Load = &v
}

func (o HostMetrics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cpu != nil {
		toSerialize["cpu"] = o.Cpu
	}
	if o.Iowait != nil {
		toSerialize["iowait"] = o.Iowait
	}
	if o.Load != nil {
		toSerialize["load"] = o.Load
	}
	return json.Marshal(toSerialize)
}

type NullableHostMetrics struct {
	value *HostMetrics
	isSet bool
}

func (v NullableHostMetrics) Get() *HostMetrics {
	return v.value
}

func (v NullableHostMetrics) Set(val *HostMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableHostMetrics) IsSet() bool {
	return v.isSet
}

func (v NullableHostMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostMetrics(val *HostMetrics) *NullableHostMetrics {
	return &NullableHostMetrics{value: val, isSet: true}
}

func (v NullableHostMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
