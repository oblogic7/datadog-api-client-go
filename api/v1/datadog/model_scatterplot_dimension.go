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

// ScatterplotDimension Dimension of the Scatterplot.
type ScatterplotDimension string

// List of ScatterplotDimension
const (
	SCATTERPLOTDIMENSION_X      ScatterplotDimension = "x"
	SCATTERPLOTDIMENSION_Y      ScatterplotDimension = "y"
	SCATTERPLOTDIMENSION_RADIUS ScatterplotDimension = "radius"
	SCATTERPLOTDIMENSION_COLOR  ScatterplotDimension = "color"
)

var allowedScatterplotDimensionEnumValues = []ScatterplotDimension{
	SCATTERPLOTDIMENSION_X,
	SCATTERPLOTDIMENSION_Y,
	SCATTERPLOTDIMENSION_RADIUS,
	SCATTERPLOTDIMENSION_COLOR,
}

func (w *ScatterplotDimension) GetAllowedValues() []ScatterplotDimension {
	return allowedScatterplotDimensionEnumValues
}

func (v *ScatterplotDimension) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScatterplotDimension(value)
	return nil
}

// NewScatterplotDimensionFromValue returns a pointer to a valid ScatterplotDimension
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewScatterplotDimensionFromValue(v string) (*ScatterplotDimension, error) {
	ev := ScatterplotDimension(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ScatterplotDimension: valid values are %v", v, allowedScatterplotDimensionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ScatterplotDimension) IsValid() bool {
	for _, existing := range allowedScatterplotDimensionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScatterplotDimension value
func (v ScatterplotDimension) Ptr() *ScatterplotDimension {
	return &v
}

type NullableScatterplotDimension struct {
	value *ScatterplotDimension
	isSet bool
}

func (v NullableScatterplotDimension) Get() *ScatterplotDimension {
	return v.value
}

func (v *NullableScatterplotDimension) Set(val *ScatterplotDimension) {
	v.value = val
	v.isSet = true
}

func (v NullableScatterplotDimension) IsSet() bool {
	return v.isSet
}

func (v *NullableScatterplotDimension) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScatterplotDimension(val *ScatterplotDimension) *NullableScatterplotDimension {
	return &NullableScatterplotDimension{value: val, isSet: true}
}

func (v NullableScatterplotDimension) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScatterplotDimension) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
