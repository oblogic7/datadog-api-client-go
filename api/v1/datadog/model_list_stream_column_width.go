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

// ListStreamColumnWidth Widget column width.
type ListStreamColumnWidth string

// List of ListStreamColumnWidth
const (
	LISTSTREAMCOLUMNWIDTH_AUTO    ListStreamColumnWidth = "auto"
	LISTSTREAMCOLUMNWIDTH_COMPACT ListStreamColumnWidth = "compact"
	LISTSTREAMCOLUMNWIDTH_FULL    ListStreamColumnWidth = "full"
)

var allowedListStreamColumnWidthEnumValues = []ListStreamColumnWidth{
	LISTSTREAMCOLUMNWIDTH_AUTO,
	LISTSTREAMCOLUMNWIDTH_COMPACT,
	LISTSTREAMCOLUMNWIDTH_FULL,
}

func (w *ListStreamColumnWidth) GetAllowedValues() []ListStreamColumnWidth {
	return allowedListStreamColumnWidthEnumValues
}

func (v *ListStreamColumnWidth) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ListStreamColumnWidth(value)
	return nil
}

// NewListStreamColumnWidthFromValue returns a pointer to a valid ListStreamColumnWidth
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewListStreamColumnWidthFromValue(v string) (*ListStreamColumnWidth, error) {
	ev := ListStreamColumnWidth(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ListStreamColumnWidth: valid values are %v", v, allowedListStreamColumnWidthEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ListStreamColumnWidth) IsValid() bool {
	for _, existing := range allowedListStreamColumnWidthEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ListStreamColumnWidth value
func (v ListStreamColumnWidth) Ptr() *ListStreamColumnWidth {
	return &v
}

type NullableListStreamColumnWidth struct {
	value *ListStreamColumnWidth
	isSet bool
}

func (v NullableListStreamColumnWidth) Get() *ListStreamColumnWidth {
	return v.value
}

func (v *NullableListStreamColumnWidth) Set(val *ListStreamColumnWidth) {
	v.value = val
	v.isSet = true
}

func (v NullableListStreamColumnWidth) IsSet() bool {
	return v.isSet
}

func (v *NullableListStreamColumnWidth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListStreamColumnWidth(val *ListStreamColumnWidth) *NullableListStreamColumnWidth {
	return &NullableListStreamColumnWidth{value: val, isSet: true}
}

func (v NullableListStreamColumnWidth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListStreamColumnWidth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
