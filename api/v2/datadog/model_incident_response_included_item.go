/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// IncidentResponseIncludedItem - An object related to an incident that is included in the response.
type IncidentResponseIncludedItem struct {
	User *User

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// UserAsIncidentResponseIncludedItem is a convenience function that returns User wrapped in IncidentResponseIncludedItem
func UserAsIncidentResponseIncludedItem(v *User) IncidentResponseIncludedItem {
	return IncidentResponseIncludedItem{User: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *IncidentResponseIncludedItem) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into User
	err = json.Unmarshal(data, &dst.User)
	if err == nil {
		if dst.User != nil && dst.User.UnparsedObject == nil {
			jsonUser, _ := json.Marshal(dst.User)
			if string(jsonUser) == "{}" { // empty struct
				dst.User = nil
			} else {
				match++
			}
		} else {
			dst.User = nil
		}
	} else {
		dst.User = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		dst.User = nil
		return json.Unmarshal(data, &dst.UnparsedObject)
	} else {
		return nil // exactly one match
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IncidentResponseIncludedItem) MarshalJSON() ([]byte, error) {
	if src.User != nil {
		return json.Marshal(&src.User)
	}

	if src.UnparsedObject != nil {
		return json.Marshal(src.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IncidentResponseIncludedItem) GetActualInstance() interface{} {
	if obj.User != nil {
		return obj.User
	}

	// all schemas are nil
	return nil
}

type NullableIncidentResponseIncludedItem struct {
	value *IncidentResponseIncludedItem
	isSet bool
}

func (v NullableIncidentResponseIncludedItem) Get() *IncidentResponseIncludedItem {
	return v.value
}

func (v *NullableIncidentResponseIncludedItem) Set(val *IncidentResponseIncludedItem) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentResponseIncludedItem) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentResponseIncludedItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentResponseIncludedItem(val *IncidentResponseIncludedItem) *NullableIncidentResponseIncludedItem {
	return &NullableIncidentResponseIncludedItem{value: val, isSet: true}
}

func (v NullableIncidentResponseIncludedItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentResponseIncludedItem) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return json.Unmarshal(src, &v.value)
}
