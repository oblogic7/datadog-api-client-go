/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// SyntheticsCIBatchMetadataProvider Description of the CI provider.
type SyntheticsCIBatchMetadataProvider struct {
	// Name of the CI provider.
	Name *string `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewSyntheticsCIBatchMetadataProvider instantiates a new SyntheticsCIBatchMetadataProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticsCIBatchMetadataProvider() *SyntheticsCIBatchMetadataProvider {
	this := SyntheticsCIBatchMetadataProvider{}
	return &this
}

// NewSyntheticsCIBatchMetadataProviderWithDefaults instantiates a new SyntheticsCIBatchMetadataProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticsCIBatchMetadataProviderWithDefaults() *SyntheticsCIBatchMetadataProvider {
	this := SyntheticsCIBatchMetadataProvider{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SyntheticsCIBatchMetadataProvider) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsCIBatchMetadataProvider) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SyntheticsCIBatchMetadataProvider) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SyntheticsCIBatchMetadataProvider) SetName(v string) {
	o.Name = &v
}

func (o SyntheticsCIBatchMetadataProvider) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

func (o *SyntheticsCIBatchMetadataProvider) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Name *string `json:"name,omitempty"`
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
	o.Name = all.Name
	return nil
}
