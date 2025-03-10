/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// MetricVolumes - Possible response objects for a metric's volume.
type MetricVolumes struct {
	MetricDistinctVolume        *MetricDistinctVolume
	MetricIngestedIndexedVolume *MetricIngestedIndexedVolume

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// MetricDistinctVolumeAsMetricVolumes is a convenience function that returns MetricDistinctVolume wrapped in MetricVolumes
func MetricDistinctVolumeAsMetricVolumes(v *MetricDistinctVolume) MetricVolumes {
	return MetricVolumes{MetricDistinctVolume: v}
}

// MetricIngestedIndexedVolumeAsMetricVolumes is a convenience function that returns MetricIngestedIndexedVolume wrapped in MetricVolumes
func MetricIngestedIndexedVolumeAsMetricVolumes(v *MetricIngestedIndexedVolume) MetricVolumes {
	return MetricVolumes{MetricIngestedIndexedVolume: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *MetricVolumes) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into MetricDistinctVolume
	err = json.Unmarshal(data, &dst.MetricDistinctVolume)
	if err == nil {
		if dst.MetricDistinctVolume != nil && dst.MetricDistinctVolume.UnparsedObject == nil {
			jsonMetricDistinctVolume, _ := json.Marshal(dst.MetricDistinctVolume)
			if string(jsonMetricDistinctVolume) == "{}" { // empty struct
				dst.MetricDistinctVolume = nil
			} else {
				match++
			}
		} else {
			dst.MetricDistinctVolume = nil
		}
	} else {
		dst.MetricDistinctVolume = nil
	}

	// try to unmarshal data into MetricIngestedIndexedVolume
	err = json.Unmarshal(data, &dst.MetricIngestedIndexedVolume)
	if err == nil {
		if dst.MetricIngestedIndexedVolume != nil && dst.MetricIngestedIndexedVolume.UnparsedObject == nil {
			jsonMetricIngestedIndexedVolume, _ := json.Marshal(dst.MetricIngestedIndexedVolume)
			if string(jsonMetricIngestedIndexedVolume) == "{}" { // empty struct
				dst.MetricIngestedIndexedVolume = nil
			} else {
				match++
			}
		} else {
			dst.MetricIngestedIndexedVolume = nil
		}
	} else {
		dst.MetricIngestedIndexedVolume = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		dst.MetricDistinctVolume = nil
		dst.MetricIngestedIndexedVolume = nil
		return json.Unmarshal(data, &dst.UnparsedObject)
	} else {
		return nil // exactly one match
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MetricVolumes) MarshalJSON() ([]byte, error) {
	if src.MetricDistinctVolume != nil {
		return json.Marshal(&src.MetricDistinctVolume)
	}

	if src.MetricIngestedIndexedVolume != nil {
		return json.Marshal(&src.MetricIngestedIndexedVolume)
	}

	if src.UnparsedObject != nil {
		return json.Marshal(src.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *MetricVolumes) GetActualInstance() interface{} {
	if obj.MetricDistinctVolume != nil {
		return obj.MetricDistinctVolume
	}

	if obj.MetricIngestedIndexedVolume != nil {
		return obj.MetricIngestedIndexedVolume
	}

	// all schemas are nil
	return nil
}

type NullableMetricVolumes struct {
	value *MetricVolumes
	isSet bool
}

func (v NullableMetricVolumes) Get() *MetricVolumes {
	return v.value
}

func (v *NullableMetricVolumes) Set(val *MetricVolumes) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricVolumes) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricVolumes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricVolumes(val *MetricVolumes) *NullableMetricVolumes {
	return &NullableMetricVolumes{value: val, isSet: true}
}

func (v NullableMetricVolumes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricVolumes) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return json.Unmarshal(src, &v.value)
}
