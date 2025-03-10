/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// FormulaAndFunctionQueryDefinition - A formula and function query.
type FormulaAndFunctionQueryDefinition struct {
	FormulaAndFunctionMetricQueryDefinition             *FormulaAndFunctionMetricQueryDefinition
	FormulaAndFunctionEventQueryDefinition              *FormulaAndFunctionEventQueryDefinition
	FormulaAndFunctionProcessQueryDefinition            *FormulaAndFunctionProcessQueryDefinition
	FormulaAndFunctionApmDependencyStatsQueryDefinition *FormulaAndFunctionApmDependencyStatsQueryDefinition
	FormulaAndFunctionApmResourceStatsQueryDefinition   *FormulaAndFunctionApmResourceStatsQueryDefinition

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// FormulaAndFunctionMetricQueryDefinitionAsFormulaAndFunctionQueryDefinition is a convenience function that returns FormulaAndFunctionMetricQueryDefinition wrapped in FormulaAndFunctionQueryDefinition
func FormulaAndFunctionMetricQueryDefinitionAsFormulaAndFunctionQueryDefinition(v *FormulaAndFunctionMetricQueryDefinition) FormulaAndFunctionQueryDefinition {
	return FormulaAndFunctionQueryDefinition{FormulaAndFunctionMetricQueryDefinition: v}
}

// FormulaAndFunctionEventQueryDefinitionAsFormulaAndFunctionQueryDefinition is a convenience function that returns FormulaAndFunctionEventQueryDefinition wrapped in FormulaAndFunctionQueryDefinition
func FormulaAndFunctionEventQueryDefinitionAsFormulaAndFunctionQueryDefinition(v *FormulaAndFunctionEventQueryDefinition) FormulaAndFunctionQueryDefinition {
	return FormulaAndFunctionQueryDefinition{FormulaAndFunctionEventQueryDefinition: v}
}

// FormulaAndFunctionProcessQueryDefinitionAsFormulaAndFunctionQueryDefinition is a convenience function that returns FormulaAndFunctionProcessQueryDefinition wrapped in FormulaAndFunctionQueryDefinition
func FormulaAndFunctionProcessQueryDefinitionAsFormulaAndFunctionQueryDefinition(v *FormulaAndFunctionProcessQueryDefinition) FormulaAndFunctionQueryDefinition {
	return FormulaAndFunctionQueryDefinition{FormulaAndFunctionProcessQueryDefinition: v}
}

// FormulaAndFunctionApmDependencyStatsQueryDefinitionAsFormulaAndFunctionQueryDefinition is a convenience function that returns FormulaAndFunctionApmDependencyStatsQueryDefinition wrapped in FormulaAndFunctionQueryDefinition
func FormulaAndFunctionApmDependencyStatsQueryDefinitionAsFormulaAndFunctionQueryDefinition(v *FormulaAndFunctionApmDependencyStatsQueryDefinition) FormulaAndFunctionQueryDefinition {
	return FormulaAndFunctionQueryDefinition{FormulaAndFunctionApmDependencyStatsQueryDefinition: v}
}

// FormulaAndFunctionApmResourceStatsQueryDefinitionAsFormulaAndFunctionQueryDefinition is a convenience function that returns FormulaAndFunctionApmResourceStatsQueryDefinition wrapped in FormulaAndFunctionQueryDefinition
func FormulaAndFunctionApmResourceStatsQueryDefinitionAsFormulaAndFunctionQueryDefinition(v *FormulaAndFunctionApmResourceStatsQueryDefinition) FormulaAndFunctionQueryDefinition {
	return FormulaAndFunctionQueryDefinition{FormulaAndFunctionApmResourceStatsQueryDefinition: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *FormulaAndFunctionQueryDefinition) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into FormulaAndFunctionMetricQueryDefinition
	err = json.Unmarshal(data, &dst.FormulaAndFunctionMetricQueryDefinition)
	if err == nil {
		if dst.FormulaAndFunctionMetricQueryDefinition != nil && dst.FormulaAndFunctionMetricQueryDefinition.UnparsedObject == nil {
			jsonFormulaAndFunctionMetricQueryDefinition, _ := json.Marshal(dst.FormulaAndFunctionMetricQueryDefinition)
			if string(jsonFormulaAndFunctionMetricQueryDefinition) == "{}" { // empty struct
				dst.FormulaAndFunctionMetricQueryDefinition = nil
			} else {
				match++
			}
		} else {
			dst.FormulaAndFunctionMetricQueryDefinition = nil
		}
	} else {
		dst.FormulaAndFunctionMetricQueryDefinition = nil
	}

	// try to unmarshal data into FormulaAndFunctionEventQueryDefinition
	err = json.Unmarshal(data, &dst.FormulaAndFunctionEventQueryDefinition)
	if err == nil {
		if dst.FormulaAndFunctionEventQueryDefinition != nil && dst.FormulaAndFunctionEventQueryDefinition.UnparsedObject == nil {
			jsonFormulaAndFunctionEventQueryDefinition, _ := json.Marshal(dst.FormulaAndFunctionEventQueryDefinition)
			if string(jsonFormulaAndFunctionEventQueryDefinition) == "{}" { // empty struct
				dst.FormulaAndFunctionEventQueryDefinition = nil
			} else {
				match++
			}
		} else {
			dst.FormulaAndFunctionEventQueryDefinition = nil
		}
	} else {
		dst.FormulaAndFunctionEventQueryDefinition = nil
	}

	// try to unmarshal data into FormulaAndFunctionProcessQueryDefinition
	err = json.Unmarshal(data, &dst.FormulaAndFunctionProcessQueryDefinition)
	if err == nil {
		if dst.FormulaAndFunctionProcessQueryDefinition != nil && dst.FormulaAndFunctionProcessQueryDefinition.UnparsedObject == nil {
			jsonFormulaAndFunctionProcessQueryDefinition, _ := json.Marshal(dst.FormulaAndFunctionProcessQueryDefinition)
			if string(jsonFormulaAndFunctionProcessQueryDefinition) == "{}" { // empty struct
				dst.FormulaAndFunctionProcessQueryDefinition = nil
			} else {
				match++
			}
		} else {
			dst.FormulaAndFunctionProcessQueryDefinition = nil
		}
	} else {
		dst.FormulaAndFunctionProcessQueryDefinition = nil
	}

	// try to unmarshal data into FormulaAndFunctionApmDependencyStatsQueryDefinition
	err = json.Unmarshal(data, &dst.FormulaAndFunctionApmDependencyStatsQueryDefinition)
	if err == nil {
		if dst.FormulaAndFunctionApmDependencyStatsQueryDefinition != nil && dst.FormulaAndFunctionApmDependencyStatsQueryDefinition.UnparsedObject == nil {
			jsonFormulaAndFunctionApmDependencyStatsQueryDefinition, _ := json.Marshal(dst.FormulaAndFunctionApmDependencyStatsQueryDefinition)
			if string(jsonFormulaAndFunctionApmDependencyStatsQueryDefinition) == "{}" { // empty struct
				dst.FormulaAndFunctionApmDependencyStatsQueryDefinition = nil
			} else {
				match++
			}
		} else {
			dst.FormulaAndFunctionApmDependencyStatsQueryDefinition = nil
		}
	} else {
		dst.FormulaAndFunctionApmDependencyStatsQueryDefinition = nil
	}

	// try to unmarshal data into FormulaAndFunctionApmResourceStatsQueryDefinition
	err = json.Unmarshal(data, &dst.FormulaAndFunctionApmResourceStatsQueryDefinition)
	if err == nil {
		if dst.FormulaAndFunctionApmResourceStatsQueryDefinition != nil && dst.FormulaAndFunctionApmResourceStatsQueryDefinition.UnparsedObject == nil {
			jsonFormulaAndFunctionApmResourceStatsQueryDefinition, _ := json.Marshal(dst.FormulaAndFunctionApmResourceStatsQueryDefinition)
			if string(jsonFormulaAndFunctionApmResourceStatsQueryDefinition) == "{}" { // empty struct
				dst.FormulaAndFunctionApmResourceStatsQueryDefinition = nil
			} else {
				match++
			}
		} else {
			dst.FormulaAndFunctionApmResourceStatsQueryDefinition = nil
		}
	} else {
		dst.FormulaAndFunctionApmResourceStatsQueryDefinition = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		dst.FormulaAndFunctionMetricQueryDefinition = nil
		dst.FormulaAndFunctionEventQueryDefinition = nil
		dst.FormulaAndFunctionProcessQueryDefinition = nil
		dst.FormulaAndFunctionApmDependencyStatsQueryDefinition = nil
		dst.FormulaAndFunctionApmResourceStatsQueryDefinition = nil
		return json.Unmarshal(data, &dst.UnparsedObject)
	} else {
		return nil // exactly one match
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src FormulaAndFunctionQueryDefinition) MarshalJSON() ([]byte, error) {
	if src.FormulaAndFunctionMetricQueryDefinition != nil {
		return json.Marshal(&src.FormulaAndFunctionMetricQueryDefinition)
	}

	if src.FormulaAndFunctionEventQueryDefinition != nil {
		return json.Marshal(&src.FormulaAndFunctionEventQueryDefinition)
	}

	if src.FormulaAndFunctionProcessQueryDefinition != nil {
		return json.Marshal(&src.FormulaAndFunctionProcessQueryDefinition)
	}

	if src.FormulaAndFunctionApmDependencyStatsQueryDefinition != nil {
		return json.Marshal(&src.FormulaAndFunctionApmDependencyStatsQueryDefinition)
	}

	if src.FormulaAndFunctionApmResourceStatsQueryDefinition != nil {
		return json.Marshal(&src.FormulaAndFunctionApmResourceStatsQueryDefinition)
	}

	if src.UnparsedObject != nil {
		return json.Marshal(src.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *FormulaAndFunctionQueryDefinition) GetActualInstance() interface{} {
	if obj.FormulaAndFunctionMetricQueryDefinition != nil {
		return obj.FormulaAndFunctionMetricQueryDefinition
	}

	if obj.FormulaAndFunctionEventQueryDefinition != nil {
		return obj.FormulaAndFunctionEventQueryDefinition
	}

	if obj.FormulaAndFunctionProcessQueryDefinition != nil {
		return obj.FormulaAndFunctionProcessQueryDefinition
	}

	if obj.FormulaAndFunctionApmDependencyStatsQueryDefinition != nil {
		return obj.FormulaAndFunctionApmDependencyStatsQueryDefinition
	}

	if obj.FormulaAndFunctionApmResourceStatsQueryDefinition != nil {
		return obj.FormulaAndFunctionApmResourceStatsQueryDefinition
	}

	// all schemas are nil
	return nil
}

type NullableFormulaAndFunctionQueryDefinition struct {
	value *FormulaAndFunctionQueryDefinition
	isSet bool
}

func (v NullableFormulaAndFunctionQueryDefinition) Get() *FormulaAndFunctionQueryDefinition {
	return v.value
}

func (v *NullableFormulaAndFunctionQueryDefinition) Set(val *FormulaAndFunctionQueryDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableFormulaAndFunctionQueryDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableFormulaAndFunctionQueryDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormulaAndFunctionQueryDefinition(val *FormulaAndFunctionQueryDefinition) *NullableFormulaAndFunctionQueryDefinition {
	return &NullableFormulaAndFunctionQueryDefinition{value: val, isSet: true}
}

func (v NullableFormulaAndFunctionQueryDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormulaAndFunctionQueryDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return json.Unmarshal(src, &v.value)
}
