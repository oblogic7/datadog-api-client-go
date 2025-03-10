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

// ListStreamResponseFormat Widget response format.
type ListStreamResponseFormat string

// List of ListStreamResponseFormat
const (
	LISTSTREAMRESPONSEFORMAT_EVENT_LIST ListStreamResponseFormat = "event_list"
)

var allowedListStreamResponseFormatEnumValues = []ListStreamResponseFormat{
	LISTSTREAMRESPONSEFORMAT_EVENT_LIST,
}

func (w *ListStreamResponseFormat) GetAllowedValues() []ListStreamResponseFormat {
	return allowedListStreamResponseFormatEnumValues
}

func (v *ListStreamResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ListStreamResponseFormat(value)
	return nil
}

// NewListStreamResponseFormatFromValue returns a pointer to a valid ListStreamResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewListStreamResponseFormatFromValue(v string) (*ListStreamResponseFormat, error) {
	ev := ListStreamResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ListStreamResponseFormat: valid values are %v", v, allowedListStreamResponseFormatEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ListStreamResponseFormat) IsValid() bool {
	for _, existing := range allowedListStreamResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ListStreamResponseFormat value
func (v ListStreamResponseFormat) Ptr() *ListStreamResponseFormat {
	return &v
}

type NullableListStreamResponseFormat struct {
	value *ListStreamResponseFormat
	isSet bool
}

func (v NullableListStreamResponseFormat) Get() *ListStreamResponseFormat {
	return v.value
}

func (v *NullableListStreamResponseFormat) Set(val *ListStreamResponseFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableListStreamResponseFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableListStreamResponseFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListStreamResponseFormat(val *ListStreamResponseFormat) *NullableListStreamResponseFormat {
	return &NullableListStreamResponseFormat{value: val, isSet: true}
}

func (v NullableListStreamResponseFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListStreamResponseFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
