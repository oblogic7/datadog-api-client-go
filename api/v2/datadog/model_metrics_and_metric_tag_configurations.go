/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// MetricsAndMetricTagConfigurations - Object for a metrics and metric tag configurations.
type MetricsAndMetricTagConfigurations struct {
	Metric                 *Metric
	MetricTagConfiguration *MetricTagConfiguration

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// MetricAsMetricsAndMetricTagConfigurations is a convenience function that returns Metric wrapped in MetricsAndMetricTagConfigurations
func MetricAsMetricsAndMetricTagConfigurations(v *Metric) MetricsAndMetricTagConfigurations {
	return MetricsAndMetricTagConfigurations{Metric: v}
}

// MetricTagConfigurationAsMetricsAndMetricTagConfigurations is a convenience function that returns MetricTagConfiguration wrapped in MetricsAndMetricTagConfigurations
func MetricTagConfigurationAsMetricsAndMetricTagConfigurations(v *MetricTagConfiguration) MetricsAndMetricTagConfigurations {
	return MetricsAndMetricTagConfigurations{MetricTagConfiguration: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *MetricsAndMetricTagConfigurations) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Metric
	err = json.Unmarshal(data, &dst.Metric)
	if err == nil {
		if dst.Metric != nil && dst.Metric.UnparsedObject == nil {
			jsonMetric, _ := json.Marshal(dst.Metric)
			if string(jsonMetric) == "{}" { // empty struct
				dst.Metric = nil
			} else {
				match++
			}
		} else {
			dst.Metric = nil
		}
	} else {
		dst.Metric = nil
	}

	// try to unmarshal data into MetricTagConfiguration
	err = json.Unmarshal(data, &dst.MetricTagConfiguration)
	if err == nil {
		if dst.MetricTagConfiguration != nil && dst.MetricTagConfiguration.UnparsedObject == nil {
			jsonMetricTagConfiguration, _ := json.Marshal(dst.MetricTagConfiguration)
			if string(jsonMetricTagConfiguration) == "{}" { // empty struct
				dst.MetricTagConfiguration = nil
			} else {
				match++
			}
		} else {
			dst.MetricTagConfiguration = nil
		}
	} else {
		dst.MetricTagConfiguration = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		dst.Metric = nil
		dst.MetricTagConfiguration = nil
		return json.Unmarshal(data, &dst.UnparsedObject)
	} else {
		return nil // exactly one match
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MetricsAndMetricTagConfigurations) MarshalJSON() ([]byte, error) {
	if src.Metric != nil {
		return json.Marshal(&src.Metric)
	}

	if src.MetricTagConfiguration != nil {
		return json.Marshal(&src.MetricTagConfiguration)
	}

	if src.UnparsedObject != nil {
		return json.Marshal(src.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *MetricsAndMetricTagConfigurations) GetActualInstance() interface{} {
	if obj.Metric != nil {
		return obj.Metric
	}

	if obj.MetricTagConfiguration != nil {
		return obj.MetricTagConfiguration
	}

	// all schemas are nil
	return nil
}

type NullableMetricsAndMetricTagConfigurations struct {
	value *MetricsAndMetricTagConfigurations
	isSet bool
}

func (v NullableMetricsAndMetricTagConfigurations) Get() *MetricsAndMetricTagConfigurations {
	return v.value
}

func (v *NullableMetricsAndMetricTagConfigurations) Set(val *MetricsAndMetricTagConfigurations) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsAndMetricTagConfigurations) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsAndMetricTagConfigurations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsAndMetricTagConfigurations(val *MetricsAndMetricTagConfigurations) *NullableMetricsAndMetricTagConfigurations {
	return &NullableMetricsAndMetricTagConfigurations{value: val, isSet: true}
}

func (v NullableMetricsAndMetricTagConfigurations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsAndMetricTagConfigurations) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return json.Unmarshal(src, &v.value)
}
