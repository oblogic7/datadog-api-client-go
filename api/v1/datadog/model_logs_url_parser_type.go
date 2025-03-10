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

// LogsURLParserType Type of logs URL parser.
type LogsURLParserType string

// List of LogsURLParserType
const (
	LOGSURLPARSERTYPE_URL_PARSER LogsURLParserType = "url-parser"
)

var allowedLogsURLParserTypeEnumValues = []LogsURLParserType{
	LOGSURLPARSERTYPE_URL_PARSER,
}

func (w *LogsURLParserType) GetAllowedValues() []LogsURLParserType {
	return allowedLogsURLParserTypeEnumValues
}

func (v *LogsURLParserType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsURLParserType(value)
	return nil
}

// NewLogsURLParserTypeFromValue returns a pointer to a valid LogsURLParserType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLogsURLParserTypeFromValue(v string) (*LogsURLParserType, error) {
	ev := LogsURLParserType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LogsURLParserType: valid values are %v", v, allowedLogsURLParserTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LogsURLParserType) IsValid() bool {
	for _, existing := range allowedLogsURLParserTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsURLParserType value
func (v LogsURLParserType) Ptr() *LogsURLParserType {
	return &v
}

type NullableLogsURLParserType struct {
	value *LogsURLParserType
	isSet bool
}

func (v NullableLogsURLParserType) Get() *LogsURLParserType {
	return v.value
}

func (v *NullableLogsURLParserType) Set(val *LogsURLParserType) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsURLParserType) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsURLParserType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsURLParserType(val *LogsURLParserType) *NullableLogsURLParserType {
	return &NullableLogsURLParserType{value: val, isSet: true}
}

func (v NullableLogsURLParserType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsURLParserType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
