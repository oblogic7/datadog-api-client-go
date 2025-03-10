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

// UsersType Users resource type.
type UsersType string

// List of UsersType
const (
	USERSTYPE_USERS UsersType = "users"
)

var allowedUsersTypeEnumValues = []UsersType{
	USERSTYPE_USERS,
}

func (w *UsersType) GetAllowedValues() []UsersType {
	return allowedUsersTypeEnumValues
}

func (v *UsersType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UsersType(value)
	return nil
}

// NewUsersTypeFromValue returns a pointer to a valid UsersType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUsersTypeFromValue(v string) (*UsersType, error) {
	ev := UsersType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UsersType: valid values are %v", v, allowedUsersTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UsersType) IsValid() bool {
	for _, existing := range allowedUsersTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UsersType value
func (v UsersType) Ptr() *UsersType {
	return &v
}

type NullableUsersType struct {
	value *UsersType
	isSet bool
}

func (v NullableUsersType) Get() *UsersType {
	return v.value
}

func (v *NullableUsersType) Set(val *UsersType) {
	v.value = val
	v.isSet = true
}

func (v NullableUsersType) IsSet() bool {
	return v.isSet
}

func (v *NullableUsersType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsersType(val *UsersType) *NullableUsersType {
	return &NullableUsersType{value: val, isSet: true}
}

func (v NullableUsersType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsersType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
