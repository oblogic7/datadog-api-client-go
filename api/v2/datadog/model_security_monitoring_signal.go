/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// SecurityMonitoringSignal Object description of a security signal.
type SecurityMonitoringSignal struct {
	// The object containing all signal attributes and their
	// associated values.
	Attributes *SecurityMonitoringSignalAttributes `json:"attributes,omitempty"`
	// The unique ID of the security signal.
	Id *string `json:"id,omitempty"`
	// The type of event.
	Type *SecurityMonitoringSignalType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewSecurityMonitoringSignal instantiates a new SecurityMonitoringSignal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityMonitoringSignal() *SecurityMonitoringSignal {
	this := SecurityMonitoringSignal{}
	var type_ SecurityMonitoringSignalType = SECURITYMONITORINGSIGNALTYPE_SIGNAL
	this.Type = &type_
	return &this
}

// NewSecurityMonitoringSignalWithDefaults instantiates a new SecurityMonitoringSignal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityMonitoringSignalWithDefaults() *SecurityMonitoringSignal {
	this := SecurityMonitoringSignal{}
	var type_ SecurityMonitoringSignalType = SECURITYMONITORINGSIGNALTYPE_SIGNAL
	this.Type = &type_
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *SecurityMonitoringSignal) GetAttributes() SecurityMonitoringSignalAttributes {
	if o == nil || o.Attributes == nil {
		var ret SecurityMonitoringSignalAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignal) GetAttributesOk() (*SecurityMonitoringSignalAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SecurityMonitoringSignal) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given SecurityMonitoringSignalAttributes and assigns it to the Attributes field.
func (o *SecurityMonitoringSignal) SetAttributes(v SecurityMonitoringSignalAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SecurityMonitoringSignal) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignal) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SecurityMonitoringSignal) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SecurityMonitoringSignal) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SecurityMonitoringSignal) GetType() SecurityMonitoringSignalType {
	if o == nil || o.Type == nil {
		var ret SecurityMonitoringSignalType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignal) GetTypeOk() (*SecurityMonitoringSignalType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SecurityMonitoringSignal) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given SecurityMonitoringSignalType and assigns it to the Type field.
func (o *SecurityMonitoringSignal) SetType(v SecurityMonitoringSignalType) {
	o.Type = &v
}

func (o SecurityMonitoringSignal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

func (o *SecurityMonitoringSignal) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Attributes *SecurityMonitoringSignalAttributes `json:"attributes,omitempty"`
		Id         *string                             `json:"id,omitempty"`
		Type       *SecurityMonitoringSignalType       `json:"type,omitempty"`
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
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Attributes = all.Attributes
	o.Id = all.Id
	o.Type = all.Type
	return nil
}
