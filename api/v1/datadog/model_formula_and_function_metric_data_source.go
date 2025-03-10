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

// FormulaAndFunctionMetricDataSource Data source for metrics queries.
type FormulaAndFunctionMetricDataSource string

// List of FormulaAndFunctionMetricDataSource
const (
	FORMULAANDFUNCTIONMETRICDATASOURCE_METRICS FormulaAndFunctionMetricDataSource = "metrics"
)

var allowedFormulaAndFunctionMetricDataSourceEnumValues = []FormulaAndFunctionMetricDataSource{
	FORMULAANDFUNCTIONMETRICDATASOURCE_METRICS,
}

func (w *FormulaAndFunctionMetricDataSource) GetAllowedValues() []FormulaAndFunctionMetricDataSource {
	return allowedFormulaAndFunctionMetricDataSourceEnumValues
}

func (v *FormulaAndFunctionMetricDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaAndFunctionMetricDataSource(value)
	return nil
}

// NewFormulaAndFunctionMetricDataSourceFromValue returns a pointer to a valid FormulaAndFunctionMetricDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFormulaAndFunctionMetricDataSourceFromValue(v string) (*FormulaAndFunctionMetricDataSource, error) {
	ev := FormulaAndFunctionMetricDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FormulaAndFunctionMetricDataSource: valid values are %v", v, allowedFormulaAndFunctionMetricDataSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FormulaAndFunctionMetricDataSource) IsValid() bool {
	for _, existing := range allowedFormulaAndFunctionMetricDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaAndFunctionMetricDataSource value
func (v FormulaAndFunctionMetricDataSource) Ptr() *FormulaAndFunctionMetricDataSource {
	return &v
}

type NullableFormulaAndFunctionMetricDataSource struct {
	value *FormulaAndFunctionMetricDataSource
	isSet bool
}

func (v NullableFormulaAndFunctionMetricDataSource) Get() *FormulaAndFunctionMetricDataSource {
	return v.value
}

func (v *NullableFormulaAndFunctionMetricDataSource) Set(val *FormulaAndFunctionMetricDataSource) {
	v.value = val
	v.isSet = true
}

func (v NullableFormulaAndFunctionMetricDataSource) IsSet() bool {
	return v.isSet
}

func (v *NullableFormulaAndFunctionMetricDataSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormulaAndFunctionMetricDataSource(val *FormulaAndFunctionMetricDataSource) *NullableFormulaAndFunctionMetricDataSource {
	return &NullableFormulaAndFunctionMetricDataSource{value: val, isSet: true}
}

func (v NullableFormulaAndFunctionMetricDataSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormulaAndFunctionMetricDataSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
