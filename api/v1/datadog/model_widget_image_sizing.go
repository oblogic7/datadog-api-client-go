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

// WidgetImageSizing How to size the image on the widget. The values are based on the image `object-fit` CSS properties.
// **Note**: `zoom`, `fit` and `center` values are deprecated.
type WidgetImageSizing string

// List of WidgetImageSizing
const (
	WIDGETIMAGESIZING_FILL      WidgetImageSizing = "fill"
	WIDGETIMAGESIZING_CONTAIN   WidgetImageSizing = "contain"
	WIDGETIMAGESIZING_COVER     WidgetImageSizing = "cover"
	WIDGETIMAGESIZING_NONE      WidgetImageSizing = "none"
	WIDGETIMAGESIZING_SCALEDOWN WidgetImageSizing = "scale-down"
	WIDGETIMAGESIZING_ZOOM      WidgetImageSizing = "zoom"
	WIDGETIMAGESIZING_FIT       WidgetImageSizing = "fit"
	WIDGETIMAGESIZING_CENTER    WidgetImageSizing = "center"
)

var allowedWidgetImageSizingEnumValues = []WidgetImageSizing{
	WIDGETIMAGESIZING_FILL,
	WIDGETIMAGESIZING_CONTAIN,
	WIDGETIMAGESIZING_COVER,
	WIDGETIMAGESIZING_NONE,
	WIDGETIMAGESIZING_SCALEDOWN,
	WIDGETIMAGESIZING_ZOOM,
	WIDGETIMAGESIZING_FIT,
	WIDGETIMAGESIZING_CENTER,
}

func (w *WidgetImageSizing) GetAllowedValues() []WidgetImageSizing {
	return allowedWidgetImageSizingEnumValues
}

func (v *WidgetImageSizing) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WidgetImageSizing(value)
	return nil
}

// NewWidgetImageSizingFromValue returns a pointer to a valid WidgetImageSizing
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWidgetImageSizingFromValue(v string) (*WidgetImageSizing, error) {
	ev := WidgetImageSizing(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WidgetImageSizing: valid values are %v", v, allowedWidgetImageSizingEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WidgetImageSizing) IsValid() bool {
	for _, existing := range allowedWidgetImageSizingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WidgetImageSizing value
func (v WidgetImageSizing) Ptr() *WidgetImageSizing {
	return &v
}

type NullableWidgetImageSizing struct {
	value *WidgetImageSizing
	isSet bool
}

func (v NullableWidgetImageSizing) Get() *WidgetImageSizing {
	return v.value
}

func (v *NullableWidgetImageSizing) Set(val *WidgetImageSizing) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetImageSizing) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetImageSizing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetImageSizing(val *WidgetImageSizing) *NullableWidgetImageSizing {
	return &NullableWidgetImageSizing{value: val, isSet: true}
}

func (v NullableWidgetImageSizing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetImageSizing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
