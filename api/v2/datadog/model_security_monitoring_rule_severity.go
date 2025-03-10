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

// SecurityMonitoringRuleSeverity Severity of the Security Signal.
type SecurityMonitoringRuleSeverity string

// List of SecurityMonitoringRuleSeverity
const (
	SECURITYMONITORINGRULESEVERITY_INFO     SecurityMonitoringRuleSeverity = "info"
	SECURITYMONITORINGRULESEVERITY_LOW      SecurityMonitoringRuleSeverity = "low"
	SECURITYMONITORINGRULESEVERITY_MEDIUM   SecurityMonitoringRuleSeverity = "medium"
	SECURITYMONITORINGRULESEVERITY_HIGH     SecurityMonitoringRuleSeverity = "high"
	SECURITYMONITORINGRULESEVERITY_CRITICAL SecurityMonitoringRuleSeverity = "critical"
)

var allowedSecurityMonitoringRuleSeverityEnumValues = []SecurityMonitoringRuleSeverity{
	SECURITYMONITORINGRULESEVERITY_INFO,
	SECURITYMONITORINGRULESEVERITY_LOW,
	SECURITYMONITORINGRULESEVERITY_MEDIUM,
	SECURITYMONITORINGRULESEVERITY_HIGH,
	SECURITYMONITORINGRULESEVERITY_CRITICAL,
}

func (w *SecurityMonitoringRuleSeverity) GetAllowedValues() []SecurityMonitoringRuleSeverity {
	return allowedSecurityMonitoringRuleSeverityEnumValues
}

func (v *SecurityMonitoringRuleSeverity) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringRuleSeverity(value)
	return nil
}

// NewSecurityMonitoringRuleSeverityFromValue returns a pointer to a valid SecurityMonitoringRuleSeverity
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSecurityMonitoringRuleSeverityFromValue(v string) (*SecurityMonitoringRuleSeverity, error) {
	ev := SecurityMonitoringRuleSeverity(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringRuleSeverity: valid values are %v", v, allowedSecurityMonitoringRuleSeverityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SecurityMonitoringRuleSeverity) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringRuleSeverityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringRuleSeverity value
func (v SecurityMonitoringRuleSeverity) Ptr() *SecurityMonitoringRuleSeverity {
	return &v
}

type NullableSecurityMonitoringRuleSeverity struct {
	value *SecurityMonitoringRuleSeverity
	isSet bool
}

func (v NullableSecurityMonitoringRuleSeverity) Get() *SecurityMonitoringRuleSeverity {
	return v.value
}

func (v *NullableSecurityMonitoringRuleSeverity) Set(val *SecurityMonitoringRuleSeverity) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityMonitoringRuleSeverity) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityMonitoringRuleSeverity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityMonitoringRuleSeverity(val *SecurityMonitoringRuleSeverity) *NullableSecurityMonitoringRuleSeverity {
	return &NullableSecurityMonitoringRuleSeverity{value: val, isSet: true}
}

func (v NullableSecurityMonitoringRuleSeverity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityMonitoringRuleSeverity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
