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

// IncidentServiceUpdateRequest Update request with an incident service payload.
type IncidentServiceUpdateRequest struct {
	// Incident Service payload for update requests.
	Data IncidentServiceUpdateData `json:"data"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewIncidentServiceUpdateRequest instantiates a new IncidentServiceUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentServiceUpdateRequest(data IncidentServiceUpdateData) *IncidentServiceUpdateRequest {
	this := IncidentServiceUpdateRequest{}
	this.Data = data
	return &this
}

// NewIncidentServiceUpdateRequestWithDefaults instantiates a new IncidentServiceUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentServiceUpdateRequestWithDefaults() *IncidentServiceUpdateRequest {
	this := IncidentServiceUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *IncidentServiceUpdateRequest) GetData() IncidentServiceUpdateData {
	if o == nil {
		var ret IncidentServiceUpdateData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *IncidentServiceUpdateRequest) GetDataOk() (*IncidentServiceUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *IncidentServiceUpdateRequest) SetData(v IncidentServiceUpdateData) {
	o.Data = v
}

func (o IncidentServiceUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

func (o *IncidentServiceUpdateRequest) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Data *IncidentServiceUpdateData `json:"data"`
	}{}
	all := struct {
		Data IncidentServiceUpdateData `json:"data"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Data == nil {
		return fmt.Errorf("Required field data missing")
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
	if all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Data = all.Data
	return nil
}
