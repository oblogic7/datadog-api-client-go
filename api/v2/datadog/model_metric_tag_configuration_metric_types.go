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

// MetricTagConfigurationMetricTypes The metric's type.
type MetricTagConfigurationMetricTypes string

// List of MetricTagConfigurationMetricTypes
const (
	METRICTAGCONFIGURATIONMETRICTYPES_GAUGE        MetricTagConfigurationMetricTypes = "gauge"
	METRICTAGCONFIGURATIONMETRICTYPES_COUNT        MetricTagConfigurationMetricTypes = "count"
	METRICTAGCONFIGURATIONMETRICTYPES_RATE         MetricTagConfigurationMetricTypes = "rate"
	METRICTAGCONFIGURATIONMETRICTYPES_DISTRIBUTION MetricTagConfigurationMetricTypes = "distribution"
)

var allowedMetricTagConfigurationMetricTypesEnumValues = []MetricTagConfigurationMetricTypes{
	METRICTAGCONFIGURATIONMETRICTYPES_GAUGE,
	METRICTAGCONFIGURATIONMETRICTYPES_COUNT,
	METRICTAGCONFIGURATIONMETRICTYPES_RATE,
	METRICTAGCONFIGURATIONMETRICTYPES_DISTRIBUTION,
}

func (w *MetricTagConfigurationMetricTypes) GetAllowedValues() []MetricTagConfigurationMetricTypes {
	return allowedMetricTagConfigurationMetricTypesEnumValues
}

func (v *MetricTagConfigurationMetricTypes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MetricTagConfigurationMetricTypes(value)
	return nil
}

// NewMetricTagConfigurationMetricTypesFromValue returns a pointer to a valid MetricTagConfigurationMetricTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMetricTagConfigurationMetricTypesFromValue(v string) (*MetricTagConfigurationMetricTypes, error) {
	ev := MetricTagConfigurationMetricTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MetricTagConfigurationMetricTypes: valid values are %v", v, allowedMetricTagConfigurationMetricTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MetricTagConfigurationMetricTypes) IsValid() bool {
	for _, existing := range allowedMetricTagConfigurationMetricTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MetricTagConfigurationMetricTypes value
func (v MetricTagConfigurationMetricTypes) Ptr() *MetricTagConfigurationMetricTypes {
	return &v
}

type NullableMetricTagConfigurationMetricTypes struct {
	value *MetricTagConfigurationMetricTypes
	isSet bool
}

func (v NullableMetricTagConfigurationMetricTypes) Get() *MetricTagConfigurationMetricTypes {
	return v.value
}

func (v *NullableMetricTagConfigurationMetricTypes) Set(val *MetricTagConfigurationMetricTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricTagConfigurationMetricTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricTagConfigurationMetricTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricTagConfigurationMetricTypes(val *MetricTagConfigurationMetricTypes) *NullableMetricTagConfigurationMetricTypes {
	return &NullableMetricTagConfigurationMetricTypes{value: val, isSet: true}
}

func (v NullableMetricTagConfigurationMetricTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricTagConfigurationMetricTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
