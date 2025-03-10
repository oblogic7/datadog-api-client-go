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

// FormulaAndFunctionEventAggregation Aggregation methods for event platform queries.
type FormulaAndFunctionEventAggregation string

// List of FormulaAndFunctionEventAggregation
const (
	FORMULAANDFUNCTIONEVENTAGGREGATION_COUNT       FormulaAndFunctionEventAggregation = "count"
	FORMULAANDFUNCTIONEVENTAGGREGATION_CARDINALITY FormulaAndFunctionEventAggregation = "cardinality"
	FORMULAANDFUNCTIONEVENTAGGREGATION_MEDIAN      FormulaAndFunctionEventAggregation = "median"
	FORMULAANDFUNCTIONEVENTAGGREGATION_PC75        FormulaAndFunctionEventAggregation = "pc75"
	FORMULAANDFUNCTIONEVENTAGGREGATION_PC90        FormulaAndFunctionEventAggregation = "pc90"
	FORMULAANDFUNCTIONEVENTAGGREGATION_PC95        FormulaAndFunctionEventAggregation = "pc95"
	FORMULAANDFUNCTIONEVENTAGGREGATION_PC98        FormulaAndFunctionEventAggregation = "pc98"
	FORMULAANDFUNCTIONEVENTAGGREGATION_PC99        FormulaAndFunctionEventAggregation = "pc99"
	FORMULAANDFUNCTIONEVENTAGGREGATION_SUM         FormulaAndFunctionEventAggregation = "sum"
	FORMULAANDFUNCTIONEVENTAGGREGATION_MIN         FormulaAndFunctionEventAggregation = "min"
	FORMULAANDFUNCTIONEVENTAGGREGATION_MAX         FormulaAndFunctionEventAggregation = "max"
	FORMULAANDFUNCTIONEVENTAGGREGATION_AVG         FormulaAndFunctionEventAggregation = "avg"
)

var allowedFormulaAndFunctionEventAggregationEnumValues = []FormulaAndFunctionEventAggregation{
	FORMULAANDFUNCTIONEVENTAGGREGATION_COUNT,
	FORMULAANDFUNCTIONEVENTAGGREGATION_CARDINALITY,
	FORMULAANDFUNCTIONEVENTAGGREGATION_MEDIAN,
	FORMULAANDFUNCTIONEVENTAGGREGATION_PC75,
	FORMULAANDFUNCTIONEVENTAGGREGATION_PC90,
	FORMULAANDFUNCTIONEVENTAGGREGATION_PC95,
	FORMULAANDFUNCTIONEVENTAGGREGATION_PC98,
	FORMULAANDFUNCTIONEVENTAGGREGATION_PC99,
	FORMULAANDFUNCTIONEVENTAGGREGATION_SUM,
	FORMULAANDFUNCTIONEVENTAGGREGATION_MIN,
	FORMULAANDFUNCTIONEVENTAGGREGATION_MAX,
	FORMULAANDFUNCTIONEVENTAGGREGATION_AVG,
}

func (w *FormulaAndFunctionEventAggregation) GetAllowedValues() []FormulaAndFunctionEventAggregation {
	return allowedFormulaAndFunctionEventAggregationEnumValues
}

func (v *FormulaAndFunctionEventAggregation) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaAndFunctionEventAggregation(value)
	return nil
}

// NewFormulaAndFunctionEventAggregationFromValue returns a pointer to a valid FormulaAndFunctionEventAggregation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFormulaAndFunctionEventAggregationFromValue(v string) (*FormulaAndFunctionEventAggregation, error) {
	ev := FormulaAndFunctionEventAggregation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FormulaAndFunctionEventAggregation: valid values are %v", v, allowedFormulaAndFunctionEventAggregationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FormulaAndFunctionEventAggregation) IsValid() bool {
	for _, existing := range allowedFormulaAndFunctionEventAggregationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaAndFunctionEventAggregation value
func (v FormulaAndFunctionEventAggregation) Ptr() *FormulaAndFunctionEventAggregation {
	return &v
}

type NullableFormulaAndFunctionEventAggregation struct {
	value *FormulaAndFunctionEventAggregation
	isSet bool
}

func (v NullableFormulaAndFunctionEventAggregation) Get() *FormulaAndFunctionEventAggregation {
	return v.value
}

func (v *NullableFormulaAndFunctionEventAggregation) Set(val *FormulaAndFunctionEventAggregation) {
	v.value = val
	v.isSet = true
}

func (v NullableFormulaAndFunctionEventAggregation) IsSet() bool {
	return v.isSet
}

func (v *NullableFormulaAndFunctionEventAggregation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormulaAndFunctionEventAggregation(val *FormulaAndFunctionEventAggregation) *NullableFormulaAndFunctionEventAggregation {
	return &NullableFormulaAndFunctionEventAggregation{value: val, isSet: true}
}

func (v NullableFormulaAndFunctionEventAggregation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormulaAndFunctionEventAggregation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
