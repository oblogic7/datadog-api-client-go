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

// DashboardLayoutType Layout type of the dashboard.
type DashboardLayoutType string

// List of DashboardLayoutType
const (
	DASHBOARDLAYOUTTYPE_ORDERED DashboardLayoutType = "ordered"
	DASHBOARDLAYOUTTYPE_FREE    DashboardLayoutType = "free"
)

var allowedDashboardLayoutTypeEnumValues = []DashboardLayoutType{
	DASHBOARDLAYOUTTYPE_ORDERED,
	DASHBOARDLAYOUTTYPE_FREE,
}

func (w *DashboardLayoutType) GetAllowedValues() []DashboardLayoutType {
	return allowedDashboardLayoutTypeEnumValues
}

func (v *DashboardLayoutType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardLayoutType(value)
	return nil
}

// NewDashboardLayoutTypeFromValue returns a pointer to a valid DashboardLayoutType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDashboardLayoutTypeFromValue(v string) (*DashboardLayoutType, error) {
	ev := DashboardLayoutType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DashboardLayoutType: valid values are %v", v, allowedDashboardLayoutTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DashboardLayoutType) IsValid() bool {
	for _, existing := range allowedDashboardLayoutTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardLayoutType value
func (v DashboardLayoutType) Ptr() *DashboardLayoutType {
	return &v
}

type NullableDashboardLayoutType struct {
	value *DashboardLayoutType
	isSet bool
}

func (v NullableDashboardLayoutType) Get() *DashboardLayoutType {
	return v.value
}

func (v *NullableDashboardLayoutType) Set(val *DashboardLayoutType) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardLayoutType) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardLayoutType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardLayoutType(val *DashboardLayoutType) *NullableDashboardLayoutType {
	return &NullableDashboardLayoutType{value: val, isSet: true}
}

func (v NullableDashboardLayoutType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardLayoutType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
