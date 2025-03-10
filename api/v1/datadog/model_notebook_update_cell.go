/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// NotebookUpdateCell - Updating a notebook can either insert new cell(s) or update existing cell(s) by including the cell `id`.
// To delete existing cell(s), simply omit it from the list of cells.
type NotebookUpdateCell struct {
	NotebookCellCreateRequest *NotebookCellCreateRequest
	NotebookCellUpdateRequest *NotebookCellUpdateRequest

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// NotebookCellCreateRequestAsNotebookUpdateCell is a convenience function that returns NotebookCellCreateRequest wrapped in NotebookUpdateCell
func NotebookCellCreateRequestAsNotebookUpdateCell(v *NotebookCellCreateRequest) NotebookUpdateCell {
	return NotebookUpdateCell{NotebookCellCreateRequest: v}
}

// NotebookCellUpdateRequestAsNotebookUpdateCell is a convenience function that returns NotebookCellUpdateRequest wrapped in NotebookUpdateCell
func NotebookCellUpdateRequestAsNotebookUpdateCell(v *NotebookCellUpdateRequest) NotebookUpdateCell {
	return NotebookUpdateCell{NotebookCellUpdateRequest: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *NotebookUpdateCell) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into NotebookCellCreateRequest
	err = json.Unmarshal(data, &dst.NotebookCellCreateRequest)
	if err == nil {
		if dst.NotebookCellCreateRequest != nil && dst.NotebookCellCreateRequest.UnparsedObject == nil {
			jsonNotebookCellCreateRequest, _ := json.Marshal(dst.NotebookCellCreateRequest)
			if string(jsonNotebookCellCreateRequest) == "{}" { // empty struct
				dst.NotebookCellCreateRequest = nil
			} else {
				match++
			}
		} else {
			dst.NotebookCellCreateRequest = nil
		}
	} else {
		dst.NotebookCellCreateRequest = nil
	}

	// try to unmarshal data into NotebookCellUpdateRequest
	err = json.Unmarshal(data, &dst.NotebookCellUpdateRequest)
	if err == nil {
		if dst.NotebookCellUpdateRequest != nil && dst.NotebookCellUpdateRequest.UnparsedObject == nil {
			jsonNotebookCellUpdateRequest, _ := json.Marshal(dst.NotebookCellUpdateRequest)
			if string(jsonNotebookCellUpdateRequest) == "{}" { // empty struct
				dst.NotebookCellUpdateRequest = nil
			} else {
				match++
			}
		} else {
			dst.NotebookCellUpdateRequest = nil
		}
	} else {
		dst.NotebookCellUpdateRequest = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		dst.NotebookCellCreateRequest = nil
		dst.NotebookCellUpdateRequest = nil
		return json.Unmarshal(data, &dst.UnparsedObject)
	} else {
		return nil // exactly one match
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NotebookUpdateCell) MarshalJSON() ([]byte, error) {
	if src.NotebookCellCreateRequest != nil {
		return json.Marshal(&src.NotebookCellCreateRequest)
	}

	if src.NotebookCellUpdateRequest != nil {
		return json.Marshal(&src.NotebookCellUpdateRequest)
	}

	if src.UnparsedObject != nil {
		return json.Marshal(src.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *NotebookUpdateCell) GetActualInstance() interface{} {
	if obj.NotebookCellCreateRequest != nil {
		return obj.NotebookCellCreateRequest
	}

	if obj.NotebookCellUpdateRequest != nil {
		return obj.NotebookCellUpdateRequest
	}

	// all schemas are nil
	return nil
}

type NullableNotebookUpdateCell struct {
	value *NotebookUpdateCell
	isSet bool
}

func (v NullableNotebookUpdateCell) Get() *NotebookUpdateCell {
	return v.value
}

func (v *NullableNotebookUpdateCell) Set(val *NotebookUpdateCell) {
	v.value = val
	v.isSet = true
}

func (v NullableNotebookUpdateCell) IsSet() bool {
	return v.isSet
}

func (v *NullableNotebookUpdateCell) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotebookUpdateCell(val *NotebookUpdateCell) *NullableNotebookUpdateCell {
	return &NullableNotebookUpdateCell{value: val, isSet: true}
}

func (v NullableNotebookUpdateCell) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotebookUpdateCell) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return json.Unmarshal(src, &v.value)
}
