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

// DashboardReflowType Reflow type for a **new dashboard layout** dashboard. Set this only when layout type is 'ordered'.
// If set to 'fixed', the dashboard expects all widgets to have a layout, and if it's set to 'auto',
// widgets should not have layouts.
type DashboardReflowType string

// List of DashboardReflowType
const (
	DASHBOARDREFLOWTYPE_AUTO  DashboardReflowType = "auto"
	DASHBOARDREFLOWTYPE_FIXED DashboardReflowType = "fixed"
)

var allowedDashboardReflowTypeEnumValues = []DashboardReflowType{
	DASHBOARDREFLOWTYPE_AUTO,
	DASHBOARDREFLOWTYPE_FIXED,
}

func (w *DashboardReflowType) GetAllowedValues() []DashboardReflowType {
	return allowedDashboardReflowTypeEnumValues
}

func (v *DashboardReflowType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardReflowType(value)
	return nil
}

// NewDashboardReflowTypeFromValue returns a pointer to a valid DashboardReflowType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDashboardReflowTypeFromValue(v string) (*DashboardReflowType, error) {
	ev := DashboardReflowType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DashboardReflowType: valid values are %v", v, allowedDashboardReflowTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DashboardReflowType) IsValid() bool {
	for _, existing := range allowedDashboardReflowTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardReflowType value
func (v DashboardReflowType) Ptr() *DashboardReflowType {
	return &v
}

type NullableDashboardReflowType struct {
	value *DashboardReflowType
	isSet bool
}

func (v NullableDashboardReflowType) Get() *DashboardReflowType {
	return v.value
}

func (v *NullableDashboardReflowType) Set(val *DashboardReflowType) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardReflowType) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardReflowType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardReflowType(val *DashboardReflowType) *NullableDashboardReflowType {
	return &NullableDashboardReflowType{value: val, isSet: true}
}

func (v NullableDashboardReflowType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardReflowType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
