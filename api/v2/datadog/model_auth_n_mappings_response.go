/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// AuthNMappingsResponse Array of AuthN Mappings response.
type AuthNMappingsResponse struct {
	// Array of returned AuthN Mappings.
	Data []AuthNMapping `json:"data,omitempty"`
	// Object describing meta attributes of response.
	Meta *ResponseMetaAttributes `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewAuthNMappingsResponse instantiates a new AuthNMappingsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthNMappingsResponse() *AuthNMappingsResponse {
	this := AuthNMappingsResponse{}
	return &this
}

// NewAuthNMappingsResponseWithDefaults instantiates a new AuthNMappingsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthNMappingsResponseWithDefaults() *AuthNMappingsResponse {
	this := AuthNMappingsResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AuthNMappingsResponse) GetData() []AuthNMapping {
	if o == nil || o.Data == nil {
		var ret []AuthNMapping
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthNMappingsResponse) GetDataOk() (*[]AuthNMapping, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AuthNMappingsResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []AuthNMapping and assigns it to the Data field.
func (o *AuthNMappingsResponse) SetData(v []AuthNMapping) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AuthNMappingsResponse) GetMeta() ResponseMetaAttributes {
	if o == nil || o.Meta == nil {
		var ret ResponseMetaAttributes
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthNMappingsResponse) GetMetaOk() (*ResponseMetaAttributes, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AuthNMappingsResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ResponseMetaAttributes and assigns it to the Meta field.
func (o *AuthNMappingsResponse) SetMeta(v ResponseMetaAttributes) {
	o.Meta = &v
}

func (o AuthNMappingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

func (o *AuthNMappingsResponse) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Data []AuthNMapping          `json:"data,omitempty"`
		Meta *ResponseMetaAttributes `json:"meta,omitempty"`
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
	o.Data = all.Data
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Meta = all.Meta
	return nil
}
