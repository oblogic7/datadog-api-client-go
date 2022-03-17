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

// DistributionPointsType The type of the distribution point.
type DistributionPointsType string

// List of DistributionPointsType
const (
	DISTRIBUTIONPOINTSTYPE_DISTRIBUTION DistributionPointsType = "distribution"
)

var allowedDistributionPointsTypeEnumValues = []DistributionPointsType{
	DISTRIBUTIONPOINTSTYPE_DISTRIBUTION,
}

func (w *DistributionPointsType) GetAllowedValues() []DistributionPointsType {
	return allowedDistributionPointsTypeEnumValues
}

func (v *DistributionPointsType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DistributionPointsType(value)
	return nil
}

// NewDistributionPointsTypeFromValue returns a pointer to a valid DistributionPointsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDistributionPointsTypeFromValue(v string) (*DistributionPointsType, error) {
	ev := DistributionPointsType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DistributionPointsType: valid values are %v", v, allowedDistributionPointsTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DistributionPointsType) IsValid() bool {
	for _, existing := range allowedDistributionPointsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DistributionPointsType value
func (v DistributionPointsType) Ptr() *DistributionPointsType {
	return &v
}

type NullableDistributionPointsType struct {
	value *DistributionPointsType
	isSet bool
}

func (v NullableDistributionPointsType) Get() *DistributionPointsType {
	return v.value
}

func (v *NullableDistributionPointsType) Set(val *DistributionPointsType) {
	v.value = val
	v.isSet = true
}

func (v NullableDistributionPointsType) IsSet() bool {
	return v.isSet
}

func (v *NullableDistributionPointsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDistributionPointsType(val *DistributionPointsType) *NullableDistributionPointsType {
	return &NullableDistributionPointsType{value: val, isSet: true}
}

func (v NullableDistributionPointsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDistributionPointsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
