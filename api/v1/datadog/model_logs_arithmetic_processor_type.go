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

// LogsArithmeticProcessorType Type of logs arithmetic processor.
type LogsArithmeticProcessorType string

// List of LogsArithmeticProcessorType
const (
	LOGSARITHMETICPROCESSORTYPE_ARITHMETIC_PROCESSOR LogsArithmeticProcessorType = "arithmetic-processor"
)

var allowedLogsArithmeticProcessorTypeEnumValues = []LogsArithmeticProcessorType{
	LOGSARITHMETICPROCESSORTYPE_ARITHMETIC_PROCESSOR,
}

func (w *LogsArithmeticProcessorType) GetAllowedValues() []LogsArithmeticProcessorType {
	return allowedLogsArithmeticProcessorTypeEnumValues
}

func (v *LogsArithmeticProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsArithmeticProcessorType(value)
	return nil
}

// NewLogsArithmeticProcessorTypeFromValue returns a pointer to a valid LogsArithmeticProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLogsArithmeticProcessorTypeFromValue(v string) (*LogsArithmeticProcessorType, error) {
	ev := LogsArithmeticProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LogsArithmeticProcessorType: valid values are %v", v, allowedLogsArithmeticProcessorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LogsArithmeticProcessorType) IsValid() bool {
	for _, existing := range allowedLogsArithmeticProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsArithmeticProcessorType value
func (v LogsArithmeticProcessorType) Ptr() *LogsArithmeticProcessorType {
	return &v
}

type NullableLogsArithmeticProcessorType struct {
	value *LogsArithmeticProcessorType
	isSet bool
}

func (v NullableLogsArithmeticProcessorType) Get() *LogsArithmeticProcessorType {
	return v.value
}

func (v *NullableLogsArithmeticProcessorType) Set(val *LogsArithmeticProcessorType) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsArithmeticProcessorType) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsArithmeticProcessorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsArithmeticProcessorType(val *LogsArithmeticProcessorType) *NullableLogsArithmeticProcessorType {
	return &NullableLogsArithmeticProcessorType{value: val, isSet: true}
}

func (v NullableLogsArithmeticProcessorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsArithmeticProcessorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
