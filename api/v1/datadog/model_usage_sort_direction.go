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

// UsageSortDirection The direction to sort by.
type UsageSortDirection string

// List of UsageSortDirection
const (
	USAGESORTDIRECTION_DESC UsageSortDirection = "desc"
	USAGESORTDIRECTION_ASC  UsageSortDirection = "asc"
)

var allowedUsageSortDirectionEnumValues = []UsageSortDirection{
	USAGESORTDIRECTION_DESC,
	USAGESORTDIRECTION_ASC,
}

func (w *UsageSortDirection) GetAllowedValues() []UsageSortDirection {
	return allowedUsageSortDirectionEnumValues
}

func (v *UsageSortDirection) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UsageSortDirection(value)
	return nil
}

// NewUsageSortDirectionFromValue returns a pointer to a valid UsageSortDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUsageSortDirectionFromValue(v string) (*UsageSortDirection, error) {
	ev := UsageSortDirection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UsageSortDirection: valid values are %v", v, allowedUsageSortDirectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UsageSortDirection) IsValid() bool {
	for _, existing := range allowedUsageSortDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UsageSortDirection value
func (v UsageSortDirection) Ptr() *UsageSortDirection {
	return &v
}

type NullableUsageSortDirection struct {
	value *UsageSortDirection
	isSet bool
}

func (v NullableUsageSortDirection) Get() *UsageSortDirection {
	return v.value
}

func (v *NullableUsageSortDirection) Set(val *UsageSortDirection) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageSortDirection) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageSortDirection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageSortDirection(val *UsageSortDirection) *NullableUsageSortDirection {
	return &NullableUsageSortDirection{value: val, isSet: true}
}

func (v NullableUsageSortDirection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageSortDirection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
