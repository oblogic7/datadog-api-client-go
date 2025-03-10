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

// WidgetChangeType Show the absolute or the relative change.
type WidgetChangeType string

// List of WidgetChangeType
const (
	WIDGETCHANGETYPE_ABSOLUTE WidgetChangeType = "absolute"
	WIDGETCHANGETYPE_RELATIVE WidgetChangeType = "relative"
)

var allowedWidgetChangeTypeEnumValues = []WidgetChangeType{
	WIDGETCHANGETYPE_ABSOLUTE,
	WIDGETCHANGETYPE_RELATIVE,
}

func (w *WidgetChangeType) GetAllowedValues() []WidgetChangeType {
	return allowedWidgetChangeTypeEnumValues
}

func (v *WidgetChangeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WidgetChangeType(value)
	return nil
}

// NewWidgetChangeTypeFromValue returns a pointer to a valid WidgetChangeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWidgetChangeTypeFromValue(v string) (*WidgetChangeType, error) {
	ev := WidgetChangeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WidgetChangeType: valid values are %v", v, allowedWidgetChangeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WidgetChangeType) IsValid() bool {
	for _, existing := range allowedWidgetChangeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WidgetChangeType value
func (v WidgetChangeType) Ptr() *WidgetChangeType {
	return &v
}

type NullableWidgetChangeType struct {
	value *WidgetChangeType
	isSet bool
}

func (v NullableWidgetChangeType) Get() *WidgetChangeType {
	return v.value
}

func (v *NullableWidgetChangeType) Set(val *WidgetChangeType) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetChangeType) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetChangeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetChangeType(val *WidgetChangeType) *NullableWidgetChangeType {
	return &NullableWidgetChangeType{value: val, isSet: true}
}

func (v NullableWidgetChangeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetChangeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
