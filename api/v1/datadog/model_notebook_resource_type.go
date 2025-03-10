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

// NotebookResourceType Type of the Notebook resource.
type NotebookResourceType string

// List of NotebookResourceType
const (
	NOTEBOOKRESOURCETYPE_NOTEBOOKS NotebookResourceType = "notebooks"
)

var allowedNotebookResourceTypeEnumValues = []NotebookResourceType{
	NOTEBOOKRESOURCETYPE_NOTEBOOKS,
}

func (w *NotebookResourceType) GetAllowedValues() []NotebookResourceType {
	return allowedNotebookResourceTypeEnumValues
}

func (v *NotebookResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NotebookResourceType(value)
	return nil
}

// NewNotebookResourceTypeFromValue returns a pointer to a valid NotebookResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNotebookResourceTypeFromValue(v string) (*NotebookResourceType, error) {
	ev := NotebookResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NotebookResourceType: valid values are %v", v, allowedNotebookResourceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NotebookResourceType) IsValid() bool {
	for _, existing := range allowedNotebookResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotebookResourceType value
func (v NotebookResourceType) Ptr() *NotebookResourceType {
	return &v
}

type NullableNotebookResourceType struct {
	value *NotebookResourceType
	isSet bool
}

func (v NullableNotebookResourceType) Get() *NotebookResourceType {
	return v.value
}

func (v *NullableNotebookResourceType) Set(val *NotebookResourceType) {
	v.value = val
	v.isSet = true
}

func (v NullableNotebookResourceType) IsSet() bool {
	return v.isSet
}

func (v *NullableNotebookResourceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotebookResourceType(val *NotebookResourceType) *NullableNotebookResourceType {
	return &NullableNotebookResourceType{value: val, isSet: true}
}

func (v NullableNotebookResourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotebookResourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
