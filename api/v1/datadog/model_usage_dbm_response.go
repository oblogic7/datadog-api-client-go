/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// UsageDBMResponse Response containing the Database Monitoring usage for each hour for a given organization.
type UsageDBMResponse struct {
	// Get hourly usage for Database Monitoring
	Usage []UsageDBMHour `json:"usage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewUsageDBMResponse instantiates a new UsageDBMResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageDBMResponse() *UsageDBMResponse {
	this := UsageDBMResponse{}
	return &this
}

// NewUsageDBMResponseWithDefaults instantiates a new UsageDBMResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageDBMResponseWithDefaults() *UsageDBMResponse {
	this := UsageDBMResponse{}
	return &this
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *UsageDBMResponse) GetUsage() []UsageDBMHour {
	if o == nil || o.Usage == nil {
		var ret []UsageDBMHour
		return ret
	}
	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageDBMResponse) GetUsageOk() (*[]UsageDBMHour, bool) {
	if o == nil || o.Usage == nil {
		return nil, false
	}
	return &o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *UsageDBMResponse) HasUsage() bool {
	if o != nil && o.Usage != nil {
		return true
	}

	return false
}

// SetUsage gets a reference to the given []UsageDBMHour and assigns it to the Usage field.
func (o *UsageDBMResponse) SetUsage(v []UsageDBMHour) {
	o.Usage = v
}

func (o UsageDBMResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Usage != nil {
		toSerialize["usage"] = o.Usage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

func (o *UsageDBMResponse) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Usage []UsageDBMHour `json:"usage,omitempty"`
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
	o.Usage = all.Usage
	return nil
}
