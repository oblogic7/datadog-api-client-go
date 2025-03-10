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

// LogsMetricResponseComputeAggregationType The type of aggregation to use.
type LogsMetricResponseComputeAggregationType string

// List of LogsMetricResponseComputeAggregationType
const (
	LOGSMETRICRESPONSECOMPUTEAGGREGATIONTYPE_COUNT        LogsMetricResponseComputeAggregationType = "count"
	LOGSMETRICRESPONSECOMPUTEAGGREGATIONTYPE_DISTRIBUTION LogsMetricResponseComputeAggregationType = "distribution"
)

var allowedLogsMetricResponseComputeAggregationTypeEnumValues = []LogsMetricResponseComputeAggregationType{
	LOGSMETRICRESPONSECOMPUTEAGGREGATIONTYPE_COUNT,
	LOGSMETRICRESPONSECOMPUTEAGGREGATIONTYPE_DISTRIBUTION,
}

func (w *LogsMetricResponseComputeAggregationType) GetAllowedValues() []LogsMetricResponseComputeAggregationType {
	return allowedLogsMetricResponseComputeAggregationTypeEnumValues
}

func (v *LogsMetricResponseComputeAggregationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsMetricResponseComputeAggregationType(value)
	return nil
}

// NewLogsMetricResponseComputeAggregationTypeFromValue returns a pointer to a valid LogsMetricResponseComputeAggregationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLogsMetricResponseComputeAggregationTypeFromValue(v string) (*LogsMetricResponseComputeAggregationType, error) {
	ev := LogsMetricResponseComputeAggregationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LogsMetricResponseComputeAggregationType: valid values are %v", v, allowedLogsMetricResponseComputeAggregationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LogsMetricResponseComputeAggregationType) IsValid() bool {
	for _, existing := range allowedLogsMetricResponseComputeAggregationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsMetricResponseComputeAggregationType value
func (v LogsMetricResponseComputeAggregationType) Ptr() *LogsMetricResponseComputeAggregationType {
	return &v
}

type NullableLogsMetricResponseComputeAggregationType struct {
	value *LogsMetricResponseComputeAggregationType
	isSet bool
}

func (v NullableLogsMetricResponseComputeAggregationType) Get() *LogsMetricResponseComputeAggregationType {
	return v.value
}

func (v *NullableLogsMetricResponseComputeAggregationType) Set(val *LogsMetricResponseComputeAggregationType) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsMetricResponseComputeAggregationType) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsMetricResponseComputeAggregationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsMetricResponseComputeAggregationType(val *LogsMetricResponseComputeAggregationType) *NullableLogsMetricResponseComputeAggregationType {
	return &NullableLogsMetricResponseComputeAggregationType{value: val, isSet: true}
}

func (v NullableLogsMetricResponseComputeAggregationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsMetricResponseComputeAggregationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
