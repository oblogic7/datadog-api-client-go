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

// LogsArchiveDestinationAzureType Type of the Azure archive destination.
type LogsArchiveDestinationAzureType string

// List of LogsArchiveDestinationAzureType
const (
	LOGSARCHIVEDESTINATIONAZURETYPE_AZURE LogsArchiveDestinationAzureType = "azure"
)

var allowedLogsArchiveDestinationAzureTypeEnumValues = []LogsArchiveDestinationAzureType{
	LOGSARCHIVEDESTINATIONAZURETYPE_AZURE,
}

func (w *LogsArchiveDestinationAzureType) GetAllowedValues() []LogsArchiveDestinationAzureType {
	return allowedLogsArchiveDestinationAzureTypeEnumValues
}

func (v *LogsArchiveDestinationAzureType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsArchiveDestinationAzureType(value)
	return nil
}

// NewLogsArchiveDestinationAzureTypeFromValue returns a pointer to a valid LogsArchiveDestinationAzureType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLogsArchiveDestinationAzureTypeFromValue(v string) (*LogsArchiveDestinationAzureType, error) {
	ev := LogsArchiveDestinationAzureType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LogsArchiveDestinationAzureType: valid values are %v", v, allowedLogsArchiveDestinationAzureTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LogsArchiveDestinationAzureType) IsValid() bool {
	for _, existing := range allowedLogsArchiveDestinationAzureTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsArchiveDestinationAzureType value
func (v LogsArchiveDestinationAzureType) Ptr() *LogsArchiveDestinationAzureType {
	return &v
}

type NullableLogsArchiveDestinationAzureType struct {
	value *LogsArchiveDestinationAzureType
	isSet bool
}

func (v NullableLogsArchiveDestinationAzureType) Get() *LogsArchiveDestinationAzureType {
	return v.value
}

func (v *NullableLogsArchiveDestinationAzureType) Set(val *LogsArchiveDestinationAzureType) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsArchiveDestinationAzureType) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsArchiveDestinationAzureType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsArchiveDestinationAzureType(val *LogsArchiveDestinationAzureType) *NullableLogsArchiveDestinationAzureType {
	return &NullableLogsArchiveDestinationAzureType{value: val, isSet: true}
}

func (v NullableLogsArchiveDestinationAzureType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsArchiveDestinationAzureType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
