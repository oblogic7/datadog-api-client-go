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

// MonitorRenotifyStatusType The different statuses for which renotification is supported.
type MonitorRenotifyStatusType string

// List of MonitorRenotifyStatusType
const (
	MONITORRENOTIFYSTATUSTYPE_ALERT   MonitorRenotifyStatusType = "alert"
	MONITORRENOTIFYSTATUSTYPE_WARN    MonitorRenotifyStatusType = "warn"
	MONITORRENOTIFYSTATUSTYPE_NO_DATA MonitorRenotifyStatusType = "no data"
)

var allowedMonitorRenotifyStatusTypeEnumValues = []MonitorRenotifyStatusType{
	MONITORRENOTIFYSTATUSTYPE_ALERT,
	MONITORRENOTIFYSTATUSTYPE_WARN,
	MONITORRENOTIFYSTATUSTYPE_NO_DATA,
}

func (w *MonitorRenotifyStatusType) GetAllowedValues() []MonitorRenotifyStatusType {
	return allowedMonitorRenotifyStatusTypeEnumValues
}

func (v *MonitorRenotifyStatusType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MonitorRenotifyStatusType(value)
	return nil
}

// NewMonitorRenotifyStatusTypeFromValue returns a pointer to a valid MonitorRenotifyStatusType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMonitorRenotifyStatusTypeFromValue(v string) (*MonitorRenotifyStatusType, error) {
	ev := MonitorRenotifyStatusType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MonitorRenotifyStatusType: valid values are %v", v, allowedMonitorRenotifyStatusTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MonitorRenotifyStatusType) IsValid() bool {
	for _, existing := range allowedMonitorRenotifyStatusTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MonitorRenotifyStatusType value
func (v MonitorRenotifyStatusType) Ptr() *MonitorRenotifyStatusType {
	return &v
}

type NullableMonitorRenotifyStatusType struct {
	value *MonitorRenotifyStatusType
	isSet bool
}

func (v NullableMonitorRenotifyStatusType) Get() *MonitorRenotifyStatusType {
	return v.value
}

func (v *NullableMonitorRenotifyStatusType) Set(val *MonitorRenotifyStatusType) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorRenotifyStatusType) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorRenotifyStatusType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorRenotifyStatusType(val *MonitorRenotifyStatusType) *NullableMonitorRenotifyStatusType {
	return &NullableMonitorRenotifyStatusType{value: val, isSet: true}
}

func (v NullableMonitorRenotifyStatusType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorRenotifyStatusType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
