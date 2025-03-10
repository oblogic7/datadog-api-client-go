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

// WidgetLineWidth Width of line displayed.
type WidgetLineWidth string

// List of WidgetLineWidth
const (
	WIDGETLINEWIDTH_NORMAL WidgetLineWidth = "normal"
	WIDGETLINEWIDTH_THICK  WidgetLineWidth = "thick"
	WIDGETLINEWIDTH_THIN   WidgetLineWidth = "thin"
)

var allowedWidgetLineWidthEnumValues = []WidgetLineWidth{
	WIDGETLINEWIDTH_NORMAL,
	WIDGETLINEWIDTH_THICK,
	WIDGETLINEWIDTH_THIN,
}

func (w *WidgetLineWidth) GetAllowedValues() []WidgetLineWidth {
	return allowedWidgetLineWidthEnumValues
}

func (v *WidgetLineWidth) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WidgetLineWidth(value)
	return nil
}

// NewWidgetLineWidthFromValue returns a pointer to a valid WidgetLineWidth
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWidgetLineWidthFromValue(v string) (*WidgetLineWidth, error) {
	ev := WidgetLineWidth(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WidgetLineWidth: valid values are %v", v, allowedWidgetLineWidthEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WidgetLineWidth) IsValid() bool {
	for _, existing := range allowedWidgetLineWidthEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WidgetLineWidth value
func (v WidgetLineWidth) Ptr() *WidgetLineWidth {
	return &v
}

type NullableWidgetLineWidth struct {
	value *WidgetLineWidth
	isSet bool
}

func (v NullableWidgetLineWidth) Get() *WidgetLineWidth {
	return v.value
}

func (v *NullableWidgetLineWidth) Set(val *WidgetLineWidth) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetLineWidth) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetLineWidth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetLineWidth(val *WidgetLineWidth) *NullableWidgetLineWidth {
	return &NullableWidgetLineWidth{value: val, isSet: true}
}

func (v NullableWidgetLineWidth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetLineWidth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
