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

// MetricCustomSpaceAggregation A space aggregation for use in query.
type MetricCustomSpaceAggregation string

// List of MetricCustomSpaceAggregation
const (
	METRICCUSTOMSPACEAGGREGATION_AVG MetricCustomSpaceAggregation = "avg"
	METRICCUSTOMSPACEAGGREGATION_MAX MetricCustomSpaceAggregation = "max"
	METRICCUSTOMSPACEAGGREGATION_MIN MetricCustomSpaceAggregation = "min"
	METRICCUSTOMSPACEAGGREGATION_SUM MetricCustomSpaceAggregation = "sum"
)

var allowedMetricCustomSpaceAggregationEnumValues = []MetricCustomSpaceAggregation{
	METRICCUSTOMSPACEAGGREGATION_AVG,
	METRICCUSTOMSPACEAGGREGATION_MAX,
	METRICCUSTOMSPACEAGGREGATION_MIN,
	METRICCUSTOMSPACEAGGREGATION_SUM,
}

func (w *MetricCustomSpaceAggregation) GetAllowedValues() []MetricCustomSpaceAggregation {
	return allowedMetricCustomSpaceAggregationEnumValues
}

func (v *MetricCustomSpaceAggregation) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MetricCustomSpaceAggregation(value)
	return nil
}

// NewMetricCustomSpaceAggregationFromValue returns a pointer to a valid MetricCustomSpaceAggregation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMetricCustomSpaceAggregationFromValue(v string) (*MetricCustomSpaceAggregation, error) {
	ev := MetricCustomSpaceAggregation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MetricCustomSpaceAggregation: valid values are %v", v, allowedMetricCustomSpaceAggregationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MetricCustomSpaceAggregation) IsValid() bool {
	for _, existing := range allowedMetricCustomSpaceAggregationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MetricCustomSpaceAggregation value
func (v MetricCustomSpaceAggregation) Ptr() *MetricCustomSpaceAggregation {
	return &v
}

type NullableMetricCustomSpaceAggregation struct {
	value *MetricCustomSpaceAggregation
	isSet bool
}

func (v NullableMetricCustomSpaceAggregation) Get() *MetricCustomSpaceAggregation {
	return v.value
}

func (v *NullableMetricCustomSpaceAggregation) Set(val *MetricCustomSpaceAggregation) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricCustomSpaceAggregation) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricCustomSpaceAggregation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricCustomSpaceAggregation(val *MetricCustomSpaceAggregation) *NullableMetricCustomSpaceAggregation {
	return &NullableMetricCustomSpaceAggregation{value: val, isSet: true}
}

func (v NullableMetricCustomSpaceAggregation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricCustomSpaceAggregation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
