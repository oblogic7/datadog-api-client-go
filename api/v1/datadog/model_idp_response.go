/*
 * (C) Datadog, Inc. 2019
 * All rights reserved
 * Licensed under a 3-clause BSD style license (see LICENSE)
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"bytes"
	"encoding/json"
)

// IdpResponse struct for IdpResponse
type IdpResponse struct {
	Message string `json:"message"`
}

// GetMessage returns the Message field value
func (o *IdpResponse) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// SetMessage sets field value
func (o *IdpResponse) SetMessage(v string) {
	o.Message = v
}

type NullableIdpResponse struct {
	Value        IdpResponse
	ExplicitNull bool
}

func (v NullableIdpResponse) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableIdpResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
