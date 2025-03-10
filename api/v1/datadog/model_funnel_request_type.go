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

// FunnelRequestType Widget request type.
type FunnelRequestType string

// List of FunnelRequestType
const (
	FUNNELREQUESTTYPE_FUNNEL FunnelRequestType = "funnel"
)

var allowedFunnelRequestTypeEnumValues = []FunnelRequestType{
	FUNNELREQUESTTYPE_FUNNEL,
}

func (w *FunnelRequestType) GetAllowedValues() []FunnelRequestType {
	return allowedFunnelRequestTypeEnumValues
}

func (v *FunnelRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FunnelRequestType(value)
	return nil
}

// NewFunnelRequestTypeFromValue returns a pointer to a valid FunnelRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFunnelRequestTypeFromValue(v string) (*FunnelRequestType, error) {
	ev := FunnelRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FunnelRequestType: valid values are %v", v, allowedFunnelRequestTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FunnelRequestType) IsValid() bool {
	for _, existing := range allowedFunnelRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FunnelRequestType value
func (v FunnelRequestType) Ptr() *FunnelRequestType {
	return &v
}

type NullableFunnelRequestType struct {
	value *FunnelRequestType
	isSet bool
}

func (v NullableFunnelRequestType) Get() *FunnelRequestType {
	return v.value
}

func (v *NullableFunnelRequestType) Set(val *FunnelRequestType) {
	v.value = val
	v.isSet = true
}

func (v NullableFunnelRequestType) IsSet() bool {
	return v.isSet
}

func (v *NullableFunnelRequestType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunnelRequestType(val *FunnelRequestType) *NullableFunnelRequestType {
	return &NullableFunnelRequestType{value: val, isSet: true}
}

func (v NullableFunnelRequestType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunnelRequestType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
