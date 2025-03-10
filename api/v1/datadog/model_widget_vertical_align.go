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

// WidgetVerticalAlign Vertical alignment.
type WidgetVerticalAlign string

// List of WidgetVerticalAlign
const (
	WIDGETVERTICALALIGN_CENTER WidgetVerticalAlign = "center"
	WIDGETVERTICALALIGN_TOP    WidgetVerticalAlign = "top"
	WIDGETVERTICALALIGN_BOTTOM WidgetVerticalAlign = "bottom"
)

var allowedWidgetVerticalAlignEnumValues = []WidgetVerticalAlign{
	WIDGETVERTICALALIGN_CENTER,
	WIDGETVERTICALALIGN_TOP,
	WIDGETVERTICALALIGN_BOTTOM,
}

func (w *WidgetVerticalAlign) GetAllowedValues() []WidgetVerticalAlign {
	return allowedWidgetVerticalAlignEnumValues
}

func (v *WidgetVerticalAlign) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WidgetVerticalAlign(value)
	return nil
}

// NewWidgetVerticalAlignFromValue returns a pointer to a valid WidgetVerticalAlign
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWidgetVerticalAlignFromValue(v string) (*WidgetVerticalAlign, error) {
	ev := WidgetVerticalAlign(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WidgetVerticalAlign: valid values are %v", v, allowedWidgetVerticalAlignEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WidgetVerticalAlign) IsValid() bool {
	for _, existing := range allowedWidgetVerticalAlignEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WidgetVerticalAlign value
func (v WidgetVerticalAlign) Ptr() *WidgetVerticalAlign {
	return &v
}

type NullableWidgetVerticalAlign struct {
	value *WidgetVerticalAlign
	isSet bool
}

func (v NullableWidgetVerticalAlign) Get() *WidgetVerticalAlign {
	return v.value
}

func (v *NullableWidgetVerticalAlign) Set(val *WidgetVerticalAlign) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetVerticalAlign) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetVerticalAlign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetVerticalAlign(val *WidgetVerticalAlign) *NullableWidgetVerticalAlign {
	return &NullableWidgetVerticalAlign{value: val, isSet: true}
}

func (v NullableWidgetVerticalAlign) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetVerticalAlign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
