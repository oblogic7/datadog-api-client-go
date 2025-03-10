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

// APIKeysType API Keys resource type.
type APIKeysType string

// List of APIKeysType
const (
	APIKEYSTYPE_API_KEYS APIKeysType = "api_keys"
)

var allowedAPIKeysTypeEnumValues = []APIKeysType{
	APIKEYSTYPE_API_KEYS,
}

func (w *APIKeysType) GetAllowedValues() []APIKeysType {
	return allowedAPIKeysTypeEnumValues
}

func (v *APIKeysType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = APIKeysType(value)
	return nil
}

// NewAPIKeysTypeFromValue returns a pointer to a valid APIKeysType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAPIKeysTypeFromValue(v string) (*APIKeysType, error) {
	ev := APIKeysType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for APIKeysType: valid values are %v", v, allowedAPIKeysTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v APIKeysType) IsValid() bool {
	for _, existing := range allowedAPIKeysTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to APIKeysType value
func (v APIKeysType) Ptr() *APIKeysType {
	return &v
}

type NullableAPIKeysType struct {
	value *APIKeysType
	isSet bool
}

func (v NullableAPIKeysType) Get() *APIKeysType {
	return v.value
}

func (v *NullableAPIKeysType) Set(val *APIKeysType) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIKeysType) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIKeysType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIKeysType(val *APIKeysType) *NullableAPIKeysType {
	return &NullableAPIKeysType{value: val, isSet: true}
}

func (v NullableAPIKeysType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIKeysType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
