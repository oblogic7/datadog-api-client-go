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

// RolesType Roles type.
type RolesType string

// List of RolesType
const (
	ROLESTYPE_ROLES RolesType = "roles"
)

var allowedRolesTypeEnumValues = []RolesType{
	ROLESTYPE_ROLES,
}

func (w *RolesType) GetAllowedValues() []RolesType {
	return allowedRolesTypeEnumValues
}

func (v *RolesType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RolesType(value)
	return nil
}

// NewRolesTypeFromValue returns a pointer to a valid RolesType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRolesTypeFromValue(v string) (*RolesType, error) {
	ev := RolesType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RolesType: valid values are %v", v, allowedRolesTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RolesType) IsValid() bool {
	for _, existing := range allowedRolesTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RolesType value
func (v RolesType) Ptr() *RolesType {
	return &v
}

type NullableRolesType struct {
	value *RolesType
	isSet bool
}

func (v NullableRolesType) Get() *RolesType {
	return v.value
}

func (v *NullableRolesType) Set(val *RolesType) {
	v.value = val
	v.isSet = true
}

func (v NullableRolesType) IsSet() bool {
	return v.isSet
}

func (v *NullableRolesType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRolesType(val *RolesType) *NullableRolesType {
	return &NullableRolesType{value: val, isSet: true}
}

func (v NullableRolesType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRolesType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
