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

// LogsArchiveState The state of the archive.
type LogsArchiveState string

// List of LogsArchiveState
const (
	LOGSARCHIVESTATE_UNKNOWN             LogsArchiveState = "UNKNOWN"
	LOGSARCHIVESTATE_WORKING             LogsArchiveState = "WORKING"
	LOGSARCHIVESTATE_FAILING             LogsArchiveState = "FAILING"
	LOGSARCHIVESTATE_WORKING_AUTH_LEGACY LogsArchiveState = "WORKING_AUTH_LEGACY"
)

var allowedLogsArchiveStateEnumValues = []LogsArchiveState{
	LOGSARCHIVESTATE_UNKNOWN,
	LOGSARCHIVESTATE_WORKING,
	LOGSARCHIVESTATE_FAILING,
	LOGSARCHIVESTATE_WORKING_AUTH_LEGACY,
}

func (w *LogsArchiveState) GetAllowedValues() []LogsArchiveState {
	return allowedLogsArchiveStateEnumValues
}

func (v *LogsArchiveState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsArchiveState(value)
	return nil
}

// NewLogsArchiveStateFromValue returns a pointer to a valid LogsArchiveState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLogsArchiveStateFromValue(v string) (*LogsArchiveState, error) {
	ev := LogsArchiveState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LogsArchiveState: valid values are %v", v, allowedLogsArchiveStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LogsArchiveState) IsValid() bool {
	for _, existing := range allowedLogsArchiveStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsArchiveState value
func (v LogsArchiveState) Ptr() *LogsArchiveState {
	return &v
}

type NullableLogsArchiveState struct {
	value *LogsArchiveState
	isSet bool
}

func (v NullableLogsArchiveState) Get() *LogsArchiveState {
	return v.value
}

func (v *NullableLogsArchiveState) Set(val *LogsArchiveState) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsArchiveState) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsArchiveState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsArchiveState(val *LogsArchiveState) *NullableLogsArchiveState {
	return &NullableLogsArchiveState{value: val, isSet: true}
}

func (v NullableLogsArchiveState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsArchiveState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
