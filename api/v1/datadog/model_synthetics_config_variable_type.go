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

// SyntheticsConfigVariableType Type of the configuration variable.
type SyntheticsConfigVariableType string

// List of SyntheticsConfigVariableType
const (
	SYNTHETICSCONFIGVARIABLETYPE_GLOBAL SyntheticsConfigVariableType = "global"
	SYNTHETICSCONFIGVARIABLETYPE_TEXT   SyntheticsConfigVariableType = "text"
)

var allowedSyntheticsConfigVariableTypeEnumValues = []SyntheticsConfigVariableType{
	SYNTHETICSCONFIGVARIABLETYPE_GLOBAL,
	SYNTHETICSCONFIGVARIABLETYPE_TEXT,
}

func (w *SyntheticsConfigVariableType) GetAllowedValues() []SyntheticsConfigVariableType {
	return allowedSyntheticsConfigVariableTypeEnumValues
}

func (v *SyntheticsConfigVariableType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsConfigVariableType(value)
	return nil
}

// NewSyntheticsConfigVariableTypeFromValue returns a pointer to a valid SyntheticsConfigVariableType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSyntheticsConfigVariableTypeFromValue(v string) (*SyntheticsConfigVariableType, error) {
	ev := SyntheticsConfigVariableType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SyntheticsConfigVariableType: valid values are %v", v, allowedSyntheticsConfigVariableTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SyntheticsConfigVariableType) IsValid() bool {
	for _, existing := range allowedSyntheticsConfigVariableTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsConfigVariableType value
func (v SyntheticsConfigVariableType) Ptr() *SyntheticsConfigVariableType {
	return &v
}

type NullableSyntheticsConfigVariableType struct {
	value *SyntheticsConfigVariableType
	isSet bool
}

func (v NullableSyntheticsConfigVariableType) Get() *SyntheticsConfigVariableType {
	return v.value
}

func (v *NullableSyntheticsConfigVariableType) Set(val *SyntheticsConfigVariableType) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsConfigVariableType) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsConfigVariableType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsConfigVariableType(val *SyntheticsConfigVariableType) *NullableSyntheticsConfigVariableType {
	return &NullableSyntheticsConfigVariableType{value: val, isSet: true}
}

func (v NullableSyntheticsConfigVariableType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsConfigVariableType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
