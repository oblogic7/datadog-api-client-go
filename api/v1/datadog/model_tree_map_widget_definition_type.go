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

// TreeMapWidgetDefinitionType Type of the treemap widget.
type TreeMapWidgetDefinitionType string

// List of TreeMapWidgetDefinitionType
const (
	TREEMAPWIDGETDEFINITIONTYPE_TREEMAP TreeMapWidgetDefinitionType = "treemap"
)

var allowedTreeMapWidgetDefinitionTypeEnumValues = []TreeMapWidgetDefinitionType{
	TREEMAPWIDGETDEFINITIONTYPE_TREEMAP,
}

func (w *TreeMapWidgetDefinitionType) GetAllowedValues() []TreeMapWidgetDefinitionType {
	return allowedTreeMapWidgetDefinitionTypeEnumValues
}

func (v *TreeMapWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TreeMapWidgetDefinitionType(value)
	return nil
}

// NewTreeMapWidgetDefinitionTypeFromValue returns a pointer to a valid TreeMapWidgetDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTreeMapWidgetDefinitionTypeFromValue(v string) (*TreeMapWidgetDefinitionType, error) {
	ev := TreeMapWidgetDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TreeMapWidgetDefinitionType: valid values are %v", v, allowedTreeMapWidgetDefinitionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TreeMapWidgetDefinitionType) IsValid() bool {
	for _, existing := range allowedTreeMapWidgetDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TreeMapWidgetDefinitionType value
func (v TreeMapWidgetDefinitionType) Ptr() *TreeMapWidgetDefinitionType {
	return &v
}

type NullableTreeMapWidgetDefinitionType struct {
	value *TreeMapWidgetDefinitionType
	isSet bool
}

func (v NullableTreeMapWidgetDefinitionType) Get() *TreeMapWidgetDefinitionType {
	return v.value
}

func (v *NullableTreeMapWidgetDefinitionType) Set(val *TreeMapWidgetDefinitionType) {
	v.value = val
	v.isSet = true
}

func (v NullableTreeMapWidgetDefinitionType) IsSet() bool {
	return v.isSet
}

func (v *NullableTreeMapWidgetDefinitionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTreeMapWidgetDefinitionType(val *TreeMapWidgetDefinitionType) *NullableTreeMapWidgetDefinitionType {
	return &NullableTreeMapWidgetDefinitionType{value: val, isSet: true}
}

func (v NullableTreeMapWidgetDefinitionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTreeMapWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
