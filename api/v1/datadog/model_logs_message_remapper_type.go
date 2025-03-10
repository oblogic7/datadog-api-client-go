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

// LogsMessageRemapperType Type of logs message remapper.
type LogsMessageRemapperType string

// List of LogsMessageRemapperType
const (
	LOGSMESSAGEREMAPPERTYPE_MESSAGE_REMAPPER LogsMessageRemapperType = "message-remapper"
)

var allowedLogsMessageRemapperTypeEnumValues = []LogsMessageRemapperType{
	LOGSMESSAGEREMAPPERTYPE_MESSAGE_REMAPPER,
}

func (w *LogsMessageRemapperType) GetAllowedValues() []LogsMessageRemapperType {
	return allowedLogsMessageRemapperTypeEnumValues
}

func (v *LogsMessageRemapperType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsMessageRemapperType(value)
	return nil
}

// NewLogsMessageRemapperTypeFromValue returns a pointer to a valid LogsMessageRemapperType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLogsMessageRemapperTypeFromValue(v string) (*LogsMessageRemapperType, error) {
	ev := LogsMessageRemapperType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LogsMessageRemapperType: valid values are %v", v, allowedLogsMessageRemapperTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LogsMessageRemapperType) IsValid() bool {
	for _, existing := range allowedLogsMessageRemapperTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsMessageRemapperType value
func (v LogsMessageRemapperType) Ptr() *LogsMessageRemapperType {
	return &v
}

type NullableLogsMessageRemapperType struct {
	value *LogsMessageRemapperType
	isSet bool
}

func (v NullableLogsMessageRemapperType) Get() *LogsMessageRemapperType {
	return v.value
}

func (v *NullableLogsMessageRemapperType) Set(val *LogsMessageRemapperType) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsMessageRemapperType) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsMessageRemapperType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsMessageRemapperType(val *LogsMessageRemapperType) *NullableLogsMessageRemapperType {
	return &NullableLogsMessageRemapperType{value: val, isSet: true}
}

func (v NullableLogsMessageRemapperType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsMessageRemapperType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
