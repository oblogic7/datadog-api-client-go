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

// SecurityMonitoringRuleEvaluationWindow A time window is specified to match when at least one of the cases matches true. This is a sliding window
// and evaluates in real time.
type SecurityMonitoringRuleEvaluationWindow int32

// List of SecurityMonitoringRuleEvaluationWindow
const (
	SECURITYMONITORINGRULEEVALUATIONWINDOW_ZERO_MINUTES    SecurityMonitoringRuleEvaluationWindow = 0
	SECURITYMONITORINGRULEEVALUATIONWINDOW_ONE_MINUTE      SecurityMonitoringRuleEvaluationWindow = 60
	SECURITYMONITORINGRULEEVALUATIONWINDOW_FIVE_MINUTES    SecurityMonitoringRuleEvaluationWindow = 300
	SECURITYMONITORINGRULEEVALUATIONWINDOW_TEN_MINUTES     SecurityMonitoringRuleEvaluationWindow = 600
	SECURITYMONITORINGRULEEVALUATIONWINDOW_FIFTEEN_MINUTES SecurityMonitoringRuleEvaluationWindow = 900
	SECURITYMONITORINGRULEEVALUATIONWINDOW_THIRTY_MINUTES  SecurityMonitoringRuleEvaluationWindow = 1800
	SECURITYMONITORINGRULEEVALUATIONWINDOW_ONE_HOUR        SecurityMonitoringRuleEvaluationWindow = 3600
	SECURITYMONITORINGRULEEVALUATIONWINDOW_TWO_HOURS       SecurityMonitoringRuleEvaluationWindow = 7200
)

var allowedSecurityMonitoringRuleEvaluationWindowEnumValues = []SecurityMonitoringRuleEvaluationWindow{
	SECURITYMONITORINGRULEEVALUATIONWINDOW_ZERO_MINUTES,
	SECURITYMONITORINGRULEEVALUATIONWINDOW_ONE_MINUTE,
	SECURITYMONITORINGRULEEVALUATIONWINDOW_FIVE_MINUTES,
	SECURITYMONITORINGRULEEVALUATIONWINDOW_TEN_MINUTES,
	SECURITYMONITORINGRULEEVALUATIONWINDOW_FIFTEEN_MINUTES,
	SECURITYMONITORINGRULEEVALUATIONWINDOW_THIRTY_MINUTES,
	SECURITYMONITORINGRULEEVALUATIONWINDOW_ONE_HOUR,
	SECURITYMONITORINGRULEEVALUATIONWINDOW_TWO_HOURS,
}

func (w *SecurityMonitoringRuleEvaluationWindow) GetAllowedValues() []SecurityMonitoringRuleEvaluationWindow {
	return allowedSecurityMonitoringRuleEvaluationWindowEnumValues
}

func (v *SecurityMonitoringRuleEvaluationWindow) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringRuleEvaluationWindow(value)
	return nil
}

// NewSecurityMonitoringRuleEvaluationWindowFromValue returns a pointer to a valid SecurityMonitoringRuleEvaluationWindow
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSecurityMonitoringRuleEvaluationWindowFromValue(v int32) (*SecurityMonitoringRuleEvaluationWindow, error) {
	ev := SecurityMonitoringRuleEvaluationWindow(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringRuleEvaluationWindow: valid values are %v", v, allowedSecurityMonitoringRuleEvaluationWindowEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SecurityMonitoringRuleEvaluationWindow) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringRuleEvaluationWindowEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringRuleEvaluationWindow value
func (v SecurityMonitoringRuleEvaluationWindow) Ptr() *SecurityMonitoringRuleEvaluationWindow {
	return &v
}

type NullableSecurityMonitoringRuleEvaluationWindow struct {
	value *SecurityMonitoringRuleEvaluationWindow
	isSet bool
}

func (v NullableSecurityMonitoringRuleEvaluationWindow) Get() *SecurityMonitoringRuleEvaluationWindow {
	return v.value
}

func (v *NullableSecurityMonitoringRuleEvaluationWindow) Set(val *SecurityMonitoringRuleEvaluationWindow) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityMonitoringRuleEvaluationWindow) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityMonitoringRuleEvaluationWindow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityMonitoringRuleEvaluationWindow(val *SecurityMonitoringRuleEvaluationWindow) *NullableSecurityMonitoringRuleEvaluationWindow {
	return &NullableSecurityMonitoringRuleEvaluationWindow{value: val, isSet: true}
}

func (v NullableSecurityMonitoringRuleEvaluationWindow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityMonitoringRuleEvaluationWindow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
