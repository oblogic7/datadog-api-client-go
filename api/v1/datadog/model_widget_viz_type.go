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

// WidgetVizType Whether to display the Alert Graph as a timeseries or a top list.
type WidgetVizType string

// List of WidgetVizType
const (
	WIDGETVIZTYPE_TIMESERIES WidgetVizType = "timeseries"
	WIDGETVIZTYPE_TOPLIST    WidgetVizType = "toplist"
)

var allowedWidgetVizTypeEnumValues = []WidgetVizType{
	WIDGETVIZTYPE_TIMESERIES,
	WIDGETVIZTYPE_TOPLIST,
}

func (w *WidgetVizType) GetAllowedValues() []WidgetVizType {
	return allowedWidgetVizTypeEnumValues
}

func (v *WidgetVizType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WidgetVizType(value)
	return nil
}

// NewWidgetVizTypeFromValue returns a pointer to a valid WidgetVizType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWidgetVizTypeFromValue(v string) (*WidgetVizType, error) {
	ev := WidgetVizType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WidgetVizType: valid values are %v", v, allowedWidgetVizTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WidgetVizType) IsValid() bool {
	for _, existing := range allowedWidgetVizTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WidgetVizType value
func (v WidgetVizType) Ptr() *WidgetVizType {
	return &v
}

type NullableWidgetVizType struct {
	value *WidgetVizType
	isSet bool
}

func (v NullableWidgetVizType) Get() *WidgetVizType {
	return v.value
}

func (v *NullableWidgetVizType) Set(val *WidgetVizType) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetVizType) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetVizType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetVizType(val *WidgetVizType) *NullableWidgetVizType {
	return &NullableWidgetVizType{value: val, isSet: true}
}

func (v NullableWidgetVizType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetVizType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
