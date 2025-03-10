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

// ChangeWidgetDefinitionType Type of the change widget.
type ChangeWidgetDefinitionType string

// List of ChangeWidgetDefinitionType
const (
	CHANGEWIDGETDEFINITIONTYPE_CHANGE ChangeWidgetDefinitionType = "change"
)

var allowedChangeWidgetDefinitionTypeEnumValues = []ChangeWidgetDefinitionType{
	CHANGEWIDGETDEFINITIONTYPE_CHANGE,
}

func (w *ChangeWidgetDefinitionType) GetAllowedValues() []ChangeWidgetDefinitionType {
	return allowedChangeWidgetDefinitionTypeEnumValues
}

func (v *ChangeWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ChangeWidgetDefinitionType(value)
	return nil
}

// NewChangeWidgetDefinitionTypeFromValue returns a pointer to a valid ChangeWidgetDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewChangeWidgetDefinitionTypeFromValue(v string) (*ChangeWidgetDefinitionType, error) {
	ev := ChangeWidgetDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ChangeWidgetDefinitionType: valid values are %v", v, allowedChangeWidgetDefinitionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ChangeWidgetDefinitionType) IsValid() bool {
	for _, existing := range allowedChangeWidgetDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ChangeWidgetDefinitionType value
func (v ChangeWidgetDefinitionType) Ptr() *ChangeWidgetDefinitionType {
	return &v
}

type NullableChangeWidgetDefinitionType struct {
	value *ChangeWidgetDefinitionType
	isSet bool
}

func (v NullableChangeWidgetDefinitionType) Get() *ChangeWidgetDefinitionType {
	return v.value
}

func (v *NullableChangeWidgetDefinitionType) Set(val *ChangeWidgetDefinitionType) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeWidgetDefinitionType) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeWidgetDefinitionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeWidgetDefinitionType(val *ChangeWidgetDefinitionType) *NullableChangeWidgetDefinitionType {
	return &NullableChangeWidgetDefinitionType{value: val, isSet: true}
}

func (v NullableChangeWidgetDefinitionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
