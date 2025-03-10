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

// LogsUserAgentParserType Type of logs User-Agent parser.
type LogsUserAgentParserType string

// List of LogsUserAgentParserType
const (
	LOGSUSERAGENTPARSERTYPE_USER_AGENT_PARSER LogsUserAgentParserType = "user-agent-parser"
)

var allowedLogsUserAgentParserTypeEnumValues = []LogsUserAgentParserType{
	LOGSUSERAGENTPARSERTYPE_USER_AGENT_PARSER,
}

func (w *LogsUserAgentParserType) GetAllowedValues() []LogsUserAgentParserType {
	return allowedLogsUserAgentParserTypeEnumValues
}

func (v *LogsUserAgentParserType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsUserAgentParserType(value)
	return nil
}

// NewLogsUserAgentParserTypeFromValue returns a pointer to a valid LogsUserAgentParserType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLogsUserAgentParserTypeFromValue(v string) (*LogsUserAgentParserType, error) {
	ev := LogsUserAgentParserType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LogsUserAgentParserType: valid values are %v", v, allowedLogsUserAgentParserTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LogsUserAgentParserType) IsValid() bool {
	for _, existing := range allowedLogsUserAgentParserTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsUserAgentParserType value
func (v LogsUserAgentParserType) Ptr() *LogsUserAgentParserType {
	return &v
}

type NullableLogsUserAgentParserType struct {
	value *LogsUserAgentParserType
	isSet bool
}

func (v NullableLogsUserAgentParserType) Get() *LogsUserAgentParserType {
	return v.value
}

func (v *NullableLogsUserAgentParserType) Set(val *LogsUserAgentParserType) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsUserAgentParserType) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsUserAgentParserType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsUserAgentParserType(val *LogsUserAgentParserType) *NullableLogsUserAgentParserType {
	return &NullableLogsUserAgentParserType{value: val, isSet: true}
}

func (v NullableLogsUserAgentParserType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsUserAgentParserType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
