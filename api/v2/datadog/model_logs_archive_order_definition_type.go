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

// LogsArchiveOrderDefinitionType Type of the archive order definition.
type LogsArchiveOrderDefinitionType string

// List of LogsArchiveOrderDefinitionType
const (
	LOGSARCHIVEORDERDEFINITIONTYPE_ARCHIVE_ORDER LogsArchiveOrderDefinitionType = "archive_order"
)

var allowedLogsArchiveOrderDefinitionTypeEnumValues = []LogsArchiveOrderDefinitionType{
	LOGSARCHIVEORDERDEFINITIONTYPE_ARCHIVE_ORDER,
}

func (w *LogsArchiveOrderDefinitionType) GetAllowedValues() []LogsArchiveOrderDefinitionType {
	return allowedLogsArchiveOrderDefinitionTypeEnumValues
}

func (v *LogsArchiveOrderDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsArchiveOrderDefinitionType(value)
	return nil
}

// NewLogsArchiveOrderDefinitionTypeFromValue returns a pointer to a valid LogsArchiveOrderDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLogsArchiveOrderDefinitionTypeFromValue(v string) (*LogsArchiveOrderDefinitionType, error) {
	ev := LogsArchiveOrderDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LogsArchiveOrderDefinitionType: valid values are %v", v, allowedLogsArchiveOrderDefinitionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LogsArchiveOrderDefinitionType) IsValid() bool {
	for _, existing := range allowedLogsArchiveOrderDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsArchiveOrderDefinitionType value
func (v LogsArchiveOrderDefinitionType) Ptr() *LogsArchiveOrderDefinitionType {
	return &v
}

type NullableLogsArchiveOrderDefinitionType struct {
	value *LogsArchiveOrderDefinitionType
	isSet bool
}

func (v NullableLogsArchiveOrderDefinitionType) Get() *LogsArchiveOrderDefinitionType {
	return v.value
}

func (v *NullableLogsArchiveOrderDefinitionType) Set(val *LogsArchiveOrderDefinitionType) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsArchiveOrderDefinitionType) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsArchiveOrderDefinitionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsArchiveOrderDefinitionType(val *LogsArchiveOrderDefinitionType) *NullableLogsArchiveOrderDefinitionType {
	return &NullableLogsArchiveOrderDefinitionType{value: val, isSet: true}
}

func (v NullableLogsArchiveOrderDefinitionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsArchiveOrderDefinitionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
