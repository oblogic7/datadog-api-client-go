/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// SecurityMonitoringRuleResponse Rule.
type SecurityMonitoringRuleResponse struct {
	// Cases for generating signals.
	Cases []SecurityMonitoringRuleCase `json:"cases,omitempty"`
	// When the rule was created, timestamp in milliseconds.
	CreatedAt *int64 `json:"createdAt,omitempty"`
	// User ID of the user who created the rule.
	CreationAuthorId *int64 `json:"creationAuthorId,omitempty"`
	// Additional queries to filter matched events before they are processed.
	Filters []SecurityMonitoringFilter `json:"filters,omitempty"`
	// Whether the notifications include the triggering group-by values in their title.
	HasExtendedTitle *bool `json:"hasExtendedTitle,omitempty"`
	// The ID of the rule.
	Id *string `json:"id,omitempty"`
	// Whether the rule is included by default.
	IsDefault *bool `json:"isDefault,omitempty"`
	// Whether the rule has been deleted.
	IsDeleted *bool `json:"isDeleted,omitempty"`
	// Whether the rule is enabled.
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// Message for generated signals.
	Message *string `json:"message,omitempty"`
	// The name of the rule.
	Name *string `json:"name,omitempty"`
	// Options on rules.
	Options *SecurityMonitoringRuleOptions `json:"options,omitempty"`
	// Queries for selecting logs which are part of the rule.
	Queries []SecurityMonitoringRuleQuery `json:"queries,omitempty"`
	// Tags for generated signals.
	Tags []string `json:"tags,omitempty"`
	// The rule type.
	Type *SecurityMonitoringRuleTypeRead `json:"type,omitempty"`
	// User ID of the user who updated the rule.
	UpdateAuthorId *int64 `json:"updateAuthorId,omitempty"`
	// The version of the rule.
	Version *int64 `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewSecurityMonitoringRuleResponse instantiates a new SecurityMonitoringRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityMonitoringRuleResponse() *SecurityMonitoringRuleResponse {
	this := SecurityMonitoringRuleResponse{}
	return &this
}

// NewSecurityMonitoringRuleResponseWithDefaults instantiates a new SecurityMonitoringRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityMonitoringRuleResponseWithDefaults() *SecurityMonitoringRuleResponse {
	this := SecurityMonitoringRuleResponse{}
	return &this
}

// GetCases returns the Cases field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleResponse) GetCases() []SecurityMonitoringRuleCase {
	if o == nil || o.Cases == nil {
		var ret []SecurityMonitoringRuleCase
		return ret
	}
	return o.Cases
}

// GetCasesOk returns a tuple with the Cases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleResponse) GetCasesOk() (*[]SecurityMonitoringRuleCase, bool) {
	if o == nil || o.Cases == nil {
		return nil, false
	}
	return &o.Cases, true
}

// HasCases returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleResponse) HasCases() bool {
	if o != nil && o.Cases != nil {
		return true
	}

	return false
}

// SetCases gets a reference to the given []SecurityMonitoringRuleCase and assigns it to the Cases field.
func (o *SecurityMonitoringRuleResponse) SetCases(v []SecurityMonitoringRuleCase) {
	o.Cases = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleResponse) GetCreatedAt() int64 {
	if o == nil || o.CreatedAt == nil {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleResponse) GetCreatedAtOk() (*int64, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *SecurityMonitoringRuleResponse) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetCreationAuthorId returns the CreationAuthorId field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleResponse) GetCreationAuthorId() int64 {
	if o == nil || o.CreationAuthorId == nil {
		var ret int64
		return ret
	}
	return *o.CreationAuthorId
}

// GetCreationAuthorIdOk returns a tuple with the CreationAuthorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleResponse) GetCreationAuthorIdOk() (*int64, bool) {
	if o == nil || o.CreationAuthorId == nil {
		return nil, false
	}
	return o.CreationAuthorId, true
}

// HasCreationAuthorId returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleResponse) HasCreationAuthorId() bool {
	if o != nil && o.CreationAuthorId != nil {
		return true
	}

	return false
}

// SetCreationAuthorId gets a reference to the given int64 and assigns it to the CreationAuthorId field.
func (o *SecurityMonitoringRuleResponse) SetCreationAuthorId(v int64) {
	o.CreationAuthorId = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleResponse) GetFilters() []SecurityMonitoringFilter {
	if o == nil || o.Filters == nil {
		var ret []SecurityMonitoringFilter
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleResponse) GetFiltersOk() (*[]SecurityMonitoringFilter, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return &o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleResponse) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []SecurityMonitoringFilter and assigns it to the Filters field.
func (o *SecurityMonitoringRuleResponse) SetFilters(v []SecurityMonitoringFilter) {
	o.Filters = v
}

// GetHasExtendedTitle returns the HasExtendedTitle field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleResponse) GetHasExtendedTitle() bool {
	if o == nil || o.HasExtendedTitle == nil {
		var ret bool
		return ret
	}
	return *o.HasExtendedTitle
}

// GetHasExtendedTitleOk returns a tuple with the HasExtendedTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleResponse) GetHasExtendedTitleOk() (*bool, bool) {
	if o == nil || o.HasExtendedTitle == nil {
		return nil, false
	}
	return o.HasExtendedTitle, true
}

// HasHasExtendedTitle returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleResponse) HasHasExtendedTitle() bool {
	if o != nil && o.HasExtendedTitle != nil {
		return true
	}

	return false
}

// SetHasExtendedTitle gets a reference to the given bool and assigns it to the HasExtendedTitle field.
func (o *SecurityMonitoringRuleResponse) SetHasExtendedTitle(v bool) {
	o.HasExtendedTitle = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SecurityMonitoringRuleResponse) SetId(v string) {
	o.Id = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleResponse) GetIsDefault() bool {
	if o == nil || o.IsDefault == nil {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleResponse) GetIsDefaultOk() (*bool, bool) {
	if o == nil || o.IsDefault == nil {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleResponse) HasIsDefault() bool {
	if o != nil && o.IsDefault != nil {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *SecurityMonitoringRuleResponse) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleResponse) GetIsDeleted() bool {
	if o == nil || o.IsDeleted == nil {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleResponse) GetIsDeletedOk() (*bool, bool) {
	if o == nil || o.IsDeleted == nil {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleResponse) HasIsDeleted() bool {
	if o != nil && o.IsDeleted != nil {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *SecurityMonitoringRuleResponse) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleResponse) GetIsEnabled() bool {
	if o == nil || o.IsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleResponse) GetIsEnabledOk() (*bool, bool) {
	if o == nil || o.IsEnabled == nil {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleResponse) HasIsEnabled() bool {
	if o != nil && o.IsEnabled != nil {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *SecurityMonitoringRuleResponse) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleResponse) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleResponse) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleResponse) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SecurityMonitoringRuleResponse) SetMessage(v string) {
	o.Message = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SecurityMonitoringRuleResponse) SetName(v string) {
	o.Name = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleResponse) GetOptions() SecurityMonitoringRuleOptions {
	if o == nil || o.Options == nil {
		var ret SecurityMonitoringRuleOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleResponse) GetOptionsOk() (*SecurityMonitoringRuleOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleResponse) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given SecurityMonitoringRuleOptions and assigns it to the Options field.
func (o *SecurityMonitoringRuleResponse) SetOptions(v SecurityMonitoringRuleOptions) {
	o.Options = &v
}

// GetQueries returns the Queries field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleResponse) GetQueries() []SecurityMonitoringRuleQuery {
	if o == nil || o.Queries == nil {
		var ret []SecurityMonitoringRuleQuery
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleResponse) GetQueriesOk() (*[]SecurityMonitoringRuleQuery, bool) {
	if o == nil || o.Queries == nil {
		return nil, false
	}
	return &o.Queries, true
}

// HasQueries returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleResponse) HasQueries() bool {
	if o != nil && o.Queries != nil {
		return true
	}

	return false
}

// SetQueries gets a reference to the given []SecurityMonitoringRuleQuery and assigns it to the Queries field.
func (o *SecurityMonitoringRuleResponse) SetQueries(v []SecurityMonitoringRuleQuery) {
	o.Queries = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleResponse) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleResponse) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleResponse) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *SecurityMonitoringRuleResponse) SetTags(v []string) {
	o.Tags = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleResponse) GetType() SecurityMonitoringRuleTypeRead {
	if o == nil || o.Type == nil {
		var ret SecurityMonitoringRuleTypeRead
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleResponse) GetTypeOk() (*SecurityMonitoringRuleTypeRead, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleResponse) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given SecurityMonitoringRuleTypeRead and assigns it to the Type field.
func (o *SecurityMonitoringRuleResponse) SetType(v SecurityMonitoringRuleTypeRead) {
	o.Type = &v
}

// GetUpdateAuthorId returns the UpdateAuthorId field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleResponse) GetUpdateAuthorId() int64 {
	if o == nil || o.UpdateAuthorId == nil {
		var ret int64
		return ret
	}
	return *o.UpdateAuthorId
}

// GetUpdateAuthorIdOk returns a tuple with the UpdateAuthorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleResponse) GetUpdateAuthorIdOk() (*int64, bool) {
	if o == nil || o.UpdateAuthorId == nil {
		return nil, false
	}
	return o.UpdateAuthorId, true
}

// HasUpdateAuthorId returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleResponse) HasUpdateAuthorId() bool {
	if o != nil && o.UpdateAuthorId != nil {
		return true
	}

	return false
}

// SetUpdateAuthorId gets a reference to the given int64 and assigns it to the UpdateAuthorId field.
func (o *SecurityMonitoringRuleResponse) SetUpdateAuthorId(v int64) {
	o.UpdateAuthorId = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleResponse) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleResponse) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleResponse) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *SecurityMonitoringRuleResponse) SetVersion(v int64) {
	o.Version = &v
}

func (o SecurityMonitoringRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Cases != nil {
		toSerialize["cases"] = o.Cases
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.CreationAuthorId != nil {
		toSerialize["creationAuthorId"] = o.CreationAuthorId
	}
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}
	if o.HasExtendedTitle != nil {
		toSerialize["hasExtendedTitle"] = o.HasExtendedTitle
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsDefault != nil {
		toSerialize["isDefault"] = o.IsDefault
	}
	if o.IsDeleted != nil {
		toSerialize["isDeleted"] = o.IsDeleted
	}
	if o.IsEnabled != nil {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.Queries != nil {
		toSerialize["queries"] = o.Queries
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.UpdateAuthorId != nil {
		toSerialize["updateAuthorId"] = o.UpdateAuthorId
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

func (o *SecurityMonitoringRuleResponse) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Cases            []SecurityMonitoringRuleCase    `json:"cases,omitempty"`
		CreatedAt        *int64                          `json:"createdAt,omitempty"`
		CreationAuthorId *int64                          `json:"creationAuthorId,omitempty"`
		Filters          []SecurityMonitoringFilter      `json:"filters,omitempty"`
		HasExtendedTitle *bool                           `json:"hasExtendedTitle,omitempty"`
		Id               *string                         `json:"id,omitempty"`
		IsDefault        *bool                           `json:"isDefault,omitempty"`
		IsDeleted        *bool                           `json:"isDeleted,omitempty"`
		IsEnabled        *bool                           `json:"isEnabled,omitempty"`
		Message          *string                         `json:"message,omitempty"`
		Name             *string                         `json:"name,omitempty"`
		Options          *SecurityMonitoringRuleOptions  `json:"options,omitempty"`
		Queries          []SecurityMonitoringRuleQuery   `json:"queries,omitempty"`
		Tags             []string                        `json:"tags,omitempty"`
		Type             *SecurityMonitoringRuleTypeRead `json:"type,omitempty"`
		UpdateAuthorId   *int64                          `json:"updateAuthorId,omitempty"`
		Version          *int64                          `json:"version,omitempty"`
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
	o.Cases = all.Cases
	o.CreatedAt = all.CreatedAt
	o.CreationAuthorId = all.CreationAuthorId
	o.Filters = all.Filters
	o.HasExtendedTitle = all.HasExtendedTitle
	o.Id = all.Id
	o.IsDefault = all.IsDefault
	o.IsDeleted = all.IsDeleted
	o.IsEnabled = all.IsEnabled
	o.Message = all.Message
	o.Name = all.Name
	if all.Options != nil && all.Options.UnparsedObject != nil && o.UnparsedObject == nil {
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
	o.UpdateAuthorId = all.UpdateAuthorId
	o.Version = all.Version
	return nil
}
