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

// SunburstWidgetLegendInlineAutomaticType Whether to show the legend inline or let it be automatically generated.
type SunburstWidgetLegendInlineAutomaticType string

// List of SunburstWidgetLegendInlineAutomaticType
const (
	SUNBURSTWIDGETLEGENDINLINEAUTOMATICTYPE_INLINE    SunburstWidgetLegendInlineAutomaticType = "inline"
	SUNBURSTWIDGETLEGENDINLINEAUTOMATICTYPE_AUTOMATIC SunburstWidgetLegendInlineAutomaticType = "automatic"
)

var allowedSunburstWidgetLegendInlineAutomaticTypeEnumValues = []SunburstWidgetLegendInlineAutomaticType{
	SUNBURSTWIDGETLEGENDINLINEAUTOMATICTYPE_INLINE,
	SUNBURSTWIDGETLEGENDINLINEAUTOMATICTYPE_AUTOMATIC,
}

func (w *SunburstWidgetLegendInlineAutomaticType) GetAllowedValues() []SunburstWidgetLegendInlineAutomaticType {
	return allowedSunburstWidgetLegendInlineAutomaticTypeEnumValues
}

func (v *SunburstWidgetLegendInlineAutomaticType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SunburstWidgetLegendInlineAutomaticType(value)
	return nil
}

// NewSunburstWidgetLegendInlineAutomaticTypeFromValue returns a pointer to a valid SunburstWidgetLegendInlineAutomaticType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSunburstWidgetLegendInlineAutomaticTypeFromValue(v string) (*SunburstWidgetLegendInlineAutomaticType, error) {
	ev := SunburstWidgetLegendInlineAutomaticType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SunburstWidgetLegendInlineAutomaticType: valid values are %v", v, allowedSunburstWidgetLegendInlineAutomaticTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SunburstWidgetLegendInlineAutomaticType) IsValid() bool {
	for _, existing := range allowedSunburstWidgetLegendInlineAutomaticTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SunburstWidgetLegendInlineAutomaticType value
func (v SunburstWidgetLegendInlineAutomaticType) Ptr() *SunburstWidgetLegendInlineAutomaticType {
	return &v
}

type NullableSunburstWidgetLegendInlineAutomaticType struct {
	value *SunburstWidgetLegendInlineAutomaticType
	isSet bool
}

func (v NullableSunburstWidgetLegendInlineAutomaticType) Get() *SunburstWidgetLegendInlineAutomaticType {
	return v.value
}

func (v *NullableSunburstWidgetLegendInlineAutomaticType) Set(val *SunburstWidgetLegendInlineAutomaticType) {
	v.value = val
	v.isSet = true
}

func (v NullableSunburstWidgetLegendInlineAutomaticType) IsSet() bool {
	return v.isSet
}

func (v *NullableSunburstWidgetLegendInlineAutomaticType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSunburstWidgetLegendInlineAutomaticType(val *SunburstWidgetLegendInlineAutomaticType) *NullableSunburstWidgetLegendInlineAutomaticType {
	return &NullableSunburstWidgetLegendInlineAutomaticType{value: val, isSet: true}
}

func (v NullableSunburstWidgetLegendInlineAutomaticType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSunburstWidgetLegendInlineAutomaticType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
