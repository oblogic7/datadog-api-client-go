/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
	"fmt"
)

// MetricType The metric resource type.
type MetricType string

// List of MetricType
const (
	METRICTYPE_METRICS MetricType = "metrics"
)

var allowedMetricTypeEnumValues = []MetricType{
	METRICTYPE_METRICS,
}

func (w *MetricType) GetAllowedValues() []MetricType {
	return allowedMetricTypeEnumValues
}

func (v *MetricType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MetricType(value)
	return nil
}

// NewMetricTypeFromValue returns a pointer to a valid MetricType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMetricTypeFromValue(v string) (*MetricType, error) {
	ev := MetricType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MetricType: valid values are %v", v, allowedMetricTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MetricType) IsValid() bool {
	for _, existing := range allowedMetricTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MetricType value
func (v MetricType) Ptr() *MetricType {
	return &v
}

type NullableMetricType struct {
	value *MetricType
	isSet bool
}

func (v NullableMetricType) Get() *MetricType {
	return v.value
}

func (v *NullableMetricType) Set(val *MetricType) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricType) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricType(val *MetricType) *NullableMetricType {
	return &NullableMetricType{value: val, isSet: true}
}

func (v NullableMetricType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
