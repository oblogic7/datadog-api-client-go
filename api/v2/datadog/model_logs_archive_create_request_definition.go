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

// LogsArchiveCreateRequestDefinition The definition of an archive.
type LogsArchiveCreateRequestDefinition struct {
	// The attributes associated with the archive.
	Attributes *LogsArchiveCreateRequestAttributes `json:"attributes,omitempty"`
	// The type of the resource. The value should always be archives.
	Type string `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewLogsArchiveCreateRequestDefinition instantiates a new LogsArchiveCreateRequestDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogsArchiveCreateRequestDefinition(type_ string) *LogsArchiveCreateRequestDefinition {
	this := LogsArchiveCreateRequestDefinition{}
	this.Type = type_
	return &this
}

// NewLogsArchiveCreateRequestDefinitionWithDefaults instantiates a new LogsArchiveCreateRequestDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogsArchiveCreateRequestDefinitionWithDefaults() *LogsArchiveCreateRequestDefinition {
	this := LogsArchiveCreateRequestDefinition{}
	var type_ string = "archives"
	this.Type = type_
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *LogsArchiveCreateRequestDefinition) GetAttributes() LogsArchiveCreateRequestAttributes {
	if o == nil || o.Attributes == nil {
		var ret LogsArchiveCreateRequestAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsArchiveCreateRequestDefinition) GetAttributesOk() (*LogsArchiveCreateRequestAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *LogsArchiveCreateRequestDefinition) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given LogsArchiveCreateRequestAttributes and assigns it to the Attributes field.
func (o *LogsArchiveCreateRequestDefinition) SetAttributes(v LogsArchiveCreateRequestAttributes) {
	o.Attributes = &v
}

// GetType returns the Type field value
func (o *LogsArchiveCreateRequestDefinition) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LogsArchiveCreateRequestDefinition) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LogsArchiveCreateRequestDefinition) SetType(v string) {
	o.Type = v
}

func (o LogsArchiveCreateRequestDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

func (o *LogsArchiveCreateRequestDefinition) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Type *string `json:"type"`
	}{}
	all := struct {
		Attributes *LogsArchiveCreateRequestAttributes `json:"attributes,omitempty"`
		Type       string                              `json:"type"`
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
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
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
