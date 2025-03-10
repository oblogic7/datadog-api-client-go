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

// WidgetLayoutType Layout type of the group.
type WidgetLayoutType string

// List of WidgetLayoutType
const (
	WIDGETLAYOUTTYPE_ORDERED WidgetLayoutType = "ordered"
)

var allowedWidgetLayoutTypeEnumValues = []WidgetLayoutType{
	WIDGETLAYOUTTYPE_ORDERED,
}

func (w *WidgetLayoutType) GetAllowedValues() []WidgetLayoutType {
	return allowedWidgetLayoutTypeEnumValues
}

func (v *WidgetLayoutType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WidgetLayoutType(value)
	return nil
}

// NewWidgetLayoutTypeFromValue returns a pointer to a valid WidgetLayoutType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWidgetLayoutTypeFromValue(v string) (*WidgetLayoutType, error) {
	ev := WidgetLayoutType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WidgetLayoutType: valid values are %v", v, allowedWidgetLayoutTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WidgetLayoutType) IsValid() bool {
	for _, existing := range allowedWidgetLayoutTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WidgetLayoutType value
func (v WidgetLayoutType) Ptr() *WidgetLayoutType {
	return &v
}

type NullableWidgetLayoutType struct {
	value *WidgetLayoutType
	isSet bool
}

func (v NullableWidgetLayoutType) Get() *WidgetLayoutType {
	return v.value
}

func (v *NullableWidgetLayoutType) Set(val *WidgetLayoutType) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetLayoutType) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetLayoutType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetLayoutType(val *WidgetLayoutType) *NullableWidgetLayoutType {
	return &NullableWidgetLayoutType{value: val, isSet: true}
}

func (v NullableWidgetLayoutType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetLayoutType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
