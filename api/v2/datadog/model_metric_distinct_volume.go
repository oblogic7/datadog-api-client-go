/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// MetricDistinctVolume Object for a single metric's distinct volume.
type MetricDistinctVolume struct {
	// Object containing the definition of a metric's distinct volume.
	Attributes *MetricDistinctVolumeAttributes `json:"attributes,omitempty"`
	// The metric name for this resource.
	Id *string `json:"id,omitempty"`
	// The metric distinct volume type.
	Type *MetricDistinctVolumeType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewMetricDistinctVolume instantiates a new MetricDistinctVolume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricDistinctVolume() *MetricDistinctVolume {
	this := MetricDistinctVolume{}
	var type_ MetricDistinctVolumeType = METRICDISTINCTVOLUMETYPE_DISTINCT_METRIC_VOLUMES
	this.Type = &type_
	return &this
}

// NewMetricDistinctVolumeWithDefaults instantiates a new MetricDistinctVolume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricDistinctVolumeWithDefaults() *MetricDistinctVolume {
	this := MetricDistinctVolume{}
	var type_ MetricDistinctVolumeType = METRICDISTINCTVOLUMETYPE_DISTINCT_METRIC_VOLUMES
	this.Type = &type_
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *MetricDistinctVolume) GetAttributes() MetricDistinctVolumeAttributes {
	if o == nil || o.Attributes == nil {
		var ret MetricDistinctVolumeAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDistinctVolume) GetAttributesOk() (*MetricDistinctVolumeAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *MetricDistinctVolume) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given MetricDistinctVolumeAttributes and assigns it to the Attributes field.
func (o *MetricDistinctVolume) SetAttributes(v MetricDistinctVolumeAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MetricDistinctVolume) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDistinctVolume) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MetricDistinctVolume) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MetricDistinctVolume) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MetricDistinctVolume) GetType() MetricDistinctVolumeType {
	if o == nil || o.Type == nil {
		var ret MetricDistinctVolumeType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDistinctVolume) GetTypeOk() (*MetricDistinctVolumeType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MetricDistinctVolume) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given MetricDistinctVolumeType and assigns it to the Type field.
func (o *MetricDistinctVolume) SetType(v MetricDistinctVolumeType) {
	o.Type = &v
}

func (o MetricDistinctVolume) MarshalJSON() ([]byte, error) {
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
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

func (o *MetricDistinctVolume) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Attributes *MetricDistinctVolumeAttributes `json:"attributes,omitempty"`
		Id         *string                         `json:"id,omitempty"`
		Type       *MetricDistinctVolumeType       `json:"type,omitempty"`
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
	o.Type = all.Type
	return nil
}
