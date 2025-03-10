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

// SecurityMonitoringRuleMaxSignalDuration A signal will “close” regardless of the query being matched once the time exceeds the maximum duration.
// This time is calculated from the first seen timestamp.
type SecurityMonitoringRuleMaxSignalDuration int32

// List of SecurityMonitoringRuleMaxSignalDuration
const (
	SECURITYMONITORINGRULEMAXSIGNALDURATION_ZERO_MINUTES    SecurityMonitoringRuleMaxSignalDuration = 0
	SECURITYMONITORINGRULEMAXSIGNALDURATION_ONE_MINUTE      SecurityMonitoringRuleMaxSignalDuration = 60
	SECURITYMONITORINGRULEMAXSIGNALDURATION_FIVE_MINUTES    SecurityMonitoringRuleMaxSignalDuration = 300
	SECURITYMONITORINGRULEMAXSIGNALDURATION_TEN_MINUTES     SecurityMonitoringRuleMaxSignalDuration = 600
	SECURITYMONITORINGRULEMAXSIGNALDURATION_FIFTEEN_MINUTES SecurityMonitoringRuleMaxSignalDuration = 900
	SECURITYMONITORINGRULEMAXSIGNALDURATION_THIRTY_MINUTES  SecurityMonitoringRuleMaxSignalDuration = 1800
	SECURITYMONITORINGRULEMAXSIGNALDURATION_ONE_HOUR        SecurityMonitoringRuleMaxSignalDuration = 3600
	SECURITYMONITORINGRULEMAXSIGNALDURATION_TWO_HOURS       SecurityMonitoringRuleMaxSignalDuration = 7200
	SECURITYMONITORINGRULEMAXSIGNALDURATION_THREE_HOURS     SecurityMonitoringRuleMaxSignalDuration = 10800
	SECURITYMONITORINGRULEMAXSIGNALDURATION_SIX_HOURS       SecurityMonitoringRuleMaxSignalDuration = 21600
	SECURITYMONITORINGRULEMAXSIGNALDURATION_TWELVE_HOURS    SecurityMonitoringRuleMaxSignalDuration = 43200
	SECURITYMONITORINGRULEMAXSIGNALDURATION_ONE_DAY         SecurityMonitoringRuleMaxSignalDuration = 86400
)

var allowedSecurityMonitoringRuleMaxSignalDurationEnumValues = []SecurityMonitoringRuleMaxSignalDuration{
	SECURITYMONITORINGRULEMAXSIGNALDURATION_ZERO_MINUTES,
	SECURITYMONITORINGRULEMAXSIGNALDURATION_ONE_MINUTE,
	SECURITYMONITORINGRULEMAXSIGNALDURATION_FIVE_MINUTES,
	SECURITYMONITORINGRULEMAXSIGNALDURATION_TEN_MINUTES,
	SECURITYMONITORINGRULEMAXSIGNALDURATION_FIFTEEN_MINUTES,
	SECURITYMONITORINGRULEMAXSIGNALDURATION_THIRTY_MINUTES,
	SECURITYMONITORINGRULEMAXSIGNALDURATION_ONE_HOUR,
	SECURITYMONITORINGRULEMAXSIGNALDURATION_TWO_HOURS,
	SECURITYMONITORINGRULEMAXSIGNALDURATION_THREE_HOURS,
	SECURITYMONITORINGRULEMAXSIGNALDURATION_SIX_HOURS,
	SECURITYMONITORINGRULEMAXSIGNALDURATION_TWELVE_HOURS,
	SECURITYMONITORINGRULEMAXSIGNALDURATION_ONE_DAY,
}

func (w *SecurityMonitoringRuleMaxSignalDuration) GetAllowedValues() []SecurityMonitoringRuleMaxSignalDuration {
	return allowedSecurityMonitoringRuleMaxSignalDurationEnumValues
}

func (v *SecurityMonitoringRuleMaxSignalDuration) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringRuleMaxSignalDuration(value)
	return nil
}

// NewSecurityMonitoringRuleMaxSignalDurationFromValue returns a pointer to a valid SecurityMonitoringRuleMaxSignalDuration
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSecurityMonitoringRuleMaxSignalDurationFromValue(v int32) (*SecurityMonitoringRuleMaxSignalDuration, error) {
	ev := SecurityMonitoringRuleMaxSignalDuration(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringRuleMaxSignalDuration: valid values are %v", v, allowedSecurityMonitoringRuleMaxSignalDurationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SecurityMonitoringRuleMaxSignalDuration) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringRuleMaxSignalDurationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringRuleMaxSignalDuration value
func (v SecurityMonitoringRuleMaxSignalDuration) Ptr() *SecurityMonitoringRuleMaxSignalDuration {
	return &v
}

type NullableSecurityMonitoringRuleMaxSignalDuration struct {
	value *SecurityMonitoringRuleMaxSignalDuration
	isSet bool
}

func (v NullableSecurityMonitoringRuleMaxSignalDuration) Get() *SecurityMonitoringRuleMaxSignalDuration {
	return v.value
}

func (v *NullableSecurityMonitoringRuleMaxSignalDuration) Set(val *SecurityMonitoringRuleMaxSignalDuration) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityMonitoringRuleMaxSignalDuration) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityMonitoringRuleMaxSignalDuration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityMonitoringRuleMaxSignalDuration(val *SecurityMonitoringRuleMaxSignalDuration) *NullableSecurityMonitoringRuleMaxSignalDuration {
	return &NullableSecurityMonitoringRuleMaxSignalDuration{value: val, isSet: true}
}

func (v NullableSecurityMonitoringRuleMaxSignalDuration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityMonitoringRuleMaxSignalDuration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
