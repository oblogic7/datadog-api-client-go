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

// SyntheticsAssertionJSONPathOperator Assertion operator to apply.
type SyntheticsAssertionJSONPathOperator string

// List of SyntheticsAssertionJSONPathOperator
const (
	SYNTHETICSASSERTIONJSONPATHOPERATOR_VALIDATES_JSON_PATH SyntheticsAssertionJSONPathOperator = "validatesJSONPath"
)

var allowedSyntheticsAssertionJSONPathOperatorEnumValues = []SyntheticsAssertionJSONPathOperator{
	SYNTHETICSASSERTIONJSONPATHOPERATOR_VALIDATES_JSON_PATH,
}

func (w *SyntheticsAssertionJSONPathOperator) GetAllowedValues() []SyntheticsAssertionJSONPathOperator {
	return allowedSyntheticsAssertionJSONPathOperatorEnumValues
}

func (v *SyntheticsAssertionJSONPathOperator) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsAssertionJSONPathOperator(value)
	return nil
}

// NewSyntheticsAssertionJSONPathOperatorFromValue returns a pointer to a valid SyntheticsAssertionJSONPathOperator
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSyntheticsAssertionJSONPathOperatorFromValue(v string) (*SyntheticsAssertionJSONPathOperator, error) {
	ev := SyntheticsAssertionJSONPathOperator(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SyntheticsAssertionJSONPathOperator: valid values are %v", v, allowedSyntheticsAssertionJSONPathOperatorEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SyntheticsAssertionJSONPathOperator) IsValid() bool {
	for _, existing := range allowedSyntheticsAssertionJSONPathOperatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsAssertionJSONPathOperator value
func (v SyntheticsAssertionJSONPathOperator) Ptr() *SyntheticsAssertionJSONPathOperator {
	return &v
}

type NullableSyntheticsAssertionJSONPathOperator struct {
	value *SyntheticsAssertionJSONPathOperator
	isSet bool
}

func (v NullableSyntheticsAssertionJSONPathOperator) Get() *SyntheticsAssertionJSONPathOperator {
	return v.value
}

func (v *NullableSyntheticsAssertionJSONPathOperator) Set(val *SyntheticsAssertionJSONPathOperator) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsAssertionJSONPathOperator) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsAssertionJSONPathOperator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsAssertionJSONPathOperator(val *SyntheticsAssertionJSONPathOperator) *NullableSyntheticsAssertionJSONPathOperator {
	return &NullableSyntheticsAssertionJSONPathOperator{value: val, isSet: true}
}

func (v NullableSyntheticsAssertionJSONPathOperator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsAssertionJSONPathOperator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
