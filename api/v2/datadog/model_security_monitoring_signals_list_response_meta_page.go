/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// SecurityMonitoringSignalsListResponseMetaPage Paging attributes.
type SecurityMonitoringSignalsListResponseMetaPage struct {
	// The cursor used to get the next results, if any. To make the next request, use the same
	// parameters with the addition of the `page[cursor]`.
	After *string `json:"after,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewSecurityMonitoringSignalsListResponseMetaPage instantiates a new SecurityMonitoringSignalsListResponseMetaPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityMonitoringSignalsListResponseMetaPage() *SecurityMonitoringSignalsListResponseMetaPage {
	this := SecurityMonitoringSignalsListResponseMetaPage{}
	return &this
}

// NewSecurityMonitoringSignalsListResponseMetaPageWithDefaults instantiates a new SecurityMonitoringSignalsListResponseMetaPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityMonitoringSignalsListResponseMetaPageWithDefaults() *SecurityMonitoringSignalsListResponseMetaPage {
	this := SecurityMonitoringSignalsListResponseMetaPage{}
	return &this
}

// GetAfter returns the After field value if set, zero value otherwise.
func (o *SecurityMonitoringSignalsListResponseMetaPage) GetAfter() string {
	if o == nil || o.After == nil {
		var ret string
		return ret
	}
	return *o.After
}

// GetAfterOk returns a tuple with the After field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalsListResponseMetaPage) GetAfterOk() (*string, bool) {
	if o == nil || o.After == nil {
		return nil, false
	}
	return o.After, true
}

// HasAfter returns a boolean if a field has been set.
func (o *SecurityMonitoringSignalsListResponseMetaPage) HasAfter() bool {
	if o != nil && o.After != nil {
		return true
	}

	return false
}

// SetAfter gets a reference to the given string and assigns it to the After field.
func (o *SecurityMonitoringSignalsListResponseMetaPage) SetAfter(v string) {
	o.After = &v
}

func (o SecurityMonitoringSignalsListResponseMetaPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.After != nil {
		toSerialize["after"] = o.After
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

func (o *SecurityMonitoringSignalsListResponseMetaPage) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		After *string `json:"after,omitempty"`
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
	o.After = all.After
	return nil
}
