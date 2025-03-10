/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// AuthNMappingIncluded - Included data in the AuthN Mapping response.
type AuthNMappingIncluded struct {
	SAMLAssertionAttribute *SAMLAssertionAttribute
	Role                   *Role

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SAMLAssertionAttributeAsAuthNMappingIncluded is a convenience function that returns SAMLAssertionAttribute wrapped in AuthNMappingIncluded
func SAMLAssertionAttributeAsAuthNMappingIncluded(v *SAMLAssertionAttribute) AuthNMappingIncluded {
	return AuthNMappingIncluded{SAMLAssertionAttribute: v}
}

// RoleAsAuthNMappingIncluded is a convenience function that returns Role wrapped in AuthNMappingIncluded
func RoleAsAuthNMappingIncluded(v *Role) AuthNMappingIncluded {
	return AuthNMappingIncluded{Role: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AuthNMappingIncluded) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SAMLAssertionAttribute
	err = json.Unmarshal(data, &dst.SAMLAssertionAttribute)
	if err == nil {
		if dst.SAMLAssertionAttribute != nil && dst.SAMLAssertionAttribute.UnparsedObject == nil {
			jsonSAMLAssertionAttribute, _ := json.Marshal(dst.SAMLAssertionAttribute)
			if string(jsonSAMLAssertionAttribute) == "{}" { // empty struct
				dst.SAMLAssertionAttribute = nil
			} else {
				match++
			}
		} else {
			dst.SAMLAssertionAttribute = nil
		}
	} else {
		dst.SAMLAssertionAttribute = nil
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
		dst.SAMLAssertionAttribute = nil
		dst.Role = nil
		return json.Unmarshal(data, &dst.UnparsedObject)
	} else {
		return nil // exactly one match
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AuthNMappingIncluded) MarshalJSON() ([]byte, error) {
	if src.SAMLAssertionAttribute != nil {
		return json.Marshal(&src.SAMLAssertionAttribute)
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
func (obj *AuthNMappingIncluded) GetActualInstance() interface{} {
	if obj.SAMLAssertionAttribute != nil {
		return obj.SAMLAssertionAttribute
	}

	if obj.Role != nil {
		return obj.Role
	}

	// all schemas are nil
	return nil
}

type NullableAuthNMappingIncluded struct {
	value *AuthNMappingIncluded
	isSet bool
}

func (v NullableAuthNMappingIncluded) Get() *AuthNMappingIncluded {
	return v.value
}

func (v *NullableAuthNMappingIncluded) Set(val *AuthNMappingIncluded) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthNMappingIncluded) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthNMappingIncluded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthNMappingIncluded(val *AuthNMappingIncluded) *NullableAuthNMappingIncluded {
	return &NullableAuthNMappingIncluded{value: val, isSet: true}
}

func (v NullableAuthNMappingIncluded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthNMappingIncluded) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return json.Unmarshal(src, &v.value)
}
