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

// WidgetMonitorSummaryDisplayFormat What to display on the widget.
type WidgetMonitorSummaryDisplayFormat string

// List of WidgetMonitorSummaryDisplayFormat
const (
	WIDGETMONITORSUMMARYDISPLAYFORMAT_COUNTS          WidgetMonitorSummaryDisplayFormat = "counts"
	WIDGETMONITORSUMMARYDISPLAYFORMAT_COUNTS_AND_LIST WidgetMonitorSummaryDisplayFormat = "countsAndList"
	WIDGETMONITORSUMMARYDISPLAYFORMAT_LIST            WidgetMonitorSummaryDisplayFormat = "list"
)

var allowedWidgetMonitorSummaryDisplayFormatEnumValues = []WidgetMonitorSummaryDisplayFormat{
	WIDGETMONITORSUMMARYDISPLAYFORMAT_COUNTS,
	WIDGETMONITORSUMMARYDISPLAYFORMAT_COUNTS_AND_LIST,
	WIDGETMONITORSUMMARYDISPLAYFORMAT_LIST,
}

func (w *WidgetMonitorSummaryDisplayFormat) GetAllowedValues() []WidgetMonitorSummaryDisplayFormat {
	return allowedWidgetMonitorSummaryDisplayFormatEnumValues
}

func (v *WidgetMonitorSummaryDisplayFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WidgetMonitorSummaryDisplayFormat(value)
	return nil
}

// NewWidgetMonitorSummaryDisplayFormatFromValue returns a pointer to a valid WidgetMonitorSummaryDisplayFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWidgetMonitorSummaryDisplayFormatFromValue(v string) (*WidgetMonitorSummaryDisplayFormat, error) {
	ev := WidgetMonitorSummaryDisplayFormat(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WidgetMonitorSummaryDisplayFormat: valid values are %v", v, allowedWidgetMonitorSummaryDisplayFormatEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WidgetMonitorSummaryDisplayFormat) IsValid() bool {
	for _, existing := range allowedWidgetMonitorSummaryDisplayFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WidgetMonitorSummaryDisplayFormat value
func (v WidgetMonitorSummaryDisplayFormat) Ptr() *WidgetMonitorSummaryDisplayFormat {
	return &v
}

type NullableWidgetMonitorSummaryDisplayFormat struct {
	value *WidgetMonitorSummaryDisplayFormat
	isSet bool
}

func (v NullableWidgetMonitorSummaryDisplayFormat) Get() *WidgetMonitorSummaryDisplayFormat {
	return v.value
}

func (v *NullableWidgetMonitorSummaryDisplayFormat) Set(val *WidgetMonitorSummaryDisplayFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetMonitorSummaryDisplayFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetMonitorSummaryDisplayFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetMonitorSummaryDisplayFormat(val *WidgetMonitorSummaryDisplayFormat) *NullableWidgetMonitorSummaryDisplayFormat {
	return &NullableWidgetMonitorSummaryDisplayFormat{value: val, isSet: true}
}

func (v NullableWidgetMonitorSummaryDisplayFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetMonitorSummaryDisplayFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
