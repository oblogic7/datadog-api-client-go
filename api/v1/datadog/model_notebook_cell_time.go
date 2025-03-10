/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// NotebookCellTime - Timeframe for the notebook cell. When 'null', the notebook global time is used.
type NotebookCellTime struct {
	NotebookRelativeTime *NotebookRelativeTime
	NotebookAbsoluteTime *NotebookAbsoluteTime

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// NotebookRelativeTimeAsNotebookCellTime is a convenience function that returns NotebookRelativeTime wrapped in NotebookCellTime
func NotebookRelativeTimeAsNotebookCellTime(v *NotebookRelativeTime) NotebookCellTime {
	return NotebookCellTime{NotebookRelativeTime: v}
}

// NotebookAbsoluteTimeAsNotebookCellTime is a convenience function that returns NotebookAbsoluteTime wrapped in NotebookCellTime
func NotebookAbsoluteTimeAsNotebookCellTime(v *NotebookAbsoluteTime) NotebookCellTime {
	return NotebookCellTime{NotebookAbsoluteTime: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *NotebookCellTime) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into NotebookRelativeTime
	err = json.Unmarshal(data, &dst.NotebookRelativeTime)
	if err == nil {
		if dst.NotebookRelativeTime != nil && dst.NotebookRelativeTime.UnparsedObject == nil {
			jsonNotebookRelativeTime, _ := json.Marshal(dst.NotebookRelativeTime)
			if string(jsonNotebookRelativeTime) == "{}" { // empty struct
				dst.NotebookRelativeTime = nil
			} else {
				match++
			}
		} else {
			dst.NotebookRelativeTime = nil
		}
	} else {
		dst.NotebookRelativeTime = nil
	}

	// try to unmarshal data into NotebookAbsoluteTime
	err = json.Unmarshal(data, &dst.NotebookAbsoluteTime)
	if err == nil {
		if dst.NotebookAbsoluteTime != nil && dst.NotebookAbsoluteTime.UnparsedObject == nil {
			jsonNotebookAbsoluteTime, _ := json.Marshal(dst.NotebookAbsoluteTime)
			if string(jsonNotebookAbsoluteTime) == "{}" { // empty struct
				dst.NotebookAbsoluteTime = nil
			} else {
				match++
			}
		} else {
			dst.NotebookAbsoluteTime = nil
		}
	} else {
		dst.NotebookAbsoluteTime = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		dst.NotebookRelativeTime = nil
		dst.NotebookAbsoluteTime = nil
		return json.Unmarshal(data, &dst.UnparsedObject)
	} else {
		return nil // exactly one match
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NotebookCellTime) MarshalJSON() ([]byte, error) {
	if src.NotebookRelativeTime != nil {
		return json.Marshal(&src.NotebookRelativeTime)
	}

	if src.NotebookAbsoluteTime != nil {
		return json.Marshal(&src.NotebookAbsoluteTime)
	}

	if src.UnparsedObject != nil {
		return json.Marshal(src.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *NotebookCellTime) GetActualInstance() interface{} {
	if obj.NotebookRelativeTime != nil {
		return obj.NotebookRelativeTime
	}

	if obj.NotebookAbsoluteTime != nil {
		return obj.NotebookAbsoluteTime
	}

	// all schemas are nil
	return nil
}

type NullableNotebookCellTime struct {
	value *NotebookCellTime
	isSet bool
}

func (v NullableNotebookCellTime) Get() *NotebookCellTime {
	return v.value
}

func (v *NullableNotebookCellTime) Set(val *NotebookCellTime) {
	v.value = val
	v.isSet = true
}

func (v NullableNotebookCellTime) IsSet() bool {
	return v.isSet
}

func (v *NullableNotebookCellTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotebookCellTime(val *NotebookCellTime) *NullableNotebookCellTime {
	return &NullableNotebookCellTime{value: val, isSet: true}
}

func (v NullableNotebookCellTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotebookCellTime) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return json.Unmarshal(src, &v.value)
}
