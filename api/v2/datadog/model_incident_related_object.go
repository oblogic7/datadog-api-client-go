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

// IncidentRelatedObject Object related to an incident.
type IncidentRelatedObject string

// List of IncidentRelatedObject
const (
	INCIDENTRELATEDOBJECT_USERS IncidentRelatedObject = "users"
)

var allowedIncidentRelatedObjectEnumValues = []IncidentRelatedObject{
	INCIDENTRELATEDOBJECT_USERS,
}

func (w *IncidentRelatedObject) GetAllowedValues() []IncidentRelatedObject {
	return allowedIncidentRelatedObjectEnumValues
}

func (v *IncidentRelatedObject) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentRelatedObject(value)
	return nil
}

// NewIncidentRelatedObjectFromValue returns a pointer to a valid IncidentRelatedObject
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIncidentRelatedObjectFromValue(v string) (*IncidentRelatedObject, error) {
	ev := IncidentRelatedObject(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IncidentRelatedObject: valid values are %v", v, allowedIncidentRelatedObjectEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IncidentRelatedObject) IsValid() bool {
	for _, existing := range allowedIncidentRelatedObjectEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentRelatedObject value
func (v IncidentRelatedObject) Ptr() *IncidentRelatedObject {
	return &v
}

type NullableIncidentRelatedObject struct {
	value *IncidentRelatedObject
	isSet bool
}

func (v NullableIncidentRelatedObject) Get() *IncidentRelatedObject {
	return v.value
}

func (v *NullableIncidentRelatedObject) Set(val *IncidentRelatedObject) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentRelatedObject) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentRelatedObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentRelatedObject(val *IncidentRelatedObject) *NullableIncidentRelatedObject {
	return &NullableIncidentRelatedObject{value: val, isSet: true}
}

func (v NullableIncidentRelatedObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentRelatedObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
