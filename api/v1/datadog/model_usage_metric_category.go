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

// UsageMetricCategory Contains the metric category.
type UsageMetricCategory string

// List of UsageMetricCategory
const (
	USAGEMETRICCATEGORY_STANDARD UsageMetricCategory = "standard"
	USAGEMETRICCATEGORY_CUSTOM   UsageMetricCategory = "custom"
)

var allowedUsageMetricCategoryEnumValues = []UsageMetricCategory{
	USAGEMETRICCATEGORY_STANDARD,
	USAGEMETRICCATEGORY_CUSTOM,
}

func (w *UsageMetricCategory) GetAllowedValues() []UsageMetricCategory {
	return allowedUsageMetricCategoryEnumValues
}

func (v *UsageMetricCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UsageMetricCategory(value)
	return nil
}

// NewUsageMetricCategoryFromValue returns a pointer to a valid UsageMetricCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUsageMetricCategoryFromValue(v string) (*UsageMetricCategory, error) {
	ev := UsageMetricCategory(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UsageMetricCategory: valid values are %v", v, allowedUsageMetricCategoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UsageMetricCategory) IsValid() bool {
	for _, existing := range allowedUsageMetricCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UsageMetricCategory value
func (v UsageMetricCategory) Ptr() *UsageMetricCategory {
	return &v
}

type NullableUsageMetricCategory struct {
	value *UsageMetricCategory
	isSet bool
}

func (v NullableUsageMetricCategory) Get() *UsageMetricCategory {
	return v.value
}

func (v *NullableUsageMetricCategory) Set(val *UsageMetricCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageMetricCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageMetricCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageMetricCategory(val *UsageMetricCategory) *NullableUsageMetricCategory {
	return &NullableUsageMetricCategory{value: val, isSet: true}
}

func (v NullableUsageMetricCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageMetricCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
