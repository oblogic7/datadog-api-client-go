/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// UsageTopAvgMetricsResponse Response containing the number of hourly recorded custom metrics for a given organization.
type UsageTopAvgMetricsResponse struct {
	// The object containing document metadata.
	Metadata *UsageTopAvgMetricsMetadata `json:"metadata,omitempty"`
	// Number of hourly recorded custom metrics for a given organization.
	Usage []UsageTopAvgMetricsHour `json:"usage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewUsageTopAvgMetricsResponse instantiates a new UsageTopAvgMetricsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageTopAvgMetricsResponse() *UsageTopAvgMetricsResponse {
	this := UsageTopAvgMetricsResponse{}
	return &this
}

// NewUsageTopAvgMetricsResponseWithDefaults instantiates a new UsageTopAvgMetricsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageTopAvgMetricsResponseWithDefaults() *UsageTopAvgMetricsResponse {
	this := UsageTopAvgMetricsResponse{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UsageTopAvgMetricsResponse) GetMetadata() UsageTopAvgMetricsMetadata {
	if o == nil || o.Metadata == nil {
		var ret UsageTopAvgMetricsMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageTopAvgMetricsResponse) GetMetadataOk() (*UsageTopAvgMetricsMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UsageTopAvgMetricsResponse) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given UsageTopAvgMetricsMetadata and assigns it to the Metadata field.
func (o *UsageTopAvgMetricsResponse) SetMetadata(v UsageTopAvgMetricsMetadata) {
	o.Metadata = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *UsageTopAvgMetricsResponse) GetUsage() []UsageTopAvgMetricsHour {
	if o == nil || o.Usage == nil {
		var ret []UsageTopAvgMetricsHour
		return ret
	}
	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageTopAvgMetricsResponse) GetUsageOk() (*[]UsageTopAvgMetricsHour, bool) {
	if o == nil || o.Usage == nil {
		return nil, false
	}
	return &o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *UsageTopAvgMetricsResponse) HasUsage() bool {
	if o != nil && o.Usage != nil {
		return true
	}

	return false
}

// SetUsage gets a reference to the given []UsageTopAvgMetricsHour and assigns it to the Usage field.
func (o *UsageTopAvgMetricsResponse) SetUsage(v []UsageTopAvgMetricsHour) {
	o.Usage = v
}

func (o UsageTopAvgMetricsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Usage != nil {
		toSerialize["usage"] = o.Usage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

func (o *UsageTopAvgMetricsResponse) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Metadata *UsageTopAvgMetricsMetadata `json:"metadata,omitempty"`
		Usage    []UsageTopAvgMetricsHour    `json:"usage,omitempty"`
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
	if all.Metadata != nil && all.Metadata.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Metadata = all.Metadata
	o.Usage = all.Usage
	return nil
}
