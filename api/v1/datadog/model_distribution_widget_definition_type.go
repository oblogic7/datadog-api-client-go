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

// DistributionWidgetDefinitionType Type of the distribution widget.
type DistributionWidgetDefinitionType string

// List of DistributionWidgetDefinitionType
const (
	DISTRIBUTIONWIDGETDEFINITIONTYPE_DISTRIBUTION DistributionWidgetDefinitionType = "distribution"
)

var allowedDistributionWidgetDefinitionTypeEnumValues = []DistributionWidgetDefinitionType{
	DISTRIBUTIONWIDGETDEFINITIONTYPE_DISTRIBUTION,
}

func (w *DistributionWidgetDefinitionType) GetAllowedValues() []DistributionWidgetDefinitionType {
	return allowedDistributionWidgetDefinitionTypeEnumValues
}

func (v *DistributionWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DistributionWidgetDefinitionType(value)
	return nil
}

// NewDistributionWidgetDefinitionTypeFromValue returns a pointer to a valid DistributionWidgetDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDistributionWidgetDefinitionTypeFromValue(v string) (*DistributionWidgetDefinitionType, error) {
	ev := DistributionWidgetDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DistributionWidgetDefinitionType: valid values are %v", v, allowedDistributionWidgetDefinitionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DistributionWidgetDefinitionType) IsValid() bool {
	for _, existing := range allowedDistributionWidgetDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DistributionWidgetDefinitionType value
func (v DistributionWidgetDefinitionType) Ptr() *DistributionWidgetDefinitionType {
	return &v
}

type NullableDistributionWidgetDefinitionType struct {
	value *DistributionWidgetDefinitionType
	isSet bool
}

func (v NullableDistributionWidgetDefinitionType) Get() *DistributionWidgetDefinitionType {
	return v.value
}

func (v *NullableDistributionWidgetDefinitionType) Set(val *DistributionWidgetDefinitionType) {
	v.value = val
	v.isSet = true
}

func (v NullableDistributionWidgetDefinitionType) IsSet() bool {
	return v.isSet
}

func (v *NullableDistributionWidgetDefinitionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDistributionWidgetDefinitionType(val *DistributionWidgetDefinitionType) *NullableDistributionWidgetDefinitionType {
	return &NullableDistributionWidgetDefinitionType{value: val, isSet: true}
}

func (v NullableDistributionWidgetDefinitionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDistributionWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
