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

// TargetFormatType If the `target_type` of the remapper is `attribute`, try to cast the value to a new specific type.
// If the cast is not possible, the original type is kept. `string`, `integer`, or `double` are the possible types.
// If the `target_type` is `tag`, this parameter may not be specified.
type TargetFormatType string

// List of TargetFormatType
const (
	TARGETFORMATTYPE_AUTO    TargetFormatType = "auto"
	TARGETFORMATTYPE_STRING  TargetFormatType = "string"
	TARGETFORMATTYPE_INTEGER TargetFormatType = "integer"
	TARGETFORMATTYPE_DOUBLE  TargetFormatType = "double"
)

var allowedTargetFormatTypeEnumValues = []TargetFormatType{
	TARGETFORMATTYPE_AUTO,
	TARGETFORMATTYPE_STRING,
	TARGETFORMATTYPE_INTEGER,
	TARGETFORMATTYPE_DOUBLE,
}

func (w *TargetFormatType) GetAllowedValues() []TargetFormatType {
	return allowedTargetFormatTypeEnumValues
}

func (v *TargetFormatType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TargetFormatType(value)
	return nil
}

// NewTargetFormatTypeFromValue returns a pointer to a valid TargetFormatType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTargetFormatTypeFromValue(v string) (*TargetFormatType, error) {
	ev := TargetFormatType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TargetFormatType: valid values are %v", v, allowedTargetFormatTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TargetFormatType) IsValid() bool {
	for _, existing := range allowedTargetFormatTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TargetFormatType value
func (v TargetFormatType) Ptr() *TargetFormatType {
	return &v
}

type NullableTargetFormatType struct {
	value *TargetFormatType
	isSet bool
}

func (v NullableTargetFormatType) Get() *TargetFormatType {
	return v.value
}

func (v *NullableTargetFormatType) Set(val *TargetFormatType) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetFormatType) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetFormatType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetFormatType(val *TargetFormatType) *NullableTargetFormatType {
	return &NullableTargetFormatType{value: val, isSet: true}
}

func (v NullableTargetFormatType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetFormatType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
