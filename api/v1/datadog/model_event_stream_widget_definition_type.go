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

// EventStreamWidgetDefinitionType Type of the event stream widget.
type EventStreamWidgetDefinitionType string

// List of EventStreamWidgetDefinitionType
const (
	EVENTSTREAMWIDGETDEFINITIONTYPE_EVENT_STREAM EventStreamWidgetDefinitionType = "event_stream"
)

var allowedEventStreamWidgetDefinitionTypeEnumValues = []EventStreamWidgetDefinitionType{
	EVENTSTREAMWIDGETDEFINITIONTYPE_EVENT_STREAM,
}

func (w *EventStreamWidgetDefinitionType) GetAllowedValues() []EventStreamWidgetDefinitionType {
	return allowedEventStreamWidgetDefinitionTypeEnumValues
}

func (v *EventStreamWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EventStreamWidgetDefinitionType(value)
	return nil
}

// NewEventStreamWidgetDefinitionTypeFromValue returns a pointer to a valid EventStreamWidgetDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventStreamWidgetDefinitionTypeFromValue(v string) (*EventStreamWidgetDefinitionType, error) {
	ev := EventStreamWidgetDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventStreamWidgetDefinitionType: valid values are %v", v, allowedEventStreamWidgetDefinitionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventStreamWidgetDefinitionType) IsValid() bool {
	for _, existing := range allowedEventStreamWidgetDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventStreamWidgetDefinitionType value
func (v EventStreamWidgetDefinitionType) Ptr() *EventStreamWidgetDefinitionType {
	return &v
}

type NullableEventStreamWidgetDefinitionType struct {
	value *EventStreamWidgetDefinitionType
	isSet bool
}

func (v NullableEventStreamWidgetDefinitionType) Get() *EventStreamWidgetDefinitionType {
	return v.value
}

func (v *NullableEventStreamWidgetDefinitionType) Set(val *EventStreamWidgetDefinitionType) {
	v.value = val
	v.isSet = true
}

func (v NullableEventStreamWidgetDefinitionType) IsSet() bool {
	return v.isSet
}

func (v *NullableEventStreamWidgetDefinitionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventStreamWidgetDefinitionType(val *EventStreamWidgetDefinitionType) *NullableEventStreamWidgetDefinitionType {
	return &NullableEventStreamWidgetDefinitionType{value: val, isSet: true}
}

func (v NullableEventStreamWidgetDefinitionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventStreamWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
