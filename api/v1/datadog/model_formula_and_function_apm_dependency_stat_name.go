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

// FormulaAndFunctionApmDependencyStatName APM statistic.
type FormulaAndFunctionApmDependencyStatName string

// List of FormulaAndFunctionApmDependencyStatName
const (
	FORMULAANDFUNCTIONAPMDEPENDENCYSTATNAME_AVG_DURATION        FormulaAndFunctionApmDependencyStatName = "avg_duration"
	FORMULAANDFUNCTIONAPMDEPENDENCYSTATNAME_AVG_ROOT_DURATION   FormulaAndFunctionApmDependencyStatName = "avg_root_duration"
	FORMULAANDFUNCTIONAPMDEPENDENCYSTATNAME_AVG_SPANS_PER_TRACE FormulaAndFunctionApmDependencyStatName = "avg_spans_per_trace"
	FORMULAANDFUNCTIONAPMDEPENDENCYSTATNAME_ERROR_RATE          FormulaAndFunctionApmDependencyStatName = "error_rate"
	FORMULAANDFUNCTIONAPMDEPENDENCYSTATNAME_PCT_EXEC_TIME       FormulaAndFunctionApmDependencyStatName = "pct_exec_time"
	FORMULAANDFUNCTIONAPMDEPENDENCYSTATNAME_PCT_OF_TRACES       FormulaAndFunctionApmDependencyStatName = "pct_of_traces"
	FORMULAANDFUNCTIONAPMDEPENDENCYSTATNAME_TOTAL_TRACES_COUNT  FormulaAndFunctionApmDependencyStatName = "total_traces_count"
)

var allowedFormulaAndFunctionApmDependencyStatNameEnumValues = []FormulaAndFunctionApmDependencyStatName{
	FORMULAANDFUNCTIONAPMDEPENDENCYSTATNAME_AVG_DURATION,
	FORMULAANDFUNCTIONAPMDEPENDENCYSTATNAME_AVG_ROOT_DURATION,
	FORMULAANDFUNCTIONAPMDEPENDENCYSTATNAME_AVG_SPANS_PER_TRACE,
	FORMULAANDFUNCTIONAPMDEPENDENCYSTATNAME_ERROR_RATE,
	FORMULAANDFUNCTIONAPMDEPENDENCYSTATNAME_PCT_EXEC_TIME,
	FORMULAANDFUNCTIONAPMDEPENDENCYSTATNAME_PCT_OF_TRACES,
	FORMULAANDFUNCTIONAPMDEPENDENCYSTATNAME_TOTAL_TRACES_COUNT,
}

func (w *FormulaAndFunctionApmDependencyStatName) GetAllowedValues() []FormulaAndFunctionApmDependencyStatName {
	return allowedFormulaAndFunctionApmDependencyStatNameEnumValues
}

func (v *FormulaAndFunctionApmDependencyStatName) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaAndFunctionApmDependencyStatName(value)
	return nil
}

// NewFormulaAndFunctionApmDependencyStatNameFromValue returns a pointer to a valid FormulaAndFunctionApmDependencyStatName
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFormulaAndFunctionApmDependencyStatNameFromValue(v string) (*FormulaAndFunctionApmDependencyStatName, error) {
	ev := FormulaAndFunctionApmDependencyStatName(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FormulaAndFunctionApmDependencyStatName: valid values are %v", v, allowedFormulaAndFunctionApmDependencyStatNameEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FormulaAndFunctionApmDependencyStatName) IsValid() bool {
	for _, existing := range allowedFormulaAndFunctionApmDependencyStatNameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaAndFunctionApmDependencyStatName value
func (v FormulaAndFunctionApmDependencyStatName) Ptr() *FormulaAndFunctionApmDependencyStatName {
	return &v
}

type NullableFormulaAndFunctionApmDependencyStatName struct {
	value *FormulaAndFunctionApmDependencyStatName
	isSet bool
}

func (v NullableFormulaAndFunctionApmDependencyStatName) Get() *FormulaAndFunctionApmDependencyStatName {
	return v.value
}

func (v *NullableFormulaAndFunctionApmDependencyStatName) Set(val *FormulaAndFunctionApmDependencyStatName) {
	v.value = val
	v.isSet = true
}

func (v NullableFormulaAndFunctionApmDependencyStatName) IsSet() bool {
	return v.isSet
}

func (v *NullableFormulaAndFunctionApmDependencyStatName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormulaAndFunctionApmDependencyStatName(val *FormulaAndFunctionApmDependencyStatName) *NullableFormulaAndFunctionApmDependencyStatName {
	return &NullableFormulaAndFunctionApmDependencyStatName{value: val, isSet: true}
}

func (v NullableFormulaAndFunctionApmDependencyStatName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormulaAndFunctionApmDependencyStatName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
