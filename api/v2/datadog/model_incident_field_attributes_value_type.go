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

// IncidentFieldAttributesValueType Type of the multiple value field definitions.
type IncidentFieldAttributesValueType string

// List of IncidentFieldAttributesValueType
const (
	INCIDENTFIELDATTRIBUTESVALUETYPE_MULTISELECT  IncidentFieldAttributesValueType = "multiselect"
	INCIDENTFIELDATTRIBUTESVALUETYPE_TEXTARRAY    IncidentFieldAttributesValueType = "textarray"
	INCIDENTFIELDATTRIBUTESVALUETYPE_METRICTAG    IncidentFieldAttributesValueType = "metrictag"
	INCIDENTFIELDATTRIBUTESVALUETYPE_AUTOCOMPLETE IncidentFieldAttributesValueType = "autocomplete"
)

var allowedIncidentFieldAttributesValueTypeEnumValues = []IncidentFieldAttributesValueType{
	INCIDENTFIELDATTRIBUTESVALUETYPE_MULTISELECT,
	INCIDENTFIELDATTRIBUTESVALUETYPE_TEXTARRAY,
	INCIDENTFIELDATTRIBUTESVALUETYPE_METRICTAG,
	INCIDENTFIELDATTRIBUTESVALUETYPE_AUTOCOMPLETE,
}

func (w *IncidentFieldAttributesValueType) GetAllowedValues() []IncidentFieldAttributesValueType {
	return allowedIncidentFieldAttributesValueTypeEnumValues
}

func (v *IncidentFieldAttributesValueType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentFieldAttributesValueType(value)
	return nil
}

// NewIncidentFieldAttributesValueTypeFromValue returns a pointer to a valid IncidentFieldAttributesValueType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIncidentFieldAttributesValueTypeFromValue(v string) (*IncidentFieldAttributesValueType, error) {
	ev := IncidentFieldAttributesValueType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IncidentFieldAttributesValueType: valid values are %v", v, allowedIncidentFieldAttributesValueTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IncidentFieldAttributesValueType) IsValid() bool {
	for _, existing := range allowedIncidentFieldAttributesValueTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentFieldAttributesValueType value
func (v IncidentFieldAttributesValueType) Ptr() *IncidentFieldAttributesValueType {
	return &v
}

type NullableIncidentFieldAttributesValueType struct {
	value *IncidentFieldAttributesValueType
	isSet bool
}

func (v NullableIncidentFieldAttributesValueType) Get() *IncidentFieldAttributesValueType {
	return v.value
}

func (v *NullableIncidentFieldAttributesValueType) Set(val *IncidentFieldAttributesValueType) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentFieldAttributesValueType) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentFieldAttributesValueType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentFieldAttributesValueType(val *IncidentFieldAttributesValueType) *NullableIncidentFieldAttributesValueType {
	return &NullableIncidentFieldAttributesValueType{value: val, isSet: true}
}

func (v NullableIncidentFieldAttributesValueType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentFieldAttributesValueType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
