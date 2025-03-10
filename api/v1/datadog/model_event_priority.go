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

// EventPriority The priority of the event. For example, `normal` or `low`.
type EventPriority string

// List of EventPriority
const (
	EVENTPRIORITY_NORMAL EventPriority = "normal"
	EVENTPRIORITY_LOW    EventPriority = "low"
)

var allowedEventPriorityEnumValues = []EventPriority{
	EVENTPRIORITY_NORMAL,
	EVENTPRIORITY_LOW,
}

func (w *EventPriority) GetAllowedValues() []EventPriority {
	return allowedEventPriorityEnumValues
}

func (v *EventPriority) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EventPriority(value)
	return nil
}

// NewEventPriorityFromValue returns a pointer to a valid EventPriority
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventPriorityFromValue(v string) (*EventPriority, error) {
	ev := EventPriority(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventPriority: valid values are %v", v, allowedEventPriorityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventPriority) IsValid() bool {
	for _, existing := range allowedEventPriorityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventPriority value
func (v EventPriority) Ptr() *EventPriority {
	return &v
}

type NullableEventPriority struct {
	value *EventPriority
	isSet bool
}

func (v NullableEventPriority) Get() *EventPriority {
	return v.value
}

func (v *NullableEventPriority) Set(val *EventPriority) {
	v.value = val
	v.isSet = true
}

func (v NullableEventPriority) IsSet() bool {
	return v.isSet
}

func (v *NullableEventPriority) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventPriority(val *EventPriority) *NullableEventPriority {
	return &NullableEventPriority{value: val, isSet: true}
}

func (v NullableEventPriority) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventPriority) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
