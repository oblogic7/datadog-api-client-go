/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// LogsMetricUpdateAttributes The log-based metric properties that will be updated.
type LogsMetricUpdateAttributes struct {
	// The log-based metric filter. Logs matching this filter will be aggregated in this metric.
	Filter *LogsMetricFilter `json:"filter,omitempty"`
	// The rules for the group by.
	GroupBy []LogsMetricGroupBy `json:"group_by,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewLogsMetricUpdateAttributes instantiates a new LogsMetricUpdateAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogsMetricUpdateAttributes() *LogsMetricUpdateAttributes {
	this := LogsMetricUpdateAttributes{}
	return &this
}

// NewLogsMetricUpdateAttributesWithDefaults instantiates a new LogsMetricUpdateAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogsMetricUpdateAttributesWithDefaults() *LogsMetricUpdateAttributes {
	this := LogsMetricUpdateAttributes{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *LogsMetricUpdateAttributes) GetFilter() LogsMetricFilter {
	if o == nil || o.Filter == nil {
		var ret LogsMetricFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsMetricUpdateAttributes) GetFilterOk() (*LogsMetricFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *LogsMetricUpdateAttributes) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given LogsMetricFilter and assigns it to the Filter field.
func (o *LogsMetricUpdateAttributes) SetFilter(v LogsMetricFilter) {
	o.Filter = &v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *LogsMetricUpdateAttributes) GetGroupBy() []LogsMetricGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []LogsMetricGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsMetricUpdateAttributes) GetGroupByOk() (*[]LogsMetricGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *LogsMetricUpdateAttributes) HasGroupBy() bool {
	if o != nil && o.GroupBy != nil {
		return true
	}

	return false
}

// SetGroupBy gets a reference to the given []LogsMetricGroupBy and assigns it to the GroupBy field.
func (o *LogsMetricUpdateAttributes) SetGroupBy(v []LogsMetricGroupBy) {
	o.GroupBy = v
}

func (o LogsMetricUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

func (o *LogsMetricUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Filter  *LogsMetricFilter   `json:"filter,omitempty"`
		GroupBy []LogsMetricGroupBy `json:"group_by,omitempty"`
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
	if all.Filter != nil && all.Filter.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Filter = all.Filter
	o.GroupBy = all.GroupBy
	return nil
}
