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

// SyntheticsBrowserVariableType Type of browser test variable.
type SyntheticsBrowserVariableType string

// List of SyntheticsBrowserVariableType
const (
	SYNTHETICSBROWSERVARIABLETYPE_ELEMENT    SyntheticsBrowserVariableType = "element"
	SYNTHETICSBROWSERVARIABLETYPE_EMAIL      SyntheticsBrowserVariableType = "email"
	SYNTHETICSBROWSERVARIABLETYPE_GLOBAL     SyntheticsBrowserVariableType = "global"
	SYNTHETICSBROWSERVARIABLETYPE_JAVASCRIPT SyntheticsBrowserVariableType = "javascript"
	SYNTHETICSBROWSERVARIABLETYPE_TEXT       SyntheticsBrowserVariableType = "text"
)

var allowedSyntheticsBrowserVariableTypeEnumValues = []SyntheticsBrowserVariableType{
	SYNTHETICSBROWSERVARIABLETYPE_ELEMENT,
	SYNTHETICSBROWSERVARIABLETYPE_EMAIL,
	SYNTHETICSBROWSERVARIABLETYPE_GLOBAL,
	SYNTHETICSBROWSERVARIABLETYPE_JAVASCRIPT,
	SYNTHETICSBROWSERVARIABLETYPE_TEXT,
}

func (w *SyntheticsBrowserVariableType) GetAllowedValues() []SyntheticsBrowserVariableType {
	return allowedSyntheticsBrowserVariableTypeEnumValues
}

func (v *SyntheticsBrowserVariableType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsBrowserVariableType(value)
	return nil
}

// NewSyntheticsBrowserVariableTypeFromValue returns a pointer to a valid SyntheticsBrowserVariableType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSyntheticsBrowserVariableTypeFromValue(v string) (*SyntheticsBrowserVariableType, error) {
	ev := SyntheticsBrowserVariableType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SyntheticsBrowserVariableType: valid values are %v", v, allowedSyntheticsBrowserVariableTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SyntheticsBrowserVariableType) IsValid() bool {
	for _, existing := range allowedSyntheticsBrowserVariableTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsBrowserVariableType value
func (v SyntheticsBrowserVariableType) Ptr() *SyntheticsBrowserVariableType {
	return &v
}

type NullableSyntheticsBrowserVariableType struct {
	value *SyntheticsBrowserVariableType
	isSet bool
}

func (v NullableSyntheticsBrowserVariableType) Get() *SyntheticsBrowserVariableType {
	return v.value
}

func (v *NullableSyntheticsBrowserVariableType) Set(val *SyntheticsBrowserVariableType) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsBrowserVariableType) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsBrowserVariableType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsBrowserVariableType(val *SyntheticsBrowserVariableType) *NullableSyntheticsBrowserVariableType {
	return &NullableSyntheticsBrowserVariableType{value: val, isSet: true}
}

func (v NullableSyntheticsBrowserVariableType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsBrowserVariableType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
