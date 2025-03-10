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

// LogsSortOrder The order to use, ascending or descending
type LogsSortOrder string

// List of LogsSortOrder
const (
	LOGSSORTORDER_ASCENDING  LogsSortOrder = "asc"
	LOGSSORTORDER_DESCENDING LogsSortOrder = "desc"
)

var allowedLogsSortOrderEnumValues = []LogsSortOrder{
	LOGSSORTORDER_ASCENDING,
	LOGSSORTORDER_DESCENDING,
}

func (w *LogsSortOrder) GetAllowedValues() []LogsSortOrder {
	return allowedLogsSortOrderEnumValues
}

func (v *LogsSortOrder) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsSortOrder(value)
	return nil
}

// NewLogsSortOrderFromValue returns a pointer to a valid LogsSortOrder
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLogsSortOrderFromValue(v string) (*LogsSortOrder, error) {
	ev := LogsSortOrder(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LogsSortOrder: valid values are %v", v, allowedLogsSortOrderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LogsSortOrder) IsValid() bool {
	for _, existing := range allowedLogsSortOrderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsSortOrder value
func (v LogsSortOrder) Ptr() *LogsSortOrder {
	return &v
}

type NullableLogsSortOrder struct {
	value *LogsSortOrder
	isSet bool
}

func (v NullableLogsSortOrder) Get() *LogsSortOrder {
	return v.value
}

func (v *NullableLogsSortOrder) Set(val *LogsSortOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsSortOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsSortOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsSortOrder(val *LogsSortOrder) *NullableLogsSortOrder {
	return &NullableLogsSortOrder{value: val, isSet: true}
}

func (v NullableLogsSortOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsSortOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
