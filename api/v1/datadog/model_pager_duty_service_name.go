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

// PagerDutyServiceName PagerDuty service object name.
type PagerDutyServiceName struct {
	// Your service name associated service key in PagerDuty.
	ServiceName string `json:"service_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewPagerDutyServiceName instantiates a new PagerDutyServiceName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagerDutyServiceName(serviceName string) *PagerDutyServiceName {
	this := PagerDutyServiceName{}
	this.ServiceName = serviceName
	return &this
}

// NewPagerDutyServiceNameWithDefaults instantiates a new PagerDutyServiceName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagerDutyServiceNameWithDefaults() *PagerDutyServiceName {
	this := PagerDutyServiceName{}
	return &this
}

// GetServiceName returns the ServiceName field value
func (o *PagerDutyServiceName) GetServiceName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value
// and a boolean to check if the value has been set.
func (o *PagerDutyServiceName) GetServiceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceName, true
}

// SetServiceName sets field value
func (o *PagerDutyServiceName) SetServiceName(v string) {
	o.ServiceName = v
}

func (o PagerDutyServiceName) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["service_name"] = o.ServiceName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

func (o *PagerDutyServiceName) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		ServiceName *string `json:"service_name"`
	}{}
	all := struct {
		ServiceName string `json:"service_name"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.ServiceName == nil {
		return fmt.Errorf("Required field service_name missing")
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
	o.ServiceName = all.ServiceName
	return nil
}
