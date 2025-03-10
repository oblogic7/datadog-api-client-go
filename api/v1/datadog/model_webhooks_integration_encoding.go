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

// WebhooksIntegrationEncoding Encoding type. Can be given either `json` or `form`.
type WebhooksIntegrationEncoding string

// List of WebhooksIntegrationEncoding
const (
	WEBHOOKSINTEGRATIONENCODING_JSON WebhooksIntegrationEncoding = "json"
	WEBHOOKSINTEGRATIONENCODING_FORM WebhooksIntegrationEncoding = "form"
)

var allowedWebhooksIntegrationEncodingEnumValues = []WebhooksIntegrationEncoding{
	WEBHOOKSINTEGRATIONENCODING_JSON,
	WEBHOOKSINTEGRATIONENCODING_FORM,
}

func (w *WebhooksIntegrationEncoding) GetAllowedValues() []WebhooksIntegrationEncoding {
	return allowedWebhooksIntegrationEncodingEnumValues
}

func (v *WebhooksIntegrationEncoding) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WebhooksIntegrationEncoding(value)
	return nil
}

// NewWebhooksIntegrationEncodingFromValue returns a pointer to a valid WebhooksIntegrationEncoding
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWebhooksIntegrationEncodingFromValue(v string) (*WebhooksIntegrationEncoding, error) {
	ev := WebhooksIntegrationEncoding(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WebhooksIntegrationEncoding: valid values are %v", v, allowedWebhooksIntegrationEncodingEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WebhooksIntegrationEncoding) IsValid() bool {
	for _, existing := range allowedWebhooksIntegrationEncodingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WebhooksIntegrationEncoding value
func (v WebhooksIntegrationEncoding) Ptr() *WebhooksIntegrationEncoding {
	return &v
}

type NullableWebhooksIntegrationEncoding struct {
	value *WebhooksIntegrationEncoding
	isSet bool
}

func (v NullableWebhooksIntegrationEncoding) Get() *WebhooksIntegrationEncoding {
	return v.value
}

func (v *NullableWebhooksIntegrationEncoding) Set(val *WebhooksIntegrationEncoding) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhooksIntegrationEncoding) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhooksIntegrationEncoding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhooksIntegrationEncoding(val *WebhooksIntegrationEncoding) *NullableWebhooksIntegrationEncoding {
	return &NullableWebhooksIntegrationEncoding{value: val, isSet: true}
}

func (v NullableWebhooksIntegrationEncoding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhooksIntegrationEncoding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
