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

// OrganizationsType Organizations resource type.
type OrganizationsType string

// List of OrganizationsType
const (
	ORGANIZATIONSTYPE_ORGS OrganizationsType = "orgs"
)

var allowedOrganizationsTypeEnumValues = []OrganizationsType{
	ORGANIZATIONSTYPE_ORGS,
}

func (w *OrganizationsType) GetAllowedValues() []OrganizationsType {
	return allowedOrganizationsTypeEnumValues
}

func (v *OrganizationsType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OrganizationsType(value)
	return nil
}

// NewOrganizationsTypeFromValue returns a pointer to a valid OrganizationsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrganizationsTypeFromValue(v string) (*OrganizationsType, error) {
	ev := OrganizationsType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrganizationsType: valid values are %v", v, allowedOrganizationsTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrganizationsType) IsValid() bool {
	for _, existing := range allowedOrganizationsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrganizationsType value
func (v OrganizationsType) Ptr() *OrganizationsType {
	return &v
}

type NullableOrganizationsType struct {
	value *OrganizationsType
	isSet bool
}

func (v NullableOrganizationsType) Get() *OrganizationsType {
	return v.value
}

func (v *NullableOrganizationsType) Set(val *OrganizationsType) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsType) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsType(val *OrganizationsType) *NullableOrganizationsType {
	return &NullableOrganizationsType{value: val, isSet: true}
}

func (v NullableOrganizationsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
