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

// AuditLogsSort Sort parameters when querying events.
type AuditLogsSort string

// List of AuditLogsSort
const (
	AUDITLOGSSORT_TIMESTAMP_ASCENDING  AuditLogsSort = "timestamp"
	AUDITLOGSSORT_TIMESTAMP_DESCENDING AuditLogsSort = "-timestamp"
)

var allowedAuditLogsSortEnumValues = []AuditLogsSort{
	AUDITLOGSSORT_TIMESTAMP_ASCENDING,
	AUDITLOGSSORT_TIMESTAMP_DESCENDING,
}

func (w *AuditLogsSort) GetAllowedValues() []AuditLogsSort {
	return allowedAuditLogsSortEnumValues
}

func (v *AuditLogsSort) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AuditLogsSort(value)
	return nil
}

// NewAuditLogsSortFromValue returns a pointer to a valid AuditLogsSort
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAuditLogsSortFromValue(v string) (*AuditLogsSort, error) {
	ev := AuditLogsSort(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AuditLogsSort: valid values are %v", v, allowedAuditLogsSortEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AuditLogsSort) IsValid() bool {
	for _, existing := range allowedAuditLogsSortEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AuditLogsSort value
func (v AuditLogsSort) Ptr() *AuditLogsSort {
	return &v
}

type NullableAuditLogsSort struct {
	value *AuditLogsSort
	isSet bool
}

func (v NullableAuditLogsSort) Get() *AuditLogsSort {
	return v.value
}

func (v *NullableAuditLogsSort) Set(val *AuditLogsSort) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditLogsSort) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditLogsSort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditLogsSort(val *AuditLogsSort) *NullableAuditLogsSort {
	return &NullableAuditLogsSort{value: val, isSet: true}
}

func (v NullableAuditLogsSort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditLogsSort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
