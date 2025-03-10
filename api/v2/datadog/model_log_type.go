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

// LogType Type of the event.
type LogType string

// List of LogType
const (
	LOGTYPE_LOG LogType = "log"
)

var allowedLogTypeEnumValues = []LogType{
	LOGTYPE_LOG,
}

func (w *LogType) GetAllowedValues() []LogType {
	return allowedLogTypeEnumValues
}

func (v *LogType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogType(value)
	return nil
}

// NewLogTypeFromValue returns a pointer to a valid LogType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLogTypeFromValue(v string) (*LogType, error) {
	ev := LogType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LogType: valid values are %v", v, allowedLogTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LogType) IsValid() bool {
	for _, existing := range allowedLogTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogType value
func (v LogType) Ptr() *LogType {
	return &v
}

type NullableLogType struct {
	value *LogType
	isSet bool
}

func (v NullableLogType) Get() *LogType {
	return v.value
}

func (v *NullableLogType) Set(val *LogType) {
	v.value = val
	v.isSet = true
}

func (v NullableLogType) IsSet() bool {
	return v.isSet
}

func (v *NullableLogType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogType(val *LogType) *NullableLogType {
	return &NullableLogType{value: val, isSet: true}
}

func (v NullableLogType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
