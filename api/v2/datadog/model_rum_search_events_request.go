/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// RUMSearchEventsRequest The request for a RUM events list.
type RUMSearchEventsRequest struct {
	// The search and filter query settings.
	Filter *RUMQueryFilter `json:"filter,omitempty"`
	// Global query options that are used during the query.
	// Note: Only supply timezone or time offset, not both. Otherwise, the query fails.
	Options *RUMQueryOptions `json:"options,omitempty"`
	// Paging attributes for listing events.
	Page *RUMQueryPageOptions `json:"page,omitempty"`
	// Sort parameters when querying events.
	Sort *RUMSort `json:"sort,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewRUMSearchEventsRequest instantiates a new RUMSearchEventsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRUMSearchEventsRequest() *RUMSearchEventsRequest {
	this := RUMSearchEventsRequest{}
	return &this
}

// NewRUMSearchEventsRequestWithDefaults instantiates a new RUMSearchEventsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRUMSearchEventsRequestWithDefaults() *RUMSearchEventsRequest {
	this := RUMSearchEventsRequest{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *RUMSearchEventsRequest) GetFilter() RUMQueryFilter {
	if o == nil || o.Filter == nil {
		var ret RUMQueryFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMSearchEventsRequest) GetFilterOk() (*RUMQueryFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *RUMSearchEventsRequest) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given RUMQueryFilter and assigns it to the Filter field.
func (o *RUMSearchEventsRequest) SetFilter(v RUMQueryFilter) {
	o.Filter = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *RUMSearchEventsRequest) GetOptions() RUMQueryOptions {
	if o == nil || o.Options == nil {
		var ret RUMQueryOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMSearchEventsRequest) GetOptionsOk() (*RUMQueryOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *RUMSearchEventsRequest) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given RUMQueryOptions and assigns it to the Options field.
func (o *RUMSearchEventsRequest) SetOptions(v RUMQueryOptions) {
	o.Options = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *RUMSearchEventsRequest) GetPage() RUMQueryPageOptions {
	if o == nil || o.Page == nil {
		var ret RUMQueryPageOptions
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMSearchEventsRequest) GetPageOk() (*RUMQueryPageOptions, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *RUMSearchEventsRequest) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given RUMQueryPageOptions and assigns it to the Page field.
func (o *RUMSearchEventsRequest) SetPage(v RUMQueryPageOptions) {
	o.Page = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *RUMSearchEventsRequest) GetSort() RUMSort {
	if o == nil || o.Sort == nil {
		var ret RUMSort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMSearchEventsRequest) GetSortOk() (*RUMSort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *RUMSearchEventsRequest) HasSort() bool {
	if o != nil && o.Sort != nil {
		return true
	}

	return false
}

// SetSort gets a reference to the given RUMSort and assigns it to the Sort field.
func (o *RUMSearchEventsRequest) SetSort(v RUMSort) {
	o.Sort = &v
}

func (o RUMSearchEventsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

func (o *RUMSearchEventsRequest) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Filter  *RUMQueryFilter      `json:"filter,omitempty"`
		Options *RUMQueryOptions     `json:"options,omitempty"`
		Page    *RUMQueryPageOptions `json:"page,omitempty"`
		Sort    *RUMSort             `json:"sort,omitempty"`
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
	if v := all.Sort; v != nil && !v.IsValid() {
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
	if all.Options != nil && all.Options.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Options = all.Options
	if all.Page != nil && all.Page.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Page = all.Page
	o.Sort = all.Sort
	return nil
}
