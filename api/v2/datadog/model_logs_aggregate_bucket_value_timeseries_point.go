/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// LogsAggregateBucketValueTimeseriesPoint A timeseries point
type LogsAggregateBucketValueTimeseriesPoint struct {
	// The time value for this point
	Time *string `json:"time,omitempty"`
	// The value for this point
	Value *float64 `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewLogsAggregateBucketValueTimeseriesPoint instantiates a new LogsAggregateBucketValueTimeseriesPoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogsAggregateBucketValueTimeseriesPoint() *LogsAggregateBucketValueTimeseriesPoint {
	this := LogsAggregateBucketValueTimeseriesPoint{}
	return &this
}

// NewLogsAggregateBucketValueTimeseriesPointWithDefaults instantiates a new LogsAggregateBucketValueTimeseriesPoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogsAggregateBucketValueTimeseriesPointWithDefaults() *LogsAggregateBucketValueTimeseriesPoint {
	this := LogsAggregateBucketValueTimeseriesPoint{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *LogsAggregateBucketValueTimeseriesPoint) GetTime() string {
	if o == nil || o.Time == nil {
		var ret string
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsAggregateBucketValueTimeseriesPoint) GetTimeOk() (*string, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *LogsAggregateBucketValueTimeseriesPoint) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given string and assigns it to the Time field.
func (o *LogsAggregateBucketValueTimeseriesPoint) SetTime(v string) {
	o.Time = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *LogsAggregateBucketValueTimeseriesPoint) GetValue() float64 {
	if o == nil || o.Value == nil {
		var ret float64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsAggregateBucketValueTimeseriesPoint) GetValueOk() (*float64, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *LogsAggregateBucketValueTimeseriesPoint) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given float64 and assigns it to the Value field.
func (o *LogsAggregateBucketValueTimeseriesPoint) SetValue(v float64) {
	o.Value = &v
}

func (o LogsAggregateBucketValueTimeseriesPoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

func (o *LogsAggregateBucketValueTimeseriesPoint) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Time  *string  `json:"time,omitempty"`
		Value *float64 `json:"value,omitempty"`
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
	o.Time = all.Time
	o.Value = all.Value
	return nil
}
