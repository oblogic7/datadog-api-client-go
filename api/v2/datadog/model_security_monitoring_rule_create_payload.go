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

// SecurityMonitoringRuleCreatePayload Create a new rule.
type SecurityMonitoringRuleCreatePayload struct {
	// Cases for generating signals.
	Cases []SecurityMonitoringRuleCaseCreate `json:"cases"`
	// Additional queries to filter matched events before they are processed.
	Filters []SecurityMonitoringFilter `json:"filters,omitempty"`
	// Whether the notifications include the triggering group-by values in their title.
	HasExtendedTitle *bool `json:"hasExtendedTitle,omitempty"`
	// Whether the rule is enabled.
	IsEnabled bool `json:"isEnabled"`
	// Message for generated signals.
	Message string `json:"message"`
	// The name of the rule.
	Name string `json:"name"`
	// Options on rules.
	Options SecurityMonitoringRuleOptions `json:"options"`
	// Queries for selecting logs which are part of the rule.
	Queries []SecurityMonitoringRuleQueryCreate `json:"queries"`
	// Tags for generated signals.
	Tags []string `json:"tags,omitempty"`
	// The rule type.
	Type *SecurityMonitoringRuleTypeCreate `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewSecurityMonitoringRuleCreatePayload instantiates a new SecurityMonitoringRuleCreatePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityMonitoringRuleCreatePayload(cases []SecurityMonitoringRuleCaseCreate, isEnabled bool, message string, name string, options SecurityMonitoringRuleOptions, queries []SecurityMonitoringRuleQueryCreate) *SecurityMonitoringRuleCreatePayload {
	this := SecurityMonitoringRuleCreatePayload{}
	this.Cases = cases
	this.IsEnabled = isEnabled
	this.Message = message
	this.Name = name
	this.Options = options
	this.Queries = queries
	return &this
}

// NewSecurityMonitoringRuleCreatePayloadWithDefaults instantiates a new SecurityMonitoringRuleCreatePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityMonitoringRuleCreatePayloadWithDefaults() *SecurityMonitoringRuleCreatePayload {
	this := SecurityMonitoringRuleCreatePayload{}
	return &this
}

// GetCases returns the Cases field value
func (o *SecurityMonitoringRuleCreatePayload) GetCases() []SecurityMonitoringRuleCaseCreate {
	if o == nil {
		var ret []SecurityMonitoringRuleCaseCreate
		return ret
	}
	return o.Cases
}

// GetCasesOk returns a tuple with the Cases field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleCreatePayload) GetCasesOk() (*[]SecurityMonitoringRuleCaseCreate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cases, true
}

// SetCases sets field value
func (o *SecurityMonitoringRuleCreatePayload) SetCases(v []SecurityMonitoringRuleCaseCreate) {
	o.Cases = v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleCreatePayload) GetFilters() []SecurityMonitoringFilter {
	if o == nil || o.Filters == nil {
		var ret []SecurityMonitoringFilter
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleCreatePayload) GetFiltersOk() (*[]SecurityMonitoringFilter, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return &o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleCreatePayload) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []SecurityMonitoringFilter and assigns it to the Filters field.
func (o *SecurityMonitoringRuleCreatePayload) SetFilters(v []SecurityMonitoringFilter) {
	o.Filters = v
}

// GetHasExtendedTitle returns the HasExtendedTitle field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleCreatePayload) GetHasExtendedTitle() bool {
	if o == nil || o.HasExtendedTitle == nil {
		var ret bool
		return ret
	}
	return *o.HasExtendedTitle
}

// GetHasExtendedTitleOk returns a tuple with the HasExtendedTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleCreatePayload) GetHasExtendedTitleOk() (*bool, bool) {
	if o == nil || o.HasExtendedTitle == nil {
		return nil, false
	}
	return o.HasExtendedTitle, true
}

// HasHasExtendedTitle returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleCreatePayload) HasHasExtendedTitle() bool {
	if o != nil && o.HasExtendedTitle != nil {
		return true
	}

	return false
}

// SetHasExtendedTitle gets a reference to the given bool and assigns it to the HasExtendedTitle field.
func (o *SecurityMonitoringRuleCreatePayload) SetHasExtendedTitle(v bool) {
	o.HasExtendedTitle = &v
}

// GetIsEnabled returns the IsEnabled field value
func (o *SecurityMonitoringRuleCreatePayload) GetIsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleCreatePayload) GetIsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value
func (o *SecurityMonitoringRuleCreatePayload) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

// GetMessage returns the Message field value
func (o *SecurityMonitoringRuleCreatePayload) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleCreatePayload) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SecurityMonitoringRuleCreatePayload) SetMessage(v string) {
	o.Message = v
}

// GetName returns the Name field value
func (o *SecurityMonitoringRuleCreatePayload) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleCreatePayload) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SecurityMonitoringRuleCreatePayload) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value
func (o *SecurityMonitoringRuleCreatePayload) GetOptions() SecurityMonitoringRuleOptions {
	if o == nil {
		var ret SecurityMonitoringRuleOptions
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleCreatePayload) GetOptionsOk() (*SecurityMonitoringRuleOptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value
func (o *SecurityMonitoringRuleCreatePayload) SetOptions(v SecurityMonitoringRuleOptions) {
	o.Options = v
}

// GetQueries returns the Queries field value
func (o *SecurityMonitoringRuleCreatePayload) GetQueries() []SecurityMonitoringRuleQueryCreate {
	if o == nil {
		var ret []SecurityMonitoringRuleQueryCreate
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleCreatePayload) GetQueriesOk() (*[]SecurityMonitoringRuleQueryCreate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Queries, true
}

// SetQueries sets field value
func (o *SecurityMonitoringRuleCreatePayload) SetQueries(v []SecurityMonitoringRuleQueryCreate) {
	o.Queries = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleCreatePayload) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleCreatePayload) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleCreatePayload) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *SecurityMonitoringRuleCreatePayload) SetTags(v []string) {
	o.Tags = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleCreatePayload) GetType() SecurityMonitoringRuleTypeCreate {
	if o == nil || o.Type == nil {
		var ret SecurityMonitoringRuleTypeCreate
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleCreatePayload) GetTypeOk() (*SecurityMonitoringRuleTypeCreate, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleCreatePayload) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given SecurityMonitoringRuleTypeCreate and assigns it to the Type field.
func (o *SecurityMonitoringRuleCreatePayload) SetType(v SecurityMonitoringRuleTypeCreate) {
	o.Type = &v
}

func (o SecurityMonitoringRuleCreatePayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["cases"] = o.Cases
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}
	if o.HasExtendedTitle != nil {
		toSerialize["hasExtendedTitle"] = o.HasExtendedTitle
	}
	toSerialize["isEnabled"] = o.IsEnabled
	toSerialize["message"] = o.Message
	toSerialize["name"] = o.Name
	toSerialize["options"] = o.Options
	toSerialize["queries"] = o.Queries
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

func (o *SecurityMonitoringRuleCreatePayload) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Cases     *[]SecurityMonitoringRuleCaseCreate  `json:"cases"`
		IsEnabled *bool                                `json:"isEnabled"`
		Message   *string                              `json:"message"`
		Name      *string                              `json:"name"`
		Options   *SecurityMonitoringRuleOptions       `json:"options"`
		Queries   *[]SecurityMonitoringRuleQueryCreate `json:"queries"`
	}{}
	all := struct {
		Cases            []SecurityMonitoringRuleCaseCreate  `json:"cases"`
		Filters          []SecurityMonitoringFilter          `json:"filters,omitempty"`
		HasExtendedTitle *bool                               `json:"hasExtendedTitle,omitempty"`
		IsEnabled        bool                                `json:"isEnabled"`
		Message          string                              `json:"message"`
		Name             string                              `json:"name"`
		Options          SecurityMonitoringRuleOptions       `json:"options"`
		Queries          []SecurityMonitoringRuleQueryCreate `json:"queries"`
		Tags             []string                            `json:"tags,omitempty"`
		Type             *SecurityMonitoringRuleTypeCreate   `json:"type,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Cases == nil {
		return fmt.Errorf("Required field cases missing")
	}
	if required.IsEnabled == nil {
		return fmt.Errorf("Required field isEnabled missing")
	}
	if required.Message == nil {
		return fmt.Errorf("Required field message missing")
	}
	if required.Name == nil {
		return fmt.Errorf("Required field name missing")
	}
	if required.Options == nil {
		return fmt.Errorf("Required field options missing")
	}
	if required.Queries == nil {
		return fmt.Errorf("Required field queries missing")
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
	if v := all.Type; v != nil && !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.Cases = all.Cases
	o.Filters = all.Filters
	o.HasExtendedTitle = all.HasExtendedTitle
	o.IsEnabled = all.IsEnabled
	o.Message = all.Message
	o.Name = all.Name
	if all.Options.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Options = all.Options
	o.Queries = all.Queries
	o.Tags = all.Tags
	o.Type = all.Type
	return nil
}
