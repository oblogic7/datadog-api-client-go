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

// FunnelWidgetDefinitionType Type of funnel widget.
type FunnelWidgetDefinitionType string

// List of FunnelWidgetDefinitionType
const (
	FUNNELWIDGETDEFINITIONTYPE_FUNNEL FunnelWidgetDefinitionType = "funnel"
)

var allowedFunnelWidgetDefinitionTypeEnumValues = []FunnelWidgetDefinitionType{
	FUNNELWIDGETDEFINITIONTYPE_FUNNEL,
}

func (w *FunnelWidgetDefinitionType) GetAllowedValues() []FunnelWidgetDefinitionType {
	return allowedFunnelWidgetDefinitionTypeEnumValues
}

func (v *FunnelWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FunnelWidgetDefinitionType(value)
	return nil
}

// NewFunnelWidgetDefinitionTypeFromValue returns a pointer to a valid FunnelWidgetDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFunnelWidgetDefinitionTypeFromValue(v string) (*FunnelWidgetDefinitionType, error) {
	ev := FunnelWidgetDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FunnelWidgetDefinitionType: valid values are %v", v, allowedFunnelWidgetDefinitionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FunnelWidgetDefinitionType) IsValid() bool {
	for _, existing := range allowedFunnelWidgetDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FunnelWidgetDefinitionType value
func (v FunnelWidgetDefinitionType) Ptr() *FunnelWidgetDefinitionType {
	return &v
}

type NullableFunnelWidgetDefinitionType struct {
	value *FunnelWidgetDefinitionType
	isSet bool
}

func (v NullableFunnelWidgetDefinitionType) Get() *FunnelWidgetDefinitionType {
	return v.value
}

func (v *NullableFunnelWidgetDefinitionType) Set(val *FunnelWidgetDefinitionType) {
	v.value = val
	v.isSet = true
}

func (v NullableFunnelWidgetDefinitionType) IsSet() bool {
	return v.isSet
}

func (v *NullableFunnelWidgetDefinitionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunnelWidgetDefinitionType(val *FunnelWidgetDefinitionType) *NullableFunnelWidgetDefinitionType {
	return &NullableFunnelWidgetDefinitionType{value: val, isSet: true}
}

func (v NullableFunnelWidgetDefinitionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunnelWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
