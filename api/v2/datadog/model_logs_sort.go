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

// LogsSort Sort parameters when querying logs.
type LogsSort string

// List of LogsSort
const (
	LOGSSORT_TIMESTAMP_ASCENDING  LogsSort = "timestamp"
	LOGSSORT_TIMESTAMP_DESCENDING LogsSort = "-timestamp"
)

var allowedLogsSortEnumValues = []LogsSort{
	LOGSSORT_TIMESTAMP_ASCENDING,
	LOGSSORT_TIMESTAMP_DESCENDING,
}

func (w *LogsSort) GetAllowedValues() []LogsSort {
	return allowedLogsSortEnumValues
}

func (v *LogsSort) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsSort(value)
	return nil
}

// NewLogsSortFromValue returns a pointer to a valid LogsSort
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLogsSortFromValue(v string) (*LogsSort, error) {
	ev := LogsSort(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LogsSort: valid values are %v", v, allowedLogsSortEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LogsSort) IsValid() bool {
	for _, existing := range allowedLogsSortEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsSort value
func (v LogsSort) Ptr() *LogsSort {
	return &v
}

type NullableLogsSort struct {
	value *LogsSort
	isSet bool
}

func (v NullableLogsSort) Get() *LogsSort {
	return v.value
}

func (v *NullableLogsSort) Set(val *LogsSort) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsSort) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsSort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsSort(val *LogsSort) *NullableLogsSort {
	return &NullableLogsSort{value: val, isSet: true}
}

func (v NullableLogsSort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsSort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
