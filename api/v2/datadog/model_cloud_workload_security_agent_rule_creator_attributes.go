/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// CloudWorkloadSecurityAgentRuleCreatorAttributes The attributes of the user who created the Agent rule.
type CloudWorkloadSecurityAgentRuleCreatorAttributes struct {
	// The handle of the user.
	Handle *string `json:"handle,omitempty"`
	// The name of the user.
	Name *string `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewCloudWorkloadSecurityAgentRuleCreatorAttributes instantiates a new CloudWorkloadSecurityAgentRuleCreatorAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudWorkloadSecurityAgentRuleCreatorAttributes() *CloudWorkloadSecurityAgentRuleCreatorAttributes {
	this := CloudWorkloadSecurityAgentRuleCreatorAttributes{}
	return &this
}

// NewCloudWorkloadSecurityAgentRuleCreatorAttributesWithDefaults instantiates a new CloudWorkloadSecurityAgentRuleCreatorAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudWorkloadSecurityAgentRuleCreatorAttributesWithDefaults() *CloudWorkloadSecurityAgentRuleCreatorAttributes {
	this := CloudWorkloadSecurityAgentRuleCreatorAttributes{}
	return &this
}

// GetHandle returns the Handle field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleCreatorAttributes) GetHandle() string {
	if o == nil || o.Handle == nil {
		var ret string
		return ret
	}
	return *o.Handle
}

// GetHandleOk returns a tuple with the Handle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleCreatorAttributes) GetHandleOk() (*string, bool) {
	if o == nil || o.Handle == nil {
		return nil, false
	}
	return o.Handle, true
}

// HasHandle returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleCreatorAttributes) HasHandle() bool {
	if o != nil && o.Handle != nil {
		return true
	}

	return false
}

// SetHandle gets a reference to the given string and assigns it to the Handle field.
func (o *CloudWorkloadSecurityAgentRuleCreatorAttributes) SetHandle(v string) {
	o.Handle = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleCreatorAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleCreatorAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleCreatorAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CloudWorkloadSecurityAgentRuleCreatorAttributes) SetName(v string) {
	o.Name = &v
}

func (o CloudWorkloadSecurityAgentRuleCreatorAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Handle != nil {
		toSerialize["handle"] = o.Handle
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

func (o *CloudWorkloadSecurityAgentRuleCreatorAttributes) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Handle *string `json:"handle,omitempty"`
		Name   *string `json:"name,omitempty"`
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
	o.Handle = all.Handle
	o.Name = all.Name
	return nil
}
