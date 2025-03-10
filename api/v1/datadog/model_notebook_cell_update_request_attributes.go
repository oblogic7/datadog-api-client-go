/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// NotebookCellUpdateRequestAttributes - The attributes of a notebook cell in update cell request. Valid cell types are `markdown`, `timeseries`, `toplist`, `heatmap`, `distribution`,
// `log_stream`. [More information on each graph visualization type.](https://docs.datadoghq.com/dashboards/widgets/)
type NotebookCellUpdateRequestAttributes struct {
	NotebookMarkdownCellAttributes     *NotebookMarkdownCellAttributes
	NotebookTimeseriesCellAttributes   *NotebookTimeseriesCellAttributes
	NotebookToplistCellAttributes      *NotebookToplistCellAttributes
	NotebookHeatMapCellAttributes      *NotebookHeatMapCellAttributes
	NotebookDistributionCellAttributes *NotebookDistributionCellAttributes
	NotebookLogStreamCellAttributes    *NotebookLogStreamCellAttributes

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// NotebookMarkdownCellAttributesAsNotebookCellUpdateRequestAttributes is a convenience function that returns NotebookMarkdownCellAttributes wrapped in NotebookCellUpdateRequestAttributes
func NotebookMarkdownCellAttributesAsNotebookCellUpdateRequestAttributes(v *NotebookMarkdownCellAttributes) NotebookCellUpdateRequestAttributes {
	return NotebookCellUpdateRequestAttributes{NotebookMarkdownCellAttributes: v}
}

// NotebookTimeseriesCellAttributesAsNotebookCellUpdateRequestAttributes is a convenience function that returns NotebookTimeseriesCellAttributes wrapped in NotebookCellUpdateRequestAttributes
func NotebookTimeseriesCellAttributesAsNotebookCellUpdateRequestAttributes(v *NotebookTimeseriesCellAttributes) NotebookCellUpdateRequestAttributes {
	return NotebookCellUpdateRequestAttributes{NotebookTimeseriesCellAttributes: v}
}

// NotebookToplistCellAttributesAsNotebookCellUpdateRequestAttributes is a convenience function that returns NotebookToplistCellAttributes wrapped in NotebookCellUpdateRequestAttributes
func NotebookToplistCellAttributesAsNotebookCellUpdateRequestAttributes(v *NotebookToplistCellAttributes) NotebookCellUpdateRequestAttributes {
	return NotebookCellUpdateRequestAttributes{NotebookToplistCellAttributes: v}
}

// NotebookHeatMapCellAttributesAsNotebookCellUpdateRequestAttributes is a convenience function that returns NotebookHeatMapCellAttributes wrapped in NotebookCellUpdateRequestAttributes
func NotebookHeatMapCellAttributesAsNotebookCellUpdateRequestAttributes(v *NotebookHeatMapCellAttributes) NotebookCellUpdateRequestAttributes {
	return NotebookCellUpdateRequestAttributes{NotebookHeatMapCellAttributes: v}
}

// NotebookDistributionCellAttributesAsNotebookCellUpdateRequestAttributes is a convenience function that returns NotebookDistributionCellAttributes wrapped in NotebookCellUpdateRequestAttributes
func NotebookDistributionCellAttributesAsNotebookCellUpdateRequestAttributes(v *NotebookDistributionCellAttributes) NotebookCellUpdateRequestAttributes {
	return NotebookCellUpdateRequestAttributes{NotebookDistributionCellAttributes: v}
}

// NotebookLogStreamCellAttributesAsNotebookCellUpdateRequestAttributes is a convenience function that returns NotebookLogStreamCellAttributes wrapped in NotebookCellUpdateRequestAttributes
func NotebookLogStreamCellAttributesAsNotebookCellUpdateRequestAttributes(v *NotebookLogStreamCellAttributes) NotebookCellUpdateRequestAttributes {
	return NotebookCellUpdateRequestAttributes{NotebookLogStreamCellAttributes: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *NotebookCellUpdateRequestAttributes) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into NotebookMarkdownCellAttributes
	err = json.Unmarshal(data, &dst.NotebookMarkdownCellAttributes)
	if err == nil {
		if dst.NotebookMarkdownCellAttributes != nil && dst.NotebookMarkdownCellAttributes.UnparsedObject == nil {
			jsonNotebookMarkdownCellAttributes, _ := json.Marshal(dst.NotebookMarkdownCellAttributes)
			if string(jsonNotebookMarkdownCellAttributes) == "{}" { // empty struct
				dst.NotebookMarkdownCellAttributes = nil
			} else {
				match++
			}
		} else {
			dst.NotebookMarkdownCellAttributes = nil
		}
	} else {
		dst.NotebookMarkdownCellAttributes = nil
	}

	// try to unmarshal data into NotebookTimeseriesCellAttributes
	err = json.Unmarshal(data, &dst.NotebookTimeseriesCellAttributes)
	if err == nil {
		if dst.NotebookTimeseriesCellAttributes != nil && dst.NotebookTimeseriesCellAttributes.UnparsedObject == nil {
			jsonNotebookTimeseriesCellAttributes, _ := json.Marshal(dst.NotebookTimeseriesCellAttributes)
			if string(jsonNotebookTimeseriesCellAttributes) == "{}" { // empty struct
				dst.NotebookTimeseriesCellAttributes = nil
			} else {
				match++
			}
		} else {
			dst.NotebookTimeseriesCellAttributes = nil
		}
	} else {
		dst.NotebookTimeseriesCellAttributes = nil
	}

	// try to unmarshal data into NotebookToplistCellAttributes
	err = json.Unmarshal(data, &dst.NotebookToplistCellAttributes)
	if err == nil {
		if dst.NotebookToplistCellAttributes != nil && dst.NotebookToplistCellAttributes.UnparsedObject == nil {
			jsonNotebookToplistCellAttributes, _ := json.Marshal(dst.NotebookToplistCellAttributes)
			if string(jsonNotebookToplistCellAttributes) == "{}" { // empty struct
				dst.NotebookToplistCellAttributes = nil
			} else {
				match++
			}
		} else {
			dst.NotebookToplistCellAttributes = nil
		}
	} else {
		dst.NotebookToplistCellAttributes = nil
	}

	// try to unmarshal data into NotebookHeatMapCellAttributes
	err = json.Unmarshal(data, &dst.NotebookHeatMapCellAttributes)
	if err == nil {
		if dst.NotebookHeatMapCellAttributes != nil && dst.NotebookHeatMapCellAttributes.UnparsedObject == nil {
			jsonNotebookHeatMapCellAttributes, _ := json.Marshal(dst.NotebookHeatMapCellAttributes)
			if string(jsonNotebookHeatMapCellAttributes) == "{}" { // empty struct
				dst.NotebookHeatMapCellAttributes = nil
			} else {
				match++
			}
		} else {
			dst.NotebookHeatMapCellAttributes = nil
		}
	} else {
		dst.NotebookHeatMapCellAttributes = nil
	}

	// try to unmarshal data into NotebookDistributionCellAttributes
	err = json.Unmarshal(data, &dst.NotebookDistributionCellAttributes)
	if err == nil {
		if dst.NotebookDistributionCellAttributes != nil && dst.NotebookDistributionCellAttributes.UnparsedObject == nil {
			jsonNotebookDistributionCellAttributes, _ := json.Marshal(dst.NotebookDistributionCellAttributes)
			if string(jsonNotebookDistributionCellAttributes) == "{}" { // empty struct
				dst.NotebookDistributionCellAttributes = nil
			} else {
				match++
			}
		} else {
			dst.NotebookDistributionCellAttributes = nil
		}
	} else {
		dst.NotebookDistributionCellAttributes = nil
	}

	// try to unmarshal data into NotebookLogStreamCellAttributes
	err = json.Unmarshal(data, &dst.NotebookLogStreamCellAttributes)
	if err == nil {
		if dst.NotebookLogStreamCellAttributes != nil && dst.NotebookLogStreamCellAttributes.UnparsedObject == nil {
			jsonNotebookLogStreamCellAttributes, _ := json.Marshal(dst.NotebookLogStreamCellAttributes)
			if string(jsonNotebookLogStreamCellAttributes) == "{}" { // empty struct
				dst.NotebookLogStreamCellAttributes = nil
			} else {
				match++
			}
		} else {
			dst.NotebookLogStreamCellAttributes = nil
		}
	} else {
		dst.NotebookLogStreamCellAttributes = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		dst.NotebookMarkdownCellAttributes = nil
		dst.NotebookTimeseriesCellAttributes = nil
		dst.NotebookToplistCellAttributes = nil
		dst.NotebookHeatMapCellAttributes = nil
		dst.NotebookDistributionCellAttributes = nil
		dst.NotebookLogStreamCellAttributes = nil
		return json.Unmarshal(data, &dst.UnparsedObject)
	} else {
		return nil // exactly one match
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NotebookCellUpdateRequestAttributes) MarshalJSON() ([]byte, error) {
	if src.NotebookMarkdownCellAttributes != nil {
		return json.Marshal(&src.NotebookMarkdownCellAttributes)
	}

	if src.NotebookTimeseriesCellAttributes != nil {
		return json.Marshal(&src.NotebookTimeseriesCellAttributes)
	}

	if src.NotebookToplistCellAttributes != nil {
		return json.Marshal(&src.NotebookToplistCellAttributes)
	}

	if src.NotebookHeatMapCellAttributes != nil {
		return json.Marshal(&src.NotebookHeatMapCellAttributes)
	}

	if src.NotebookDistributionCellAttributes != nil {
		return json.Marshal(&src.NotebookDistributionCellAttributes)
	}

	if src.NotebookLogStreamCellAttributes != nil {
		return json.Marshal(&src.NotebookLogStreamCellAttributes)
	}

	if src.UnparsedObject != nil {
		return json.Marshal(src.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *NotebookCellUpdateRequestAttributes) GetActualInstance() interface{} {
	if obj.NotebookMarkdownCellAttributes != nil {
		return obj.NotebookMarkdownCellAttributes
	}

	if obj.NotebookTimeseriesCellAttributes != nil {
		return obj.NotebookTimeseriesCellAttributes
	}

	if obj.NotebookToplistCellAttributes != nil {
		return obj.NotebookToplistCellAttributes
	}

	if obj.NotebookHeatMapCellAttributes != nil {
		return obj.NotebookHeatMapCellAttributes
	}

	if obj.NotebookDistributionCellAttributes != nil {
		return obj.NotebookDistributionCellAttributes
	}

	if obj.NotebookLogStreamCellAttributes != nil {
		return obj.NotebookLogStreamCellAttributes
	}

	// all schemas are nil
	return nil
}

type NullableNotebookCellUpdateRequestAttributes struct {
	value *NotebookCellUpdateRequestAttributes
	isSet bool
}

func (v NullableNotebookCellUpdateRequestAttributes) Get() *NotebookCellUpdateRequestAttributes {
	return v.value
}

func (v *NullableNotebookCellUpdateRequestAttributes) Set(val *NotebookCellUpdateRequestAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableNotebookCellUpdateRequestAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableNotebookCellUpdateRequestAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotebookCellUpdateRequestAttributes(val *NotebookCellUpdateRequestAttributes) *NullableNotebookCellUpdateRequestAttributes {
	return &NullableNotebookCellUpdateRequestAttributes{value: val, isSet: true}
}

func (v NullableNotebookCellUpdateRequestAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotebookCellUpdateRequestAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return json.Unmarshal(src, &v.value)
}
