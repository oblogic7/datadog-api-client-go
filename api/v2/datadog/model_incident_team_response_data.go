/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// IncidentTeamResponseData Incident Team data from a response.
type IncidentTeamResponseData struct {
	// The incident team's attributes from a response.
	Attributes *IncidentTeamResponseAttributes `json:"attributes,omitempty"`
	// The incident team's ID.
	Id *string `json:"id,omitempty"`
	// The incident team's relationships.
	Relationships *IncidentTeamRelationships `json:"relationships,omitempty"`
	// Incident Team resource type.
	Type *IncidentTeamType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewIncidentTeamResponseData instantiates a new IncidentTeamResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentTeamResponseData() *IncidentTeamResponseData {
	this := IncidentTeamResponseData{}
	var type_ IncidentTeamType = INCIDENTTEAMTYPE_TEAMS
	this.Type = &type_
	return &this
}

// NewIncidentTeamResponseDataWithDefaults instantiates a new IncidentTeamResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentTeamResponseDataWithDefaults() *IncidentTeamResponseData {
	this := IncidentTeamResponseData{}
	var type_ IncidentTeamType = INCIDENTTEAMTYPE_TEAMS
	this.Type = &type_
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *IncidentTeamResponseData) GetAttributes() IncidentTeamResponseAttributes {
	if o == nil || o.Attributes == nil {
		var ret IncidentTeamResponseAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTeamResponseData) GetAttributesOk() (*IncidentTeamResponseAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *IncidentTeamResponseData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given IncidentTeamResponseAttributes and assigns it to the Attributes field.
func (o *IncidentTeamResponseData) SetAttributes(v IncidentTeamResponseAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IncidentTeamResponseData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTeamResponseData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IncidentTeamResponseData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IncidentTeamResponseData) SetId(v string) {
	o.Id = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *IncidentTeamResponseData) GetRelationships() IncidentTeamRelationships {
	if o == nil || o.Relationships == nil {
		var ret IncidentTeamRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTeamResponseData) GetRelationshipsOk() (*IncidentTeamRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *IncidentTeamResponseData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given IncidentTeamRelationships and assigns it to the Relationships field.
func (o *IncidentTeamResponseData) SetRelationships(v IncidentTeamRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IncidentTeamResponseData) GetType() IncidentTeamType {
	if o == nil || o.Type == nil {
		var ret IncidentTeamType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTeamResponseData) GetTypeOk() (*IncidentTeamType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IncidentTeamResponseData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given IncidentTeamType and assigns it to the Type field.
func (o *IncidentTeamResponseData) SetType(v IncidentTeamType) {
	o.Type = &v
}

func (o IncidentTeamResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

func (o *IncidentTeamResponseData) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Attributes    *IncidentTeamResponseAttributes `json:"attributes,omitempty"`
		Id            *string                         `json:"id,omitempty"`
		Relationships *IncidentTeamRelationships      `json:"relationships,omitempty"`
		Type          *IncidentTeamType               `json:"type,omitempty"`
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
	if v := all.Type; v != nil && !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Attributes = all.Attributes
	o.Id = all.Id
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
