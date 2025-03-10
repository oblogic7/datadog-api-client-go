/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// SunburstWidgetLegend - Configuration of the legend.
type SunburstWidgetLegend struct {
	SunburstWidgetLegendTable           *SunburstWidgetLegendTable
	SunburstWidgetLegendInlineAutomatic *SunburstWidgetLegendInlineAutomatic

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SunburstWidgetLegendTableAsSunburstWidgetLegend is a convenience function that returns SunburstWidgetLegendTable wrapped in SunburstWidgetLegend
func SunburstWidgetLegendTableAsSunburstWidgetLegend(v *SunburstWidgetLegendTable) SunburstWidgetLegend {
	return SunburstWidgetLegend{SunburstWidgetLegendTable: v}
}

// SunburstWidgetLegendInlineAutomaticAsSunburstWidgetLegend is a convenience function that returns SunburstWidgetLegendInlineAutomatic wrapped in SunburstWidgetLegend
func SunburstWidgetLegendInlineAutomaticAsSunburstWidgetLegend(v *SunburstWidgetLegendInlineAutomatic) SunburstWidgetLegend {
	return SunburstWidgetLegend{SunburstWidgetLegendInlineAutomatic: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SunburstWidgetLegend) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SunburstWidgetLegendTable
	err = json.Unmarshal(data, &dst.SunburstWidgetLegendTable)
	if err == nil {
		if dst.SunburstWidgetLegendTable != nil && dst.SunburstWidgetLegendTable.UnparsedObject == nil {
			jsonSunburstWidgetLegendTable, _ := json.Marshal(dst.SunburstWidgetLegendTable)
			if string(jsonSunburstWidgetLegendTable) == "{}" { // empty struct
				dst.SunburstWidgetLegendTable = nil
			} else {
				match++
			}
		} else {
			dst.SunburstWidgetLegendTable = nil
		}
	} else {
		dst.SunburstWidgetLegendTable = nil
	}

	// try to unmarshal data into SunburstWidgetLegendInlineAutomatic
	err = json.Unmarshal(data, &dst.SunburstWidgetLegendInlineAutomatic)
	if err == nil {
		if dst.SunburstWidgetLegendInlineAutomatic != nil && dst.SunburstWidgetLegendInlineAutomatic.UnparsedObject == nil {
			jsonSunburstWidgetLegendInlineAutomatic, _ := json.Marshal(dst.SunburstWidgetLegendInlineAutomatic)
			if string(jsonSunburstWidgetLegendInlineAutomatic) == "{}" { // empty struct
				dst.SunburstWidgetLegendInlineAutomatic = nil
			} else {
				match++
			}
		} else {
			dst.SunburstWidgetLegendInlineAutomatic = nil
		}
	} else {
		dst.SunburstWidgetLegendInlineAutomatic = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		dst.SunburstWidgetLegendTable = nil
		dst.SunburstWidgetLegendInlineAutomatic = nil
		return json.Unmarshal(data, &dst.UnparsedObject)
	} else {
		return nil // exactly one match
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SunburstWidgetLegend) MarshalJSON() ([]byte, error) {
	if src.SunburstWidgetLegendTable != nil {
		return json.Marshal(&src.SunburstWidgetLegendTable)
	}

	if src.SunburstWidgetLegendInlineAutomatic != nil {
		return json.Marshal(&src.SunburstWidgetLegendInlineAutomatic)
	}

	if src.UnparsedObject != nil {
		return json.Marshal(src.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SunburstWidgetLegend) GetActualInstance() interface{} {
	if obj.SunburstWidgetLegendTable != nil {
		return obj.SunburstWidgetLegendTable
	}

	if obj.SunburstWidgetLegendInlineAutomatic != nil {
		return obj.SunburstWidgetLegendInlineAutomatic
	}

	// all schemas are nil
	return nil
}

type NullableSunburstWidgetLegend struct {
	value *SunburstWidgetLegend
	isSet bool
}

func (v NullableSunburstWidgetLegend) Get() *SunburstWidgetLegend {
	return v.value
}

func (v *NullableSunburstWidgetLegend) Set(val *SunburstWidgetLegend) {
	v.value = val
	v.isSet = true
}

func (v NullableSunburstWidgetLegend) IsSet() bool {
	return v.isSet
}

func (v *NullableSunburstWidgetLegend) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSunburstWidgetLegend(val *SunburstWidgetLegend) *NullableSunburstWidgetLegend {
	return &NullableSunburstWidgetLegend{value: val, isSet: true}
}

func (v NullableSunburstWidgetLegend) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSunburstWidgetLegend) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return json.Unmarshal(src, &v.value)
}
