/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// LogsProcessor - Definition of a logs processor.
type LogsProcessor struct {
	LogsGrokParser             *LogsGrokParser
	LogsDateRemapper           *LogsDateRemapper
	LogsStatusRemapper         *LogsStatusRemapper
	LogsServiceRemapper        *LogsServiceRemapper
	LogsMessageRemapper        *LogsMessageRemapper
	LogsAttributeRemapper      *LogsAttributeRemapper
	LogsURLParser              *LogsURLParser
	LogsUserAgentParser        *LogsUserAgentParser
	LogsCategoryProcessor      *LogsCategoryProcessor
	LogsArithmeticProcessor    *LogsArithmeticProcessor
	LogsStringBuilderProcessor *LogsStringBuilderProcessor
	LogsPipelineProcessor      *LogsPipelineProcessor
	LogsGeoIPParser            *LogsGeoIPParser
	LogsLookupProcessor        *LogsLookupProcessor
	LogsTraceRemapper          *LogsTraceRemapper

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// LogsGrokParserAsLogsProcessor is a convenience function that returns LogsGrokParser wrapped in LogsProcessor
func LogsGrokParserAsLogsProcessor(v *LogsGrokParser) LogsProcessor {
	return LogsProcessor{LogsGrokParser: v}
}

// LogsDateRemapperAsLogsProcessor is a convenience function that returns LogsDateRemapper wrapped in LogsProcessor
func LogsDateRemapperAsLogsProcessor(v *LogsDateRemapper) LogsProcessor {
	return LogsProcessor{LogsDateRemapper: v}
}

// LogsStatusRemapperAsLogsProcessor is a convenience function that returns LogsStatusRemapper wrapped in LogsProcessor
func LogsStatusRemapperAsLogsProcessor(v *LogsStatusRemapper) LogsProcessor {
	return LogsProcessor{LogsStatusRemapper: v}
}

// LogsServiceRemapperAsLogsProcessor is a convenience function that returns LogsServiceRemapper wrapped in LogsProcessor
func LogsServiceRemapperAsLogsProcessor(v *LogsServiceRemapper) LogsProcessor {
	return LogsProcessor{LogsServiceRemapper: v}
}

// LogsMessageRemapperAsLogsProcessor is a convenience function that returns LogsMessageRemapper wrapped in LogsProcessor
func LogsMessageRemapperAsLogsProcessor(v *LogsMessageRemapper) LogsProcessor {
	return LogsProcessor{LogsMessageRemapper: v}
}

// LogsAttributeRemapperAsLogsProcessor is a convenience function that returns LogsAttributeRemapper wrapped in LogsProcessor
func LogsAttributeRemapperAsLogsProcessor(v *LogsAttributeRemapper) LogsProcessor {
	return LogsProcessor{LogsAttributeRemapper: v}
}

// LogsURLParserAsLogsProcessor is a convenience function that returns LogsURLParser wrapped in LogsProcessor
func LogsURLParserAsLogsProcessor(v *LogsURLParser) LogsProcessor {
	return LogsProcessor{LogsURLParser: v}
}

// LogsUserAgentParserAsLogsProcessor is a convenience function that returns LogsUserAgentParser wrapped in LogsProcessor
func LogsUserAgentParserAsLogsProcessor(v *LogsUserAgentParser) LogsProcessor {
	return LogsProcessor{LogsUserAgentParser: v}
}

// LogsCategoryProcessorAsLogsProcessor is a convenience function that returns LogsCategoryProcessor wrapped in LogsProcessor
func LogsCategoryProcessorAsLogsProcessor(v *LogsCategoryProcessor) LogsProcessor {
	return LogsProcessor{LogsCategoryProcessor: v}
}

// LogsArithmeticProcessorAsLogsProcessor is a convenience function that returns LogsArithmeticProcessor wrapped in LogsProcessor
func LogsArithmeticProcessorAsLogsProcessor(v *LogsArithmeticProcessor) LogsProcessor {
	return LogsProcessor{LogsArithmeticProcessor: v}
}

// LogsStringBuilderProcessorAsLogsProcessor is a convenience function that returns LogsStringBuilderProcessor wrapped in LogsProcessor
func LogsStringBuilderProcessorAsLogsProcessor(v *LogsStringBuilderProcessor) LogsProcessor {
	return LogsProcessor{LogsStringBuilderProcessor: v}
}

// LogsPipelineProcessorAsLogsProcessor is a convenience function that returns LogsPipelineProcessor wrapped in LogsProcessor
func LogsPipelineProcessorAsLogsProcessor(v *LogsPipelineProcessor) LogsProcessor {
	return LogsProcessor{LogsPipelineProcessor: v}
}

// LogsGeoIPParserAsLogsProcessor is a convenience function that returns LogsGeoIPParser wrapped in LogsProcessor
func LogsGeoIPParserAsLogsProcessor(v *LogsGeoIPParser) LogsProcessor {
	return LogsProcessor{LogsGeoIPParser: v}
}

// LogsLookupProcessorAsLogsProcessor is a convenience function that returns LogsLookupProcessor wrapped in LogsProcessor
func LogsLookupProcessorAsLogsProcessor(v *LogsLookupProcessor) LogsProcessor {
	return LogsProcessor{LogsLookupProcessor: v}
}

// LogsTraceRemapperAsLogsProcessor is a convenience function that returns LogsTraceRemapper wrapped in LogsProcessor
func LogsTraceRemapperAsLogsProcessor(v *LogsTraceRemapper) LogsProcessor {
	return LogsProcessor{LogsTraceRemapper: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *LogsProcessor) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into LogsGrokParser
	err = json.Unmarshal(data, &dst.LogsGrokParser)
	if err == nil {
		if dst.LogsGrokParser != nil && dst.LogsGrokParser.UnparsedObject == nil {
			jsonLogsGrokParser, _ := json.Marshal(dst.LogsGrokParser)
			if string(jsonLogsGrokParser) == "{}" { // empty struct
				dst.LogsGrokParser = nil
			} else {
				match++
			}
		} else {
			dst.LogsGrokParser = nil
		}
	} else {
		dst.LogsGrokParser = nil
	}

	// try to unmarshal data into LogsDateRemapper
	err = json.Unmarshal(data, &dst.LogsDateRemapper)
	if err == nil {
		if dst.LogsDateRemapper != nil && dst.LogsDateRemapper.UnparsedObject == nil {
			jsonLogsDateRemapper, _ := json.Marshal(dst.LogsDateRemapper)
			if string(jsonLogsDateRemapper) == "{}" { // empty struct
				dst.LogsDateRemapper = nil
			} else {
				match++
			}
		} else {
			dst.LogsDateRemapper = nil
		}
	} else {
		dst.LogsDateRemapper = nil
	}

	// try to unmarshal data into LogsStatusRemapper
	err = json.Unmarshal(data, &dst.LogsStatusRemapper)
	if err == nil {
		if dst.LogsStatusRemapper != nil && dst.LogsStatusRemapper.UnparsedObject == nil {
			jsonLogsStatusRemapper, _ := json.Marshal(dst.LogsStatusRemapper)
			if string(jsonLogsStatusRemapper) == "{}" { // empty struct
				dst.LogsStatusRemapper = nil
			} else {
				match++
			}
		} else {
			dst.LogsStatusRemapper = nil
		}
	} else {
		dst.LogsStatusRemapper = nil
	}

	// try to unmarshal data into LogsServiceRemapper
	err = json.Unmarshal(data, &dst.LogsServiceRemapper)
	if err == nil {
		if dst.LogsServiceRemapper != nil && dst.LogsServiceRemapper.UnparsedObject == nil {
			jsonLogsServiceRemapper, _ := json.Marshal(dst.LogsServiceRemapper)
			if string(jsonLogsServiceRemapper) == "{}" { // empty struct
				dst.LogsServiceRemapper = nil
			} else {
				match++
			}
		} else {
			dst.LogsServiceRemapper = nil
		}
	} else {
		dst.LogsServiceRemapper = nil
	}

	// try to unmarshal data into LogsMessageRemapper
	err = json.Unmarshal(data, &dst.LogsMessageRemapper)
	if err == nil {
		if dst.LogsMessageRemapper != nil && dst.LogsMessageRemapper.UnparsedObject == nil {
			jsonLogsMessageRemapper, _ := json.Marshal(dst.LogsMessageRemapper)
			if string(jsonLogsMessageRemapper) == "{}" { // empty struct
				dst.LogsMessageRemapper = nil
			} else {
				match++
			}
		} else {
			dst.LogsMessageRemapper = nil
		}
	} else {
		dst.LogsMessageRemapper = nil
	}

	// try to unmarshal data into LogsAttributeRemapper
	err = json.Unmarshal(data, &dst.LogsAttributeRemapper)
	if err == nil {
		if dst.LogsAttributeRemapper != nil && dst.LogsAttributeRemapper.UnparsedObject == nil {
			jsonLogsAttributeRemapper, _ := json.Marshal(dst.LogsAttributeRemapper)
			if string(jsonLogsAttributeRemapper) == "{}" { // empty struct
				dst.LogsAttributeRemapper = nil
			} else {
				match++
			}
		} else {
			dst.LogsAttributeRemapper = nil
		}
	} else {
		dst.LogsAttributeRemapper = nil
	}

	// try to unmarshal data into LogsURLParser
	err = json.Unmarshal(data, &dst.LogsURLParser)
	if err == nil {
		if dst.LogsURLParser != nil && dst.LogsURLParser.UnparsedObject == nil {
			jsonLogsURLParser, _ := json.Marshal(dst.LogsURLParser)
			if string(jsonLogsURLParser) == "{}" { // empty struct
				dst.LogsURLParser = nil
			} else {
				match++
			}
		} else {
			dst.LogsURLParser = nil
		}
	} else {
		dst.LogsURLParser = nil
	}

	// try to unmarshal data into LogsUserAgentParser
	err = json.Unmarshal(data, &dst.LogsUserAgentParser)
	if err == nil {
		if dst.LogsUserAgentParser != nil && dst.LogsUserAgentParser.UnparsedObject == nil {
			jsonLogsUserAgentParser, _ := json.Marshal(dst.LogsUserAgentParser)
			if string(jsonLogsUserAgentParser) == "{}" { // empty struct
				dst.LogsUserAgentParser = nil
			} else {
				match++
			}
		} else {
			dst.LogsUserAgentParser = nil
		}
	} else {
		dst.LogsUserAgentParser = nil
	}

	// try to unmarshal data into LogsCategoryProcessor
	err = json.Unmarshal(data, &dst.LogsCategoryProcessor)
	if err == nil {
		if dst.LogsCategoryProcessor != nil && dst.LogsCategoryProcessor.UnparsedObject == nil {
			jsonLogsCategoryProcessor, _ := json.Marshal(dst.LogsCategoryProcessor)
			if string(jsonLogsCategoryProcessor) == "{}" { // empty struct
				dst.LogsCategoryProcessor = nil
			} else {
				match++
			}
		} else {
			dst.LogsCategoryProcessor = nil
		}
	} else {
		dst.LogsCategoryProcessor = nil
	}

	// try to unmarshal data into LogsArithmeticProcessor
	err = json.Unmarshal(data, &dst.LogsArithmeticProcessor)
	if err == nil {
		if dst.LogsArithmeticProcessor != nil && dst.LogsArithmeticProcessor.UnparsedObject == nil {
			jsonLogsArithmeticProcessor, _ := json.Marshal(dst.LogsArithmeticProcessor)
			if string(jsonLogsArithmeticProcessor) == "{}" { // empty struct
				dst.LogsArithmeticProcessor = nil
			} else {
				match++
			}
		} else {
			dst.LogsArithmeticProcessor = nil
		}
	} else {
		dst.LogsArithmeticProcessor = nil
	}

	// try to unmarshal data into LogsStringBuilderProcessor
	err = json.Unmarshal(data, &dst.LogsStringBuilderProcessor)
	if err == nil {
		if dst.LogsStringBuilderProcessor != nil && dst.LogsStringBuilderProcessor.UnparsedObject == nil {
			jsonLogsStringBuilderProcessor, _ := json.Marshal(dst.LogsStringBuilderProcessor)
			if string(jsonLogsStringBuilderProcessor) == "{}" { // empty struct
				dst.LogsStringBuilderProcessor = nil
			} else {
				match++
			}
		} else {
			dst.LogsStringBuilderProcessor = nil
		}
	} else {
		dst.LogsStringBuilderProcessor = nil
	}

	// try to unmarshal data into LogsPipelineProcessor
	err = json.Unmarshal(data, &dst.LogsPipelineProcessor)
	if err == nil {
		if dst.LogsPipelineProcessor != nil && dst.LogsPipelineProcessor.UnparsedObject == nil {
			jsonLogsPipelineProcessor, _ := json.Marshal(dst.LogsPipelineProcessor)
			if string(jsonLogsPipelineProcessor) == "{}" { // empty struct
				dst.LogsPipelineProcessor = nil
			} else {
				match++
			}
		} else {
			dst.LogsPipelineProcessor = nil
		}
	} else {
		dst.LogsPipelineProcessor = nil
	}

	// try to unmarshal data into LogsGeoIPParser
	err = json.Unmarshal(data, &dst.LogsGeoIPParser)
	if err == nil {
		if dst.LogsGeoIPParser != nil && dst.LogsGeoIPParser.UnparsedObject == nil {
			jsonLogsGeoIPParser, _ := json.Marshal(dst.LogsGeoIPParser)
			if string(jsonLogsGeoIPParser) == "{}" { // empty struct
				dst.LogsGeoIPParser = nil
			} else {
				match++
			}
		} else {
			dst.LogsGeoIPParser = nil
		}
	} else {
		dst.LogsGeoIPParser = nil
	}

	// try to unmarshal data into LogsLookupProcessor
	err = json.Unmarshal(data, &dst.LogsLookupProcessor)
	if err == nil {
		if dst.LogsLookupProcessor != nil && dst.LogsLookupProcessor.UnparsedObject == nil {
			jsonLogsLookupProcessor, _ := json.Marshal(dst.LogsLookupProcessor)
			if string(jsonLogsLookupProcessor) == "{}" { // empty struct
				dst.LogsLookupProcessor = nil
			} else {
				match++
			}
		} else {
			dst.LogsLookupProcessor = nil
		}
	} else {
		dst.LogsLookupProcessor = nil
	}

	// try to unmarshal data into LogsTraceRemapper
	err = json.Unmarshal(data, &dst.LogsTraceRemapper)
	if err == nil {
		if dst.LogsTraceRemapper != nil && dst.LogsTraceRemapper.UnparsedObject == nil {
			jsonLogsTraceRemapper, _ := json.Marshal(dst.LogsTraceRemapper)
			if string(jsonLogsTraceRemapper) == "{}" { // empty struct
				dst.LogsTraceRemapper = nil
			} else {
				match++
			}
		} else {
			dst.LogsTraceRemapper = nil
		}
	} else {
		dst.LogsTraceRemapper = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		dst.LogsGrokParser = nil
		dst.LogsDateRemapper = nil
		dst.LogsStatusRemapper = nil
		dst.LogsServiceRemapper = nil
		dst.LogsMessageRemapper = nil
		dst.LogsAttributeRemapper = nil
		dst.LogsURLParser = nil
		dst.LogsUserAgentParser = nil
		dst.LogsCategoryProcessor = nil
		dst.LogsArithmeticProcessor = nil
		dst.LogsStringBuilderProcessor = nil
		dst.LogsPipelineProcessor = nil
		dst.LogsGeoIPParser = nil
		dst.LogsLookupProcessor = nil
		dst.LogsTraceRemapper = nil
		return json.Unmarshal(data, &dst.UnparsedObject)
	} else {
		return nil // exactly one match
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src LogsProcessor) MarshalJSON() ([]byte, error) {
	if src.LogsGrokParser != nil {
		return json.Marshal(&src.LogsGrokParser)
	}

	if src.LogsDateRemapper != nil {
		return json.Marshal(&src.LogsDateRemapper)
	}

	if src.LogsStatusRemapper != nil {
		return json.Marshal(&src.LogsStatusRemapper)
	}

	if src.LogsServiceRemapper != nil {
		return json.Marshal(&src.LogsServiceRemapper)
	}

	if src.LogsMessageRemapper != nil {
		return json.Marshal(&src.LogsMessageRemapper)
	}

	if src.LogsAttributeRemapper != nil {
		return json.Marshal(&src.LogsAttributeRemapper)
	}

	if src.LogsURLParser != nil {
		return json.Marshal(&src.LogsURLParser)
	}

	if src.LogsUserAgentParser != nil {
		return json.Marshal(&src.LogsUserAgentParser)
	}

	if src.LogsCategoryProcessor != nil {
		return json.Marshal(&src.LogsCategoryProcessor)
	}

	if src.LogsArithmeticProcessor != nil {
		return json.Marshal(&src.LogsArithmeticProcessor)
	}

	if src.LogsStringBuilderProcessor != nil {
		return json.Marshal(&src.LogsStringBuilderProcessor)
	}

	if src.LogsPipelineProcessor != nil {
		return json.Marshal(&src.LogsPipelineProcessor)
	}

	if src.LogsGeoIPParser != nil {
		return json.Marshal(&src.LogsGeoIPParser)
	}

	if src.LogsLookupProcessor != nil {
		return json.Marshal(&src.LogsLookupProcessor)
	}

	if src.LogsTraceRemapper != nil {
		return json.Marshal(&src.LogsTraceRemapper)
	}

	if src.UnparsedObject != nil {
		return json.Marshal(src.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *LogsProcessor) GetActualInstance() interface{} {
	if obj.LogsGrokParser != nil {
		return obj.LogsGrokParser
	}

	if obj.LogsDateRemapper != nil {
		return obj.LogsDateRemapper
	}

	if obj.LogsStatusRemapper != nil {
		return obj.LogsStatusRemapper
	}

	if obj.LogsServiceRemapper != nil {
		return obj.LogsServiceRemapper
	}

	if obj.LogsMessageRemapper != nil {
		return obj.LogsMessageRemapper
	}

	if obj.LogsAttributeRemapper != nil {
		return obj.LogsAttributeRemapper
	}

	if obj.LogsURLParser != nil {
		return obj.LogsURLParser
	}

	if obj.LogsUserAgentParser != nil {
		return obj.LogsUserAgentParser
	}

	if obj.LogsCategoryProcessor != nil {
		return obj.LogsCategoryProcessor
	}

	if obj.LogsArithmeticProcessor != nil {
		return obj.LogsArithmeticProcessor
	}

	if obj.LogsStringBuilderProcessor != nil {
		return obj.LogsStringBuilderProcessor
	}

	if obj.LogsPipelineProcessor != nil {
		return obj.LogsPipelineProcessor
	}

	if obj.LogsGeoIPParser != nil {
		return obj.LogsGeoIPParser
	}

	if obj.LogsLookupProcessor != nil {
		return obj.LogsLookupProcessor
	}

	if obj.LogsTraceRemapper != nil {
		return obj.LogsTraceRemapper
	}

	// all schemas are nil
	return nil
}

type NullableLogsProcessor struct {
	value *LogsProcessor
	isSet bool
}

func (v NullableLogsProcessor) Get() *LogsProcessor {
	return v.value
}

func (v *NullableLogsProcessor) Set(val *LogsProcessor) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsProcessor) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsProcessor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsProcessor(val *LogsProcessor) *NullableLogsProcessor {
	return &NullableLogsProcessor{value: val, isSet: true}
}

func (v NullableLogsProcessor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsProcessor) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return json.Unmarshal(src, &v.value)
}
