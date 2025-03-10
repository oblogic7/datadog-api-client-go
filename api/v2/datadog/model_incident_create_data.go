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

// IncidentCreateData Incident data for a create request.
type IncidentCreateData struct {
	// The incident's attributes for a create request.
	Attributes IncidentCreateAttributes `json:"attributes"`
	// The relationships the incident will have with other resources once created.
	Relationships *IncidentCreateRelationships `json:"relationships,omitempty"`
	// Incident resource type.
	Type IncidentType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewIncidentCreateData instantiates a new IncidentCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentCreateData(attributes IncidentCreateAttributes, type_ IncidentType) *IncidentCreateData {
	this := IncidentCreateData{}
	this.Attributes = attributes
	this.Type = type_
	return &this
}

// NewIncidentCreateDataWithDefaults instantiates a new IncidentCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentCreateDataWithDefaults() *IncidentCreateData {
	this := IncidentCreateData{}
	var type_ IncidentType = INCIDENTTYPE_INCIDENTS
	this.Type = type_
	return &this
}

// GetAttributes returns the Attributes field value
func (o *IncidentCreateData) GetAttributes() IncidentCreateAttributes {
	if o == nil {
		var ret IncidentCreateAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *IncidentCreateData) GetAttributesOk() (*IncidentCreateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *IncidentCreateData) SetAttributes(v IncidentCreateAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *IncidentCreateData) GetRelationships() IncidentCreateRelationships {
	if o == nil || o.Relationships == nil {
		var ret IncidentCreateRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentCreateData) GetRelationshipsOk() (*IncidentCreateRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *IncidentCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given IncidentCreateRelationships and assigns it to the Relationships field.
func (o *IncidentCreateData) SetRelationships(v IncidentCreateRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value
func (o *IncidentCreateData) GetType() IncidentType {
	if o == nil {
		var ret IncidentType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IncidentCreateData) GetTypeOk() (*IncidentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *IncidentCreateData) SetType(v IncidentType) {
	o.Type = v
}

func (o IncidentCreateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

func (o *IncidentCreateData) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Attributes *IncidentCreateAttributes `json:"attributes"`
		Type       *IncidentType             `json:"type"`
	}{}
	all := struct {
		Attributes    IncidentCreateAttributes     `json:"attributes"`
		Relationships *IncidentCreateRelationships `json:"relationships,omitempty"`
		Type          IncidentType                 `json:"type"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Attributes == nil {
		return fmt.Errorf("Required field attributes missing")
	}
	if required.Type == nil {
		return fmt.Errorf("Required field type missing")
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
	if v := all.Type; !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Attributes = all.Attributes
	if all.Relationships != nil && all.Relationships.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Relationships = all.Relationships
	o.Type = all.Type
	return nil
}
