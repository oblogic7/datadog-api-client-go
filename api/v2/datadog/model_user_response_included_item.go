/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// UserResponseIncludedItem - An object related to a user.
type UserResponseIncludedItem struct {
	Organization *Organization
	Permission   *Permission
	Role         *Role

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// OrganizationAsUserResponseIncludedItem is a convenience function that returns Organization wrapped in UserResponseIncludedItem
func OrganizationAsUserResponseIncludedItem(v *Organization) UserResponseIncludedItem {
	return UserResponseIncludedItem{Organization: v}
}

// PermissionAsUserResponseIncludedItem is a convenience function that returns Permission wrapped in UserResponseIncludedItem
func PermissionAsUserResponseIncludedItem(v *Permission) UserResponseIncludedItem {
	return UserResponseIncludedItem{Permission: v}
}

// RoleAsUserResponseIncludedItem is a convenience function that returns Role wrapped in UserResponseIncludedItem
func RoleAsUserResponseIncludedItem(v *Role) UserResponseIncludedItem {
	return UserResponseIncludedItem{Role: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *UserResponseIncludedItem) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Organization
	err = json.Unmarshal(data, &dst.Organization)
	if err == nil {
		if dst.Organization != nil && dst.Organization.UnparsedObject == nil {
			jsonOrganization, _ := json.Marshal(dst.Organization)
			if string(jsonOrganization) == "{}" { // empty struct
				dst.Organization = nil
			} else {
				match++
			}
		} else {
			dst.Organization = nil
		}
	} else {
		dst.Organization = nil
	}

	// try to unmarshal data into Permission
	err = json.Unmarshal(data, &dst.Permission)
	if err == nil {
		if dst.Permission != nil && dst.Permission.UnparsedObject == nil {
			jsonPermission, _ := json.Marshal(dst.Permission)
			if string(jsonPermission) == "{}" { // empty struct
				dst.Permission = nil
			} else {
				match++
			}
		} else {
			dst.Permission = nil
		}
	} else {
		dst.Permission = nil
	}

	// try to unmarshal data into Role
	err = json.Unmarshal(data, &dst.Role)
	if err == nil {
		if dst.Role != nil && dst.Role.UnparsedObject == nil {
			jsonRole, _ := json.Marshal(dst.Role)
			if string(jsonRole) == "{}" { // empty struct
				dst.Role = nil
			} else {
				match++
			}
		} else {
			dst.Role = nil
		}
	} else {
		dst.Role = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		dst.Organization = nil
		dst.Permission = nil
		dst.Role = nil
		return json.Unmarshal(data, &dst.UnparsedObject)
	} else {
		return nil // exactly one match
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UserResponseIncludedItem) MarshalJSON() ([]byte, error) {
	if src.Organization != nil {
		return json.Marshal(&src.Organization)
	}

	if src.Permission != nil {
		return json.Marshal(&src.Permission)
	}

	if src.Role != nil {
		return json.Marshal(&src.Role)
	}

	if src.UnparsedObject != nil {
		return json.Marshal(src.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UserResponseIncludedItem) GetActualInstance() interface{} {
	if obj.Organization != nil {
		return obj.Organization
	}

	if obj.Permission != nil {
		return obj.Permission
	}

	if obj.Role != nil {
		return obj.Role
	}

	// all schemas are nil
	return nil
}

type NullableUserResponseIncludedItem struct {
	value *UserResponseIncludedItem
	isSet bool
}

func (v NullableUserResponseIncludedItem) Get() *UserResponseIncludedItem {
	return v.value
}

func (v *NullableUserResponseIncludedItem) Set(val *UserResponseIncludedItem) {
	v.value = val
	v.isSet = true
}

func (v NullableUserResponseIncludedItem) IsSet() bool {
	return v.isSet
}

func (v *NullableUserResponseIncludedItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserResponseIncludedItem(val *UserResponseIncludedItem) *NullableUserResponseIncludedItem {
	return &NullableUserResponseIncludedItem{value: val, isSet: true}
}

func (v NullableUserResponseIncludedItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserResponseIncludedItem) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return json.Unmarshal(src, &v.value)
}
