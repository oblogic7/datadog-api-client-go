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

// WidgetHorizontalAlign Horizontal alignment.
type WidgetHorizontalAlign string

// List of WidgetHorizontalAlign
const (
	WIDGETHORIZONTALALIGN_CENTER WidgetHorizontalAlign = "center"
	WIDGETHORIZONTALALIGN_LEFT   WidgetHorizontalAlign = "left"
	WIDGETHORIZONTALALIGN_RIGHT  WidgetHorizontalAlign = "right"
)

var allowedWidgetHorizontalAlignEnumValues = []WidgetHorizontalAlign{
	WIDGETHORIZONTALALIGN_CENTER,
	WIDGETHORIZONTALALIGN_LEFT,
	WIDGETHORIZONTALALIGN_RIGHT,
}

func (w *WidgetHorizontalAlign) GetAllowedValues() []WidgetHorizontalAlign {
	return allowedWidgetHorizontalAlignEnumValues
}

func (v *WidgetHorizontalAlign) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WidgetHorizontalAlign(value)
	return nil
}

// NewWidgetHorizontalAlignFromValue returns a pointer to a valid WidgetHorizontalAlign
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWidgetHorizontalAlignFromValue(v string) (*WidgetHorizontalAlign, error) {
	ev := WidgetHorizontalAlign(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WidgetHorizontalAlign: valid values are %v", v, allowedWidgetHorizontalAlignEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WidgetHorizontalAlign) IsValid() bool {
	for _, existing := range allowedWidgetHorizontalAlignEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WidgetHorizontalAlign value
func (v WidgetHorizontalAlign) Ptr() *WidgetHorizontalAlign {
	return &v
}

type NullableWidgetHorizontalAlign struct {
	value *WidgetHorizontalAlign
	isSet bool
}

func (v NullableWidgetHorizontalAlign) Get() *WidgetHorizontalAlign {
	return v.value
}

func (v *NullableWidgetHorizontalAlign) Set(val *WidgetHorizontalAlign) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetHorizontalAlign) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetHorizontalAlign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetHorizontalAlign(val *WidgetHorizontalAlign) *NullableWidgetHorizontalAlign {
	return &NullableWidgetHorizontalAlign{value: val, isSet: true}
}

func (v NullableWidgetHorizontalAlign) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetHorizontalAlign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
