/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
	"time"
)

// UsageSyntheticsHour The number of synthetics tests run for each hour for a given organization.
type UsageSyntheticsHour struct {
	// Contains the number of Synthetics API tests run.
	CheckCallsCount *int64 `json:"check_calls_count,omitempty"`
	// The hour for the usage.
	Hour *time.Time `json:"hour,omitempty"`
	// The organization name.
	OrgName *string `json:"org_name,omitempty"`
	// The organization public ID.
	PublicId *string `json:"public_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewUsageSyntheticsHour instantiates a new UsageSyntheticsHour object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageSyntheticsHour() *UsageSyntheticsHour {
	this := UsageSyntheticsHour{}
	return &this
}

// NewUsageSyntheticsHourWithDefaults instantiates a new UsageSyntheticsHour object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageSyntheticsHourWithDefaults() *UsageSyntheticsHour {
	this := UsageSyntheticsHour{}
	return &this
}

// GetCheckCallsCount returns the CheckCallsCount field value if set, zero value otherwise.
func (o *UsageSyntheticsHour) GetCheckCallsCount() int64 {
	if o == nil || o.CheckCallsCount == nil {
		var ret int64
		return ret
	}
	return *o.CheckCallsCount
}

// GetCheckCallsCountOk returns a tuple with the CheckCallsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSyntheticsHour) GetCheckCallsCountOk() (*int64, bool) {
	if o == nil || o.CheckCallsCount == nil {
		return nil, false
	}
	return o.CheckCallsCount, true
}

// HasCheckCallsCount returns a boolean if a field has been set.
func (o *UsageSyntheticsHour) HasCheckCallsCount() bool {
	if o != nil && o.CheckCallsCount != nil {
		return true
	}

	return false
}

// SetCheckCallsCount gets a reference to the given int64 and assigns it to the CheckCallsCount field.
func (o *UsageSyntheticsHour) SetCheckCallsCount(v int64) {
	o.CheckCallsCount = &v
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *UsageSyntheticsHour) GetHour() time.Time {
	if o == nil || o.Hour == nil {
		var ret time.Time
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSyntheticsHour) GetHourOk() (*time.Time, bool) {
	if o == nil || o.Hour == nil {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *UsageSyntheticsHour) HasHour() bool {
	if o != nil && o.Hour != nil {
		return true
	}

	return false
}

// SetHour gets a reference to the given time.Time and assigns it to the Hour field.
func (o *UsageSyntheticsHour) SetHour(v time.Time) {
	o.Hour = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *UsageSyntheticsHour) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSyntheticsHour) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *UsageSyntheticsHour) HasOrgName() bool {
	if o != nil && o.OrgName != nil {
		return true
	}

	return false
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *UsageSyntheticsHour) SetOrgName(v string) {
	o.OrgName = &v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *UsageSyntheticsHour) GetPublicId() string {
	if o == nil || o.PublicId == nil {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSyntheticsHour) GetPublicIdOk() (*string, bool) {
	if o == nil || o.PublicId == nil {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *UsageSyntheticsHour) HasPublicId() bool {
	if o != nil && o.PublicId != nil {
		return true
	}

	return false
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *UsageSyntheticsHour) SetPublicId(v string) {
	o.PublicId = &v
}

func (o UsageSyntheticsHour) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.CheckCallsCount != nil {
		toSerialize["check_calls_count"] = o.CheckCallsCount
	}
	if o.Hour != nil {
		if o.Hour.Nanosecond() == 0 {
			toSerialize["hour"] = o.Hour.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["hour"] = o.Hour.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.OrgName != nil {
		toSerialize["org_name"] = o.OrgName
	}
	if o.PublicId != nil {
		toSerialize["public_id"] = o.PublicId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

func (o *UsageSyntheticsHour) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		CheckCallsCount *int64     `json:"check_calls_count,omitempty"`
		Hour            *time.Time `json:"hour,omitempty"`
		OrgName         *string    `json:"org_name,omitempty"`
		PublicId        *string    `json:"public_id,omitempty"`
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
	o.CheckCallsCount = all.CheckCallsCount
	o.Hour = all.Hour
	o.OrgName = all.OrgName
	o.PublicId = all.PublicId
	return nil
}
