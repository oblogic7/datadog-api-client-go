/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// IncidentFieldAttributes - Dynamic fields for which selections can be made, with field names as keys.
type IncidentFieldAttributes struct {
	IncidentFieldAttributesSingleValue   *IncidentFieldAttributesSingleValue
	IncidentFieldAttributesMultipleValue *IncidentFieldAttributesMultipleValue

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// IncidentFieldAttributesSingleValueAsIncidentFieldAttributes is a convenience function that returns IncidentFieldAttributesSingleValue wrapped in IncidentFieldAttributes
func IncidentFieldAttributesSingleValueAsIncidentFieldAttributes(v *IncidentFieldAttributesSingleValue) IncidentFieldAttributes {
	return IncidentFieldAttributes{IncidentFieldAttributesSingleValue: v}
}

// IncidentFieldAttributesMultipleValueAsIncidentFieldAttributes is a convenience function that returns IncidentFieldAttributesMultipleValue wrapped in IncidentFieldAttributes
func IncidentFieldAttributesMultipleValueAsIncidentFieldAttributes(v *IncidentFieldAttributesMultipleValue) IncidentFieldAttributes {
	return IncidentFieldAttributes{IncidentFieldAttributesMultipleValue: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *IncidentFieldAttributes) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into IncidentFieldAttributesSingleValue
	err = json.Unmarshal(data, &dst.IncidentFieldAttributesSingleValue)
	if err == nil {
		if dst.IncidentFieldAttributesSingleValue != nil && dst.IncidentFieldAttributesSingleValue.UnparsedObject == nil {
			jsonIncidentFieldAttributesSingleValue, _ := json.Marshal(dst.IncidentFieldAttributesSingleValue)
			if string(jsonIncidentFieldAttributesSingleValue) == "{}" { // empty struct
				dst.IncidentFieldAttributesSingleValue = nil
			} else {
				match++
			}
		} else {
			dst.IncidentFieldAttributesSingleValue = nil
		}
	} else {
		dst.IncidentFieldAttributesSingleValue = nil
	}

	// try to unmarshal data into IncidentFieldAttributesMultipleValue
	err = json.Unmarshal(data, &dst.IncidentFieldAttributesMultipleValue)
	if err == nil {
		if dst.IncidentFieldAttributesMultipleValue != nil && dst.IncidentFieldAttributesMultipleValue.UnparsedObject == nil {
			jsonIncidentFieldAttributesMultipleValue, _ := json.Marshal(dst.IncidentFieldAttributesMultipleValue)
			if string(jsonIncidentFieldAttributesMultipleValue) == "{}" { // empty struct
				dst.IncidentFieldAttributesMultipleValue = nil
			} else {
				match++
			}
		} else {
			dst.IncidentFieldAttributesMultipleValue = nil
		}
	} else {
		dst.IncidentFieldAttributesMultipleValue = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		dst.IncidentFieldAttributesSingleValue = nil
		dst.IncidentFieldAttributesMultipleValue = nil
		return json.Unmarshal(data, &dst.UnparsedObject)
	} else {
		return nil // exactly one match
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IncidentFieldAttributes) MarshalJSON() ([]byte, error) {
	if src.IncidentFieldAttributesSingleValue != nil {
		return json.Marshal(&src.IncidentFieldAttributesSingleValue)
	}

	if src.IncidentFieldAttributesMultipleValue != nil {
		return json.Marshal(&src.IncidentFieldAttributesMultipleValue)
	}

	if src.UnparsedObject != nil {
		return json.Marshal(src.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IncidentFieldAttributes) GetActualInstance() interface{} {
	if obj.IncidentFieldAttributesSingleValue != nil {
		return obj.IncidentFieldAttributesSingleValue
	}

	if obj.IncidentFieldAttributesMultipleValue != nil {
		return obj.IncidentFieldAttributesMultipleValue
	}

	// all schemas are nil
	return nil
}

type NullableIncidentFieldAttributes struct {
	value *IncidentFieldAttributes
	isSet bool
}

func (v NullableIncidentFieldAttributes) Get() *IncidentFieldAttributes {
	return v.value
}

func (v *NullableIncidentFieldAttributes) Set(val *IncidentFieldAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentFieldAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentFieldAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentFieldAttributes(val *IncidentFieldAttributes) *NullableIncidentFieldAttributes {
	return &NullableIncidentFieldAttributes{value: val, isSet: true}
}

func (v NullableIncidentFieldAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentFieldAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return json.Unmarshal(src, &v.value)
}
