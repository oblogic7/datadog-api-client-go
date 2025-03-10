/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// SLOBulkDeleteResponse The bulk partial delete service level objective object endpoint
// response.
//
// This endpoint operates on multiple service level objective objects, so
// it may be partially successful. In such cases, the "data" and "error"
// fields in this response indicate which deletions succeeded and failed.
type SLOBulkDeleteResponse struct {
	// An array of service level objective objects.
	Data *SLOBulkDeleteResponseData `json:"data,omitempty"`
	// Array of errors object returned.
	Errors []SLOBulkDeleteError `json:"errors,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewSLOBulkDeleteResponse instantiates a new SLOBulkDeleteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSLOBulkDeleteResponse() *SLOBulkDeleteResponse {
	this := SLOBulkDeleteResponse{}
	return &this
}

// NewSLOBulkDeleteResponseWithDefaults instantiates a new SLOBulkDeleteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSLOBulkDeleteResponseWithDefaults() *SLOBulkDeleteResponse {
	this := SLOBulkDeleteResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SLOBulkDeleteResponse) GetData() SLOBulkDeleteResponseData {
	if o == nil || o.Data == nil {
		var ret SLOBulkDeleteResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOBulkDeleteResponse) GetDataOk() (*SLOBulkDeleteResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SLOBulkDeleteResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given SLOBulkDeleteResponseData and assigns it to the Data field.
func (o *SLOBulkDeleteResponse) SetData(v SLOBulkDeleteResponseData) {
	o.Data = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SLOBulkDeleteResponse) GetErrors() []SLOBulkDeleteError {
	if o == nil || o.Errors == nil {
		var ret []SLOBulkDeleteError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOBulkDeleteResponse) GetErrorsOk() (*[]SLOBulkDeleteError, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return &o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SLOBulkDeleteResponse) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SLOBulkDeleteError and assigns it to the Errors field.
func (o *SLOBulkDeleteResponse) SetErrors(v []SLOBulkDeleteError) {
	o.Errors = v
}

func (o SLOBulkDeleteResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

func (o *SLOBulkDeleteResponse) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Data   *SLOBulkDeleteResponseData `json:"data,omitempty"`
		Errors []SLOBulkDeleteError       `json:"errors,omitempty"`
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
	if all.Data != nil && all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Data = all.Data
	o.Errors = all.Errors
	return nil
}
