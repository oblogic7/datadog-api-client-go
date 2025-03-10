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

// ListStreamWidgetDefinitionType Type of the list stream widget.
type ListStreamWidgetDefinitionType string

// List of ListStreamWidgetDefinitionType
const (
	LISTSTREAMWIDGETDEFINITIONTYPE_LIST_STREAM ListStreamWidgetDefinitionType = "list_stream"
)

var allowedListStreamWidgetDefinitionTypeEnumValues = []ListStreamWidgetDefinitionType{
	LISTSTREAMWIDGETDEFINITIONTYPE_LIST_STREAM,
}

func (w *ListStreamWidgetDefinitionType) GetAllowedValues() []ListStreamWidgetDefinitionType {
	return allowedListStreamWidgetDefinitionTypeEnumValues
}

func (v *ListStreamWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ListStreamWidgetDefinitionType(value)
	return nil
}

// NewListStreamWidgetDefinitionTypeFromValue returns a pointer to a valid ListStreamWidgetDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewListStreamWidgetDefinitionTypeFromValue(v string) (*ListStreamWidgetDefinitionType, error) {
	ev := ListStreamWidgetDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ListStreamWidgetDefinitionType: valid values are %v", v, allowedListStreamWidgetDefinitionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ListStreamWidgetDefinitionType) IsValid() bool {
	for _, existing := range allowedListStreamWidgetDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ListStreamWidgetDefinitionType value
func (v ListStreamWidgetDefinitionType) Ptr() *ListStreamWidgetDefinitionType {
	return &v
}

type NullableListStreamWidgetDefinitionType struct {
	value *ListStreamWidgetDefinitionType
	isSet bool
}

func (v NullableListStreamWidgetDefinitionType) Get() *ListStreamWidgetDefinitionType {
	return v.value
}

func (v *NullableListStreamWidgetDefinitionType) Set(val *ListStreamWidgetDefinitionType) {
	v.value = val
	v.isSet = true
}

func (v NullableListStreamWidgetDefinitionType) IsSet() bool {
	return v.isSet
}

func (v *NullableListStreamWidgetDefinitionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListStreamWidgetDefinitionType(val *ListStreamWidgetDefinitionType) *NullableListStreamWidgetDefinitionType {
	return &NullableListStreamWidgetDefinitionType{value: val, isSet: true}
}

func (v NullableListStreamWidgetDefinitionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListStreamWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
