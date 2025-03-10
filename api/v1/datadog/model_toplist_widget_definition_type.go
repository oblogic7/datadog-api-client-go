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

// ToplistWidgetDefinitionType Type of the top list widget.
type ToplistWidgetDefinitionType string

// List of ToplistWidgetDefinitionType
const (
	TOPLISTWIDGETDEFINITIONTYPE_TOPLIST ToplistWidgetDefinitionType = "toplist"
)

var allowedToplistWidgetDefinitionTypeEnumValues = []ToplistWidgetDefinitionType{
	TOPLISTWIDGETDEFINITIONTYPE_TOPLIST,
}

func (w *ToplistWidgetDefinitionType) GetAllowedValues() []ToplistWidgetDefinitionType {
	return allowedToplistWidgetDefinitionTypeEnumValues
}

func (v *ToplistWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ToplistWidgetDefinitionType(value)
	return nil
}

// NewToplistWidgetDefinitionTypeFromValue returns a pointer to a valid ToplistWidgetDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToplistWidgetDefinitionTypeFromValue(v string) (*ToplistWidgetDefinitionType, error) {
	ev := ToplistWidgetDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToplistWidgetDefinitionType: valid values are %v", v, allowedToplistWidgetDefinitionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToplistWidgetDefinitionType) IsValid() bool {
	for _, existing := range allowedToplistWidgetDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ToplistWidgetDefinitionType value
func (v ToplistWidgetDefinitionType) Ptr() *ToplistWidgetDefinitionType {
	return &v
}

type NullableToplistWidgetDefinitionType struct {
	value *ToplistWidgetDefinitionType
	isSet bool
}

func (v NullableToplistWidgetDefinitionType) Get() *ToplistWidgetDefinitionType {
	return v.value
}

func (v *NullableToplistWidgetDefinitionType) Set(val *ToplistWidgetDefinitionType) {
	v.value = val
	v.isSet = true
}

func (v NullableToplistWidgetDefinitionType) IsSet() bool {
	return v.isSet
}

func (v *NullableToplistWidgetDefinitionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToplistWidgetDefinitionType(val *ToplistWidgetDefinitionType) *NullableToplistWidgetDefinitionType {
	return &NullableToplistWidgetDefinitionType{value: val, isSet: true}
}

func (v NullableToplistWidgetDefinitionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToplistWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
