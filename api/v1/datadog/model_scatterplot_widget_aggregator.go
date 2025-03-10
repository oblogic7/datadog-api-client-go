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

// ScatterplotWidgetAggregator Aggregator used for the request.
type ScatterplotWidgetAggregator string

// List of ScatterplotWidgetAggregator
const (
	SCATTERPLOTWIDGETAGGREGATOR_AVERAGE ScatterplotWidgetAggregator = "avg"
	SCATTERPLOTWIDGETAGGREGATOR_LAST    ScatterplotWidgetAggregator = "last"
	SCATTERPLOTWIDGETAGGREGATOR_MAXIMUM ScatterplotWidgetAggregator = "max"
	SCATTERPLOTWIDGETAGGREGATOR_MINIMUM ScatterplotWidgetAggregator = "min"
	SCATTERPLOTWIDGETAGGREGATOR_SUM     ScatterplotWidgetAggregator = "sum"
)

var allowedScatterplotWidgetAggregatorEnumValues = []ScatterplotWidgetAggregator{
	SCATTERPLOTWIDGETAGGREGATOR_AVERAGE,
	SCATTERPLOTWIDGETAGGREGATOR_LAST,
	SCATTERPLOTWIDGETAGGREGATOR_MAXIMUM,
	SCATTERPLOTWIDGETAGGREGATOR_MINIMUM,
	SCATTERPLOTWIDGETAGGREGATOR_SUM,
}

func (w *ScatterplotWidgetAggregator) GetAllowedValues() []ScatterplotWidgetAggregator {
	return allowedScatterplotWidgetAggregatorEnumValues
}

func (v *ScatterplotWidgetAggregator) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScatterplotWidgetAggregator(value)
	return nil
}

// NewScatterplotWidgetAggregatorFromValue returns a pointer to a valid ScatterplotWidgetAggregator
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewScatterplotWidgetAggregatorFromValue(v string) (*ScatterplotWidgetAggregator, error) {
	ev := ScatterplotWidgetAggregator(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ScatterplotWidgetAggregator: valid values are %v", v, allowedScatterplotWidgetAggregatorEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ScatterplotWidgetAggregator) IsValid() bool {
	for _, existing := range allowedScatterplotWidgetAggregatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScatterplotWidgetAggregator value
func (v ScatterplotWidgetAggregator) Ptr() *ScatterplotWidgetAggregator {
	return &v
}

type NullableScatterplotWidgetAggregator struct {
	value *ScatterplotWidgetAggregator
	isSet bool
}

func (v NullableScatterplotWidgetAggregator) Get() *ScatterplotWidgetAggregator {
	return v.value
}

func (v *NullableScatterplotWidgetAggregator) Set(val *ScatterplotWidgetAggregator) {
	v.value = val
	v.isSet = true
}

func (v NullableScatterplotWidgetAggregator) IsSet() bool {
	return v.isSet
}

func (v *NullableScatterplotWidgetAggregator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScatterplotWidgetAggregator(val *ScatterplotWidgetAggregator) *NullableScatterplotWidgetAggregator {
	return &NullableScatterplotWidgetAggregator{value: val, isSet: true}
}

func (v NullableScatterplotWidgetAggregator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScatterplotWidgetAggregator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
