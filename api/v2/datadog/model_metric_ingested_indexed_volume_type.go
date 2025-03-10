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

// MetricIngestedIndexedVolumeType The metric ingested and indexed volume type.
type MetricIngestedIndexedVolumeType string

// List of MetricIngestedIndexedVolumeType
const (
	METRICINGESTEDINDEXEDVOLUMETYPE_METRIC_VOLUMES MetricIngestedIndexedVolumeType = "metric_volumes"
)

var allowedMetricIngestedIndexedVolumeTypeEnumValues = []MetricIngestedIndexedVolumeType{
	METRICINGESTEDINDEXEDVOLUMETYPE_METRIC_VOLUMES,
}

func (w *MetricIngestedIndexedVolumeType) GetAllowedValues() []MetricIngestedIndexedVolumeType {
	return allowedMetricIngestedIndexedVolumeTypeEnumValues
}

func (v *MetricIngestedIndexedVolumeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MetricIngestedIndexedVolumeType(value)
	return nil
}

// NewMetricIngestedIndexedVolumeTypeFromValue returns a pointer to a valid MetricIngestedIndexedVolumeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMetricIngestedIndexedVolumeTypeFromValue(v string) (*MetricIngestedIndexedVolumeType, error) {
	ev := MetricIngestedIndexedVolumeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MetricIngestedIndexedVolumeType: valid values are %v", v, allowedMetricIngestedIndexedVolumeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MetricIngestedIndexedVolumeType) IsValid() bool {
	for _, existing := range allowedMetricIngestedIndexedVolumeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MetricIngestedIndexedVolumeType value
func (v MetricIngestedIndexedVolumeType) Ptr() *MetricIngestedIndexedVolumeType {
	return &v
}

type NullableMetricIngestedIndexedVolumeType struct {
	value *MetricIngestedIndexedVolumeType
	isSet bool
}

func (v NullableMetricIngestedIndexedVolumeType) Get() *MetricIngestedIndexedVolumeType {
	return v.value
}

func (v *NullableMetricIngestedIndexedVolumeType) Set(val *MetricIngestedIndexedVolumeType) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricIngestedIndexedVolumeType) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricIngestedIndexedVolumeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricIngestedIndexedVolumeType(val *MetricIngestedIndexedVolumeType) *NullableMetricIngestedIndexedVolumeType {
	return &NullableMetricIngestedIndexedVolumeType{value: val, isSet: true}
}

func (v NullableMetricIngestedIndexedVolumeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricIngestedIndexedVolumeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
