/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// IPPrefixesWebhooks Available prefix information for the Webhook endpoints.
type IPPrefixesWebhooks struct {
	// List of IPv4 prefixes.
	PrefixesIpv4 []string `json:"prefixes_ipv4,omitempty"`
	// List of IPv6 prefixes.
	PrefixesIpv6 []string `json:"prefixes_ipv6,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewIPPrefixesWebhooks instantiates a new IPPrefixesWebhooks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIPPrefixesWebhooks() *IPPrefixesWebhooks {
	this := IPPrefixesWebhooks{}
	return &this
}

// NewIPPrefixesWebhooksWithDefaults instantiates a new IPPrefixesWebhooks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIPPrefixesWebhooksWithDefaults() *IPPrefixesWebhooks {
	this := IPPrefixesWebhooks{}
	return &this
}

// GetPrefixesIpv4 returns the PrefixesIpv4 field value if set, zero value otherwise.
func (o *IPPrefixesWebhooks) GetPrefixesIpv4() []string {
	if o == nil || o.PrefixesIpv4 == nil {
		var ret []string
		return ret
	}
	return o.PrefixesIpv4
}

// GetPrefixesIpv4Ok returns a tuple with the PrefixesIpv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPPrefixesWebhooks) GetPrefixesIpv4Ok() (*[]string, bool) {
	if o == nil || o.PrefixesIpv4 == nil {
		return nil, false
	}
	return &o.PrefixesIpv4, true
}

// HasPrefixesIpv4 returns a boolean if a field has been set.
func (o *IPPrefixesWebhooks) HasPrefixesIpv4() bool {
	if o != nil && o.PrefixesIpv4 != nil {
		return true
	}

	return false
}

// SetPrefixesIpv4 gets a reference to the given []string and assigns it to the PrefixesIpv4 field.
func (o *IPPrefixesWebhooks) SetPrefixesIpv4(v []string) {
	o.PrefixesIpv4 = v
}

// GetPrefixesIpv6 returns the PrefixesIpv6 field value if set, zero value otherwise.
func (o *IPPrefixesWebhooks) GetPrefixesIpv6() []string {
	if o == nil || o.PrefixesIpv6 == nil {
		var ret []string
		return ret
	}
	return o.PrefixesIpv6
}

// GetPrefixesIpv6Ok returns a tuple with the PrefixesIpv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPPrefixesWebhooks) GetPrefixesIpv6Ok() (*[]string, bool) {
	if o == nil || o.PrefixesIpv6 == nil {
		return nil, false
	}
	return &o.PrefixesIpv6, true
}

// HasPrefixesIpv6 returns a boolean if a field has been set.
func (o *IPPrefixesWebhooks) HasPrefixesIpv6() bool {
	if o != nil && o.PrefixesIpv6 != nil {
		return true
	}

	return false
}

// SetPrefixesIpv6 gets a reference to the given []string and assigns it to the PrefixesIpv6 field.
func (o *IPPrefixesWebhooks) SetPrefixesIpv6(v []string) {
	o.PrefixesIpv6 = v
}

func (o IPPrefixesWebhooks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.PrefixesIpv4 != nil {
		toSerialize["prefixes_ipv4"] = o.PrefixesIpv4
	}
	if o.PrefixesIpv6 != nil {
		toSerialize["prefixes_ipv6"] = o.PrefixesIpv6
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

func (o *IPPrefixesWebhooks) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		PrefixesIpv4 []string `json:"prefixes_ipv4,omitempty"`
		PrefixesIpv6 []string `json:"prefixes_ipv6,omitempty"`
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
	o.PrefixesIpv4 = all.PrefixesIpv4
	o.PrefixesIpv6 = all.PrefixesIpv6
	return nil
}
