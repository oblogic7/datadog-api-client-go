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

// LogsPipelineProcessorType Type of logs pipeline processor.
type LogsPipelineProcessorType string

// List of LogsPipelineProcessorType
const (
	LOGSPIPELINEPROCESSORTYPE_PIPELINE LogsPipelineProcessorType = "pipeline"
)

var allowedLogsPipelineProcessorTypeEnumValues = []LogsPipelineProcessorType{
	LOGSPIPELINEPROCESSORTYPE_PIPELINE,
}

func (w *LogsPipelineProcessorType) GetAllowedValues() []LogsPipelineProcessorType {
	return allowedLogsPipelineProcessorTypeEnumValues
}

func (v *LogsPipelineProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsPipelineProcessorType(value)
	return nil
}

// NewLogsPipelineProcessorTypeFromValue returns a pointer to a valid LogsPipelineProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLogsPipelineProcessorTypeFromValue(v string) (*LogsPipelineProcessorType, error) {
	ev := LogsPipelineProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LogsPipelineProcessorType: valid values are %v", v, allowedLogsPipelineProcessorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LogsPipelineProcessorType) IsValid() bool {
	for _, existing := range allowedLogsPipelineProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsPipelineProcessorType value
func (v LogsPipelineProcessorType) Ptr() *LogsPipelineProcessorType {
	return &v
}

type NullableLogsPipelineProcessorType struct {
	value *LogsPipelineProcessorType
	isSet bool
}

func (v NullableLogsPipelineProcessorType) Get() *LogsPipelineProcessorType {
	return v.value
}

func (v *NullableLogsPipelineProcessorType) Set(val *LogsPipelineProcessorType) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsPipelineProcessorType) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsPipelineProcessorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsPipelineProcessorType(val *LogsPipelineProcessorType) *NullableLogsPipelineProcessorType {
	return &NullableLogsPipelineProcessorType{value: val, isSet: true}
}

func (v NullableLogsPipelineProcessorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsPipelineProcessorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
