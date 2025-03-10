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

// SyntheticsBrowserTestType Type of the Synthetic test, `browser`.
type SyntheticsBrowserTestType string

// List of SyntheticsBrowserTestType
const (
	SYNTHETICSBROWSERTESTTYPE_BROWSER SyntheticsBrowserTestType = "browser"
)

var allowedSyntheticsBrowserTestTypeEnumValues = []SyntheticsBrowserTestType{
	SYNTHETICSBROWSERTESTTYPE_BROWSER,
}

func (w *SyntheticsBrowserTestType) GetAllowedValues() []SyntheticsBrowserTestType {
	return allowedSyntheticsBrowserTestTypeEnumValues
}

func (v *SyntheticsBrowserTestType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsBrowserTestType(value)
	return nil
}

// NewSyntheticsBrowserTestTypeFromValue returns a pointer to a valid SyntheticsBrowserTestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSyntheticsBrowserTestTypeFromValue(v string) (*SyntheticsBrowserTestType, error) {
	ev := SyntheticsBrowserTestType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SyntheticsBrowserTestType: valid values are %v", v, allowedSyntheticsBrowserTestTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SyntheticsBrowserTestType) IsValid() bool {
	for _, existing := range allowedSyntheticsBrowserTestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsBrowserTestType value
func (v SyntheticsBrowserTestType) Ptr() *SyntheticsBrowserTestType {
	return &v
}

type NullableSyntheticsBrowserTestType struct {
	value *SyntheticsBrowserTestType
	isSet bool
}

func (v NullableSyntheticsBrowserTestType) Get() *SyntheticsBrowserTestType {
	return v.value
}

func (v *NullableSyntheticsBrowserTestType) Set(val *SyntheticsBrowserTestType) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsBrowserTestType) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsBrowserTestType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsBrowserTestType(val *SyntheticsBrowserTestType) *NullableSyntheticsBrowserTestType {
	return &NullableSyntheticsBrowserTestType{value: val, isSet: true}
}

func (v NullableSyntheticsBrowserTestType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsBrowserTestType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
