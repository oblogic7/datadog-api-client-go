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

// LogsGeoIPParserType Type of GeoIP parser.
type LogsGeoIPParserType string

// List of LogsGeoIPParserType
const (
	LOGSGEOIPPARSERTYPE_GEO_IP_PARSER LogsGeoIPParserType = "geo-ip-parser"
)

var allowedLogsGeoIPParserTypeEnumValues = []LogsGeoIPParserType{
	LOGSGEOIPPARSERTYPE_GEO_IP_PARSER,
}

func (w *LogsGeoIPParserType) GetAllowedValues() []LogsGeoIPParserType {
	return allowedLogsGeoIPParserTypeEnumValues
}

func (v *LogsGeoIPParserType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsGeoIPParserType(value)
	return nil
}

// NewLogsGeoIPParserTypeFromValue returns a pointer to a valid LogsGeoIPParserType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLogsGeoIPParserTypeFromValue(v string) (*LogsGeoIPParserType, error) {
	ev := LogsGeoIPParserType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LogsGeoIPParserType: valid values are %v", v, allowedLogsGeoIPParserTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LogsGeoIPParserType) IsValid() bool {
	for _, existing := range allowedLogsGeoIPParserTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsGeoIPParserType value
func (v LogsGeoIPParserType) Ptr() *LogsGeoIPParserType {
	return &v
}

type NullableLogsGeoIPParserType struct {
	value *LogsGeoIPParserType
	isSet bool
}

func (v NullableLogsGeoIPParserType) Get() *LogsGeoIPParserType {
	return v.value
}

func (v *NullableLogsGeoIPParserType) Set(val *LogsGeoIPParserType) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsGeoIPParserType) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsGeoIPParserType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsGeoIPParserType(val *LogsGeoIPParserType) *NullableLogsGeoIPParserType {
	return &NullableLogsGeoIPParserType{value: val, isSet: true}
}

func (v NullableLogsGeoIPParserType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsGeoIPParserType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
