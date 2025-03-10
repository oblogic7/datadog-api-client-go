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

// MonitorType The type of the monitor. For more information about `type`, see the [monitor options](https://docs.datadoghq.com/monitors/guide/monitor_api_options/) docs.
type MonitorType string

// List of MonitorType
const (
	MONITORTYPE_COMPOSITE             MonitorType = "composite"
	MONITORTYPE_EVENT_ALERT           MonitorType = "event alert"
	MONITORTYPE_LOG_ALERT             MonitorType = "log alert"
	MONITORTYPE_METRIC_ALERT          MonitorType = "metric alert"
	MONITORTYPE_PROCESS_ALERT         MonitorType = "process alert"
	MONITORTYPE_QUERY_ALERT           MonitorType = "query alert"
	MONITORTYPE_RUM_ALERT             MonitorType = "rum alert"
	MONITORTYPE_SERVICE_CHECK         MonitorType = "service check"
	MONITORTYPE_SYNTHETICS_ALERT      MonitorType = "synthetics alert"
	MONITORTYPE_TRACE_ANALYTICS_ALERT MonitorType = "trace-analytics alert"
	MONITORTYPE_SLO_ALERT             MonitorType = "slo alert"
	MONITORTYPE_EVENT_V2_ALERT        MonitorType = "event-v2 alert"
	MONITORTYPE_AUDIT_ALERT           MonitorType = "audit alert"
	MONITORTYPE_CI_PIPELINES_ALERT    MonitorType = "ci-pipelines alert"
	MONITORTYPE_CI_TESTS_ALERT        MonitorType = "ci-tests alert"
	MONITORTYPE_ERROR_TRACKING_ALERT  MonitorType = "error-tracking alert"
)

var allowedMonitorTypeEnumValues = []MonitorType{
	MONITORTYPE_COMPOSITE,
	MONITORTYPE_EVENT_ALERT,
	MONITORTYPE_LOG_ALERT,
	MONITORTYPE_METRIC_ALERT,
	MONITORTYPE_PROCESS_ALERT,
	MONITORTYPE_QUERY_ALERT,
	MONITORTYPE_RUM_ALERT,
	MONITORTYPE_SERVICE_CHECK,
	MONITORTYPE_SYNTHETICS_ALERT,
	MONITORTYPE_TRACE_ANALYTICS_ALERT,
	MONITORTYPE_SLO_ALERT,
	MONITORTYPE_EVENT_V2_ALERT,
	MONITORTYPE_AUDIT_ALERT,
	MONITORTYPE_CI_PIPELINES_ALERT,
	MONITORTYPE_CI_TESTS_ALERT,
	MONITORTYPE_ERROR_TRACKING_ALERT,
}

func (w *MonitorType) GetAllowedValues() []MonitorType {
	return allowedMonitorTypeEnumValues
}

func (v *MonitorType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MonitorType(value)
	return nil
}

// NewMonitorTypeFromValue returns a pointer to a valid MonitorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMonitorTypeFromValue(v string) (*MonitorType, error) {
	ev := MonitorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MonitorType: valid values are %v", v, allowedMonitorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MonitorType) IsValid() bool {
	for _, existing := range allowedMonitorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MonitorType value
func (v MonitorType) Ptr() *MonitorType {
	return &v
}

type NullableMonitorType struct {
	value *MonitorType
	isSet bool
}

func (v NullableMonitorType) Get() *MonitorType {
	return v.value
}

func (v *NullableMonitorType) Set(val *MonitorType) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorType) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorType(val *MonitorType) *NullableMonitorType {
	return &NullableMonitorType{value: val, isSet: true}
}

func (v NullableMonitorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
