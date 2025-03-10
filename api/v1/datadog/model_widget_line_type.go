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

// WidgetLineType Type of lines displayed.
type WidgetLineType string

// List of WidgetLineType
const (
	WIDGETLINETYPE_DASHED WidgetLineType = "dashed"
	WIDGETLINETYPE_DOTTED WidgetLineType = "dotted"
	WIDGETLINETYPE_SOLID  WidgetLineType = "solid"
)

var allowedWidgetLineTypeEnumValues = []WidgetLineType{
	WIDGETLINETYPE_DASHED,
	WIDGETLINETYPE_DOTTED,
	WIDGETLINETYPE_SOLID,
}

func (w *WidgetLineType) GetAllowedValues() []WidgetLineType {
	return allowedWidgetLineTypeEnumValues
}

func (v *WidgetLineType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WidgetLineType(value)
	return nil
}

// NewWidgetLineTypeFromValue returns a pointer to a valid WidgetLineType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWidgetLineTypeFromValue(v string) (*WidgetLineType, error) {
	ev := WidgetLineType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WidgetLineType: valid values are %v", v, allowedWidgetLineTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WidgetLineType) IsValid() bool {
	for _, existing := range allowedWidgetLineTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WidgetLineType value
func (v WidgetLineType) Ptr() *WidgetLineType {
	return &v
}

type NullableWidgetLineType struct {
	value *WidgetLineType
	isSet bool
}

func (v NullableWidgetLineType) Get() *WidgetLineType {
	return v.value
}

func (v *NullableWidgetLineType) Set(val *WidgetLineType) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetLineType) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetLineType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetLineType(val *WidgetLineType) *NullableWidgetLineType {
	return &NullableWidgetLineType{value: val, isSet: true}
}

func (v NullableWidgetLineType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetLineType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
