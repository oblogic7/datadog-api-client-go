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

// MonitorFormulaAndFunctionEventAggregation Aggregation methods for event platform queries.
type MonitorFormulaAndFunctionEventAggregation string

// List of MonitorFormulaAndFunctionEventAggregation
const (
	MONITORFORMULAANDFUNCTIONEVENTAGGREGATION_COUNT       MonitorFormulaAndFunctionEventAggregation = "count"
	MONITORFORMULAANDFUNCTIONEVENTAGGREGATION_CARDINALITY MonitorFormulaAndFunctionEventAggregation = "cardinality"
	MONITORFORMULAANDFUNCTIONEVENTAGGREGATION_MEDIAN      MonitorFormulaAndFunctionEventAggregation = "median"
	MONITORFORMULAANDFUNCTIONEVENTAGGREGATION_PC75        MonitorFormulaAndFunctionEventAggregation = "pc75"
	MONITORFORMULAANDFUNCTIONEVENTAGGREGATION_PC90        MonitorFormulaAndFunctionEventAggregation = "pc90"
	MONITORFORMULAANDFUNCTIONEVENTAGGREGATION_PC95        MonitorFormulaAndFunctionEventAggregation = "pc95"
	MONITORFORMULAANDFUNCTIONEVENTAGGREGATION_PC98        MonitorFormulaAndFunctionEventAggregation = "pc98"
	MONITORFORMULAANDFUNCTIONEVENTAGGREGATION_PC99        MonitorFormulaAndFunctionEventAggregation = "pc99"
	MONITORFORMULAANDFUNCTIONEVENTAGGREGATION_SUM         MonitorFormulaAndFunctionEventAggregation = "sum"
	MONITORFORMULAANDFUNCTIONEVENTAGGREGATION_MIN         MonitorFormulaAndFunctionEventAggregation = "min"
	MONITORFORMULAANDFUNCTIONEVENTAGGREGATION_MAX         MonitorFormulaAndFunctionEventAggregation = "max"
	MONITORFORMULAANDFUNCTIONEVENTAGGREGATION_AVG         MonitorFormulaAndFunctionEventAggregation = "avg"
)

var allowedMonitorFormulaAndFunctionEventAggregationEnumValues = []MonitorFormulaAndFunctionEventAggregation{
	MONITORFORMULAANDFUNCTIONEVENTAGGREGATION_COUNT,
	MONITORFORMULAANDFUNCTIONEVENTAGGREGATION_CARDINALITY,
	MONITORFORMULAANDFUNCTIONEVENTAGGREGATION_MEDIAN,
	MONITORFORMULAANDFUNCTIONEVENTAGGREGATION_PC75,
	MONITORFORMULAANDFUNCTIONEVENTAGGREGATION_PC90,
	MONITORFORMULAANDFUNCTIONEVENTAGGREGATION_PC95,
	MONITORFORMULAANDFUNCTIONEVENTAGGREGATION_PC98,
	MONITORFORMULAANDFUNCTIONEVENTAGGREGATION_PC99,
	MONITORFORMULAANDFUNCTIONEVENTAGGREGATION_SUM,
	MONITORFORMULAANDFUNCTIONEVENTAGGREGATION_MIN,
	MONITORFORMULAANDFUNCTIONEVENTAGGREGATION_MAX,
	MONITORFORMULAANDFUNCTIONEVENTAGGREGATION_AVG,
}

func (w *MonitorFormulaAndFunctionEventAggregation) GetAllowedValues() []MonitorFormulaAndFunctionEventAggregation {
	return allowedMonitorFormulaAndFunctionEventAggregationEnumValues
}

func (v *MonitorFormulaAndFunctionEventAggregation) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MonitorFormulaAndFunctionEventAggregation(value)
	return nil
}

// NewMonitorFormulaAndFunctionEventAggregationFromValue returns a pointer to a valid MonitorFormulaAndFunctionEventAggregation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMonitorFormulaAndFunctionEventAggregationFromValue(v string) (*MonitorFormulaAndFunctionEventAggregation, error) {
	ev := MonitorFormulaAndFunctionEventAggregation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MonitorFormulaAndFunctionEventAggregation: valid values are %v", v, allowedMonitorFormulaAndFunctionEventAggregationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MonitorFormulaAndFunctionEventAggregation) IsValid() bool {
	for _, existing := range allowedMonitorFormulaAndFunctionEventAggregationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MonitorFormulaAndFunctionEventAggregation value
func (v MonitorFormulaAndFunctionEventAggregation) Ptr() *MonitorFormulaAndFunctionEventAggregation {
	return &v
}

type NullableMonitorFormulaAndFunctionEventAggregation struct {
	value *MonitorFormulaAndFunctionEventAggregation
	isSet bool
}

func (v NullableMonitorFormulaAndFunctionEventAggregation) Get() *MonitorFormulaAndFunctionEventAggregation {
	return v.value
}

func (v *NullableMonitorFormulaAndFunctionEventAggregation) Set(val *MonitorFormulaAndFunctionEventAggregation) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorFormulaAndFunctionEventAggregation) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorFormulaAndFunctionEventAggregation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorFormulaAndFunctionEventAggregation(val *MonitorFormulaAndFunctionEventAggregation) *NullableMonitorFormulaAndFunctionEventAggregation {
	return &NullableMonitorFormulaAndFunctionEventAggregation{value: val, isSet: true}
}

func (v NullableMonitorFormulaAndFunctionEventAggregation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorFormulaAndFunctionEventAggregation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
