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

// AuditLogsResponseStatus The status of the response.
type AuditLogsResponseStatus string

// List of AuditLogsResponseStatus
const (
	AUDITLOGSRESPONSESTATUS_DONE    AuditLogsResponseStatus = "done"
	AUDITLOGSRESPONSESTATUS_TIMEOUT AuditLogsResponseStatus = "timeout"
)

var allowedAuditLogsResponseStatusEnumValues = []AuditLogsResponseStatus{
	AUDITLOGSRESPONSESTATUS_DONE,
	AUDITLOGSRESPONSESTATUS_TIMEOUT,
}

func (w *AuditLogsResponseStatus) GetAllowedValues() []AuditLogsResponseStatus {
	return allowedAuditLogsResponseStatusEnumValues
}

func (v *AuditLogsResponseStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AuditLogsResponseStatus(value)
	return nil
}

// NewAuditLogsResponseStatusFromValue returns a pointer to a valid AuditLogsResponseStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAuditLogsResponseStatusFromValue(v string) (*AuditLogsResponseStatus, error) {
	ev := AuditLogsResponseStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AuditLogsResponseStatus: valid values are %v", v, allowedAuditLogsResponseStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AuditLogsResponseStatus) IsValid() bool {
	for _, existing := range allowedAuditLogsResponseStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AuditLogsResponseStatus value
func (v AuditLogsResponseStatus) Ptr() *AuditLogsResponseStatus {
	return &v
}

type NullableAuditLogsResponseStatus struct {
	value *AuditLogsResponseStatus
	isSet bool
}

func (v NullableAuditLogsResponseStatus) Get() *AuditLogsResponseStatus {
	return v.value
}

func (v *NullableAuditLogsResponseStatus) Set(val *AuditLogsResponseStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditLogsResponseStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditLogsResponseStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditLogsResponseStatus(val *AuditLogsResponseStatus) *NullableAuditLogsResponseStatus {
	return &NullableAuditLogsResponseStatus{value: val, isSet: true}
}

func (v NullableAuditLogsResponseStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditLogsResponseStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
