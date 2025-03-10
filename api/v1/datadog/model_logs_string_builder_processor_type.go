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

// LogsStringBuilderProcessorType Type of logs string builder processor.
type LogsStringBuilderProcessorType string

// List of LogsStringBuilderProcessorType
const (
	LOGSSTRINGBUILDERPROCESSORTYPE_STRING_BUILDER_PROCESSOR LogsStringBuilderProcessorType = "string-builder-processor"
)

var allowedLogsStringBuilderProcessorTypeEnumValues = []LogsStringBuilderProcessorType{
	LOGSSTRINGBUILDERPROCESSORTYPE_STRING_BUILDER_PROCESSOR,
}

func (w *LogsStringBuilderProcessorType) GetAllowedValues() []LogsStringBuilderProcessorType {
	return allowedLogsStringBuilderProcessorTypeEnumValues
}

func (v *LogsStringBuilderProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsStringBuilderProcessorType(value)
	return nil
}

// NewLogsStringBuilderProcessorTypeFromValue returns a pointer to a valid LogsStringBuilderProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLogsStringBuilderProcessorTypeFromValue(v string) (*LogsStringBuilderProcessorType, error) {
	ev := LogsStringBuilderProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LogsStringBuilderProcessorType: valid values are %v", v, allowedLogsStringBuilderProcessorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LogsStringBuilderProcessorType) IsValid() bool {
	for _, existing := range allowedLogsStringBuilderProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsStringBuilderProcessorType value
func (v LogsStringBuilderProcessorType) Ptr() *LogsStringBuilderProcessorType {
	return &v
}

type NullableLogsStringBuilderProcessorType struct {
	value *LogsStringBuilderProcessorType
	isSet bool
}

func (v NullableLogsStringBuilderProcessorType) Get() *LogsStringBuilderProcessorType {
	return v.value
}

func (v *NullableLogsStringBuilderProcessorType) Set(val *LogsStringBuilderProcessorType) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsStringBuilderProcessorType) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsStringBuilderProcessorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsStringBuilderProcessorType(val *LogsStringBuilderProcessorType) *NullableLogsStringBuilderProcessorType {
	return &NullableLogsStringBuilderProcessorType{value: val, isSet: true}
}

func (v NullableLogsStringBuilderProcessorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsStringBuilderProcessorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
