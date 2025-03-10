/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// LogsArchiveCreateRequestDestination - An archive's destination.
type LogsArchiveCreateRequestDestination struct {
	LogsArchiveDestinationAzure *LogsArchiveDestinationAzure
	LogsArchiveDestinationGCS   *LogsArchiveDestinationGCS
	LogsArchiveDestinationS3    *LogsArchiveDestinationS3

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// LogsArchiveDestinationAzureAsLogsArchiveCreateRequestDestination is a convenience function that returns LogsArchiveDestinationAzure wrapped in LogsArchiveCreateRequestDestination
func LogsArchiveDestinationAzureAsLogsArchiveCreateRequestDestination(v *LogsArchiveDestinationAzure) LogsArchiveCreateRequestDestination {
	return LogsArchiveCreateRequestDestination{LogsArchiveDestinationAzure: v}
}

// LogsArchiveDestinationGCSAsLogsArchiveCreateRequestDestination is a convenience function that returns LogsArchiveDestinationGCS wrapped in LogsArchiveCreateRequestDestination
func LogsArchiveDestinationGCSAsLogsArchiveCreateRequestDestination(v *LogsArchiveDestinationGCS) LogsArchiveCreateRequestDestination {
	return LogsArchiveCreateRequestDestination{LogsArchiveDestinationGCS: v}
}

// LogsArchiveDestinationS3AsLogsArchiveCreateRequestDestination is a convenience function that returns LogsArchiveDestinationS3 wrapped in LogsArchiveCreateRequestDestination
func LogsArchiveDestinationS3AsLogsArchiveCreateRequestDestination(v *LogsArchiveDestinationS3) LogsArchiveCreateRequestDestination {
	return LogsArchiveCreateRequestDestination{LogsArchiveDestinationS3: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *LogsArchiveCreateRequestDestination) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into LogsArchiveDestinationAzure
	err = json.Unmarshal(data, &dst.LogsArchiveDestinationAzure)
	if err == nil {
		if dst.LogsArchiveDestinationAzure != nil && dst.LogsArchiveDestinationAzure.UnparsedObject == nil {
			jsonLogsArchiveDestinationAzure, _ := json.Marshal(dst.LogsArchiveDestinationAzure)
			if string(jsonLogsArchiveDestinationAzure) == "{}" { // empty struct
				dst.LogsArchiveDestinationAzure = nil
			} else {
				match++
			}
		} else {
			dst.LogsArchiveDestinationAzure = nil
		}
	} else {
		dst.LogsArchiveDestinationAzure = nil
	}

	// try to unmarshal data into LogsArchiveDestinationGCS
	err = json.Unmarshal(data, &dst.LogsArchiveDestinationGCS)
	if err == nil {
		if dst.LogsArchiveDestinationGCS != nil && dst.LogsArchiveDestinationGCS.UnparsedObject == nil {
			jsonLogsArchiveDestinationGCS, _ := json.Marshal(dst.LogsArchiveDestinationGCS)
			if string(jsonLogsArchiveDestinationGCS) == "{}" { // empty struct
				dst.LogsArchiveDestinationGCS = nil
			} else {
				match++
			}
		} else {
			dst.LogsArchiveDestinationGCS = nil
		}
	} else {
		dst.LogsArchiveDestinationGCS = nil
	}

	// try to unmarshal data into LogsArchiveDestinationS3
	err = json.Unmarshal(data, &dst.LogsArchiveDestinationS3)
	if err == nil {
		if dst.LogsArchiveDestinationS3 != nil && dst.LogsArchiveDestinationS3.UnparsedObject == nil {
			jsonLogsArchiveDestinationS3, _ := json.Marshal(dst.LogsArchiveDestinationS3)
			if string(jsonLogsArchiveDestinationS3) == "{}" { // empty struct
				dst.LogsArchiveDestinationS3 = nil
			} else {
				match++
			}
		} else {
			dst.LogsArchiveDestinationS3 = nil
		}
	} else {
		dst.LogsArchiveDestinationS3 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		dst.LogsArchiveDestinationAzure = nil
		dst.LogsArchiveDestinationGCS = nil
		dst.LogsArchiveDestinationS3 = nil
		return json.Unmarshal(data, &dst.UnparsedObject)
	} else {
		return nil // exactly one match
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src LogsArchiveCreateRequestDestination) MarshalJSON() ([]byte, error) {
	if src.LogsArchiveDestinationAzure != nil {
		return json.Marshal(&src.LogsArchiveDestinationAzure)
	}

	if src.LogsArchiveDestinationGCS != nil {
		return json.Marshal(&src.LogsArchiveDestinationGCS)
	}

	if src.LogsArchiveDestinationS3 != nil {
		return json.Marshal(&src.LogsArchiveDestinationS3)
	}

	if src.UnparsedObject != nil {
		return json.Marshal(src.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *LogsArchiveCreateRequestDestination) GetActualInstance() interface{} {
	if obj.LogsArchiveDestinationAzure != nil {
		return obj.LogsArchiveDestinationAzure
	}

	if obj.LogsArchiveDestinationGCS != nil {
		return obj.LogsArchiveDestinationGCS
	}

	if obj.LogsArchiveDestinationS3 != nil {
		return obj.LogsArchiveDestinationS3
	}

	// all schemas are nil
	return nil
}

type NullableLogsArchiveCreateRequestDestination struct {
	value *LogsArchiveCreateRequestDestination
	isSet bool
}

func (v NullableLogsArchiveCreateRequestDestination) Get() *LogsArchiveCreateRequestDestination {
	return v.value
}

func (v *NullableLogsArchiveCreateRequestDestination) Set(val *LogsArchiveCreateRequestDestination) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsArchiveCreateRequestDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsArchiveCreateRequestDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsArchiveCreateRequestDestination(val *LogsArchiveCreateRequestDestination) *NullableLogsArchiveCreateRequestDestination {
	return &NullableLogsArchiveCreateRequestDestination{value: val, isSet: true}
}

func (v NullableLogsArchiveCreateRequestDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsArchiveCreateRequestDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return json.Unmarshal(src, &v.value)
}
