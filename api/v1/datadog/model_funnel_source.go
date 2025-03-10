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

// FunnelSource Source from which to query items to display in the funnel.
type FunnelSource string

// List of FunnelSource
const (
	FUNNELSOURCE_RUM FunnelSource = "rum"
)

var allowedFunnelSourceEnumValues = []FunnelSource{
	FUNNELSOURCE_RUM,
}

func (w *FunnelSource) GetAllowedValues() []FunnelSource {
	return allowedFunnelSourceEnumValues
}

func (v *FunnelSource) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FunnelSource(value)
	return nil
}

// NewFunnelSourceFromValue returns a pointer to a valid FunnelSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFunnelSourceFromValue(v string) (*FunnelSource, error) {
	ev := FunnelSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FunnelSource: valid values are %v", v, allowedFunnelSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FunnelSource) IsValid() bool {
	for _, existing := range allowedFunnelSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FunnelSource value
func (v FunnelSource) Ptr() *FunnelSource {
	return &v
}

type NullableFunnelSource struct {
	value *FunnelSource
	isSet bool
}

func (v NullableFunnelSource) Get() *FunnelSource {
	return v.value
}

func (v *NullableFunnelSource) Set(val *FunnelSource) {
	v.value = val
	v.isSet = true
}

func (v NullableFunnelSource) IsSet() bool {
	return v.isSet
}

func (v *NullableFunnelSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunnelSource(val *FunnelSource) *NullableFunnelSource {
	return &NullableFunnelSource{value: val, isSet: true}
}

func (v NullableFunnelSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunnelSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
