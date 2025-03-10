/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// SyntheticsAssertion - Object describing the assertions type, their associated operator,
// which property they apply, and upon which target.
type SyntheticsAssertion struct {
	SyntheticsAssertionTarget         *SyntheticsAssertionTarget
	SyntheticsAssertionJSONPathTarget *SyntheticsAssertionJSONPathTarget

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SyntheticsAssertionTargetAsSyntheticsAssertion is a convenience function that returns SyntheticsAssertionTarget wrapped in SyntheticsAssertion
func SyntheticsAssertionTargetAsSyntheticsAssertion(v *SyntheticsAssertionTarget) SyntheticsAssertion {
	return SyntheticsAssertion{SyntheticsAssertionTarget: v}
}

// SyntheticsAssertionJSONPathTargetAsSyntheticsAssertion is a convenience function that returns SyntheticsAssertionJSONPathTarget wrapped in SyntheticsAssertion
func SyntheticsAssertionJSONPathTargetAsSyntheticsAssertion(v *SyntheticsAssertionJSONPathTarget) SyntheticsAssertion {
	return SyntheticsAssertion{SyntheticsAssertionJSONPathTarget: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SyntheticsAssertion) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SyntheticsAssertionTarget
	err = json.Unmarshal(data, &dst.SyntheticsAssertionTarget)
	if err == nil {
		if dst.SyntheticsAssertionTarget != nil && dst.SyntheticsAssertionTarget.UnparsedObject == nil {
			jsonSyntheticsAssertionTarget, _ := json.Marshal(dst.SyntheticsAssertionTarget)
			if string(jsonSyntheticsAssertionTarget) == "{}" { // empty struct
				dst.SyntheticsAssertionTarget = nil
			} else {
				match++
			}
		} else {
			dst.SyntheticsAssertionTarget = nil
		}
	} else {
		dst.SyntheticsAssertionTarget = nil
	}

	// try to unmarshal data into SyntheticsAssertionJSONPathTarget
	err = json.Unmarshal(data, &dst.SyntheticsAssertionJSONPathTarget)
	if err == nil {
		if dst.SyntheticsAssertionJSONPathTarget != nil && dst.SyntheticsAssertionJSONPathTarget.UnparsedObject == nil {
			jsonSyntheticsAssertionJSONPathTarget, _ := json.Marshal(dst.SyntheticsAssertionJSONPathTarget)
			if string(jsonSyntheticsAssertionJSONPathTarget) == "{}" { // empty struct
				dst.SyntheticsAssertionJSONPathTarget = nil
			} else {
				match++
			}
		} else {
			dst.SyntheticsAssertionJSONPathTarget = nil
		}
	} else {
		dst.SyntheticsAssertionJSONPathTarget = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		dst.SyntheticsAssertionTarget = nil
		dst.SyntheticsAssertionJSONPathTarget = nil
		return json.Unmarshal(data, &dst.UnparsedObject)
	} else {
		return nil // exactly one match
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SyntheticsAssertion) MarshalJSON() ([]byte, error) {
	if src.SyntheticsAssertionTarget != nil {
		return json.Marshal(&src.SyntheticsAssertionTarget)
	}

	if src.SyntheticsAssertionJSONPathTarget != nil {
		return json.Marshal(&src.SyntheticsAssertionJSONPathTarget)
	}

	if src.UnparsedObject != nil {
		return json.Marshal(src.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SyntheticsAssertion) GetActualInstance() interface{} {
	if obj.SyntheticsAssertionTarget != nil {
		return obj.SyntheticsAssertionTarget
	}

	if obj.SyntheticsAssertionJSONPathTarget != nil {
		return obj.SyntheticsAssertionJSONPathTarget
	}

	// all schemas are nil
	return nil
}

type NullableSyntheticsAssertion struct {
	value *SyntheticsAssertion
	isSet bool
}

func (v NullableSyntheticsAssertion) Get() *SyntheticsAssertion {
	return v.value
}

func (v *NullableSyntheticsAssertion) Set(val *SyntheticsAssertion) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsAssertion) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsAssertion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsAssertion(val *SyntheticsAssertion) *NullableSyntheticsAssertion {
	return &NullableSyntheticsAssertion{value: val, isSet: true}
}

func (v NullableSyntheticsAssertion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsAssertion) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return json.Unmarshal(src, &v.value)
}
