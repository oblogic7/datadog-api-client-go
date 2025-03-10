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

// UsageAttributionSort The field to sort by.
type UsageAttributionSort string

// List of UsageAttributionSort
const (
	USAGEATTRIBUTIONSORT_API_PERCENTAGE                    UsageAttributionSort = "api_percentage"
	USAGEATTRIBUTIONSORT_SNMP_USAGE                        UsageAttributionSort = "snmp_usage"
	USAGEATTRIBUTIONSORT_APM_HOST_USAGE                    UsageAttributionSort = "apm_host_usage"
	USAGEATTRIBUTIONSORT_API_USAGE                         UsageAttributionSort = "api_usage"
	USAGEATTRIBUTIONSORT_CONTAINER_USAGE                   UsageAttributionSort = "container_usage"
	USAGEATTRIBUTIONSORT_CUSTOM_TIMESERIES_PERCENTAGE      UsageAttributionSort = "custom_timeseries_percentage"
	USAGEATTRIBUTIONSORT_CONTAINER_PERCENTAGE              UsageAttributionSort = "container_percentage"
	USAGEATTRIBUTIONSORT_APM_HOST_PERCENTAGE               UsageAttributionSort = "apm_host_percentage"
	USAGEATTRIBUTIONSORT_NPM_HOST_PERCENTAGE               UsageAttributionSort = "npm_host_percentage"
	USAGEATTRIBUTIONSORT_BROWSER_PERCENTAGE                UsageAttributionSort = "browser_percentage"
	USAGEATTRIBUTIONSORT_BROWSER_USAGE                     UsageAttributionSort = "browser_usage"
	USAGEATTRIBUTIONSORT_INFRA_HOST_PERCENTAGE             UsageAttributionSort = "infra_host_percentage"
	USAGEATTRIBUTIONSORT_SNMP_PERCENTAGE                   UsageAttributionSort = "snmp_percentage"
	USAGEATTRIBUTIONSORT_NPM_HOST_USAGE                    UsageAttributionSort = "npm_host_usage"
	USAGEATTRIBUTIONSORT_INFRA_HOST_USAGE                  UsageAttributionSort = "infra_host_usage"
	USAGEATTRIBUTIONSORT_CUSTOM_TIMESERIES_USAGE           UsageAttributionSort = "custom_timeseries_usage"
	USAGEATTRIBUTIONSORT_LAMBDA_FUNCTIONS_USAGE            UsageAttributionSort = "lambda_functions_usage"
	USAGEATTRIBUTIONSORT_LAMBDA_FUNCTIONS_PERCENTAGE       UsageAttributionSort = "lambda_functions_percentage"
	USAGEATTRIBUTIONSORT_LAMBDA_INVOCATIONS_USAGE          UsageAttributionSort = "lambda_invocations_usage"
	USAGEATTRIBUTIONSORT_LAMBDA_INVOCATIONS_PERCENTAGE     UsageAttributionSort = "lambda_invocations_percentage"
	USAGEATTRIBUTIONSORT_ESTIMATED_INDEXED_LOGS_USAGE      UsageAttributionSort = "estimated_indexed_logs_usage"
	USAGEATTRIBUTIONSORT_ESTIMATED_INDEXED_LOGS_PERCENTAGE UsageAttributionSort = "estimated_indexed_logs_percentage"
)

var allowedUsageAttributionSortEnumValues = []UsageAttributionSort{
	USAGEATTRIBUTIONSORT_API_PERCENTAGE,
	USAGEATTRIBUTIONSORT_SNMP_USAGE,
	USAGEATTRIBUTIONSORT_APM_HOST_USAGE,
	USAGEATTRIBUTIONSORT_API_USAGE,
	USAGEATTRIBUTIONSORT_CONTAINER_USAGE,
	USAGEATTRIBUTIONSORT_CUSTOM_TIMESERIES_PERCENTAGE,
	USAGEATTRIBUTIONSORT_CONTAINER_PERCENTAGE,
	USAGEATTRIBUTIONSORT_APM_HOST_PERCENTAGE,
	USAGEATTRIBUTIONSORT_NPM_HOST_PERCENTAGE,
	USAGEATTRIBUTIONSORT_BROWSER_PERCENTAGE,
	USAGEATTRIBUTIONSORT_BROWSER_USAGE,
	USAGEATTRIBUTIONSORT_INFRA_HOST_PERCENTAGE,
	USAGEATTRIBUTIONSORT_SNMP_PERCENTAGE,
	USAGEATTRIBUTIONSORT_NPM_HOST_USAGE,
	USAGEATTRIBUTIONSORT_INFRA_HOST_USAGE,
	USAGEATTRIBUTIONSORT_CUSTOM_TIMESERIES_USAGE,
	USAGEATTRIBUTIONSORT_LAMBDA_FUNCTIONS_USAGE,
	USAGEATTRIBUTIONSORT_LAMBDA_FUNCTIONS_PERCENTAGE,
	USAGEATTRIBUTIONSORT_LAMBDA_INVOCATIONS_USAGE,
	USAGEATTRIBUTIONSORT_LAMBDA_INVOCATIONS_PERCENTAGE,
	USAGEATTRIBUTIONSORT_ESTIMATED_INDEXED_LOGS_USAGE,
	USAGEATTRIBUTIONSORT_ESTIMATED_INDEXED_LOGS_PERCENTAGE,
}

func (w *UsageAttributionSort) GetAllowedValues() []UsageAttributionSort {
	return allowedUsageAttributionSortEnumValues
}

func (v *UsageAttributionSort) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UsageAttributionSort(value)
	return nil
}

// NewUsageAttributionSortFromValue returns a pointer to a valid UsageAttributionSort
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUsageAttributionSortFromValue(v string) (*UsageAttributionSort, error) {
	ev := UsageAttributionSort(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UsageAttributionSort: valid values are %v", v, allowedUsageAttributionSortEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UsageAttributionSort) IsValid() bool {
	for _, existing := range allowedUsageAttributionSortEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UsageAttributionSort value
func (v UsageAttributionSort) Ptr() *UsageAttributionSort {
	return &v
}

type NullableUsageAttributionSort struct {
	value *UsageAttributionSort
	isSet bool
}

func (v NullableUsageAttributionSort) Get() *UsageAttributionSort {
	return v.value
}

func (v *NullableUsageAttributionSort) Set(val *UsageAttributionSort) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageAttributionSort) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageAttributionSort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageAttributionSort(val *UsageAttributionSort) *NullableUsageAttributionSort {
	return &NullableUsageAttributionSort{value: val, isSet: true}
}

func (v NullableUsageAttributionSort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageAttributionSort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
