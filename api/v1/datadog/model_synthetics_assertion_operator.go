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

// SyntheticsAssertionOperator Assertion operator to apply.
type SyntheticsAssertionOperator string

// List of SyntheticsAssertionOperator
const (
	SYNTHETICSASSERTIONOPERATOR_CONTAINS             SyntheticsAssertionOperator = "contains"
	SYNTHETICSASSERTIONOPERATOR_DOES_NOT_CONTAIN     SyntheticsAssertionOperator = "doesNotContain"
	SYNTHETICSASSERTIONOPERATOR_IS                   SyntheticsAssertionOperator = "is"
	SYNTHETICSASSERTIONOPERATOR_IS_NOT               SyntheticsAssertionOperator = "isNot"
	SYNTHETICSASSERTIONOPERATOR_LESS_THAN            SyntheticsAssertionOperator = "lessThan"
	SYNTHETICSASSERTIONOPERATOR_LESS_THAN_OR_EQUAL   SyntheticsAssertionOperator = "lessThanOrEqual"
	SYNTHETICSASSERTIONOPERATOR_MORE_THAN            SyntheticsAssertionOperator = "moreThan"
	SYNTHETICSASSERTIONOPERATOR_MORE_THAN_OR_EQUAL   SyntheticsAssertionOperator = "moreThanOrEqual"
	SYNTHETICSASSERTIONOPERATOR_MATCHES              SyntheticsAssertionOperator = "matches"
	SYNTHETICSASSERTIONOPERATOR_DOES_NOT_MATCH       SyntheticsAssertionOperator = "doesNotMatch"
	SYNTHETICSASSERTIONOPERATOR_VALIDATES            SyntheticsAssertionOperator = "validates"
	SYNTHETICSASSERTIONOPERATOR_IS_IN_MORE_DAYS_THAN SyntheticsAssertionOperator = "isInMoreThan"
	SYNTHETICSASSERTIONOPERATOR_IS_IN_LESS_DAYS_THAN SyntheticsAssertionOperator = "isInLessThan"
)

var allowedSyntheticsAssertionOperatorEnumValues = []SyntheticsAssertionOperator{
	SYNTHETICSASSERTIONOPERATOR_CONTAINS,
	SYNTHETICSASSERTIONOPERATOR_DOES_NOT_CONTAIN,
	SYNTHETICSASSERTIONOPERATOR_IS,
	SYNTHETICSASSERTIONOPERATOR_IS_NOT,
	SYNTHETICSASSERTIONOPERATOR_LESS_THAN,
	SYNTHETICSASSERTIONOPERATOR_LESS_THAN_OR_EQUAL,
	SYNTHETICSASSERTIONOPERATOR_MORE_THAN,
	SYNTHETICSASSERTIONOPERATOR_MORE_THAN_OR_EQUAL,
	SYNTHETICSASSERTIONOPERATOR_MATCHES,
	SYNTHETICSASSERTIONOPERATOR_DOES_NOT_MATCH,
	SYNTHETICSASSERTIONOPERATOR_VALIDATES,
	SYNTHETICSASSERTIONOPERATOR_IS_IN_MORE_DAYS_THAN,
	SYNTHETICSASSERTIONOPERATOR_IS_IN_LESS_DAYS_THAN,
}

func (w *SyntheticsAssertionOperator) GetAllowedValues() []SyntheticsAssertionOperator {
	return allowedSyntheticsAssertionOperatorEnumValues
}

func (v *SyntheticsAssertionOperator) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsAssertionOperator(value)
	return nil
}

// NewSyntheticsAssertionOperatorFromValue returns a pointer to a valid SyntheticsAssertionOperator
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSyntheticsAssertionOperatorFromValue(v string) (*SyntheticsAssertionOperator, error) {
	ev := SyntheticsAssertionOperator(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SyntheticsAssertionOperator: valid values are %v", v, allowedSyntheticsAssertionOperatorEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SyntheticsAssertionOperator) IsValid() bool {
	for _, existing := range allowedSyntheticsAssertionOperatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsAssertionOperator value
func (v SyntheticsAssertionOperator) Ptr() *SyntheticsAssertionOperator {
	return &v
}

type NullableSyntheticsAssertionOperator struct {
	value *SyntheticsAssertionOperator
	isSet bool
}

func (v NullableSyntheticsAssertionOperator) Get() *SyntheticsAssertionOperator {
	return v.value
}

func (v *NullableSyntheticsAssertionOperator) Set(val *SyntheticsAssertionOperator) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsAssertionOperator) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsAssertionOperator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsAssertionOperator(val *SyntheticsAssertionOperator) *NullableSyntheticsAssertionOperator {
	return &NullableSyntheticsAssertionOperator{value: val, isSet: true}
}

func (v NullableSyntheticsAssertionOperator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsAssertionOperator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
