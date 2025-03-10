/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// IncidentTimelineCellCreateAttributes - The timeline cell's attributes for a create request.
type IncidentTimelineCellCreateAttributes struct {
	IncidentTimelineCellMarkdownCreateAttributes *IncidentTimelineCellMarkdownCreateAttributes

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// IncidentTimelineCellMarkdownCreateAttributesAsIncidentTimelineCellCreateAttributes is a convenience function that returns IncidentTimelineCellMarkdownCreateAttributes wrapped in IncidentTimelineCellCreateAttributes
func IncidentTimelineCellMarkdownCreateAttributesAsIncidentTimelineCellCreateAttributes(v *IncidentTimelineCellMarkdownCreateAttributes) IncidentTimelineCellCreateAttributes {
	return IncidentTimelineCellCreateAttributes{IncidentTimelineCellMarkdownCreateAttributes: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *IncidentTimelineCellCreateAttributes) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into IncidentTimelineCellMarkdownCreateAttributes
	err = json.Unmarshal(data, &dst.IncidentTimelineCellMarkdownCreateAttributes)
	if err == nil {
		if dst.IncidentTimelineCellMarkdownCreateAttributes != nil && dst.IncidentTimelineCellMarkdownCreateAttributes.UnparsedObject == nil {
			jsonIncidentTimelineCellMarkdownCreateAttributes, _ := json.Marshal(dst.IncidentTimelineCellMarkdownCreateAttributes)
			if string(jsonIncidentTimelineCellMarkdownCreateAttributes) == "{}" { // empty struct
				dst.IncidentTimelineCellMarkdownCreateAttributes = nil
			} else {
				match++
			}
		} else {
			dst.IncidentTimelineCellMarkdownCreateAttributes = nil
		}
	} else {
		dst.IncidentTimelineCellMarkdownCreateAttributes = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		dst.IncidentTimelineCellMarkdownCreateAttributes = nil
		return json.Unmarshal(data, &dst.UnparsedObject)
	} else {
		return nil // exactly one match
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IncidentTimelineCellCreateAttributes) MarshalJSON() ([]byte, error) {
	if src.IncidentTimelineCellMarkdownCreateAttributes != nil {
		return json.Marshal(&src.IncidentTimelineCellMarkdownCreateAttributes)
	}

	if src.UnparsedObject != nil {
		return json.Marshal(src.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IncidentTimelineCellCreateAttributes) GetActualInstance() interface{} {
	if obj.IncidentTimelineCellMarkdownCreateAttributes != nil {
		return obj.IncidentTimelineCellMarkdownCreateAttributes
	}

	// all schemas are nil
	return nil
}

type NullableIncidentTimelineCellCreateAttributes struct {
	value *IncidentTimelineCellCreateAttributes
	isSet bool
}

func (v NullableIncidentTimelineCellCreateAttributes) Get() *IncidentTimelineCellCreateAttributes {
	return v.value
}

func (v *NullableIncidentTimelineCellCreateAttributes) Set(val *IncidentTimelineCellCreateAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentTimelineCellCreateAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentTimelineCellCreateAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentTimelineCellCreateAttributes(val *IncidentTimelineCellCreateAttributes) *NullableIncidentTimelineCellCreateAttributes {
	return &NullableIncidentTimelineCellCreateAttributes{value: val, isSet: true}
}

func (v NullableIncidentTimelineCellCreateAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentTimelineCellCreateAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return json.Unmarshal(src, &v.value)
}
