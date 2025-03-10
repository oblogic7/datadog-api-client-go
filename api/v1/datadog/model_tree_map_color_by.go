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

// TreeMapColorBy The attribute used to determine color in the widget.
type TreeMapColorBy string

// List of TreeMapColorBy
const (
	TREEMAPCOLORBY_USER TreeMapColorBy = "user"
)

var allowedTreeMapColorByEnumValues = []TreeMapColorBy{
	TREEMAPCOLORBY_USER,
}

func (w *TreeMapColorBy) GetAllowedValues() []TreeMapColorBy {
	return allowedTreeMapColorByEnumValues
}

func (v *TreeMapColorBy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TreeMapColorBy(value)
	return nil
}

// NewTreeMapColorByFromValue returns a pointer to a valid TreeMapColorBy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTreeMapColorByFromValue(v string) (*TreeMapColorBy, error) {
	ev := TreeMapColorBy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TreeMapColorBy: valid values are %v", v, allowedTreeMapColorByEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TreeMapColorBy) IsValid() bool {
	for _, existing := range allowedTreeMapColorByEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TreeMapColorBy value
func (v TreeMapColorBy) Ptr() *TreeMapColorBy {
	return &v
}

type NullableTreeMapColorBy struct {
	value *TreeMapColorBy
	isSet bool
}

func (v NullableTreeMapColorBy) Get() *TreeMapColorBy {
	return v.value
}

func (v *NullableTreeMapColorBy) Set(val *TreeMapColorBy) {
	v.value = val
	v.isSet = true
}

func (v NullableTreeMapColorBy) IsSet() bool {
	return v.isSet
}

func (v *NullableTreeMapColorBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTreeMapColorBy(val *TreeMapColorBy) *NullableTreeMapColorBy {
	return &NullableTreeMapColorBy{value: val, isSet: true}
}

func (v NullableTreeMapColorBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTreeMapColorBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
