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

// SyntheticsBasicAuthNTLMType The type of authentication to use when performing the test.
type SyntheticsBasicAuthNTLMType string

// List of SyntheticsBasicAuthNTLMType
const (
	SYNTHETICSBASICAUTHNTLMTYPE_NTLM SyntheticsBasicAuthNTLMType = "ntlm"
)

var allowedSyntheticsBasicAuthNTLMTypeEnumValues = []SyntheticsBasicAuthNTLMType{
	SYNTHETICSBASICAUTHNTLMTYPE_NTLM,
}

func (w *SyntheticsBasicAuthNTLMType) GetAllowedValues() []SyntheticsBasicAuthNTLMType {
	return allowedSyntheticsBasicAuthNTLMTypeEnumValues
}

func (v *SyntheticsBasicAuthNTLMType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsBasicAuthNTLMType(value)
	return nil
}

// NewSyntheticsBasicAuthNTLMTypeFromValue returns a pointer to a valid SyntheticsBasicAuthNTLMType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSyntheticsBasicAuthNTLMTypeFromValue(v string) (*SyntheticsBasicAuthNTLMType, error) {
	ev := SyntheticsBasicAuthNTLMType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SyntheticsBasicAuthNTLMType: valid values are %v", v, allowedSyntheticsBasicAuthNTLMTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SyntheticsBasicAuthNTLMType) IsValid() bool {
	for _, existing := range allowedSyntheticsBasicAuthNTLMTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsBasicAuthNTLMType value
func (v SyntheticsBasicAuthNTLMType) Ptr() *SyntheticsBasicAuthNTLMType {
	return &v
}

type NullableSyntheticsBasicAuthNTLMType struct {
	value *SyntheticsBasicAuthNTLMType
	isSet bool
}

func (v NullableSyntheticsBasicAuthNTLMType) Get() *SyntheticsBasicAuthNTLMType {
	return v.value
}

func (v *NullableSyntheticsBasicAuthNTLMType) Set(val *SyntheticsBasicAuthNTLMType) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsBasicAuthNTLMType) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsBasicAuthNTLMType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsBasicAuthNTLMType(val *SyntheticsBasicAuthNTLMType) *NullableSyntheticsBasicAuthNTLMType {
	return &NullableSyntheticsBasicAuthNTLMType{value: val, isSet: true}
}

func (v NullableSyntheticsBasicAuthNTLMType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsBasicAuthNTLMType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
