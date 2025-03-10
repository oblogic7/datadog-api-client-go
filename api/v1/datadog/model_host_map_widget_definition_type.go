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

// HostMapWidgetDefinitionType Type of the host map widget.
type HostMapWidgetDefinitionType string

// List of HostMapWidgetDefinitionType
const (
	HOSTMAPWIDGETDEFINITIONTYPE_HOSTMAP HostMapWidgetDefinitionType = "hostmap"
)

var allowedHostMapWidgetDefinitionTypeEnumValues = []HostMapWidgetDefinitionType{
	HOSTMAPWIDGETDEFINITIONTYPE_HOSTMAP,
}

func (w *HostMapWidgetDefinitionType) GetAllowedValues() []HostMapWidgetDefinitionType {
	return allowedHostMapWidgetDefinitionTypeEnumValues
}

func (v *HostMapWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HostMapWidgetDefinitionType(value)
	return nil
}

// NewHostMapWidgetDefinitionTypeFromValue returns a pointer to a valid HostMapWidgetDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHostMapWidgetDefinitionTypeFromValue(v string) (*HostMapWidgetDefinitionType, error) {
	ev := HostMapWidgetDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for HostMapWidgetDefinitionType: valid values are %v", v, allowedHostMapWidgetDefinitionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HostMapWidgetDefinitionType) IsValid() bool {
	for _, existing := range allowedHostMapWidgetDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HostMapWidgetDefinitionType value
func (v HostMapWidgetDefinitionType) Ptr() *HostMapWidgetDefinitionType {
	return &v
}

type NullableHostMapWidgetDefinitionType struct {
	value *HostMapWidgetDefinitionType
	isSet bool
}

func (v NullableHostMapWidgetDefinitionType) Get() *HostMapWidgetDefinitionType {
	return v.value
}

func (v *NullableHostMapWidgetDefinitionType) Set(val *HostMapWidgetDefinitionType) {
	v.value = val
	v.isSet = true
}

func (v NullableHostMapWidgetDefinitionType) IsSet() bool {
	return v.isSet
}

func (v *NullableHostMapWidgetDefinitionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostMapWidgetDefinitionType(val *HostMapWidgetDefinitionType) *NullableHostMapWidgetDefinitionType {
	return &NullableHostMapWidgetDefinitionType{value: val, isSet: true}
}

func (v NullableHostMapWidgetDefinitionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostMapWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
