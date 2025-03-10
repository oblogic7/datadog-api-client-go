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

// ScatterPlotWidgetDefinitionType Type of the scatter plot widget.
type ScatterPlotWidgetDefinitionType string

// List of ScatterPlotWidgetDefinitionType
const (
	SCATTERPLOTWIDGETDEFINITIONTYPE_SCATTERPLOT ScatterPlotWidgetDefinitionType = "scatterplot"
)

var allowedScatterPlotWidgetDefinitionTypeEnumValues = []ScatterPlotWidgetDefinitionType{
	SCATTERPLOTWIDGETDEFINITIONTYPE_SCATTERPLOT,
}

func (w *ScatterPlotWidgetDefinitionType) GetAllowedValues() []ScatterPlotWidgetDefinitionType {
	return allowedScatterPlotWidgetDefinitionTypeEnumValues
}

func (v *ScatterPlotWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScatterPlotWidgetDefinitionType(value)
	return nil
}

// NewScatterPlotWidgetDefinitionTypeFromValue returns a pointer to a valid ScatterPlotWidgetDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewScatterPlotWidgetDefinitionTypeFromValue(v string) (*ScatterPlotWidgetDefinitionType, error) {
	ev := ScatterPlotWidgetDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ScatterPlotWidgetDefinitionType: valid values are %v", v, allowedScatterPlotWidgetDefinitionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ScatterPlotWidgetDefinitionType) IsValid() bool {
	for _, existing := range allowedScatterPlotWidgetDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScatterPlotWidgetDefinitionType value
func (v ScatterPlotWidgetDefinitionType) Ptr() *ScatterPlotWidgetDefinitionType {
	return &v
}

type NullableScatterPlotWidgetDefinitionType struct {
	value *ScatterPlotWidgetDefinitionType
	isSet bool
}

func (v NullableScatterPlotWidgetDefinitionType) Get() *ScatterPlotWidgetDefinitionType {
	return v.value
}

func (v *NullableScatterPlotWidgetDefinitionType) Set(val *ScatterPlotWidgetDefinitionType) {
	v.value = val
	v.isSet = true
}

func (v NullableScatterPlotWidgetDefinitionType) IsSet() bool {
	return v.isSet
}

func (v *NullableScatterPlotWidgetDefinitionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScatterPlotWidgetDefinitionType(val *ScatterPlotWidgetDefinitionType) *NullableScatterPlotWidgetDefinitionType {
	return &NullableScatterPlotWidgetDefinitionType{value: val, isSet: true}
}

func (v NullableScatterPlotWidgetDefinitionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScatterPlotWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
