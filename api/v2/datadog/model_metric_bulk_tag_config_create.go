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

// MetricBulkTagConfigCreate Request object to bulk configure tags for metrics matching the given prefix.
type MetricBulkTagConfigCreate struct {
	// Optional parameters for bulk creating metric tag configurations.
	Attributes *MetricBulkTagConfigCreateAttributes `json:"attributes,omitempty"`
	// A text prefix to match against metric names.
	Id string `json:"id"`
	// The metric bulk configure tags resource.
	Type MetricBulkConfigureTagsType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewMetricBulkTagConfigCreate instantiates a new MetricBulkTagConfigCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricBulkTagConfigCreate(id string, type_ MetricBulkConfigureTagsType) *MetricBulkTagConfigCreate {
	this := MetricBulkTagConfigCreate{}
	this.Id = id
	this.Type = type_
	return &this
}

// NewMetricBulkTagConfigCreateWithDefaults instantiates a new MetricBulkTagConfigCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricBulkTagConfigCreateWithDefaults() *MetricBulkTagConfigCreate {
	this := MetricBulkTagConfigCreate{}
	var type_ MetricBulkConfigureTagsType = METRICBULKCONFIGURETAGSTYPE_BULK_MANAGE_TAGS
	this.Type = type_
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *MetricBulkTagConfigCreate) GetAttributes() MetricBulkTagConfigCreateAttributes {
	if o == nil || o.Attributes == nil {
		var ret MetricBulkTagConfigCreateAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricBulkTagConfigCreate) GetAttributesOk() (*MetricBulkTagConfigCreateAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *MetricBulkTagConfigCreate) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given MetricBulkTagConfigCreateAttributes and assigns it to the Attributes field.
func (o *MetricBulkTagConfigCreate) SetAttributes(v MetricBulkTagConfigCreateAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value
func (o *MetricBulkTagConfigCreate) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MetricBulkTagConfigCreate) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MetricBulkTagConfigCreate) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *MetricBulkTagConfigCreate) GetType() MetricBulkConfigureTagsType {
	if o == nil {
		var ret MetricBulkConfigureTagsType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MetricBulkTagConfigCreate) GetTypeOk() (*MetricBulkConfigureTagsType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MetricBulkTagConfigCreate) SetType(v MetricBulkConfigureTagsType) {
	o.Type = v
}

func (o MetricBulkTagConfigCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

func (o *MetricBulkTagConfigCreate) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Id   *string                      `json:"id"`
		Type *MetricBulkConfigureTagsType `json:"type"`
	}{}
	all := struct {
		Attributes *MetricBulkTagConfigCreateAttributes `json:"attributes,omitempty"`
		Id         string                               `json:"id"`
		Type       MetricBulkConfigureTagsType          `json:"type"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Id == nil {
		return fmt.Errorf("Required field id missing")
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
