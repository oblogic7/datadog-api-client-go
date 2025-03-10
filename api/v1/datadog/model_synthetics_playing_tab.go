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

// SyntheticsPlayingTab Navigate between different tabs for your browser test.
type SyntheticsPlayingTab int64

// List of SyntheticsPlayingTab
const (
	SYNTHETICSPLAYINGTAB_MAIN_TAB SyntheticsPlayingTab = -1
	SYNTHETICSPLAYINGTAB_NEW_TAB  SyntheticsPlayingTab = 0
	SYNTHETICSPLAYINGTAB_TAB_1    SyntheticsPlayingTab = 1
	SYNTHETICSPLAYINGTAB_TAB_2    SyntheticsPlayingTab = 2
	SYNTHETICSPLAYINGTAB_TAB_3    SyntheticsPlayingTab = 3
)

var allowedSyntheticsPlayingTabEnumValues = []SyntheticsPlayingTab{
	SYNTHETICSPLAYINGTAB_MAIN_TAB,
	SYNTHETICSPLAYINGTAB_NEW_TAB,
	SYNTHETICSPLAYINGTAB_TAB_1,
	SYNTHETICSPLAYINGTAB_TAB_2,
	SYNTHETICSPLAYINGTAB_TAB_3,
}

func (w *SyntheticsPlayingTab) GetAllowedValues() []SyntheticsPlayingTab {
	return allowedSyntheticsPlayingTabEnumValues
}

func (v *SyntheticsPlayingTab) UnmarshalJSON(src []byte) error {
	var value int64
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsPlayingTab(value)
	return nil
}

// NewSyntheticsPlayingTabFromValue returns a pointer to a valid SyntheticsPlayingTab
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSyntheticsPlayingTabFromValue(v int64) (*SyntheticsPlayingTab, error) {
	ev := SyntheticsPlayingTab(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SyntheticsPlayingTab: valid values are %v", v, allowedSyntheticsPlayingTabEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SyntheticsPlayingTab) IsValid() bool {
	for _, existing := range allowedSyntheticsPlayingTabEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsPlayingTab value
func (v SyntheticsPlayingTab) Ptr() *SyntheticsPlayingTab {
	return &v
}

type NullableSyntheticsPlayingTab struct {
	value *SyntheticsPlayingTab
	isSet bool
}

func (v NullableSyntheticsPlayingTab) Get() *SyntheticsPlayingTab {
	return v.value
}

func (v *NullableSyntheticsPlayingTab) Set(val *SyntheticsPlayingTab) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsPlayingTab) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsPlayingTab) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsPlayingTab(val *SyntheticsPlayingTab) *NullableSyntheticsPlayingTab {
	return &NullableSyntheticsPlayingTab{value: val, isSet: true}
}

func (v NullableSyntheticsPlayingTab) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsPlayingTab) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
