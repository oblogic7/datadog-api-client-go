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

// ImageWidgetDefinitionType Type of the image widget.
type ImageWidgetDefinitionType string

// List of ImageWidgetDefinitionType
const (
	IMAGEWIDGETDEFINITIONTYPE_IMAGE ImageWidgetDefinitionType = "image"
)

var allowedImageWidgetDefinitionTypeEnumValues = []ImageWidgetDefinitionType{
	IMAGEWIDGETDEFINITIONTYPE_IMAGE,
}

func (w *ImageWidgetDefinitionType) GetAllowedValues() []ImageWidgetDefinitionType {
	return allowedImageWidgetDefinitionTypeEnumValues
}

func (v *ImageWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ImageWidgetDefinitionType(value)
	return nil
}

// NewImageWidgetDefinitionTypeFromValue returns a pointer to a valid ImageWidgetDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewImageWidgetDefinitionTypeFromValue(v string) (*ImageWidgetDefinitionType, error) {
	ev := ImageWidgetDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ImageWidgetDefinitionType: valid values are %v", v, allowedImageWidgetDefinitionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ImageWidgetDefinitionType) IsValid() bool {
	for _, existing := range allowedImageWidgetDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ImageWidgetDefinitionType value
func (v ImageWidgetDefinitionType) Ptr() *ImageWidgetDefinitionType {
	return &v
}

type NullableImageWidgetDefinitionType struct {
	value *ImageWidgetDefinitionType
	isSet bool
}

func (v NullableImageWidgetDefinitionType) Get() *ImageWidgetDefinitionType {
	return v.value
}

func (v *NullableImageWidgetDefinitionType) Set(val *ImageWidgetDefinitionType) {
	v.value = val
	v.isSet = true
}

func (v NullableImageWidgetDefinitionType) IsSet() bool {
	return v.isSet
}

func (v *NullableImageWidgetDefinitionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageWidgetDefinitionType(val *ImageWidgetDefinitionType) *NullableImageWidgetDefinitionType {
	return &NullableImageWidgetDefinitionType{value: val, isSet: true}
}

func (v NullableImageWidgetDefinitionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
