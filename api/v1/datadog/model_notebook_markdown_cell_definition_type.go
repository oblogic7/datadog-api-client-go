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

// NotebookMarkdownCellDefinitionType Type of the markdown cell.
type NotebookMarkdownCellDefinitionType string

// List of NotebookMarkdownCellDefinitionType
const (
	NOTEBOOKMARKDOWNCELLDEFINITIONTYPE_MARKDOWN NotebookMarkdownCellDefinitionType = "markdown"
)

var allowedNotebookMarkdownCellDefinitionTypeEnumValues = []NotebookMarkdownCellDefinitionType{
	NOTEBOOKMARKDOWNCELLDEFINITIONTYPE_MARKDOWN,
}

func (w *NotebookMarkdownCellDefinitionType) GetAllowedValues() []NotebookMarkdownCellDefinitionType {
	return allowedNotebookMarkdownCellDefinitionTypeEnumValues
}

func (v *NotebookMarkdownCellDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NotebookMarkdownCellDefinitionType(value)
	return nil
}

// NewNotebookMarkdownCellDefinitionTypeFromValue returns a pointer to a valid NotebookMarkdownCellDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNotebookMarkdownCellDefinitionTypeFromValue(v string) (*NotebookMarkdownCellDefinitionType, error) {
	ev := NotebookMarkdownCellDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NotebookMarkdownCellDefinitionType: valid values are %v", v, allowedNotebookMarkdownCellDefinitionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NotebookMarkdownCellDefinitionType) IsValid() bool {
	for _, existing := range allowedNotebookMarkdownCellDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotebookMarkdownCellDefinitionType value
func (v NotebookMarkdownCellDefinitionType) Ptr() *NotebookMarkdownCellDefinitionType {
	return &v
}

type NullableNotebookMarkdownCellDefinitionType struct {
	value *NotebookMarkdownCellDefinitionType
	isSet bool
}

func (v NullableNotebookMarkdownCellDefinitionType) Get() *NotebookMarkdownCellDefinitionType {
	return v.value
}

func (v *NullableNotebookMarkdownCellDefinitionType) Set(val *NotebookMarkdownCellDefinitionType) {
	v.value = val
	v.isSet = true
}

func (v NullableNotebookMarkdownCellDefinitionType) IsSet() bool {
	return v.isSet
}

func (v *NullableNotebookMarkdownCellDefinitionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotebookMarkdownCellDefinitionType(val *NotebookMarkdownCellDefinitionType) *NullableNotebookMarkdownCellDefinitionType {
	return &NullableNotebookMarkdownCellDefinitionType{value: val, isSet: true}
}

func (v NullableNotebookMarkdownCellDefinitionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotebookMarkdownCellDefinitionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
