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

// ApplicationKeyCreateData Object used to create an application key.
type ApplicationKeyCreateData struct {
	// Attributes used to create an application Key.
	Attributes ApplicationKeyCreateAttributes `json:"attributes"`
	// Application Keys resource type.
	Type ApplicationKeysType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewApplicationKeyCreateData instantiates a new ApplicationKeyCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationKeyCreateData(attributes ApplicationKeyCreateAttributes, type_ ApplicationKeysType) *ApplicationKeyCreateData {
	this := ApplicationKeyCreateData{}
	this.Attributes = attributes
	this.Type = type_
	return &this
}

// NewApplicationKeyCreateDataWithDefaults instantiates a new ApplicationKeyCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationKeyCreateDataWithDefaults() *ApplicationKeyCreateData {
	this := ApplicationKeyCreateData{}
	var type_ ApplicationKeysType = APPLICATIONKEYSTYPE_APPLICATION_KEYS
	this.Type = type_
	return &this
}

// GetAttributes returns the Attributes field value
func (o *ApplicationKeyCreateData) GetAttributes() ApplicationKeyCreateAttributes {
	if o == nil {
		var ret ApplicationKeyCreateAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ApplicationKeyCreateData) GetAttributesOk() (*ApplicationKeyCreateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ApplicationKeyCreateData) SetAttributes(v ApplicationKeyCreateAttributes) {
	o.Attributes = v
}

// GetType returns the Type field value
func (o *ApplicationKeyCreateData) GetType() ApplicationKeysType {
	if o == nil {
		var ret ApplicationKeysType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApplicationKeyCreateData) GetTypeOk() (*ApplicationKeysType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApplicationKeyCreateData) SetType(v ApplicationKeysType) {
	o.Type = v
}

func (o ApplicationKeyCreateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

func (o *ApplicationKeyCreateData) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Attributes *ApplicationKeyCreateAttributes `json:"attributes"`
		Type       *ApplicationKeysType            `json:"type"`
	}{}
	all := struct {
		Attributes ApplicationKeyCreateAttributes `json:"attributes"`
		Type       ApplicationKeysType            `json:"type"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Attributes == nil {
		return fmt.Errorf("Required field attributes missing")
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
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Attributes = all.Attributes
	o.Type = all.Type
	return nil
}
