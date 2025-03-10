/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// WidgetDefinition - [Definition of the widget](https://docs.datadoghq.com/dashboards/widgets/).
type WidgetDefinition struct {
	AlertGraphWidgetDefinition     *AlertGraphWidgetDefinition
	AlertValueWidgetDefinition     *AlertValueWidgetDefinition
	ChangeWidgetDefinition         *ChangeWidgetDefinition
	CheckStatusWidgetDefinition    *CheckStatusWidgetDefinition
	DistributionWidgetDefinition   *DistributionWidgetDefinition
	EventStreamWidgetDefinition    *EventStreamWidgetDefinition
	EventTimelineWidgetDefinition  *EventTimelineWidgetDefinition
	FreeTextWidgetDefinition       *FreeTextWidgetDefinition
	GeomapWidgetDefinition         *GeomapWidgetDefinition
	GroupWidgetDefinition          *GroupWidgetDefinition
	HeatMapWidgetDefinition        *HeatMapWidgetDefinition
	HostMapWidgetDefinition        *HostMapWidgetDefinition
	IFrameWidgetDefinition         *IFrameWidgetDefinition
	ImageWidgetDefinition          *ImageWidgetDefinition
	LogStreamWidgetDefinition      *LogStreamWidgetDefinition
	MonitorSummaryWidgetDefinition *MonitorSummaryWidgetDefinition
	NoteWidgetDefinition           *NoteWidgetDefinition
	QueryValueWidgetDefinition     *QueryValueWidgetDefinition
	ScatterPlotWidgetDefinition    *ScatterPlotWidgetDefinition
	SLOWidgetDefinition            *SLOWidgetDefinition
	ServiceMapWidgetDefinition     *ServiceMapWidgetDefinition
	ServiceSummaryWidgetDefinition *ServiceSummaryWidgetDefinition
	SunburstWidgetDefinition       *SunburstWidgetDefinition
	TableWidgetDefinition          *TableWidgetDefinition
	TimeseriesWidgetDefinition     *TimeseriesWidgetDefinition
	ToplistWidgetDefinition        *ToplistWidgetDefinition
	TreeMapWidgetDefinition        *TreeMapWidgetDefinition
	ListStreamWidgetDefinition     *ListStreamWidgetDefinition
	FunnelWidgetDefinition         *FunnelWidgetDefinition

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// AlertGraphWidgetDefinitionAsWidgetDefinition is a convenience function that returns AlertGraphWidgetDefinition wrapped in WidgetDefinition
func AlertGraphWidgetDefinitionAsWidgetDefinition(v *AlertGraphWidgetDefinition) WidgetDefinition {
	return WidgetDefinition{AlertGraphWidgetDefinition: v}
}

// AlertValueWidgetDefinitionAsWidgetDefinition is a convenience function that returns AlertValueWidgetDefinition wrapped in WidgetDefinition
func AlertValueWidgetDefinitionAsWidgetDefinition(v *AlertValueWidgetDefinition) WidgetDefinition {
	return WidgetDefinition{AlertValueWidgetDefinition: v}
}

// ChangeWidgetDefinitionAsWidgetDefinition is a convenience function that returns ChangeWidgetDefinition wrapped in WidgetDefinition
func ChangeWidgetDefinitionAsWidgetDefinition(v *ChangeWidgetDefinition) WidgetDefinition {
	return WidgetDefinition{ChangeWidgetDefinition: v}
}

// CheckStatusWidgetDefinitionAsWidgetDefinition is a convenience function that returns CheckStatusWidgetDefinition wrapped in WidgetDefinition
func CheckStatusWidgetDefinitionAsWidgetDefinition(v *CheckStatusWidgetDefinition) WidgetDefinition {
	return WidgetDefinition{CheckStatusWidgetDefinition: v}
}

// DistributionWidgetDefinitionAsWidgetDefinition is a convenience function that returns DistributionWidgetDefinition wrapped in WidgetDefinition
func DistributionWidgetDefinitionAsWidgetDefinition(v *DistributionWidgetDefinition) WidgetDefinition {
	return WidgetDefinition{DistributionWidgetDefinition: v}
}

// EventStreamWidgetDefinitionAsWidgetDefinition is a convenience function that returns EventStreamWidgetDefinition wrapped in WidgetDefinition
func EventStreamWidgetDefinitionAsWidgetDefinition(v *EventStreamWidgetDefinition) WidgetDefinition {
	return WidgetDefinition{EventStreamWidgetDefinition: v}
}

// EventTimelineWidgetDefinitionAsWidgetDefinition is a convenience function that returns EventTimelineWidgetDefinition wrapped in WidgetDefinition
func EventTimelineWidgetDefinitionAsWidgetDefinition(v *EventTimelineWidgetDefinition) WidgetDefinition {
	return WidgetDefinition{EventTimelineWidgetDefinition: v}
}

// FreeTextWidgetDefinitionAsWidgetDefinition is a convenience function that returns FreeTextWidgetDefinition wrapped in WidgetDefinition
func FreeTextWidgetDefinitionAsWidgetDefinition(v *FreeTextWidgetDefinition) WidgetDefinition {
	return WidgetDefinition{FreeTextWidgetDefinition: v}
}

// GeomapWidgetDefinitionAsWidgetDefinition is a convenience function that returns GeomapWidgetDefinition wrapped in WidgetDefinition
func GeomapWidgetDefinitionAsWidgetDefinition(v *GeomapWidgetDefinition) WidgetDefinition {
	return WidgetDefinition{GeomapWidgetDefinition: v}
}

// GroupWidgetDefinitionAsWidgetDefinition is a convenience function that returns GroupWidgetDefinition wrapped in WidgetDefinition
func GroupWidgetDefinitionAsWidgetDefinition(v *GroupWidgetDefinition) WidgetDefinition {
	return WidgetDefinition{GroupWidgetDefinition: v}
}

// HeatMapWidgetDefinitionAsWidgetDefinition is a convenience function that returns HeatMapWidgetDefinition wrapped in WidgetDefinition
func HeatMapWidgetDefinitionAsWidgetDefinition(v *HeatMapWidgetDefinition) WidgetDefinition {
	return WidgetDefinition{HeatMapWidgetDefinition: v}
}

// HostMapWidgetDefinitionAsWidgetDefinition is a convenience function that returns HostMapWidgetDefinition wrapped in WidgetDefinition
func HostMapWidgetDefinitionAsWidgetDefinition(v *HostMapWidgetDefinition) WidgetDefinition {
	return WidgetDefinition{HostMapWidgetDefinition: v}
}

// IFrameWidgetDefinitionAsWidgetDefinition is a convenience function that returns IFrameWidgetDefinition wrapped in WidgetDefinition
func IFrameWidgetDefinitionAsWidgetDefinition(v *IFrameWidgetDefinition) WidgetDefinition {
	return WidgetDefinition{IFrameWidgetDefinition: v}
}

// ImageWidgetDefinitionAsWidgetDefinition is a convenience function that returns ImageWidgetDefinition wrapped in WidgetDefinition
func ImageWidgetDefinitionAsWidgetDefinition(v *ImageWidgetDefinition) WidgetDefinition {
	return WidgetDefinition{ImageWidgetDefinition: v}
}

// LogStreamWidgetDefinitionAsWidgetDefinition is a convenience function that returns LogStreamWidgetDefinition wrapped in WidgetDefinition
func LogStreamWidgetDefinitionAsWidgetDefinition(v *LogStreamWidgetDefinition) WidgetDefinition {
	return WidgetDefinition{LogStreamWidgetDefinition: v}
}

// MonitorSummaryWidgetDefinitionAsWidgetDefinition is a convenience function that returns MonitorSummaryWidgetDefinition wrapped in WidgetDefinition
func MonitorSummaryWidgetDefinitionAsWidgetDefinition(v *MonitorSummaryWidgetDefinition) WidgetDefinition {
	return WidgetDefinition{MonitorSummaryWidgetDefinition: v}
}

// NoteWidgetDefinitionAsWidgetDefinition is a convenience function that returns NoteWidgetDefinition wrapped in WidgetDefinition
func NoteWidgetDefinitionAsWidgetDefinition(v *NoteWidgetDefinition) WidgetDefinition {
	return WidgetDefinition{NoteWidgetDefinition: v}
}

// QueryValueWidgetDefinitionAsWidgetDefinition is a convenience function that returns QueryValueWidgetDefinition wrapped in WidgetDefinition
func QueryValueWidgetDefinitionAsWidgetDefinition(v *QueryValueWidgetDefinition) WidgetDefinition {
	return WidgetDefinition{QueryValueWidgetDefinition: v}
}

// ScatterPlotWidgetDefinitionAsWidgetDefinition is a convenience function that returns ScatterPlotWidgetDefinition wrapped in WidgetDefinition
func ScatterPlotWidgetDefinitionAsWidgetDefinition(v *ScatterPlotWidgetDefinition) WidgetDefinition {
	return WidgetDefinition{ScatterPlotWidgetDefinition: v}
}

// SLOWidgetDefinitionAsWidgetDefinition is a convenience function that returns SLOWidgetDefinition wrapped in WidgetDefinition
func SLOWidgetDefinitionAsWidgetDefinition(v *SLOWidgetDefinition) WidgetDefinition {
	return WidgetDefinition{SLOWidgetDefinition: v}
}

// ServiceMapWidgetDefinitionAsWidgetDefinition is a convenience function that returns ServiceMapWidgetDefinition wrapped in WidgetDefinition
func ServiceMapWidgetDefinitionAsWidgetDefinition(v *ServiceMapWidgetDefinition) WidgetDefinition {
	return WidgetDefinition{ServiceMapWidgetDefinition: v}
}

// ServiceSummaryWidgetDefinitionAsWidgetDefinition is a convenience function that returns ServiceSummaryWidgetDefinition wrapped in WidgetDefinition
func ServiceSummaryWidgetDefinitionAsWidgetDefinition(v *ServiceSummaryWidgetDefinition) WidgetDefinition {
	return WidgetDefinition{ServiceSummaryWidgetDefinition: v}
}

// SunburstWidgetDefinitionAsWidgetDefinition is a convenience function that returns SunburstWidgetDefinition wrapped in WidgetDefinition
func SunburstWidgetDefinitionAsWidgetDefinition(v *SunburstWidgetDefinition) WidgetDefinition {
	return WidgetDefinition{SunburstWidgetDefinition: v}
}

// TableWidgetDefinitionAsWidgetDefinition is a convenience function that returns TableWidgetDefinition wrapped in WidgetDefinition
func TableWidgetDefinitionAsWidgetDefinition(v *TableWidgetDefinition) WidgetDefinition {
	return WidgetDefinition{TableWidgetDefinition: v}
}

// TimeseriesWidgetDefinitionAsWidgetDefinition is a convenience function that returns TimeseriesWidgetDefinition wrapped in WidgetDefinition
func TimeseriesWidgetDefinitionAsWidgetDefinition(v *TimeseriesWidgetDefinition) WidgetDefinition {
	return WidgetDefinition{TimeseriesWidgetDefinition: v}
}

// ToplistWidgetDefinitionAsWidgetDefinition is a convenience function that returns ToplistWidgetDefinition wrapped in WidgetDefinition
func ToplistWidgetDefinitionAsWidgetDefinition(v *ToplistWidgetDefinition) WidgetDefinition {
	return WidgetDefinition{ToplistWidgetDefinition: v}
}

// TreeMapWidgetDefinitionAsWidgetDefinition is a convenience function that returns TreeMapWidgetDefinition wrapped in WidgetDefinition
func TreeMapWidgetDefinitionAsWidgetDefinition(v *TreeMapWidgetDefinition) WidgetDefinition {
	return WidgetDefinition{TreeMapWidgetDefinition: v}
}

// ListStreamWidgetDefinitionAsWidgetDefinition is a convenience function that returns ListStreamWidgetDefinition wrapped in WidgetDefinition
func ListStreamWidgetDefinitionAsWidgetDefinition(v *ListStreamWidgetDefinition) WidgetDefinition {
	return WidgetDefinition{ListStreamWidgetDefinition: v}
}

// FunnelWidgetDefinitionAsWidgetDefinition is a convenience function that returns FunnelWidgetDefinition wrapped in WidgetDefinition
func FunnelWidgetDefinitionAsWidgetDefinition(v *FunnelWidgetDefinition) WidgetDefinition {
	return WidgetDefinition{FunnelWidgetDefinition: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *WidgetDefinition) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AlertGraphWidgetDefinition
	err = json.Unmarshal(data, &dst.AlertGraphWidgetDefinition)
	if err == nil {
		if dst.AlertGraphWidgetDefinition != nil && dst.AlertGraphWidgetDefinition.UnparsedObject == nil {
			jsonAlertGraphWidgetDefinition, _ := json.Marshal(dst.AlertGraphWidgetDefinition)
			if string(jsonAlertGraphWidgetDefinition) == "{}" { // empty struct
				dst.AlertGraphWidgetDefinition = nil
			} else {
				match++
			}
		} else {
			dst.AlertGraphWidgetDefinition = nil
		}
	} else {
		dst.AlertGraphWidgetDefinition = nil
	}

	// try to unmarshal data into AlertValueWidgetDefinition
	err = json.Unmarshal(data, &dst.AlertValueWidgetDefinition)
	if err == nil {
		if dst.AlertValueWidgetDefinition != nil && dst.AlertValueWidgetDefinition.UnparsedObject == nil {
			jsonAlertValueWidgetDefinition, _ := json.Marshal(dst.AlertValueWidgetDefinition)
			if string(jsonAlertValueWidgetDefinition) == "{}" { // empty struct
				dst.AlertValueWidgetDefinition = nil
			} else {
				match++
			}
		} else {
			dst.AlertValueWidgetDefinition = nil
		}
	} else {
		dst.AlertValueWidgetDefinition = nil
	}

	// try to unmarshal data into ChangeWidgetDefinition
	err = json.Unmarshal(data, &dst.ChangeWidgetDefinition)
	if err == nil {
		if dst.ChangeWidgetDefinition != nil && dst.ChangeWidgetDefinition.UnparsedObject == nil {
			jsonChangeWidgetDefinition, _ := json.Marshal(dst.ChangeWidgetDefinition)
			if string(jsonChangeWidgetDefinition) == "{}" { // empty struct
				dst.ChangeWidgetDefinition = nil
			} else {
				match++
			}
		} else {
			dst.ChangeWidgetDefinition = nil
		}
	} else {
		dst.ChangeWidgetDefinition = nil
	}

	// try to unmarshal data into CheckStatusWidgetDefinition
	err = json.Unmarshal(data, &dst.CheckStatusWidgetDefinition)
	if err == nil {
		if dst.CheckStatusWidgetDefinition != nil && dst.CheckStatusWidgetDefinition.UnparsedObject == nil {
			jsonCheckStatusWidgetDefinition, _ := json.Marshal(dst.CheckStatusWidgetDefinition)
			if string(jsonCheckStatusWidgetDefinition) == "{}" { // empty struct
				dst.CheckStatusWidgetDefinition = nil
			} else {
				match++
			}
		} else {
			dst.CheckStatusWidgetDefinition = nil
		}
	} else {
		dst.CheckStatusWidgetDefinition = nil
	}

	// try to unmarshal data into DistributionWidgetDefinition
	err = json.Unmarshal(data, &dst.DistributionWidgetDefinition)
	if err == nil {
		if dst.DistributionWidgetDefinition != nil && dst.DistributionWidgetDefinition.UnparsedObject == nil {
			jsonDistributionWidgetDefinition, _ := json.Marshal(dst.DistributionWidgetDefinition)
			if string(jsonDistributionWidgetDefinition) == "{}" { // empty struct
				dst.DistributionWidgetDefinition = nil
			} else {
				match++
			}
		} else {
			dst.DistributionWidgetDefinition = nil
		}
	} else {
		dst.DistributionWidgetDefinition = nil
	}

	// try to unmarshal data into EventStreamWidgetDefinition
	err = json.Unmarshal(data, &dst.EventStreamWidgetDefinition)
	if err == nil {
		if dst.EventStreamWidgetDefinition != nil && dst.EventStreamWidgetDefinition.UnparsedObject == nil {
			jsonEventStreamWidgetDefinition, _ := json.Marshal(dst.EventStreamWidgetDefinition)
			if string(jsonEventStreamWidgetDefinition) == "{}" { // empty struct
				dst.EventStreamWidgetDefinition = nil
			} else {
				match++
			}
		} else {
			dst.EventStreamWidgetDefinition = nil
		}
	} else {
		dst.EventStreamWidgetDefinition = nil
	}

	// try to unmarshal data into EventTimelineWidgetDefinition
	err = json.Unmarshal(data, &dst.EventTimelineWidgetDefinition)
	if err == nil {
		if dst.EventTimelineWidgetDefinition != nil && dst.EventTimelineWidgetDefinition.UnparsedObject == nil {
			jsonEventTimelineWidgetDefinition, _ := json.Marshal(dst.EventTimelineWidgetDefinition)
			if string(jsonEventTimelineWidgetDefinition) == "{}" { // empty struct
				dst.EventTimelineWidgetDefinition = nil
			} else {
				match++
			}
		} else {
			dst.EventTimelineWidgetDefinition = nil
		}
	} else {
		dst.EventTimelineWidgetDefinition = nil
	}

	// try to unmarshal data into FreeTextWidgetDefinition
	err = json.Unmarshal(data, &dst.FreeTextWidgetDefinition)
	if err == nil {
		if dst.FreeTextWidgetDefinition != nil && dst.FreeTextWidgetDefinition.UnparsedObject == nil {
			jsonFreeTextWidgetDefinition, _ := json.Marshal(dst.FreeTextWidgetDefinition)
			if string(jsonFreeTextWidgetDefinition) == "{}" { // empty struct
				dst.FreeTextWidgetDefinition = nil
			} else {
				match++
			}
		} else {
			dst.FreeTextWidgetDefinition = nil
		}
	} else {
		dst.FreeTextWidgetDefinition = nil
	}

	// try to unmarshal data into GeomapWidgetDefinition
	err = json.Unmarshal(data, &dst.GeomapWidgetDefinition)
	if err == nil {
		if dst.GeomapWidgetDefinition != nil && dst.GeomapWidgetDefinition.UnparsedObject == nil {
			jsonGeomapWidgetDefinition, _ := json.Marshal(dst.GeomapWidgetDefinition)
			if string(jsonGeomapWidgetDefinition) == "{}" { // empty struct
				dst.GeomapWidgetDefinition = nil
			} else {
				match++
			}
		} else {
			dst.GeomapWidgetDefinition = nil
		}
	} else {
		dst.GeomapWidgetDefinition = nil
	}

	// try to unmarshal data into GroupWidgetDefinition
	err = json.Unmarshal(data, &dst.GroupWidgetDefinition)
	if err == nil {
		if dst.GroupWidgetDefinition != nil && dst.GroupWidgetDefinition.UnparsedObject == nil {
			jsonGroupWidgetDefinition, _ := json.Marshal(dst.GroupWidgetDefinition)
			if string(jsonGroupWidgetDefinition) == "{}" { // empty struct
				dst.GroupWidgetDefinition = nil
			} else {
				match++
			}
		} else {
			dst.GroupWidgetDefinition = nil
		}
	} else {
		dst.GroupWidgetDefinition = nil
	}

	// try to unmarshal data into HeatMapWidgetDefinition
	err = json.Unmarshal(data, &dst.HeatMapWidgetDefinition)
	if err == nil {
		if dst.HeatMapWidgetDefinition != nil && dst.HeatMapWidgetDefinition.UnparsedObject == nil {
			jsonHeatMapWidgetDefinition, _ := json.Marshal(dst.HeatMapWidgetDefinition)
			if string(jsonHeatMapWidgetDefinition) == "{}" { // empty struct
				dst.HeatMapWidgetDefinition = nil
			} else {
				match++
			}
		} else {
			dst.HeatMapWidgetDefinition = nil
		}
	} else {
		dst.HeatMapWidgetDefinition = nil
	}

	// try to unmarshal data into HostMapWidgetDefinition
	err = json.Unmarshal(data, &dst.HostMapWidgetDefinition)
	if err == nil {
		if dst.HostMapWidgetDefinition != nil && dst.HostMapWidgetDefinition.UnparsedObject == nil {
			jsonHostMapWidgetDefinition, _ := json.Marshal(dst.HostMapWidgetDefinition)
			if string(jsonHostMapWidgetDefinition) == "{}" { // empty struct
				dst.HostMapWidgetDefinition = nil
			} else {
				match++
			}
		} else {
			dst.HostMapWidgetDefinition = nil
		}
	} else {
		dst.HostMapWidgetDefinition = nil
	}

	// try to unmarshal data into IFrameWidgetDefinition
	err = json.Unmarshal(data, &dst.IFrameWidgetDefinition)
	if err == nil {
		if dst.IFrameWidgetDefinition != nil && dst.IFrameWidgetDefinition.UnparsedObject == nil {
			jsonIFrameWidgetDefinition, _ := json.Marshal(dst.IFrameWidgetDefinition)
			if string(jsonIFrameWidgetDefinition) == "{}" { // empty struct
				dst.IFrameWidgetDefinition = nil
			} else {
				match++
			}
		} else {
			dst.IFrameWidgetDefinition = nil
		}
	} else {
		dst.IFrameWidgetDefinition = nil
	}

	// try to unmarshal data into ImageWidgetDefinition
	err = json.Unmarshal(data, &dst.ImageWidgetDefinition)
	if err == nil {
		if dst.ImageWidgetDefinition != nil && dst.ImageWidgetDefinition.UnparsedObject == nil {
			jsonImageWidgetDefinition, _ := json.Marshal(dst.ImageWidgetDefinition)
			if string(jsonImageWidgetDefinition) == "{}" { // empty struct
				dst.ImageWidgetDefinition = nil
			} else {
				match++
			}
		} else {
			dst.ImageWidgetDefinition = nil
		}
	} else {
		dst.ImageWidgetDefinition = nil
	}

	// try to unmarshal data into LogStreamWidgetDefinition
	err = json.Unmarshal(data, &dst.LogStreamWidgetDefinition)
	if err == nil {
		if dst.LogStreamWidgetDefinition != nil && dst.LogStreamWidgetDefinition.UnparsedObject == nil {
			jsonLogStreamWidgetDefinition, _ := json.Marshal(dst.LogStreamWidgetDefinition)
			if string(jsonLogStreamWidgetDefinition) == "{}" { // empty struct
				dst.LogStreamWidgetDefinition = nil
			} else {
				match++
			}
		} else {
			dst.LogStreamWidgetDefinition = nil
		}
	} else {
		dst.LogStreamWidgetDefinition = nil
	}

	// try to unmarshal data into MonitorSummaryWidgetDefinition
	err = json.Unmarshal(data, &dst.MonitorSummaryWidgetDefinition)
	if err == nil {
		if dst.MonitorSummaryWidgetDefinition != nil && dst.MonitorSummaryWidgetDefinition.UnparsedObject == nil {
			jsonMonitorSummaryWidgetDefinition, _ := json.Marshal(dst.MonitorSummaryWidgetDefinition)
			if string(jsonMonitorSummaryWidgetDefinition) == "{}" { // empty struct
				dst.MonitorSummaryWidgetDefinition = nil
			} else {
				match++
			}
		} else {
			dst.MonitorSummaryWidgetDefinition = nil
		}
	} else {
		dst.MonitorSummaryWidgetDefinition = nil
	}

	// try to unmarshal data into NoteWidgetDefinition
	err = json.Unmarshal(data, &dst.NoteWidgetDefinition)
	if err == nil {
		if dst.NoteWidgetDefinition != nil && dst.NoteWidgetDefinition.UnparsedObject == nil {
			jsonNoteWidgetDefinition, _ := json.Marshal(dst.NoteWidgetDefinition)
			if string(jsonNoteWidgetDefinition) == "{}" { // empty struct
				dst.NoteWidgetDefinition = nil
			} else {
				match++
			}
		} else {
			dst.NoteWidgetDefinition = nil
		}
	} else {
		dst.NoteWidgetDefinition = nil
	}

	// try to unmarshal data into QueryValueWidgetDefinition
	err = json.Unmarshal(data, &dst.QueryValueWidgetDefinition)
	if err == nil {
		if dst.QueryValueWidgetDefinition != nil && dst.QueryValueWidgetDefinition.UnparsedObject == nil {
			jsonQueryValueWidgetDefinition, _ := json.Marshal(dst.QueryValueWidgetDefinition)
			if string(jsonQueryValueWidgetDefinition) == "{}" { // empty struct
				dst.QueryValueWidgetDefinition = nil
			} else {
				match++
			}
		} else {
			dst.QueryValueWidgetDefinition = nil
		}
	} else {
		dst.QueryValueWidgetDefinition = nil
	}

	// try to unmarshal data into ScatterPlotWidgetDefinition
	err = json.Unmarshal(data, &dst.ScatterPlotWidgetDefinition)
	if err == nil {
		if dst.ScatterPlotWidgetDefinition != nil && dst.ScatterPlotWidgetDefinition.UnparsedObject == nil {
			jsonScatterPlotWidgetDefinition, _ := json.Marshal(dst.ScatterPlotWidgetDefinition)
			if string(jsonScatterPlotWidgetDefinition) == "{}" { // empty struct
				dst.ScatterPlotWidgetDefinition = nil
			} else {
				match++
			}
		} else {
			dst.ScatterPlotWidgetDefinition = nil
		}
	} else {
		dst.ScatterPlotWidgetDefinition = nil
	}

	// try to unmarshal data into SLOWidgetDefinition
	err = json.Unmarshal(data, &dst.SLOWidgetDefinition)
	if err == nil {
		if dst.SLOWidgetDefinition != nil && dst.SLOWidgetDefinition.UnparsedObject == nil {
			jsonSLOWidgetDefinition, _ := json.Marshal(dst.SLOWidgetDefinition)
			if string(jsonSLOWidgetDefinition) == "{}" { // empty struct
				dst.SLOWidgetDefinition = nil
			} else {
				match++
			}
		} else {
			dst.SLOWidgetDefinition = nil
		}
	} else {
		dst.SLOWidgetDefinition = nil
	}

	// try to unmarshal data into ServiceMapWidgetDefinition
	err = json.Unmarshal(data, &dst.ServiceMapWidgetDefinition)
	if err == nil {
		if dst.ServiceMapWidgetDefinition != nil && dst.ServiceMapWidgetDefinition.UnparsedObject == nil {
			jsonServiceMapWidgetDefinition, _ := json.Marshal(dst.ServiceMapWidgetDefinition)
			if string(jsonServiceMapWidgetDefinition) == "{}" { // empty struct
				dst.ServiceMapWidgetDefinition = nil
			} else {
				match++
			}
		} else {
			dst.ServiceMapWidgetDefinition = nil
		}
	} else {
		dst.ServiceMapWidgetDefinition = nil
	}

	// try to unmarshal data into ServiceSummaryWidgetDefinition
	err = json.Unmarshal(data, &dst.ServiceSummaryWidgetDefinition)
	if err == nil {
		if dst.ServiceSummaryWidgetDefinition != nil && dst.ServiceSummaryWidgetDefinition.UnparsedObject == nil {
			jsonServiceSummaryWidgetDefinition, _ := json.Marshal(dst.ServiceSummaryWidgetDefinition)
			if string(jsonServiceSummaryWidgetDefinition) == "{}" { // empty struct
				dst.ServiceSummaryWidgetDefinition = nil
			} else {
				match++
			}
		} else {
			dst.ServiceSummaryWidgetDefinition = nil
		}
	} else {
		dst.ServiceSummaryWidgetDefinition = nil
	}

	// try to unmarshal data into SunburstWidgetDefinition
	err = json.Unmarshal(data, &dst.SunburstWidgetDefinition)
	if err == nil {
		if dst.SunburstWidgetDefinition != nil && dst.SunburstWidgetDefinition.UnparsedObject == nil {
			jsonSunburstWidgetDefinition, _ := json.Marshal(dst.SunburstWidgetDefinition)
			if string(jsonSunburstWidgetDefinition) == "{}" { // empty struct
				dst.SunburstWidgetDefinition = nil
			} else {
				match++
			}
		} else {
			dst.SunburstWidgetDefinition = nil
		}
	} else {
		dst.SunburstWidgetDefinition = nil
	}

	// try to unmarshal data into TableWidgetDefinition
	err = json.Unmarshal(data, &dst.TableWidgetDefinition)
	if err == nil {
		if dst.TableWidgetDefinition != nil && dst.TableWidgetDefinition.UnparsedObject == nil {
			jsonTableWidgetDefinition, _ := json.Marshal(dst.TableWidgetDefinition)
			if string(jsonTableWidgetDefinition) == "{}" { // empty struct
				dst.TableWidgetDefinition = nil
			} else {
				match++
			}
		} else {
			dst.TableWidgetDefinition = nil
		}
	} else {
		dst.TableWidgetDefinition = nil
	}

	// try to unmarshal data into TimeseriesWidgetDefinition
	err = json.Unmarshal(data, &dst.TimeseriesWidgetDefinition)
	if err == nil {
		if dst.TimeseriesWidgetDefinition != nil && dst.TimeseriesWidgetDefinition.UnparsedObject == nil {
			jsonTimeseriesWidgetDefinition, _ := json.Marshal(dst.TimeseriesWidgetDefinition)
			if string(jsonTimeseriesWidgetDefinition) == "{}" { // empty struct
				dst.TimeseriesWidgetDefinition = nil
			} else {
				match++
			}
		} else {
			dst.TimeseriesWidgetDefinition = nil
		}
	} else {
		dst.TimeseriesWidgetDefinition = nil
	}

	// try to unmarshal data into ToplistWidgetDefinition
	err = json.Unmarshal(data, &dst.ToplistWidgetDefinition)
	if err == nil {
		if dst.ToplistWidgetDefinition != nil && dst.ToplistWidgetDefinition.UnparsedObject == nil {
			jsonToplistWidgetDefinition, _ := json.Marshal(dst.ToplistWidgetDefinition)
			if string(jsonToplistWidgetDefinition) == "{}" { // empty struct
				dst.ToplistWidgetDefinition = nil
			} else {
				match++
			}
		} else {
			dst.ToplistWidgetDefinition = nil
		}
	} else {
		dst.ToplistWidgetDefinition = nil
	}

	// try to unmarshal data into TreeMapWidgetDefinition
	err = json.Unmarshal(data, &dst.TreeMapWidgetDefinition)
	if err == nil {
		if dst.TreeMapWidgetDefinition != nil && dst.TreeMapWidgetDefinition.UnparsedObject == nil {
			jsonTreeMapWidgetDefinition, _ := json.Marshal(dst.TreeMapWidgetDefinition)
			if string(jsonTreeMapWidgetDefinition) == "{}" { // empty struct
				dst.TreeMapWidgetDefinition = nil
			} else {
				match++
			}
		} else {
			dst.TreeMapWidgetDefinition = nil
		}
	} else {
		dst.TreeMapWidgetDefinition = nil
	}

	// try to unmarshal data into ListStreamWidgetDefinition
	err = json.Unmarshal(data, &dst.ListStreamWidgetDefinition)
	if err == nil {
		if dst.ListStreamWidgetDefinition != nil && dst.ListStreamWidgetDefinition.UnparsedObject == nil {
			jsonListStreamWidgetDefinition, _ := json.Marshal(dst.ListStreamWidgetDefinition)
			if string(jsonListStreamWidgetDefinition) == "{}" { // empty struct
				dst.ListStreamWidgetDefinition = nil
			} else {
				match++
			}
		} else {
			dst.ListStreamWidgetDefinition = nil
		}
	} else {
		dst.ListStreamWidgetDefinition = nil
	}

	// try to unmarshal data into FunnelWidgetDefinition
	err = json.Unmarshal(data, &dst.FunnelWidgetDefinition)
	if err == nil {
		if dst.FunnelWidgetDefinition != nil && dst.FunnelWidgetDefinition.UnparsedObject == nil {
			jsonFunnelWidgetDefinition, _ := json.Marshal(dst.FunnelWidgetDefinition)
			if string(jsonFunnelWidgetDefinition) == "{}" { // empty struct
				dst.FunnelWidgetDefinition = nil
			} else {
				match++
			}
		} else {
			dst.FunnelWidgetDefinition = nil
		}
	} else {
		dst.FunnelWidgetDefinition = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		dst.AlertGraphWidgetDefinition = nil
		dst.AlertValueWidgetDefinition = nil
		dst.ChangeWidgetDefinition = nil
		dst.CheckStatusWidgetDefinition = nil
		dst.DistributionWidgetDefinition = nil
		dst.EventStreamWidgetDefinition = nil
		dst.EventTimelineWidgetDefinition = nil
		dst.FreeTextWidgetDefinition = nil
		dst.GeomapWidgetDefinition = nil
		dst.GroupWidgetDefinition = nil
		dst.HeatMapWidgetDefinition = nil
		dst.HostMapWidgetDefinition = nil
		dst.IFrameWidgetDefinition = nil
		dst.ImageWidgetDefinition = nil
		dst.LogStreamWidgetDefinition = nil
		dst.MonitorSummaryWidgetDefinition = nil
		dst.NoteWidgetDefinition = nil
		dst.QueryValueWidgetDefinition = nil
		dst.ScatterPlotWidgetDefinition = nil
		dst.SLOWidgetDefinition = nil
		dst.ServiceMapWidgetDefinition = nil
		dst.ServiceSummaryWidgetDefinition = nil
		dst.SunburstWidgetDefinition = nil
		dst.TableWidgetDefinition = nil
		dst.TimeseriesWidgetDefinition = nil
		dst.ToplistWidgetDefinition = nil
		dst.TreeMapWidgetDefinition = nil
		dst.ListStreamWidgetDefinition = nil
		dst.FunnelWidgetDefinition = nil
		return json.Unmarshal(data, &dst.UnparsedObject)
	} else {
		return nil // exactly one match
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src WidgetDefinition) MarshalJSON() ([]byte, error) {
	if src.AlertGraphWidgetDefinition != nil {
		return json.Marshal(&src.AlertGraphWidgetDefinition)
	}

	if src.AlertValueWidgetDefinition != nil {
		return json.Marshal(&src.AlertValueWidgetDefinition)
	}

	if src.ChangeWidgetDefinition != nil {
		return json.Marshal(&src.ChangeWidgetDefinition)
	}

	if src.CheckStatusWidgetDefinition != nil {
		return json.Marshal(&src.CheckStatusWidgetDefinition)
	}

	if src.DistributionWidgetDefinition != nil {
		return json.Marshal(&src.DistributionWidgetDefinition)
	}

	if src.EventStreamWidgetDefinition != nil {
		return json.Marshal(&src.EventStreamWidgetDefinition)
	}

	if src.EventTimelineWidgetDefinition != nil {
		return json.Marshal(&src.EventTimelineWidgetDefinition)
	}

	if src.FreeTextWidgetDefinition != nil {
		return json.Marshal(&src.FreeTextWidgetDefinition)
	}

	if src.GeomapWidgetDefinition != nil {
		return json.Marshal(&src.GeomapWidgetDefinition)
	}

	if src.GroupWidgetDefinition != nil {
		return json.Marshal(&src.GroupWidgetDefinition)
	}

	if src.HeatMapWidgetDefinition != nil {
		return json.Marshal(&src.HeatMapWidgetDefinition)
	}

	if src.HostMapWidgetDefinition != nil {
		return json.Marshal(&src.HostMapWidgetDefinition)
	}

	if src.IFrameWidgetDefinition != nil {
		return json.Marshal(&src.IFrameWidgetDefinition)
	}

	if src.ImageWidgetDefinition != nil {
		return json.Marshal(&src.ImageWidgetDefinition)
	}

	if src.LogStreamWidgetDefinition != nil {
		return json.Marshal(&src.LogStreamWidgetDefinition)
	}

	if src.MonitorSummaryWidgetDefinition != nil {
		return json.Marshal(&src.MonitorSummaryWidgetDefinition)
	}

	if src.NoteWidgetDefinition != nil {
		return json.Marshal(&src.NoteWidgetDefinition)
	}

	if src.QueryValueWidgetDefinition != nil {
		return json.Marshal(&src.QueryValueWidgetDefinition)
	}

	if src.ScatterPlotWidgetDefinition != nil {
		return json.Marshal(&src.ScatterPlotWidgetDefinition)
	}

	if src.SLOWidgetDefinition != nil {
		return json.Marshal(&src.SLOWidgetDefinition)
	}

	if src.ServiceMapWidgetDefinition != nil {
		return json.Marshal(&src.ServiceMapWidgetDefinition)
	}

	if src.ServiceSummaryWidgetDefinition != nil {
		return json.Marshal(&src.ServiceSummaryWidgetDefinition)
	}

	if src.SunburstWidgetDefinition != nil {
		return json.Marshal(&src.SunburstWidgetDefinition)
	}

	if src.TableWidgetDefinition != nil {
		return json.Marshal(&src.TableWidgetDefinition)
	}

	if src.TimeseriesWidgetDefinition != nil {
		return json.Marshal(&src.TimeseriesWidgetDefinition)
	}

	if src.ToplistWidgetDefinition != nil {
		return json.Marshal(&src.ToplistWidgetDefinition)
	}

	if src.TreeMapWidgetDefinition != nil {
		return json.Marshal(&src.TreeMapWidgetDefinition)
	}

	if src.ListStreamWidgetDefinition != nil {
		return json.Marshal(&src.ListStreamWidgetDefinition)
	}

	if src.FunnelWidgetDefinition != nil {
		return json.Marshal(&src.FunnelWidgetDefinition)
	}

	if src.UnparsedObject != nil {
		return json.Marshal(src.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *WidgetDefinition) GetActualInstance() interface{} {
	if obj.AlertGraphWidgetDefinition != nil {
		return obj.AlertGraphWidgetDefinition
	}

	if obj.AlertValueWidgetDefinition != nil {
		return obj.AlertValueWidgetDefinition
	}

	if obj.ChangeWidgetDefinition != nil {
		return obj.ChangeWidgetDefinition
	}

	if obj.CheckStatusWidgetDefinition != nil {
		return obj.CheckStatusWidgetDefinition
	}

	if obj.DistributionWidgetDefinition != nil {
		return obj.DistributionWidgetDefinition
	}

	if obj.EventStreamWidgetDefinition != nil {
		return obj.EventStreamWidgetDefinition
	}

	if obj.EventTimelineWidgetDefinition != nil {
		return obj.EventTimelineWidgetDefinition
	}

	if obj.FreeTextWidgetDefinition != nil {
		return obj.FreeTextWidgetDefinition
	}

	if obj.GeomapWidgetDefinition != nil {
		return obj.GeomapWidgetDefinition
	}

	if obj.GroupWidgetDefinition != nil {
		return obj.GroupWidgetDefinition
	}

	if obj.HeatMapWidgetDefinition != nil {
		return obj.HeatMapWidgetDefinition
	}

	if obj.HostMapWidgetDefinition != nil {
		return obj.HostMapWidgetDefinition
	}

	if obj.IFrameWidgetDefinition != nil {
		return obj.IFrameWidgetDefinition
	}

	if obj.ImageWidgetDefinition != nil {
		return obj.ImageWidgetDefinition
	}

	if obj.LogStreamWidgetDefinition != nil {
		return obj.LogStreamWidgetDefinition
	}

	if obj.MonitorSummaryWidgetDefinition != nil {
		return obj.MonitorSummaryWidgetDefinition
	}

	if obj.NoteWidgetDefinition != nil {
		return obj.NoteWidgetDefinition
	}

	if obj.QueryValueWidgetDefinition != nil {
		return obj.QueryValueWidgetDefinition
	}

	if obj.ScatterPlotWidgetDefinition != nil {
		return obj.ScatterPlotWidgetDefinition
	}

	if obj.SLOWidgetDefinition != nil {
		return obj.SLOWidgetDefinition
	}

	if obj.ServiceMapWidgetDefinition != nil {
		return obj.ServiceMapWidgetDefinition
	}

	if obj.ServiceSummaryWidgetDefinition != nil {
		return obj.ServiceSummaryWidgetDefinition
	}

	if obj.SunburstWidgetDefinition != nil {
		return obj.SunburstWidgetDefinition
	}

	if obj.TableWidgetDefinition != nil {
		return obj.TableWidgetDefinition
	}

	if obj.TimeseriesWidgetDefinition != nil {
		return obj.TimeseriesWidgetDefinition
	}

	if obj.ToplistWidgetDefinition != nil {
		return obj.ToplistWidgetDefinition
	}

	if obj.TreeMapWidgetDefinition != nil {
		return obj.TreeMapWidgetDefinition
	}

	if obj.ListStreamWidgetDefinition != nil {
		return obj.ListStreamWidgetDefinition
	}

	if obj.FunnelWidgetDefinition != nil {
		return obj.FunnelWidgetDefinition
	}

	// all schemas are nil
	return nil
}

type NullableWidgetDefinition struct {
	value *WidgetDefinition
	isSet bool
}

func (v NullableWidgetDefinition) Get() *WidgetDefinition {
	return v.value
}

func (v *NullableWidgetDefinition) Set(val *WidgetDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetDefinition(val *WidgetDefinition) *NullableWidgetDefinition {
	return &NullableWidgetDefinition{value: val, isSet: true}
}

func (v NullableWidgetDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return json.Unmarshal(src, &v.value)
}
