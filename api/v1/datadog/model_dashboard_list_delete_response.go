/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// DashboardListDeleteResponse Deleted dashboard details.
type DashboardListDeleteResponse struct {
	// ID of the deleted dashboard list.
	DeletedDashboardListId *int64 `json:"deleted_dashboard_list_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewDashboardListDeleteResponse instantiates a new DashboardListDeleteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardListDeleteResponse() *DashboardListDeleteResponse {
	this := DashboardListDeleteResponse{}
	return &this
}

// NewDashboardListDeleteResponseWithDefaults instantiates a new DashboardListDeleteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardListDeleteResponseWithDefaults() *DashboardListDeleteResponse {
	this := DashboardListDeleteResponse{}
	return &this
}

// GetDeletedDashboardListId returns the DeletedDashboardListId field value if set, zero value otherwise.
func (o *DashboardListDeleteResponse) GetDeletedDashboardListId() int64 {
	if o == nil || o.DeletedDashboardListId == nil {
		var ret int64
		return ret
	}
	return *o.DeletedDashboardListId
}

// GetDeletedDashboardListIdOk returns a tuple with the DeletedDashboardListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardListDeleteResponse) GetDeletedDashboardListIdOk() (*int64, bool) {
	if o == nil || o.DeletedDashboardListId == nil {
		return nil, false
	}
	return o.DeletedDashboardListId, true
}

// HasDeletedDashboardListId returns a boolean if a field has been set.
func (o *DashboardListDeleteResponse) HasDeletedDashboardListId() bool {
	if o != nil && o.DeletedDashboardListId != nil {
		return true
	}

	return false
}

// SetDeletedDashboardListId gets a reference to the given int64 and assigns it to the DeletedDashboardListId field.
func (o *DashboardListDeleteResponse) SetDeletedDashboardListId(v int64) {
	o.DeletedDashboardListId = &v
}

func (o DashboardListDeleteResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.DeletedDashboardListId != nil {
		toSerialize["deleted_dashboard_list_id"] = o.DeletedDashboardListId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

func (o *DashboardListDeleteResponse) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		DeletedDashboardListId *int64 `json:"deleted_dashboard_list_id,omitempty"`
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
	o.DeletedDashboardListId = all.DeletedDashboardListId
	return nil
}
