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

// NotebookStatus Publication status of the notebook. For now, always "published".
type NotebookStatus string

// List of NotebookStatus
const (
	NOTEBOOKSTATUS_PUBLISHED NotebookStatus = "published"
)

var allowedNotebookStatusEnumValues = []NotebookStatus{
	NOTEBOOKSTATUS_PUBLISHED,
}

func (w *NotebookStatus) GetAllowedValues() []NotebookStatus {
	return allowedNotebookStatusEnumValues
}

func (v *NotebookStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NotebookStatus(value)
	return nil
}

// NewNotebookStatusFromValue returns a pointer to a valid NotebookStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNotebookStatusFromValue(v string) (*NotebookStatus, error) {
	ev := NotebookStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NotebookStatus: valid values are %v", v, allowedNotebookStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NotebookStatus) IsValid() bool {
	for _, existing := range allowedNotebookStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotebookStatus value
func (v NotebookStatus) Ptr() *NotebookStatus {
	return &v
}

type NullableNotebookStatus struct {
	value *NotebookStatus
	isSet bool
}

func (v NullableNotebookStatus) Get() *NotebookStatus {
	return v.value
}

func (v *NullableNotebookStatus) Set(val *NotebookStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableNotebookStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableNotebookStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotebookStatus(val *NotebookStatus) *NullableNotebookStatus {
	return &NullableNotebookStatus{value: val, isSet: true}
}

func (v NullableNotebookStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotebookStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
