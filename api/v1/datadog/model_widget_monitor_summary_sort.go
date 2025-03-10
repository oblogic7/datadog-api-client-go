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

// WidgetMonitorSummarySort Widget sorting methods.
type WidgetMonitorSummarySort string

// List of WidgetMonitorSummarySort
const (
	WIDGETMONITORSUMMARYSORT_NAME                 WidgetMonitorSummarySort = "name"
	WIDGETMONITORSUMMARYSORT_GROUP                WidgetMonitorSummarySort = "group"
	WIDGETMONITORSUMMARYSORT_STATUS               WidgetMonitorSummarySort = "status"
	WIDGETMONITORSUMMARYSORT_TAGS                 WidgetMonitorSummarySort = "tags"
	WIDGETMONITORSUMMARYSORT_TRIGGERED            WidgetMonitorSummarySort = "triggered"
	WIDGETMONITORSUMMARYSORT_GROUP_ASCENDING      WidgetMonitorSummarySort = "group,asc"
	WIDGETMONITORSUMMARYSORT_GROUP_DESCENDING     WidgetMonitorSummarySort = "group,desc"
	WIDGETMONITORSUMMARYSORT_NAME_ASCENDING       WidgetMonitorSummarySort = "name,asc"
	WIDGETMONITORSUMMARYSORT_NAME_DESCENDING      WidgetMonitorSummarySort = "name,desc"
	WIDGETMONITORSUMMARYSORT_STATUS_ASCENDING     WidgetMonitorSummarySort = "status,asc"
	WIDGETMONITORSUMMARYSORT_STATUS_DESCENDING    WidgetMonitorSummarySort = "status,desc"
	WIDGETMONITORSUMMARYSORT_TAGS_ASCENDING       WidgetMonitorSummarySort = "tags,asc"
	WIDGETMONITORSUMMARYSORT_TAGS_DESCENDING      WidgetMonitorSummarySort = "tags,desc"
	WIDGETMONITORSUMMARYSORT_TRIGGERED_ASCENDING  WidgetMonitorSummarySort = "triggered,asc"
	WIDGETMONITORSUMMARYSORT_TRIGGERED_DESCENDING WidgetMonitorSummarySort = "triggered,desc"
)

var allowedWidgetMonitorSummarySortEnumValues = []WidgetMonitorSummarySort{
	WIDGETMONITORSUMMARYSORT_NAME,
	WIDGETMONITORSUMMARYSORT_GROUP,
	WIDGETMONITORSUMMARYSORT_STATUS,
	WIDGETMONITORSUMMARYSORT_TAGS,
	WIDGETMONITORSUMMARYSORT_TRIGGERED,
	WIDGETMONITORSUMMARYSORT_GROUP_ASCENDING,
	WIDGETMONITORSUMMARYSORT_GROUP_DESCENDING,
	WIDGETMONITORSUMMARYSORT_NAME_ASCENDING,
	WIDGETMONITORSUMMARYSORT_NAME_DESCENDING,
	WIDGETMONITORSUMMARYSORT_STATUS_ASCENDING,
	WIDGETMONITORSUMMARYSORT_STATUS_DESCENDING,
	WIDGETMONITORSUMMARYSORT_TAGS_ASCENDING,
	WIDGETMONITORSUMMARYSORT_TAGS_DESCENDING,
	WIDGETMONITORSUMMARYSORT_TRIGGERED_ASCENDING,
	WIDGETMONITORSUMMARYSORT_TRIGGERED_DESCENDING,
}

func (w *WidgetMonitorSummarySort) GetAllowedValues() []WidgetMonitorSummarySort {
	return allowedWidgetMonitorSummarySortEnumValues
}

func (v *WidgetMonitorSummarySort) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WidgetMonitorSummarySort(value)
	return nil
}

// NewWidgetMonitorSummarySortFromValue returns a pointer to a valid WidgetMonitorSummarySort
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWidgetMonitorSummarySortFromValue(v string) (*WidgetMonitorSummarySort, error) {
	ev := WidgetMonitorSummarySort(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WidgetMonitorSummarySort: valid values are %v", v, allowedWidgetMonitorSummarySortEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WidgetMonitorSummarySort) IsValid() bool {
	for _, existing := range allowedWidgetMonitorSummarySortEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WidgetMonitorSummarySort value
func (v WidgetMonitorSummarySort) Ptr() *WidgetMonitorSummarySort {
	return &v
}

type NullableWidgetMonitorSummarySort struct {
	value *WidgetMonitorSummarySort
	isSet bool
}

func (v NullableWidgetMonitorSummarySort) Get() *WidgetMonitorSummarySort {
	return v.value
}

func (v *NullableWidgetMonitorSummarySort) Set(val *WidgetMonitorSummarySort) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetMonitorSummarySort) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetMonitorSummarySort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetMonitorSummarySort(val *WidgetMonitorSummarySort) *NullableWidgetMonitorSummarySort {
	return &NullableWidgetMonitorSummarySort{value: val, isSet: true}
}

func (v NullableWidgetMonitorSummarySort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetMonitorSummarySort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
