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

// IncidentUpdateAttributes The incident's attributes for an update request.
type IncidentUpdateAttributes struct {
	// Timestamp when customers were no longer impacted by the incident.
	CustomerImpactEnd NullableTime `json:"customer_impact_end,omitempty"`
	// A summary of the impact customers experienced during the incident.
	CustomerImpactScope *string `json:"customer_impact_scope,omitempty"`
	// Timestamp when customers began being impacted by the incident.
	CustomerImpactStart NullableTime `json:"customer_impact_start,omitempty"`
	// A flag indicating whether the incident caused customer impact.
	CustomerImpacted *bool `json:"customer_impacted,omitempty"`
	// Timestamp when the incident was detected.
	Detected NullableTime `json:"detected,omitempty"`
	// A condensed view of the user-defined fields for which to update selections.
	Fields map[string]IncidentFieldAttributes `json:"fields,omitempty"`
	// Notification handles that will be notified of the incident during update.
	NotificationHandles []IncidentNotificationHandle `json:"notification_handles,omitempty"`
	// Timestamp when the incident's state was set to resolved.
	Resolved NullableTime `json:"resolved,omitempty"`
	// The title of the incident, which summarizes what happened.
	Title *string `json:"title,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewIncidentUpdateAttributes instantiates a new IncidentUpdateAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentUpdateAttributes() *IncidentUpdateAttributes {
	this := IncidentUpdateAttributes{}
	return &this
}

// NewIncidentUpdateAttributesWithDefaults instantiates a new IncidentUpdateAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentUpdateAttributesWithDefaults() *IncidentUpdateAttributes {
	this := IncidentUpdateAttributes{}
	return &this
}

// GetCustomerImpactEnd returns the CustomerImpactEnd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentUpdateAttributes) GetCustomerImpactEnd() time.Time {
	if o == nil || o.CustomerImpactEnd.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CustomerImpactEnd.Get()
}

// GetCustomerImpactEndOk returns a tuple with the CustomerImpactEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncidentUpdateAttributes) GetCustomerImpactEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerImpactEnd.Get(), o.CustomerImpactEnd.IsSet()
}

// HasCustomerImpactEnd returns a boolean if a field has been set.
func (o *IncidentUpdateAttributes) HasCustomerImpactEnd() bool {
	if o != nil && o.CustomerImpactEnd.IsSet() {
		return true
	}

	return false
}

// SetCustomerImpactEnd gets a reference to the given NullableTime and assigns it to the CustomerImpactEnd field.
func (o *IncidentUpdateAttributes) SetCustomerImpactEnd(v time.Time) {
	o.CustomerImpactEnd.Set(&v)
}

// SetCustomerImpactEndNil sets the value for CustomerImpactEnd to be an explicit nil
func (o *IncidentUpdateAttributes) SetCustomerImpactEndNil() {
	o.CustomerImpactEnd.Set(nil)
}

// UnsetCustomerImpactEnd ensures that no value is present for CustomerImpactEnd, not even an explicit nil
func (o *IncidentUpdateAttributes) UnsetCustomerImpactEnd() {
	o.CustomerImpactEnd.Unset()
}

// GetCustomerImpactScope returns the CustomerImpactScope field value if set, zero value otherwise.
func (o *IncidentUpdateAttributes) GetCustomerImpactScope() string {
	if o == nil || o.CustomerImpactScope == nil {
		var ret string
		return ret
	}
	return *o.CustomerImpactScope
}

// GetCustomerImpactScopeOk returns a tuple with the CustomerImpactScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentUpdateAttributes) GetCustomerImpactScopeOk() (*string, bool) {
	if o == nil || o.CustomerImpactScope == nil {
		return nil, false
	}
	return o.CustomerImpactScope, true
}

// HasCustomerImpactScope returns a boolean if a field has been set.
func (o *IncidentUpdateAttributes) HasCustomerImpactScope() bool {
	if o != nil && o.CustomerImpactScope != nil {
		return true
	}

	return false
}

// SetCustomerImpactScope gets a reference to the given string and assigns it to the CustomerImpactScope field.
func (o *IncidentUpdateAttributes) SetCustomerImpactScope(v string) {
	o.CustomerImpactScope = &v
}

// GetCustomerImpactStart returns the CustomerImpactStart field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentUpdateAttributes) GetCustomerImpactStart() time.Time {
	if o == nil || o.CustomerImpactStart.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CustomerImpactStart.Get()
}

// GetCustomerImpactStartOk returns a tuple with the CustomerImpactStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncidentUpdateAttributes) GetCustomerImpactStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerImpactStart.Get(), o.CustomerImpactStart.IsSet()
}

// HasCustomerImpactStart returns a boolean if a field has been set.
func (o *IncidentUpdateAttributes) HasCustomerImpactStart() bool {
	if o != nil && o.CustomerImpactStart.IsSet() {
		return true
	}

	return false
}

// SetCustomerImpactStart gets a reference to the given NullableTime and assigns it to the CustomerImpactStart field.
func (o *IncidentUpdateAttributes) SetCustomerImpactStart(v time.Time) {
	o.CustomerImpactStart.Set(&v)
}

// SetCustomerImpactStartNil sets the value for CustomerImpactStart to be an explicit nil
func (o *IncidentUpdateAttributes) SetCustomerImpactStartNil() {
	o.CustomerImpactStart.Set(nil)
}

// UnsetCustomerImpactStart ensures that no value is present for CustomerImpactStart, not even an explicit nil
func (o *IncidentUpdateAttributes) UnsetCustomerImpactStart() {
	o.CustomerImpactStart.Unset()
}

// GetCustomerImpacted returns the CustomerImpacted field value if set, zero value otherwise.
func (o *IncidentUpdateAttributes) GetCustomerImpacted() bool {
	if o == nil || o.CustomerImpacted == nil {
		var ret bool
		return ret
	}
	return *o.CustomerImpacted
}

// GetCustomerImpactedOk returns a tuple with the CustomerImpacted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentUpdateAttributes) GetCustomerImpactedOk() (*bool, bool) {
	if o == nil || o.CustomerImpacted == nil {
		return nil, false
	}
	return o.CustomerImpacted, true
}

// HasCustomerImpacted returns a boolean if a field has been set.
func (o *IncidentUpdateAttributes) HasCustomerImpacted() bool {
	if o != nil && o.CustomerImpacted != nil {
		return true
	}

	return false
}

// SetCustomerImpacted gets a reference to the given bool and assigns it to the CustomerImpacted field.
func (o *IncidentUpdateAttributes) SetCustomerImpacted(v bool) {
	o.CustomerImpacted = &v
}

// GetDetected returns the Detected field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentUpdateAttributes) GetDetected() time.Time {
	if o == nil || o.Detected.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Detected.Get()
}

// GetDetectedOk returns a tuple with the Detected field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncidentUpdateAttributes) GetDetectedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Detected.Get(), o.Detected.IsSet()
}

// HasDetected returns a boolean if a field has been set.
func (o *IncidentUpdateAttributes) HasDetected() bool {
	if o != nil && o.Detected.IsSet() {
		return true
	}

	return false
}

// SetDetected gets a reference to the given NullableTime and assigns it to the Detected field.
func (o *IncidentUpdateAttributes) SetDetected(v time.Time) {
	o.Detected.Set(&v)
}

// SetDetectedNil sets the value for Detected to be an explicit nil
func (o *IncidentUpdateAttributes) SetDetectedNil() {
	o.Detected.Set(nil)
}

// UnsetDetected ensures that no value is present for Detected, not even an explicit nil
func (o *IncidentUpdateAttributes) UnsetDetected() {
	o.Detected.Unset()
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *IncidentUpdateAttributes) GetFields() map[string]IncidentFieldAttributes {
	if o == nil || o.Fields == nil {
		var ret map[string]IncidentFieldAttributes
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentUpdateAttributes) GetFieldsOk() (*map[string]IncidentFieldAttributes, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return &o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *IncidentUpdateAttributes) HasFields() bool {
	if o != nil && o.Fields != nil {
		return true
	}

	return false
}

// SetFields gets a reference to the given map[string]IncidentFieldAttributes and assigns it to the Fields field.
func (o *IncidentUpdateAttributes) SetFields(v map[string]IncidentFieldAttributes) {
	o.Fields = v
}

// GetNotificationHandles returns the NotificationHandles field value if set, zero value otherwise.
func (o *IncidentUpdateAttributes) GetNotificationHandles() []IncidentNotificationHandle {
	if o == nil || o.NotificationHandles == nil {
		var ret []IncidentNotificationHandle
		return ret
	}
	return o.NotificationHandles
}

// GetNotificationHandlesOk returns a tuple with the NotificationHandles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentUpdateAttributes) GetNotificationHandlesOk() (*[]IncidentNotificationHandle, bool) {
	if o == nil || o.NotificationHandles == nil {
		return nil, false
	}
	return &o.NotificationHandles, true
}

// HasNotificationHandles returns a boolean if a field has been set.
func (o *IncidentUpdateAttributes) HasNotificationHandles() bool {
	if o != nil && o.NotificationHandles != nil {
		return true
	}

	return false
}

// SetNotificationHandles gets a reference to the given []IncidentNotificationHandle and assigns it to the NotificationHandles field.
func (o *IncidentUpdateAttributes) SetNotificationHandles(v []IncidentNotificationHandle) {
	o.NotificationHandles = v
}

// GetResolved returns the Resolved field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentUpdateAttributes) GetResolved() time.Time {
	if o == nil || o.Resolved.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Resolved.Get()
}

// GetResolvedOk returns a tuple with the Resolved field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncidentUpdateAttributes) GetResolvedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resolved.Get(), o.Resolved.IsSet()
}

// HasResolved returns a boolean if a field has been set.
func (o *IncidentUpdateAttributes) HasResolved() bool {
	if o != nil && o.Resolved.IsSet() {
		return true
	}

	return false
}

// SetResolved gets a reference to the given NullableTime and assigns it to the Resolved field.
func (o *IncidentUpdateAttributes) SetResolved(v time.Time) {
	o.Resolved.Set(&v)
}

// SetResolvedNil sets the value for Resolved to be an explicit nil
func (o *IncidentUpdateAttributes) SetResolvedNil() {
	o.Resolved.Set(nil)
}

// UnsetResolved ensures that no value is present for Resolved, not even an explicit nil
func (o *IncidentUpdateAttributes) UnsetResolved() {
	o.Resolved.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *IncidentUpdateAttributes) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentUpdateAttributes) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *IncidentUpdateAttributes) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *IncidentUpdateAttributes) SetTitle(v string) {
	o.Title = &v
}

func (o IncidentUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.CustomerImpactEnd.IsSet() {
		toSerialize["customer_impact_end"] = o.CustomerImpactEnd.Get()
	}
	if o.CustomerImpactScope != nil {
		toSerialize["customer_impact_scope"] = o.CustomerImpactScope
	}
	if o.CustomerImpactStart.IsSet() {
		toSerialize["customer_impact_start"] = o.CustomerImpactStart.Get()
	}
	if o.CustomerImpacted != nil {
		toSerialize["customer_impacted"] = o.CustomerImpacted
	}
	if o.Detected.IsSet() {
		toSerialize["detected"] = o.Detected.Get()
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if o.NotificationHandles != nil {
		toSerialize["notification_handles"] = o.NotificationHandles
	}
	if o.Resolved.IsSet() {
		toSerialize["resolved"] = o.Resolved.Get()
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

func (o *IncidentUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		CustomerImpactEnd   NullableTime                       `json:"customer_impact_end,omitempty"`
		CustomerImpactScope *string                            `json:"customer_impact_scope,omitempty"`
		CustomerImpactStart NullableTime                       `json:"customer_impact_start,omitempty"`
		CustomerImpacted    *bool                              `json:"customer_impacted,omitempty"`
		Detected            NullableTime                       `json:"detected,omitempty"`
		Fields              map[string]IncidentFieldAttributes `json:"fields,omitempty"`
		NotificationHandles []IncidentNotificationHandle       `json:"notification_handles,omitempty"`
		Resolved            NullableTime                       `json:"resolved,omitempty"`
		Title               *string                            `json:"title,omitempty"`
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
	o.CustomerImpactEnd = all.CustomerImpactEnd
	o.CustomerImpactScope = all.CustomerImpactScope
	o.CustomerImpactStart = all.CustomerImpactStart
	o.CustomerImpacted = all.CustomerImpacted
	o.Detected = all.Detected
	o.Fields = all.Fields
	o.NotificationHandles = all.NotificationHandles
	o.Resolved = all.Resolved
	o.Title = all.Title
	return nil
}
