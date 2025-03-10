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

// NotebookMetadataType Metadata type of the notebook.
type NotebookMetadataType string

// List of NotebookMetadataType
const (
	NOTEBOOKMETADATATYPE_POSTMORTEM    NotebookMetadataType = "postmortem"
	NOTEBOOKMETADATATYPE_RUNBOOK       NotebookMetadataType = "runbook"
	NOTEBOOKMETADATATYPE_INVESTIGATION NotebookMetadataType = "investigation"
	NOTEBOOKMETADATATYPE_DOCUMENTATION NotebookMetadataType = "documentation"
	NOTEBOOKMETADATATYPE_REPORT        NotebookMetadataType = "report"
)

var allowedNotebookMetadataTypeEnumValues = []NotebookMetadataType{
	NOTEBOOKMETADATATYPE_POSTMORTEM,
	NOTEBOOKMETADATATYPE_RUNBOOK,
	NOTEBOOKMETADATATYPE_INVESTIGATION,
	NOTEBOOKMETADATATYPE_DOCUMENTATION,
	NOTEBOOKMETADATATYPE_REPORT,
}

func (w *NotebookMetadataType) GetAllowedValues() []NotebookMetadataType {
	return allowedNotebookMetadataTypeEnumValues
}

func (v *NotebookMetadataType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NotebookMetadataType(value)
	return nil
}

// NewNotebookMetadataTypeFromValue returns a pointer to a valid NotebookMetadataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNotebookMetadataTypeFromValue(v string) (*NotebookMetadataType, error) {
	ev := NotebookMetadataType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NotebookMetadataType: valid values are %v", v, allowedNotebookMetadataTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NotebookMetadataType) IsValid() bool {
	for _, existing := range allowedNotebookMetadataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotebookMetadataType value
func (v NotebookMetadataType) Ptr() *NotebookMetadataType {
	return &v
}

type NullableNotebookMetadataType struct {
	value *NotebookMetadataType
	isSet bool
}

func (v NullableNotebookMetadataType) Get() *NotebookMetadataType {
	return v.value
}

func (v *NullableNotebookMetadataType) Set(val *NotebookMetadataType) {
	v.value = val
	v.isSet = true
}

func (v NullableNotebookMetadataType) IsSet() bool {
	return v.isSet
}

func (v *NullableNotebookMetadataType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotebookMetadataType(val *NotebookMetadataType) *NullableNotebookMetadataType {
	return &NullableNotebookMetadataType{value: val, isSet: true}
}

func (v NullableNotebookMetadataType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotebookMetadataType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
