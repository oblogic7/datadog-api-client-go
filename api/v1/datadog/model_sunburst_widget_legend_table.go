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

// SunburstWidgetLegendTable Configuration of table-based legend.
type SunburstWidgetLegendTable struct {
	// Whether or not to show a table legend.
	Type SunburstWidgetLegendTableType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewSunburstWidgetLegendTable instantiates a new SunburstWidgetLegendTable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSunburstWidgetLegendTable(type_ SunburstWidgetLegendTableType) *SunburstWidgetLegendTable {
	this := SunburstWidgetLegendTable{}
	this.Type = type_
	return &this
}

// NewSunburstWidgetLegendTableWithDefaults instantiates a new SunburstWidgetLegendTable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSunburstWidgetLegendTableWithDefaults() *SunburstWidgetLegendTable {
	this := SunburstWidgetLegendTable{}
	return &this
}

// GetType returns the Type field value
func (o *SunburstWidgetLegendTable) GetType() SunburstWidgetLegendTableType {
	if o == nil {
		var ret SunburstWidgetLegendTableType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SunburstWidgetLegendTable) GetTypeOk() (*SunburstWidgetLegendTableType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SunburstWidgetLegendTable) SetType(v SunburstWidgetLegendTableType) {
	o.Type = v
}

func (o SunburstWidgetLegendTable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

func (o *SunburstWidgetLegendTable) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Type *SunburstWidgetLegendTableType `json:"type"`
	}{}
	all := struct {
		Type SunburstWidgetLegendTableType `json:"type"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Type == nil {
		return fmt.Errorf("Required field type missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if v := all.Type; !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.Type = all.Type
	return nil
}
