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

// HourlyUsageAttributionUsageType Supported products for hourly usage attribution requests.
type HourlyUsageAttributionUsageType string

// List of HourlyUsageAttributionUsageType
const (
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_API_USAGE                    HourlyUsageAttributionUsageType = "api_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_APM_HOST_USAGE               HourlyUsageAttributionUsageType = "apm_host_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_BROWSER_USAGE                HourlyUsageAttributionUsageType = "browser_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CONTAINER_USAGE              HourlyUsageAttributionUsageType = "container_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CUSTOM_TIMESERIES_USAGE      HourlyUsageAttributionUsageType = "custom_timeseries_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_ESTIMATED_INDEXED_LOGS_USAGE HourlyUsageAttributionUsageType = "estimated_indexed_logs_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_FARGATE_USAGE                HourlyUsageAttributionUsageType = "fargate_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_FUNCTIONS_USAGE              HourlyUsageAttributionUsageType = "functions_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_INDEXED_LOGS_USAGE           HourlyUsageAttributionUsageType = "indexed_logs_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_INFRA_HOST_USAGE             HourlyUsageAttributionUsageType = "infra_host_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_INVOCATIONS_USAGE            HourlyUsageAttributionUsageType = "invocations_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_NPM_HOST_USAGE               HourlyUsageAttributionUsageType = "npm_host_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_PROFILED_CONTAINER_USAGE     HourlyUsageAttributionUsageType = "profiled_container_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_PROFILED_HOST_USAGE          HourlyUsageAttributionUsageType = "profiled_host_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_SNMP_USAGE                   HourlyUsageAttributionUsageType = "snmp_usage"
)

var allowedHourlyUsageAttributionUsageTypeEnumValues = []HourlyUsageAttributionUsageType{
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_API_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_APM_HOST_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_BROWSER_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CONTAINER_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CUSTOM_TIMESERIES_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_ESTIMATED_INDEXED_LOGS_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_FARGATE_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_FUNCTIONS_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_INDEXED_LOGS_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_INFRA_HOST_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_INVOCATIONS_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_NPM_HOST_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_PROFILED_CONTAINER_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_PROFILED_HOST_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_SNMP_USAGE,
}

func (w *HourlyUsageAttributionUsageType) GetAllowedValues() []HourlyUsageAttributionUsageType {
	return allowedHourlyUsageAttributionUsageTypeEnumValues
}

func (v *HourlyUsageAttributionUsageType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HourlyUsageAttributionUsageType(value)
	return nil
}

// NewHourlyUsageAttributionUsageTypeFromValue returns a pointer to a valid HourlyUsageAttributionUsageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHourlyUsageAttributionUsageTypeFromValue(v string) (*HourlyUsageAttributionUsageType, error) {
	ev := HourlyUsageAttributionUsageType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for HourlyUsageAttributionUsageType: valid values are %v", v, allowedHourlyUsageAttributionUsageTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HourlyUsageAttributionUsageType) IsValid() bool {
	for _, existing := range allowedHourlyUsageAttributionUsageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HourlyUsageAttributionUsageType value
func (v HourlyUsageAttributionUsageType) Ptr() *HourlyUsageAttributionUsageType {
	return &v
}

type NullableHourlyUsageAttributionUsageType struct {
	value *HourlyUsageAttributionUsageType
	isSet bool
}

func (v NullableHourlyUsageAttributionUsageType) Get() *HourlyUsageAttributionUsageType {
	return v.value
}

func (v *NullableHourlyUsageAttributionUsageType) Set(val *HourlyUsageAttributionUsageType) {
	v.value = val
	v.isSet = true
}

func (v NullableHourlyUsageAttributionUsageType) IsSet() bool {
	return v.isSet
}

func (v *NullableHourlyUsageAttributionUsageType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHourlyUsageAttributionUsageType(val *HourlyUsageAttributionUsageType) *NullableHourlyUsageAttributionUsageType {
	return &NullableHourlyUsageAttributionUsageType{value: val, isSet: true}
}

func (v NullableHourlyUsageAttributionUsageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHourlyUsageAttributionUsageType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
