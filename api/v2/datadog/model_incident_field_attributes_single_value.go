/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// IncidentFieldAttributesSingleValue A field with a single value selected.
type IncidentFieldAttributesSingleValue struct {
	// Type of the single value field definitions.
	Type *IncidentFieldAttributesSingleValueType `json:"type,omitempty"`
	// The single value selected for this field.
	Value NullableString `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewIncidentFieldAttributesSingleValue instantiates a new IncidentFieldAttributesSingleValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentFieldAttributesSingleValue() *IncidentFieldAttributesSingleValue {
	this := IncidentFieldAttributesSingleValue{}
	var type_ IncidentFieldAttributesSingleValueType = INCIDENTFIELDATTRIBUTESSINGLEVALUETYPE_DROPDOWN
	this.Type = &type_
	return &this
}

// NewIncidentFieldAttributesSingleValueWithDefaults instantiates a new IncidentFieldAttributesSingleValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentFieldAttributesSingleValueWithDefaults() *IncidentFieldAttributesSingleValue {
	this := IncidentFieldAttributesSingleValue{}
	var type_ IncidentFieldAttributesSingleValueType = INCIDENTFIELDATTRIBUTESSINGLEVALUETYPE_DROPDOWN
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IncidentFieldAttributesSingleValue) GetType() IncidentFieldAttributesSingleValueType {
	if o == nil || o.Type == nil {
		var ret IncidentFieldAttributesSingleValueType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentFieldAttributesSingleValue) GetTypeOk() (*IncidentFieldAttributesSingleValueType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IncidentFieldAttributesSingleValue) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given IncidentFieldAttributesSingleValueType and assigns it to the Type field.
func (o *IncidentFieldAttributesSingleValue) SetType(v IncidentFieldAttributesSingleValueType) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentFieldAttributesSingleValue) GetValue() string {
	if o == nil || o.Value.Get() == nil {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncidentFieldAttributesSingleValue) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *IncidentFieldAttributesSingleValue) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *IncidentFieldAttributesSingleValue) SetValue(v string) {
	o.Value.Set(&v)
}

// SetValueNil sets the value for Value to be an explicit nil
func (o *IncidentFieldAttributesSingleValue) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *IncidentFieldAttributesSingleValue) UnsetValue() {
	o.Value.Unset()
}

func (o IncidentFieldAttributesSingleValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

func (o *IncidentFieldAttributesSingleValue) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Type  *IncidentFieldAttributesSingleValueType `json:"type,omitempty"`
		Value NullableString                          `json:"value,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if v := all.Type; v != nil && !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.Type = all.Type
	o.Value = all.Value
	return nil
}
