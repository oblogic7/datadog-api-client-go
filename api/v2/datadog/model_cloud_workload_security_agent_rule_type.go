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

// CloudWorkloadSecurityAgentRuleType The type of the resource. The value should always be `agent_rule`.
type CloudWorkloadSecurityAgentRuleType string

// List of CloudWorkloadSecurityAgentRuleType
const (
	CLOUDWORKLOADSECURITYAGENTRULETYPE_AGENT_RULE CloudWorkloadSecurityAgentRuleType = "agent_rule"
)

var allowedCloudWorkloadSecurityAgentRuleTypeEnumValues = []CloudWorkloadSecurityAgentRuleType{
	CLOUDWORKLOADSECURITYAGENTRULETYPE_AGENT_RULE,
}

func (w *CloudWorkloadSecurityAgentRuleType) GetAllowedValues() []CloudWorkloadSecurityAgentRuleType {
	return allowedCloudWorkloadSecurityAgentRuleTypeEnumValues
}

func (v *CloudWorkloadSecurityAgentRuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CloudWorkloadSecurityAgentRuleType(value)
	return nil
}

// NewCloudWorkloadSecurityAgentRuleTypeFromValue returns a pointer to a valid CloudWorkloadSecurityAgentRuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCloudWorkloadSecurityAgentRuleTypeFromValue(v string) (*CloudWorkloadSecurityAgentRuleType, error) {
	ev := CloudWorkloadSecurityAgentRuleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CloudWorkloadSecurityAgentRuleType: valid values are %v", v, allowedCloudWorkloadSecurityAgentRuleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CloudWorkloadSecurityAgentRuleType) IsValid() bool {
	for _, existing := range allowedCloudWorkloadSecurityAgentRuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CloudWorkloadSecurityAgentRuleType value
func (v CloudWorkloadSecurityAgentRuleType) Ptr() *CloudWorkloadSecurityAgentRuleType {
	return &v
}

type NullableCloudWorkloadSecurityAgentRuleType struct {
	value *CloudWorkloadSecurityAgentRuleType
	isSet bool
}

func (v NullableCloudWorkloadSecurityAgentRuleType) Get() *CloudWorkloadSecurityAgentRuleType {
	return v.value
}

func (v *NullableCloudWorkloadSecurityAgentRuleType) Set(val *CloudWorkloadSecurityAgentRuleType) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudWorkloadSecurityAgentRuleType) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudWorkloadSecurityAgentRuleType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudWorkloadSecurityAgentRuleType(val *CloudWorkloadSecurityAgentRuleType) *NullableCloudWorkloadSecurityAgentRuleType {
	return &NullableCloudWorkloadSecurityAgentRuleType{value: val, isSet: true}
}

func (v NullableCloudWorkloadSecurityAgentRuleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudWorkloadSecurityAgentRuleType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
