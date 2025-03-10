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

// WidgetOrderBy What to order by.
type WidgetOrderBy string

// List of WidgetOrderBy
const (
	WIDGETORDERBY_CHANGE  WidgetOrderBy = "change"
	WIDGETORDERBY_NAME    WidgetOrderBy = "name"
	WIDGETORDERBY_PRESENT WidgetOrderBy = "present"
	WIDGETORDERBY_PAST    WidgetOrderBy = "past"
)

var allowedWidgetOrderByEnumValues = []WidgetOrderBy{
	WIDGETORDERBY_CHANGE,
	WIDGETORDERBY_NAME,
	WIDGETORDERBY_PRESENT,
	WIDGETORDERBY_PAST,
}

func (w *WidgetOrderBy) GetAllowedValues() []WidgetOrderBy {
	return allowedWidgetOrderByEnumValues
}

func (v *WidgetOrderBy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WidgetOrderBy(value)
	return nil
}

// NewWidgetOrderByFromValue returns a pointer to a valid WidgetOrderBy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWidgetOrderByFromValue(v string) (*WidgetOrderBy, error) {
	ev := WidgetOrderBy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WidgetOrderBy: valid values are %v", v, allowedWidgetOrderByEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WidgetOrderBy) IsValid() bool {
	for _, existing := range allowedWidgetOrderByEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WidgetOrderBy value
func (v WidgetOrderBy) Ptr() *WidgetOrderBy {
	return &v
}

type NullableWidgetOrderBy struct {
	value *WidgetOrderBy
	isSet bool
}

func (v NullableWidgetOrderBy) Get() *WidgetOrderBy {
	return v.value
}

func (v *NullableWidgetOrderBy) Set(val *WidgetOrderBy) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetOrderBy) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetOrderBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetOrderBy(val *WidgetOrderBy) *NullableWidgetOrderBy {
	return &NullableWidgetOrderBy{value: val, isSet: true}
}

func (v NullableWidgetOrderBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetOrderBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
