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

// RUMEventType Type of the event.
type RUMEventType string

// List of RUMEventType
const (
	RUMEVENTTYPE_RUM RUMEventType = "rum"
)

var allowedRUMEventTypeEnumValues = []RUMEventType{
	RUMEVENTTYPE_RUM,
}

func (w *RUMEventType) GetAllowedValues() []RUMEventType {
	return allowedRUMEventTypeEnumValues
}

func (v *RUMEventType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMEventType(value)
	return nil
}

// NewRUMEventTypeFromValue returns a pointer to a valid RUMEventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRUMEventTypeFromValue(v string) (*RUMEventType, error) {
	ev := RUMEventType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RUMEventType: valid values are %v", v, allowedRUMEventTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RUMEventType) IsValid() bool {
	for _, existing := range allowedRUMEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMEventType value
func (v RUMEventType) Ptr() *RUMEventType {
	return &v
}

type NullableRUMEventType struct {
	value *RUMEventType
	isSet bool
}

func (v NullableRUMEventType) Get() *RUMEventType {
	return v.value
}

func (v *NullableRUMEventType) Set(val *RUMEventType) {
	v.value = val
	v.isSet = true
}

func (v NullableRUMEventType) IsSet() bool {
	return v.isSet
}

func (v *NullableRUMEventType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRUMEventType(val *RUMEventType) *NullableRUMEventType {
	return &NullableRUMEventType{value: val, isSet: true}
}

func (v NullableRUMEventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRUMEventType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
