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

// QuerySortOrder Direction of sort.
type QuerySortOrder string

// List of QuerySortOrder
const (
	QUERYSORTORDER_ASC  QuerySortOrder = "asc"
	QUERYSORTORDER_DESC QuerySortOrder = "desc"
)

var allowedQuerySortOrderEnumValues = []QuerySortOrder{
	QUERYSORTORDER_ASC,
	QUERYSORTORDER_DESC,
}

func (w *QuerySortOrder) GetAllowedValues() []QuerySortOrder {
	return allowedQuerySortOrderEnumValues
}

func (v *QuerySortOrder) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = QuerySortOrder(value)
	return nil
}

// NewQuerySortOrderFromValue returns a pointer to a valid QuerySortOrder
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQuerySortOrderFromValue(v string) (*QuerySortOrder, error) {
	ev := QuerySortOrder(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QuerySortOrder: valid values are %v", v, allowedQuerySortOrderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QuerySortOrder) IsValid() bool {
	for _, existing := range allowedQuerySortOrderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QuerySortOrder value
func (v QuerySortOrder) Ptr() *QuerySortOrder {
	return &v
}

type NullableQuerySortOrder struct {
	value *QuerySortOrder
	isSet bool
}

func (v NullableQuerySortOrder) Get() *QuerySortOrder {
	return v.value
}

func (v *NullableQuerySortOrder) Set(val *QuerySortOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableQuerySortOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableQuerySortOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuerySortOrder(val *QuerySortOrder) *NullableQuerySortOrder {
	return &NullableQuerySortOrder{value: val, isSet: true}
}

func (v NullableQuerySortOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuerySortOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
