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

// MetricTagConfigurationType The metric tag configuration resource type.
type MetricTagConfigurationType string

// List of MetricTagConfigurationType
const (
	METRICTAGCONFIGURATIONTYPE_MANAGE_TAGS MetricTagConfigurationType = "manage_tags"
)

var allowedMetricTagConfigurationTypeEnumValues = []MetricTagConfigurationType{
	METRICTAGCONFIGURATIONTYPE_MANAGE_TAGS,
}

func (w *MetricTagConfigurationType) GetAllowedValues() []MetricTagConfigurationType {
	return allowedMetricTagConfigurationTypeEnumValues
}

func (v *MetricTagConfigurationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MetricTagConfigurationType(value)
	return nil
}

// NewMetricTagConfigurationTypeFromValue returns a pointer to a valid MetricTagConfigurationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMetricTagConfigurationTypeFromValue(v string) (*MetricTagConfigurationType, error) {
	ev := MetricTagConfigurationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MetricTagConfigurationType: valid values are %v", v, allowedMetricTagConfigurationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MetricTagConfigurationType) IsValid() bool {
	for _, existing := range allowedMetricTagConfigurationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MetricTagConfigurationType value
func (v MetricTagConfigurationType) Ptr() *MetricTagConfigurationType {
	return &v
}

type NullableMetricTagConfigurationType struct {
	value *MetricTagConfigurationType
	isSet bool
}

func (v NullableMetricTagConfigurationType) Get() *MetricTagConfigurationType {
	return v.value
}

func (v *NullableMetricTagConfigurationType) Set(val *MetricTagConfigurationType) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricTagConfigurationType) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricTagConfigurationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricTagConfigurationType(val *MetricTagConfigurationType) *NullableMetricTagConfigurationType {
	return &NullableMetricTagConfigurationType{value: val, isSet: true}
}

func (v NullableMetricTagConfigurationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricTagConfigurationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
