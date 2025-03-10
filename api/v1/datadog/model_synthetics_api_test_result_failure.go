/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// SyntheticsApiTestResultFailure The API test failure details.
type SyntheticsApiTestResultFailure struct {
	// Error code that can be returned by a Synthetic test.
	Code *SyntheticsApiTestFailureCode `json:"code,omitempty"`
	// The API test error message.
	Message *string `json:"message,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewSyntheticsApiTestResultFailure instantiates a new SyntheticsApiTestResultFailure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticsApiTestResultFailure() *SyntheticsApiTestResultFailure {
	this := SyntheticsApiTestResultFailure{}
	return &this
}

// NewSyntheticsApiTestResultFailureWithDefaults instantiates a new SyntheticsApiTestResultFailure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticsApiTestResultFailureWithDefaults() *SyntheticsApiTestResultFailure {
	this := SyntheticsApiTestResultFailure{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SyntheticsApiTestResultFailure) GetCode() SyntheticsApiTestFailureCode {
	if o == nil || o.Code == nil {
		var ret SyntheticsApiTestFailureCode
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsApiTestResultFailure) GetCodeOk() (*SyntheticsApiTestFailureCode, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SyntheticsApiTestResultFailure) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given SyntheticsApiTestFailureCode and assigns it to the Code field.
func (o *SyntheticsApiTestResultFailure) SetCode(v SyntheticsApiTestFailureCode) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SyntheticsApiTestResultFailure) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsApiTestResultFailure) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SyntheticsApiTestResultFailure) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SyntheticsApiTestResultFailure) SetMessage(v string) {
	o.Message = &v
}

func (o SyntheticsApiTestResultFailure) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

func (o *SyntheticsApiTestResultFailure) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Code    *SyntheticsApiTestFailureCode `json:"code,omitempty"`
		Message *string                       `json:"message,omitempty"`
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
	if v := all.Code; v != nil && !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.Code = all.Code
	o.Message = all.Message
	return nil
}
