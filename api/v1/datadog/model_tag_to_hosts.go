/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// TagToHosts In this object, the key is the tag, the value is a list of host names that are reporting that tag.
type TagToHosts struct {
	// A list of tags to apply to the host.
	Tags map[string][]string `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewTagToHosts instantiates a new TagToHosts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagToHosts() *TagToHosts {
	this := TagToHosts{}
	return &this
}

// NewTagToHostsWithDefaults instantiates a new TagToHosts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagToHostsWithDefaults() *TagToHosts {
	this := TagToHosts{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *TagToHosts) GetTags() map[string][]string {
	if o == nil || o.Tags == nil {
		var ret map[string][]string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagToHosts) GetTagsOk() (*map[string][]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *TagToHosts) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string][]string and assigns it to the Tags field.
func (o *TagToHosts) SetTags(v map[string][]string) {
	o.Tags = v
}

func (o TagToHosts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

func (o *TagToHosts) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Tags map[string][]string `json:"tags,omitempty"`
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
	o.Tags = all.Tags
	return nil
}
