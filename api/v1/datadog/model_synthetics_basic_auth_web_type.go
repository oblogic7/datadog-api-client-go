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

// SyntheticsBasicAuthWebType The type of basic authentication to use when performing the test.
type SyntheticsBasicAuthWebType string

// List of SyntheticsBasicAuthWebType
const (
	SYNTHETICSBASICAUTHWEBTYPE_WEB SyntheticsBasicAuthWebType = "web"
)

var allowedSyntheticsBasicAuthWebTypeEnumValues = []SyntheticsBasicAuthWebType{
	SYNTHETICSBASICAUTHWEBTYPE_WEB,
}

func (w *SyntheticsBasicAuthWebType) GetAllowedValues() []SyntheticsBasicAuthWebType {
	return allowedSyntheticsBasicAuthWebTypeEnumValues
}

func (v *SyntheticsBasicAuthWebType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsBasicAuthWebType(value)
	return nil
}

// NewSyntheticsBasicAuthWebTypeFromValue returns a pointer to a valid SyntheticsBasicAuthWebType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSyntheticsBasicAuthWebTypeFromValue(v string) (*SyntheticsBasicAuthWebType, error) {
	ev := SyntheticsBasicAuthWebType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SyntheticsBasicAuthWebType: valid values are %v", v, allowedSyntheticsBasicAuthWebTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SyntheticsBasicAuthWebType) IsValid() bool {
	for _, existing := range allowedSyntheticsBasicAuthWebTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsBasicAuthWebType value
func (v SyntheticsBasicAuthWebType) Ptr() *SyntheticsBasicAuthWebType {
	return &v
}

type NullableSyntheticsBasicAuthWebType struct {
	value *SyntheticsBasicAuthWebType
	isSet bool
}

func (v NullableSyntheticsBasicAuthWebType) Get() *SyntheticsBasicAuthWebType {
	return v.value
}

func (v *NullableSyntheticsBasicAuthWebType) Set(val *SyntheticsBasicAuthWebType) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsBasicAuthWebType) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsBasicAuthWebType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsBasicAuthWebType(val *SyntheticsBasicAuthWebType) *NullableSyntheticsBasicAuthWebType {
	return &NullableSyntheticsBasicAuthWebType{value: val, isSet: true}
}

func (v NullableSyntheticsBasicAuthWebType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsBasicAuthWebType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
