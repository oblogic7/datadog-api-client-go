/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// UsageBillableSummaryKeys Response with aggregated usage types.
type UsageBillableSummaryKeys struct {
	// Response with properties for each aggregated usage type.
	ApmHostSum *UsageBillableSummaryBody `json:"apm_host_sum,omitempty"`
	// Response with properties for each aggregated usage type.
	ApmHostTop99p *UsageBillableSummaryBody `json:"apm_host_top99p,omitempty"`
	// Response with properties for each aggregated usage type.
	ApmTraceSearchSum *UsageBillableSummaryBody `json:"apm_trace_search_sum,omitempty"`
	// Response with properties for each aggregated usage type.
	FargateContainerAverage *UsageBillableSummaryBody `json:"fargate_container_average,omitempty"`
	// Response with properties for each aggregated usage type.
	InfraContainerSum *UsageBillableSummaryBody `json:"infra_container_sum,omitempty"`
	// Response with properties for each aggregated usage type.
	InfraHostSum *UsageBillableSummaryBody `json:"infra_host_sum,omitempty"`
	// Response with properties for each aggregated usage type.
	InfraHostTop99p *UsageBillableSummaryBody `json:"infra_host_top99p,omitempty"`
	// Response with properties for each aggregated usage type.
	IotTop99p *UsageBillableSummaryBody `json:"iot_top99p,omitempty"`
	// Response with properties for each aggregated usage type.
	LambdaFunctionAverage *UsageBillableSummaryBody `json:"lambda_function_average,omitempty"`
	// Response with properties for each aggregated usage type.
	LogsIndexed15daySum *UsageBillableSummaryBody `json:"logs_indexed_15day_sum,omitempty"`
	// Response with properties for each aggregated usage type.
	LogsIndexed180daySum *UsageBillableSummaryBody `json:"logs_indexed_180day_sum,omitempty"`
	// Response with properties for each aggregated usage type.
	LogsIndexed30daySum *UsageBillableSummaryBody `json:"logs_indexed_30day_sum,omitempty"`
	// Response with properties for each aggregated usage type.
	LogsIndexed3daySum *UsageBillableSummaryBody `json:"logs_indexed_3day_sum,omitempty"`
	// Response with properties for each aggregated usage type.
	LogsIndexed45daySum *UsageBillableSummaryBody `json:"logs_indexed_45day_sum,omitempty"`
	// Response with properties for each aggregated usage type.
	LogsIndexed60daySum *UsageBillableSummaryBody `json:"logs_indexed_60day_sum,omitempty"`
	// Response with properties for each aggregated usage type.
	LogsIndexed7daySum *UsageBillableSummaryBody `json:"logs_indexed_7day_sum,omitempty"`
	// Response with properties for each aggregated usage type.
	LogsIndexed90daySum *UsageBillableSummaryBody `json:"logs_indexed_90day_sum,omitempty"`
	// Response with properties for each aggregated usage type.
	LogsIndexedCustomRetentionSum *UsageBillableSummaryBody `json:"logs_indexed_custom_retention_sum,omitempty"`
	// Response with properties for each aggregated usage type.
	LogsIndexedSum *UsageBillableSummaryBody `json:"logs_indexed_sum,omitempty"`
	// Response with properties for each aggregated usage type.
	LogsIngestedSum *UsageBillableSummaryBody `json:"logs_ingested_sum,omitempty"`
	// Response with properties for each aggregated usage type.
	NetworkDeviceTop99p *UsageBillableSummaryBody `json:"network_device_top99p,omitempty"`
	// Response with properties for each aggregated usage type.
	NpmFlowSum *UsageBillableSummaryBody `json:"npm_flow_sum,omitempty"`
	// Response with properties for each aggregated usage type.
	NpmHostSum *UsageBillableSummaryBody `json:"npm_host_sum,omitempty"`
	// Response with properties for each aggregated usage type.
	NpmHostTop99p *UsageBillableSummaryBody `json:"npm_host_top99p,omitempty"`
	// Response with properties for each aggregated usage type.
	ProfContainerSum *UsageBillableSummaryBody `json:"prof_container_sum,omitempty"`
	// Response with properties for each aggregated usage type.
	ProfHostTop99p *UsageBillableSummaryBody `json:"prof_host_top99p,omitempty"`
	// Response with properties for each aggregated usage type.
	RumSum *UsageBillableSummaryBody `json:"rum_sum,omitempty"`
	// Response with properties for each aggregated usage type.
	ServerlessInvocationSum *UsageBillableSummaryBody `json:"serverless_invocation_sum,omitempty"`
	// Response with properties for each aggregated usage type.
	SiemSum *UsageBillableSummaryBody `json:"siem_sum,omitempty"`
	// Response with properties for each aggregated usage type.
	SyntheticsApiTestsSum *UsageBillableSummaryBody `json:"synthetics_api_tests_sum,omitempty"`
	// Response with properties for each aggregated usage type.
	SyntheticsBrowserChecksSum *UsageBillableSummaryBody `json:"synthetics_browser_checks_sum,omitempty"`
	// Response with properties for each aggregated usage type.
	TimeseriesAverage *UsageBillableSummaryBody `json:"timeseries_average,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewUsageBillableSummaryKeys instantiates a new UsageBillableSummaryKeys object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageBillableSummaryKeys() *UsageBillableSummaryKeys {
	this := UsageBillableSummaryKeys{}
	return &this
}

// NewUsageBillableSummaryKeysWithDefaults instantiates a new UsageBillableSummaryKeys object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageBillableSummaryKeysWithDefaults() *UsageBillableSummaryKeys {
	this := UsageBillableSummaryKeys{}
	return &this
}

// GetApmHostSum returns the ApmHostSum field value if set, zero value otherwise.
func (o *UsageBillableSummaryKeys) GetApmHostSum() UsageBillableSummaryBody {
	if o == nil || o.ApmHostSum == nil {
		var ret UsageBillableSummaryBody
		return ret
	}
	return *o.ApmHostSum
}

// GetApmHostSumOk returns a tuple with the ApmHostSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryKeys) GetApmHostSumOk() (*UsageBillableSummaryBody, bool) {
	if o == nil || o.ApmHostSum == nil {
		return nil, false
	}
	return o.ApmHostSum, true
}

// HasApmHostSum returns a boolean if a field has been set.
func (o *UsageBillableSummaryKeys) HasApmHostSum() bool {
	if o != nil && o.ApmHostSum != nil {
		return true
	}

	return false
}

// SetApmHostSum gets a reference to the given UsageBillableSummaryBody and assigns it to the ApmHostSum field.
func (o *UsageBillableSummaryKeys) SetApmHostSum(v UsageBillableSummaryBody) {
	o.ApmHostSum = &v
}

// GetApmHostTop99p returns the ApmHostTop99p field value if set, zero value otherwise.
func (o *UsageBillableSummaryKeys) GetApmHostTop99p() UsageBillableSummaryBody {
	if o == nil || o.ApmHostTop99p == nil {
		var ret UsageBillableSummaryBody
		return ret
	}
	return *o.ApmHostTop99p
}

// GetApmHostTop99pOk returns a tuple with the ApmHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryKeys) GetApmHostTop99pOk() (*UsageBillableSummaryBody, bool) {
	if o == nil || o.ApmHostTop99p == nil {
		return nil, false
	}
	return o.ApmHostTop99p, true
}

// HasApmHostTop99p returns a boolean if a field has been set.
func (o *UsageBillableSummaryKeys) HasApmHostTop99p() bool {
	if o != nil && o.ApmHostTop99p != nil {
		return true
	}

	return false
}

// SetApmHostTop99p gets a reference to the given UsageBillableSummaryBody and assigns it to the ApmHostTop99p field.
func (o *UsageBillableSummaryKeys) SetApmHostTop99p(v UsageBillableSummaryBody) {
	o.ApmHostTop99p = &v
}

// GetApmTraceSearchSum returns the ApmTraceSearchSum field value if set, zero value otherwise.
func (o *UsageBillableSummaryKeys) GetApmTraceSearchSum() UsageBillableSummaryBody {
	if o == nil || o.ApmTraceSearchSum == nil {
		var ret UsageBillableSummaryBody
		return ret
	}
	return *o.ApmTraceSearchSum
}

// GetApmTraceSearchSumOk returns a tuple with the ApmTraceSearchSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryKeys) GetApmTraceSearchSumOk() (*UsageBillableSummaryBody, bool) {
	if o == nil || o.ApmTraceSearchSum == nil {
		return nil, false
	}
	return o.ApmTraceSearchSum, true
}

// HasApmTraceSearchSum returns a boolean if a field has been set.
func (o *UsageBillableSummaryKeys) HasApmTraceSearchSum() bool {
	if o != nil && o.ApmTraceSearchSum != nil {
		return true
	}

	return false
}

// SetApmTraceSearchSum gets a reference to the given UsageBillableSummaryBody and assigns it to the ApmTraceSearchSum field.
func (o *UsageBillableSummaryKeys) SetApmTraceSearchSum(v UsageBillableSummaryBody) {
	o.ApmTraceSearchSum = &v
}

// GetFargateContainerAverage returns the FargateContainerAverage field value if set, zero value otherwise.
func (o *UsageBillableSummaryKeys) GetFargateContainerAverage() UsageBillableSummaryBody {
	if o == nil || o.FargateContainerAverage == nil {
		var ret UsageBillableSummaryBody
		return ret
	}
	return *o.FargateContainerAverage
}

// GetFargateContainerAverageOk returns a tuple with the FargateContainerAverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryKeys) GetFargateContainerAverageOk() (*UsageBillableSummaryBody, bool) {
	if o == nil || o.FargateContainerAverage == nil {
		return nil, false
	}
	return o.FargateContainerAverage, true
}

// HasFargateContainerAverage returns a boolean if a field has been set.
func (o *UsageBillableSummaryKeys) HasFargateContainerAverage() bool {
	if o != nil && o.FargateContainerAverage != nil {
		return true
	}

	return false
}

// SetFargateContainerAverage gets a reference to the given UsageBillableSummaryBody and assigns it to the FargateContainerAverage field.
func (o *UsageBillableSummaryKeys) SetFargateContainerAverage(v UsageBillableSummaryBody) {
	o.FargateContainerAverage = &v
}

// GetInfraContainerSum returns the InfraContainerSum field value if set, zero value otherwise.
func (o *UsageBillableSummaryKeys) GetInfraContainerSum() UsageBillableSummaryBody {
	if o == nil || o.InfraContainerSum == nil {
		var ret UsageBillableSummaryBody
		return ret
	}
	return *o.InfraContainerSum
}

// GetInfraContainerSumOk returns a tuple with the InfraContainerSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryKeys) GetInfraContainerSumOk() (*UsageBillableSummaryBody, bool) {
	if o == nil || o.InfraContainerSum == nil {
		return nil, false
	}
	return o.InfraContainerSum, true
}

// HasInfraContainerSum returns a boolean if a field has been set.
func (o *UsageBillableSummaryKeys) HasInfraContainerSum() bool {
	if o != nil && o.InfraContainerSum != nil {
		return true
	}

	return false
}

// SetInfraContainerSum gets a reference to the given UsageBillableSummaryBody and assigns it to the InfraContainerSum field.
func (o *UsageBillableSummaryKeys) SetInfraContainerSum(v UsageBillableSummaryBody) {
	o.InfraContainerSum = &v
}

// GetInfraHostSum returns the InfraHostSum field value if set, zero value otherwise.
func (o *UsageBillableSummaryKeys) GetInfraHostSum() UsageBillableSummaryBody {
	if o == nil || o.InfraHostSum == nil {
		var ret UsageBillableSummaryBody
		return ret
	}
	return *o.InfraHostSum
}

// GetInfraHostSumOk returns a tuple with the InfraHostSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryKeys) GetInfraHostSumOk() (*UsageBillableSummaryBody, bool) {
	if o == nil || o.InfraHostSum == nil {
		return nil, false
	}
	return o.InfraHostSum, true
}

// HasInfraHostSum returns a boolean if a field has been set.
func (o *UsageBillableSummaryKeys) HasInfraHostSum() bool {
	if o != nil && o.InfraHostSum != nil {
		return true
	}

	return false
}

// SetInfraHostSum gets a reference to the given UsageBillableSummaryBody and assigns it to the InfraHostSum field.
func (o *UsageBillableSummaryKeys) SetInfraHostSum(v UsageBillableSummaryBody) {
	o.InfraHostSum = &v
}

// GetInfraHostTop99p returns the InfraHostTop99p field value if set, zero value otherwise.
func (o *UsageBillableSummaryKeys) GetInfraHostTop99p() UsageBillableSummaryBody {
	if o == nil || o.InfraHostTop99p == nil {
		var ret UsageBillableSummaryBody
		return ret
	}
	return *o.InfraHostTop99p
}

// GetInfraHostTop99pOk returns a tuple with the InfraHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryKeys) GetInfraHostTop99pOk() (*UsageBillableSummaryBody, bool) {
	if o == nil || o.InfraHostTop99p == nil {
		return nil, false
	}
	return o.InfraHostTop99p, true
}

// HasInfraHostTop99p returns a boolean if a field has been set.
func (o *UsageBillableSummaryKeys) HasInfraHostTop99p() bool {
	if o != nil && o.InfraHostTop99p != nil {
		return true
	}

	return false
}

// SetInfraHostTop99p gets a reference to the given UsageBillableSummaryBody and assigns it to the InfraHostTop99p field.
func (o *UsageBillableSummaryKeys) SetInfraHostTop99p(v UsageBillableSummaryBody) {
	o.InfraHostTop99p = &v
}

// GetIotTop99p returns the IotTop99p field value if set, zero value otherwise.
func (o *UsageBillableSummaryKeys) GetIotTop99p() UsageBillableSummaryBody {
	if o == nil || o.IotTop99p == nil {
		var ret UsageBillableSummaryBody
		return ret
	}
	return *o.IotTop99p
}

// GetIotTop99pOk returns a tuple with the IotTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryKeys) GetIotTop99pOk() (*UsageBillableSummaryBody, bool) {
	if o == nil || o.IotTop99p == nil {
		return nil, false
	}
	return o.IotTop99p, true
}

// HasIotTop99p returns a boolean if a field has been set.
func (o *UsageBillableSummaryKeys) HasIotTop99p() bool {
	if o != nil && o.IotTop99p != nil {
		return true
	}

	return false
}

// SetIotTop99p gets a reference to the given UsageBillableSummaryBody and assigns it to the IotTop99p field.
func (o *UsageBillableSummaryKeys) SetIotTop99p(v UsageBillableSummaryBody) {
	o.IotTop99p = &v
}

// GetLambdaFunctionAverage returns the LambdaFunctionAverage field value if set, zero value otherwise.
func (o *UsageBillableSummaryKeys) GetLambdaFunctionAverage() UsageBillableSummaryBody {
	if o == nil || o.LambdaFunctionAverage == nil {
		var ret UsageBillableSummaryBody
		return ret
	}
	return *o.LambdaFunctionAverage
}

// GetLambdaFunctionAverageOk returns a tuple with the LambdaFunctionAverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryKeys) GetLambdaFunctionAverageOk() (*UsageBillableSummaryBody, bool) {
	if o == nil || o.LambdaFunctionAverage == nil {
		return nil, false
	}
	return o.LambdaFunctionAverage, true
}

// HasLambdaFunctionAverage returns a boolean if a field has been set.
func (o *UsageBillableSummaryKeys) HasLambdaFunctionAverage() bool {
	if o != nil && o.LambdaFunctionAverage != nil {
		return true
	}

	return false
}

// SetLambdaFunctionAverage gets a reference to the given UsageBillableSummaryBody and assigns it to the LambdaFunctionAverage field.
func (o *UsageBillableSummaryKeys) SetLambdaFunctionAverage(v UsageBillableSummaryBody) {
	o.LambdaFunctionAverage = &v
}

// GetLogsIndexed15daySum returns the LogsIndexed15daySum field value if set, zero value otherwise.
func (o *UsageBillableSummaryKeys) GetLogsIndexed15daySum() UsageBillableSummaryBody {
	if o == nil || o.LogsIndexed15daySum == nil {
		var ret UsageBillableSummaryBody
		return ret
	}
	return *o.LogsIndexed15daySum
}

// GetLogsIndexed15daySumOk returns a tuple with the LogsIndexed15daySum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryKeys) GetLogsIndexed15daySumOk() (*UsageBillableSummaryBody, bool) {
	if o == nil || o.LogsIndexed15daySum == nil {
		return nil, false
	}
	return o.LogsIndexed15daySum, true
}

// HasLogsIndexed15daySum returns a boolean if a field has been set.
func (o *UsageBillableSummaryKeys) HasLogsIndexed15daySum() bool {
	if o != nil && o.LogsIndexed15daySum != nil {
		return true
	}

	return false
}

// SetLogsIndexed15daySum gets a reference to the given UsageBillableSummaryBody and assigns it to the LogsIndexed15daySum field.
func (o *UsageBillableSummaryKeys) SetLogsIndexed15daySum(v UsageBillableSummaryBody) {
	o.LogsIndexed15daySum = &v
}

// GetLogsIndexed180daySum returns the LogsIndexed180daySum field value if set, zero value otherwise.
func (o *UsageBillableSummaryKeys) GetLogsIndexed180daySum() UsageBillableSummaryBody {
	if o == nil || o.LogsIndexed180daySum == nil {
		var ret UsageBillableSummaryBody
		return ret
	}
	return *o.LogsIndexed180daySum
}

// GetLogsIndexed180daySumOk returns a tuple with the LogsIndexed180daySum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryKeys) GetLogsIndexed180daySumOk() (*UsageBillableSummaryBody, bool) {
	if o == nil || o.LogsIndexed180daySum == nil {
		return nil, false
	}
	return o.LogsIndexed180daySum, true
}

// HasLogsIndexed180daySum returns a boolean if a field has been set.
func (o *UsageBillableSummaryKeys) HasLogsIndexed180daySum() bool {
	if o != nil && o.LogsIndexed180daySum != nil {
		return true
	}

	return false
}

// SetLogsIndexed180daySum gets a reference to the given UsageBillableSummaryBody and assigns it to the LogsIndexed180daySum field.
func (o *UsageBillableSummaryKeys) SetLogsIndexed180daySum(v UsageBillableSummaryBody) {
	o.LogsIndexed180daySum = &v
}

// GetLogsIndexed30daySum returns the LogsIndexed30daySum field value if set, zero value otherwise.
func (o *UsageBillableSummaryKeys) GetLogsIndexed30daySum() UsageBillableSummaryBody {
	if o == nil || o.LogsIndexed30daySum == nil {
		var ret UsageBillableSummaryBody
		return ret
	}
	return *o.LogsIndexed30daySum
}

// GetLogsIndexed30daySumOk returns a tuple with the LogsIndexed30daySum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryKeys) GetLogsIndexed30daySumOk() (*UsageBillableSummaryBody, bool) {
	if o == nil || o.LogsIndexed30daySum == nil {
		return nil, false
	}
	return o.LogsIndexed30daySum, true
}

// HasLogsIndexed30daySum returns a boolean if a field has been set.
func (o *UsageBillableSummaryKeys) HasLogsIndexed30daySum() bool {
	if o != nil && o.LogsIndexed30daySum != nil {
		return true
	}

	return false
}

// SetLogsIndexed30daySum gets a reference to the given UsageBillableSummaryBody and assigns it to the LogsIndexed30daySum field.
func (o *UsageBillableSummaryKeys) SetLogsIndexed30daySum(v UsageBillableSummaryBody) {
	o.LogsIndexed30daySum = &v
}

// GetLogsIndexed3daySum returns the LogsIndexed3daySum field value if set, zero value otherwise.
func (o *UsageBillableSummaryKeys) GetLogsIndexed3daySum() UsageBillableSummaryBody {
	if o == nil || o.LogsIndexed3daySum == nil {
		var ret UsageBillableSummaryBody
		return ret
	}
	return *o.LogsIndexed3daySum
}

// GetLogsIndexed3daySumOk returns a tuple with the LogsIndexed3daySum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryKeys) GetLogsIndexed3daySumOk() (*UsageBillableSummaryBody, bool) {
	if o == nil || o.LogsIndexed3daySum == nil {
		return nil, false
	}
	return o.LogsIndexed3daySum, true
}

// HasLogsIndexed3daySum returns a boolean if a field has been set.
func (o *UsageBillableSummaryKeys) HasLogsIndexed3daySum() bool {
	if o != nil && o.LogsIndexed3daySum != nil {
		return true
	}

	return false
}

// SetLogsIndexed3daySum gets a reference to the given UsageBillableSummaryBody and assigns it to the LogsIndexed3daySum field.
func (o *UsageBillableSummaryKeys) SetLogsIndexed3daySum(v UsageBillableSummaryBody) {
	o.LogsIndexed3daySum = &v
}

// GetLogsIndexed45daySum returns the LogsIndexed45daySum field value if set, zero value otherwise.
func (o *UsageBillableSummaryKeys) GetLogsIndexed45daySum() UsageBillableSummaryBody {
	if o == nil || o.LogsIndexed45daySum == nil {
		var ret UsageBillableSummaryBody
		return ret
	}
	return *o.LogsIndexed45daySum
}

// GetLogsIndexed45daySumOk returns a tuple with the LogsIndexed45daySum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryKeys) GetLogsIndexed45daySumOk() (*UsageBillableSummaryBody, bool) {
	if o == nil || o.LogsIndexed45daySum == nil {
		return nil, false
	}
	return o.LogsIndexed45daySum, true
}

// HasLogsIndexed45daySum returns a boolean if a field has been set.
func (o *UsageBillableSummaryKeys) HasLogsIndexed45daySum() bool {
	if o != nil && o.LogsIndexed45daySum != nil {
		return true
	}

	return false
}

// SetLogsIndexed45daySum gets a reference to the given UsageBillableSummaryBody and assigns it to the LogsIndexed45daySum field.
func (o *UsageBillableSummaryKeys) SetLogsIndexed45daySum(v UsageBillableSummaryBody) {
	o.LogsIndexed45daySum = &v
}

// GetLogsIndexed60daySum returns the LogsIndexed60daySum field value if set, zero value otherwise.
func (o *UsageBillableSummaryKeys) GetLogsIndexed60daySum() UsageBillableSummaryBody {
	if o == nil || o.LogsIndexed60daySum == nil {
		var ret UsageBillableSummaryBody
		return ret
	}
	return *o.LogsIndexed60daySum
}

// GetLogsIndexed60daySumOk returns a tuple with the LogsIndexed60daySum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryKeys) GetLogsIndexed60daySumOk() (*UsageBillableSummaryBody, bool) {
	if o == nil || o.LogsIndexed60daySum == nil {
		return nil, false
	}
	return o.LogsIndexed60daySum, true
}

// HasLogsIndexed60daySum returns a boolean if a field has been set.
func (o *UsageBillableSummaryKeys) HasLogsIndexed60daySum() bool {
	if o != nil && o.LogsIndexed60daySum != nil {
		return true
	}

	return false
}

// SetLogsIndexed60daySum gets a reference to the given UsageBillableSummaryBody and assigns it to the LogsIndexed60daySum field.
func (o *UsageBillableSummaryKeys) SetLogsIndexed60daySum(v UsageBillableSummaryBody) {
	o.LogsIndexed60daySum = &v
}

// GetLogsIndexed7daySum returns the LogsIndexed7daySum field value if set, zero value otherwise.
func (o *UsageBillableSummaryKeys) GetLogsIndexed7daySum() UsageBillableSummaryBody {
	if o == nil || o.LogsIndexed7daySum == nil {
		var ret UsageBillableSummaryBody
		return ret
	}
	return *o.LogsIndexed7daySum
}

// GetLogsIndexed7daySumOk returns a tuple with the LogsIndexed7daySum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryKeys) GetLogsIndexed7daySumOk() (*UsageBillableSummaryBody, bool) {
	if o == nil || o.LogsIndexed7daySum == nil {
		return nil, false
	}
	return o.LogsIndexed7daySum, true
}

// HasLogsIndexed7daySum returns a boolean if a field has been set.
func (o *UsageBillableSummaryKeys) HasLogsIndexed7daySum() bool {
	if o != nil && o.LogsIndexed7daySum != nil {
		return true
	}

	return false
}

// SetLogsIndexed7daySum gets a reference to the given UsageBillableSummaryBody and assigns it to the LogsIndexed7daySum field.
func (o *UsageBillableSummaryKeys) SetLogsIndexed7daySum(v UsageBillableSummaryBody) {
	o.LogsIndexed7daySum = &v
}

// GetLogsIndexed90daySum returns the LogsIndexed90daySum field value if set, zero value otherwise.
func (o *UsageBillableSummaryKeys) GetLogsIndexed90daySum() UsageBillableSummaryBody {
	if o == nil || o.LogsIndexed90daySum == nil {
		var ret UsageBillableSummaryBody
		return ret
	}
	return *o.LogsIndexed90daySum
}

// GetLogsIndexed90daySumOk returns a tuple with the LogsIndexed90daySum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryKeys) GetLogsIndexed90daySumOk() (*UsageBillableSummaryBody, bool) {
	if o == nil || o.LogsIndexed90daySum == nil {
		return nil, false
	}
	return o.LogsIndexed90daySum, true
}

// HasLogsIndexed90daySum returns a boolean if a field has been set.
func (o *UsageBillableSummaryKeys) HasLogsIndexed90daySum() bool {
	if o != nil && o.LogsIndexed90daySum != nil {
		return true
	}

	return false
}

// SetLogsIndexed90daySum gets a reference to the given UsageBillableSummaryBody and assigns it to the LogsIndexed90daySum field.
func (o *UsageBillableSummaryKeys) SetLogsIndexed90daySum(v UsageBillableSummaryBody) {
	o.LogsIndexed90daySum = &v
}

// GetLogsIndexedCustomRetentionSum returns the LogsIndexedCustomRetentionSum field value if set, zero value otherwise.
func (o *UsageBillableSummaryKeys) GetLogsIndexedCustomRetentionSum() UsageBillableSummaryBody {
	if o == nil || o.LogsIndexedCustomRetentionSum == nil {
		var ret UsageBillableSummaryBody
		return ret
	}
	return *o.LogsIndexedCustomRetentionSum
}

// GetLogsIndexedCustomRetentionSumOk returns a tuple with the LogsIndexedCustomRetentionSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryKeys) GetLogsIndexedCustomRetentionSumOk() (*UsageBillableSummaryBody, bool) {
	if o == nil || o.LogsIndexedCustomRetentionSum == nil {
		return nil, false
	}
	return o.LogsIndexedCustomRetentionSum, true
}

// HasLogsIndexedCustomRetentionSum returns a boolean if a field has been set.
func (o *UsageBillableSummaryKeys) HasLogsIndexedCustomRetentionSum() bool {
	if o != nil && o.LogsIndexedCustomRetentionSum != nil {
		return true
	}

	return false
}

// SetLogsIndexedCustomRetentionSum gets a reference to the given UsageBillableSummaryBody and assigns it to the LogsIndexedCustomRetentionSum field.
func (o *UsageBillableSummaryKeys) SetLogsIndexedCustomRetentionSum(v UsageBillableSummaryBody) {
	o.LogsIndexedCustomRetentionSum = &v
}

// GetLogsIndexedSum returns the LogsIndexedSum field value if set, zero value otherwise.
func (o *UsageBillableSummaryKeys) GetLogsIndexedSum() UsageBillableSummaryBody {
	if o == nil || o.LogsIndexedSum == nil {
		var ret UsageBillableSummaryBody
		return ret
	}
	return *o.LogsIndexedSum
}

// GetLogsIndexedSumOk returns a tuple with the LogsIndexedSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryKeys) GetLogsIndexedSumOk() (*UsageBillableSummaryBody, bool) {
	if o == nil || o.LogsIndexedSum == nil {
		return nil, false
	}
	return o.LogsIndexedSum, true
}

// HasLogsIndexedSum returns a boolean if a field has been set.
func (o *UsageBillableSummaryKeys) HasLogsIndexedSum() bool {
	if o != nil && o.LogsIndexedSum != nil {
		return true
	}

	return false
}

// SetLogsIndexedSum gets a reference to the given UsageBillableSummaryBody and assigns it to the LogsIndexedSum field.
func (o *UsageBillableSummaryKeys) SetLogsIndexedSum(v UsageBillableSummaryBody) {
	o.LogsIndexedSum = &v
}

// GetLogsIngestedSum returns the LogsIngestedSum field value if set, zero value otherwise.
func (o *UsageBillableSummaryKeys) GetLogsIngestedSum() UsageBillableSummaryBody {
	if o == nil || o.LogsIngestedSum == nil {
		var ret UsageBillableSummaryBody
		return ret
	}
	return *o.LogsIngestedSum
}

// GetLogsIngestedSumOk returns a tuple with the LogsIngestedSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryKeys) GetLogsIngestedSumOk() (*UsageBillableSummaryBody, bool) {
	if o == nil || o.LogsIngestedSum == nil {
		return nil, false
	}
	return o.LogsIngestedSum, true
}

// HasLogsIngestedSum returns a boolean if a field has been set.
func (o *UsageBillableSummaryKeys) HasLogsIngestedSum() bool {
	if o != nil && o.LogsIngestedSum != nil {
		return true
	}

	return false
}

// SetLogsIngestedSum gets a reference to the given UsageBillableSummaryBody and assigns it to the LogsIngestedSum field.
func (o *UsageBillableSummaryKeys) SetLogsIngestedSum(v UsageBillableSummaryBody) {
	o.LogsIngestedSum = &v
}

// GetNetworkDeviceTop99p returns the NetworkDeviceTop99p field value if set, zero value otherwise.
func (o *UsageBillableSummaryKeys) GetNetworkDeviceTop99p() UsageBillableSummaryBody {
	if o == nil || o.NetworkDeviceTop99p == nil {
		var ret UsageBillableSummaryBody
		return ret
	}
	return *o.NetworkDeviceTop99p
}

// GetNetworkDeviceTop99pOk returns a tuple with the NetworkDeviceTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryKeys) GetNetworkDeviceTop99pOk() (*UsageBillableSummaryBody, bool) {
	if o == nil || o.NetworkDeviceTop99p == nil {
		return nil, false
	}
	return o.NetworkDeviceTop99p, true
}

// HasNetworkDeviceTop99p returns a boolean if a field has been set.
func (o *UsageBillableSummaryKeys) HasNetworkDeviceTop99p() bool {
	if o != nil && o.NetworkDeviceTop99p != nil {
		return true
	}

	return false
}

// SetNetworkDeviceTop99p gets a reference to the given UsageBillableSummaryBody and assigns it to the NetworkDeviceTop99p field.
func (o *UsageBillableSummaryKeys) SetNetworkDeviceTop99p(v UsageBillableSummaryBody) {
	o.NetworkDeviceTop99p = &v
}

// GetNpmFlowSum returns the NpmFlowSum field value if set, zero value otherwise.
func (o *UsageBillableSummaryKeys) GetNpmFlowSum() UsageBillableSummaryBody {
	if o == nil || o.NpmFlowSum == nil {
		var ret UsageBillableSummaryBody
		return ret
	}
	return *o.NpmFlowSum
}

// GetNpmFlowSumOk returns a tuple with the NpmFlowSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryKeys) GetNpmFlowSumOk() (*UsageBillableSummaryBody, bool) {
	if o == nil || o.NpmFlowSum == nil {
		return nil, false
	}
	return o.NpmFlowSum, true
}

// HasNpmFlowSum returns a boolean if a field has been set.
func (o *UsageBillableSummaryKeys) HasNpmFlowSum() bool {
	if o != nil && o.NpmFlowSum != nil {
		return true
	}

	return false
}

// SetNpmFlowSum gets a reference to the given UsageBillableSummaryBody and assigns it to the NpmFlowSum field.
func (o *UsageBillableSummaryKeys) SetNpmFlowSum(v UsageBillableSummaryBody) {
	o.NpmFlowSum = &v
}

// GetNpmHostSum returns the NpmHostSum field value if set, zero value otherwise.
func (o *UsageBillableSummaryKeys) GetNpmHostSum() UsageBillableSummaryBody {
	if o == nil || o.NpmHostSum == nil {
		var ret UsageBillableSummaryBody
		return ret
	}
	return *o.NpmHostSum
}

// GetNpmHostSumOk returns a tuple with the NpmHostSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryKeys) GetNpmHostSumOk() (*UsageBillableSummaryBody, bool) {
	if o == nil || o.NpmHostSum == nil {
		return nil, false
	}
	return o.NpmHostSum, true
}

// HasNpmHostSum returns a boolean if a field has been set.
func (o *UsageBillableSummaryKeys) HasNpmHostSum() bool {
	if o != nil && o.NpmHostSum != nil {
		return true
	}

	return false
}

// SetNpmHostSum gets a reference to the given UsageBillableSummaryBody and assigns it to the NpmHostSum field.
func (o *UsageBillableSummaryKeys) SetNpmHostSum(v UsageBillableSummaryBody) {
	o.NpmHostSum = &v
}

// GetNpmHostTop99p returns the NpmHostTop99p field value if set, zero value otherwise.
func (o *UsageBillableSummaryKeys) GetNpmHostTop99p() UsageBillableSummaryBody {
	if o == nil || o.NpmHostTop99p == nil {
		var ret UsageBillableSummaryBody
		return ret
	}
	return *o.NpmHostTop99p
}

// GetNpmHostTop99pOk returns a tuple with the NpmHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryKeys) GetNpmHostTop99pOk() (*UsageBillableSummaryBody, bool) {
	if o == nil || o.NpmHostTop99p == nil {
		return nil, false
	}
	return o.NpmHostTop99p, true
}

// HasNpmHostTop99p returns a boolean if a field has been set.
func (o *UsageBillableSummaryKeys) HasNpmHostTop99p() bool {
	if o != nil && o.NpmHostTop99p != nil {
		return true
	}

	return false
}

// SetNpmHostTop99p gets a reference to the given UsageBillableSummaryBody and assigns it to the NpmHostTop99p field.
func (o *UsageBillableSummaryKeys) SetNpmHostTop99p(v UsageBillableSummaryBody) {
	o.NpmHostTop99p = &v
}

// GetProfContainerSum returns the ProfContainerSum field value if set, zero value otherwise.
func (o *UsageBillableSummaryKeys) GetProfContainerSum() UsageBillableSummaryBody {
	if o == nil || o.ProfContainerSum == nil {
		var ret UsageBillableSummaryBody
		return ret
	}
	return *o.ProfContainerSum
}

// GetProfContainerSumOk returns a tuple with the ProfContainerSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryKeys) GetProfContainerSumOk() (*UsageBillableSummaryBody, bool) {
	if o == nil || o.ProfContainerSum == nil {
		return nil, false
	}
	return o.ProfContainerSum, true
}

// HasProfContainerSum returns a boolean if a field has been set.
func (o *UsageBillableSummaryKeys) HasProfContainerSum() bool {
	if o != nil && o.ProfContainerSum != nil {
		return true
	}

	return false
}

// SetProfContainerSum gets a reference to the given UsageBillableSummaryBody and assigns it to the ProfContainerSum field.
func (o *UsageBillableSummaryKeys) SetProfContainerSum(v UsageBillableSummaryBody) {
	o.ProfContainerSum = &v
}

// GetProfHostTop99p returns the ProfHostTop99p field value if set, zero value otherwise.
func (o *UsageBillableSummaryKeys) GetProfHostTop99p() UsageBillableSummaryBody {
	if o == nil || o.ProfHostTop99p == nil {
		var ret UsageBillableSummaryBody
		return ret
	}
	return *o.ProfHostTop99p
}

// GetProfHostTop99pOk returns a tuple with the ProfHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryKeys) GetProfHostTop99pOk() (*UsageBillableSummaryBody, bool) {
	if o == nil || o.ProfHostTop99p == nil {
		return nil, false
	}
	return o.ProfHostTop99p, true
}

// HasProfHostTop99p returns a boolean if a field has been set.
func (o *UsageBillableSummaryKeys) HasProfHostTop99p() bool {
	if o != nil && o.ProfHostTop99p != nil {
		return true
	}

	return false
}

// SetProfHostTop99p gets a reference to the given UsageBillableSummaryBody and assigns it to the ProfHostTop99p field.
func (o *UsageBillableSummaryKeys) SetProfHostTop99p(v UsageBillableSummaryBody) {
	o.ProfHostTop99p = &v
}

// GetRumSum returns the RumSum field value if set, zero value otherwise.
func (o *UsageBillableSummaryKeys) GetRumSum() UsageBillableSummaryBody {
	if o == nil || o.RumSum == nil {
		var ret UsageBillableSummaryBody
		return ret
	}
	return *o.RumSum
}

// GetRumSumOk returns a tuple with the RumSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryKeys) GetRumSumOk() (*UsageBillableSummaryBody, bool) {
	if o == nil || o.RumSum == nil {
		return nil, false
	}
	return o.RumSum, true
}

// HasRumSum returns a boolean if a field has been set.
func (o *UsageBillableSummaryKeys) HasRumSum() bool {
	if o != nil && o.RumSum != nil {
		return true
	}

	return false
}

// SetRumSum gets a reference to the given UsageBillableSummaryBody and assigns it to the RumSum field.
func (o *UsageBillableSummaryKeys) SetRumSum(v UsageBillableSummaryBody) {
	o.RumSum = &v
}

// GetServerlessInvocationSum returns the ServerlessInvocationSum field value if set, zero value otherwise.
func (o *UsageBillableSummaryKeys) GetServerlessInvocationSum() UsageBillableSummaryBody {
	if o == nil || o.ServerlessInvocationSum == nil {
		var ret UsageBillableSummaryBody
		return ret
	}
	return *o.ServerlessInvocationSum
}

// GetServerlessInvocationSumOk returns a tuple with the ServerlessInvocationSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryKeys) GetServerlessInvocationSumOk() (*UsageBillableSummaryBody, bool) {
	if o == nil || o.ServerlessInvocationSum == nil {
		return nil, false
	}
	return o.ServerlessInvocationSum, true
}

// HasServerlessInvocationSum returns a boolean if a field has been set.
func (o *UsageBillableSummaryKeys) HasServerlessInvocationSum() bool {
	if o != nil && o.ServerlessInvocationSum != nil {
		return true
	}

	return false
}

// SetServerlessInvocationSum gets a reference to the given UsageBillableSummaryBody and assigns it to the ServerlessInvocationSum field.
func (o *UsageBillableSummaryKeys) SetServerlessInvocationSum(v UsageBillableSummaryBody) {
	o.ServerlessInvocationSum = &v
}

// GetSiemSum returns the SiemSum field value if set, zero value otherwise.
func (o *UsageBillableSummaryKeys) GetSiemSum() UsageBillableSummaryBody {
	if o == nil || o.SiemSum == nil {
		var ret UsageBillableSummaryBody
		return ret
	}
	return *o.SiemSum
}

// GetSiemSumOk returns a tuple with the SiemSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryKeys) GetSiemSumOk() (*UsageBillableSummaryBody, bool) {
	if o == nil || o.SiemSum == nil {
		return nil, false
	}
	return o.SiemSum, true
}

// HasSiemSum returns a boolean if a field has been set.
func (o *UsageBillableSummaryKeys) HasSiemSum() bool {
	if o != nil && o.SiemSum != nil {
		return true
	}

	return false
}

// SetSiemSum gets a reference to the given UsageBillableSummaryBody and assigns it to the SiemSum field.
func (o *UsageBillableSummaryKeys) SetSiemSum(v UsageBillableSummaryBody) {
	o.SiemSum = &v
}

// GetSyntheticsApiTestsSum returns the SyntheticsApiTestsSum field value if set, zero value otherwise.
func (o *UsageBillableSummaryKeys) GetSyntheticsApiTestsSum() UsageBillableSummaryBody {
	if o == nil || o.SyntheticsApiTestsSum == nil {
		var ret UsageBillableSummaryBody
		return ret
	}
	return *o.SyntheticsApiTestsSum
}

// GetSyntheticsApiTestsSumOk returns a tuple with the SyntheticsApiTestsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryKeys) GetSyntheticsApiTestsSumOk() (*UsageBillableSummaryBody, bool) {
	if o == nil || o.SyntheticsApiTestsSum == nil {
		return nil, false
	}
	return o.SyntheticsApiTestsSum, true
}

// HasSyntheticsApiTestsSum returns a boolean if a field has been set.
func (o *UsageBillableSummaryKeys) HasSyntheticsApiTestsSum() bool {
	if o != nil && o.SyntheticsApiTestsSum != nil {
		return true
	}

	return false
}

// SetSyntheticsApiTestsSum gets a reference to the given UsageBillableSummaryBody and assigns it to the SyntheticsApiTestsSum field.
func (o *UsageBillableSummaryKeys) SetSyntheticsApiTestsSum(v UsageBillableSummaryBody) {
	o.SyntheticsApiTestsSum = &v
}

// GetSyntheticsBrowserChecksSum returns the SyntheticsBrowserChecksSum field value if set, zero value otherwise.
func (o *UsageBillableSummaryKeys) GetSyntheticsBrowserChecksSum() UsageBillableSummaryBody {
	if o == nil || o.SyntheticsBrowserChecksSum == nil {
		var ret UsageBillableSummaryBody
		return ret
	}
	return *o.SyntheticsBrowserChecksSum
}

// GetSyntheticsBrowserChecksSumOk returns a tuple with the SyntheticsBrowserChecksSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryKeys) GetSyntheticsBrowserChecksSumOk() (*UsageBillableSummaryBody, bool) {
	if o == nil || o.SyntheticsBrowserChecksSum == nil {
		return nil, false
	}
	return o.SyntheticsBrowserChecksSum, true
}

// HasSyntheticsBrowserChecksSum returns a boolean if a field has been set.
func (o *UsageBillableSummaryKeys) HasSyntheticsBrowserChecksSum() bool {
	if o != nil && o.SyntheticsBrowserChecksSum != nil {
		return true
	}

	return false
}

// SetSyntheticsBrowserChecksSum gets a reference to the given UsageBillableSummaryBody and assigns it to the SyntheticsBrowserChecksSum field.
func (o *UsageBillableSummaryKeys) SetSyntheticsBrowserChecksSum(v UsageBillableSummaryBody) {
	o.SyntheticsBrowserChecksSum = &v
}

// GetTimeseriesAverage returns the TimeseriesAverage field value if set, zero value otherwise.
func (o *UsageBillableSummaryKeys) GetTimeseriesAverage() UsageBillableSummaryBody {
	if o == nil || o.TimeseriesAverage == nil {
		var ret UsageBillableSummaryBody
		return ret
	}
	return *o.TimeseriesAverage
}

// GetTimeseriesAverageOk returns a tuple with the TimeseriesAverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryKeys) GetTimeseriesAverageOk() (*UsageBillableSummaryBody, bool) {
	if o == nil || o.TimeseriesAverage == nil {
		return nil, false
	}
	return o.TimeseriesAverage, true
}

// HasTimeseriesAverage returns a boolean if a field has been set.
func (o *UsageBillableSummaryKeys) HasTimeseriesAverage() bool {
	if o != nil && o.TimeseriesAverage != nil {
		return true
	}

	return false
}

// SetTimeseriesAverage gets a reference to the given UsageBillableSummaryBody and assigns it to the TimeseriesAverage field.
func (o *UsageBillableSummaryKeys) SetTimeseriesAverage(v UsageBillableSummaryBody) {
	o.TimeseriesAverage = &v
}

func (o UsageBillableSummaryKeys) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.ApmHostSum != nil {
		toSerialize["apm_host_sum"] = o.ApmHostSum
	}
	if o.ApmHostTop99p != nil {
		toSerialize["apm_host_top99p"] = o.ApmHostTop99p
	}
	if o.ApmTraceSearchSum != nil {
		toSerialize["apm_trace_search_sum"] = o.ApmTraceSearchSum
	}
	if o.FargateContainerAverage != nil {
		toSerialize["fargate_container_average"] = o.FargateContainerAverage
	}
	if o.InfraContainerSum != nil {
		toSerialize["infra_container_sum"] = o.InfraContainerSum
	}
	if o.InfraHostSum != nil {
		toSerialize["infra_host_sum"] = o.InfraHostSum
	}
	if o.InfraHostTop99p != nil {
		toSerialize["infra_host_top99p"] = o.InfraHostTop99p
	}
	if o.IotTop99p != nil {
		toSerialize["iot_top99p"] = o.IotTop99p
	}
	if o.LambdaFunctionAverage != nil {
		toSerialize["lambda_function_average"] = o.LambdaFunctionAverage
	}
	if o.LogsIndexed15daySum != nil {
		toSerialize["logs_indexed_15day_sum"] = o.LogsIndexed15daySum
	}
	if o.LogsIndexed180daySum != nil {
		toSerialize["logs_indexed_180day_sum"] = o.LogsIndexed180daySum
	}
	if o.LogsIndexed30daySum != nil {
		toSerialize["logs_indexed_30day_sum"] = o.LogsIndexed30daySum
	}
	if o.LogsIndexed3daySum != nil {
		toSerialize["logs_indexed_3day_sum"] = o.LogsIndexed3daySum
	}
	if o.LogsIndexed45daySum != nil {
		toSerialize["logs_indexed_45day_sum"] = o.LogsIndexed45daySum
	}
	if o.LogsIndexed60daySum != nil {
		toSerialize["logs_indexed_60day_sum"] = o.LogsIndexed60daySum
	}
	if o.LogsIndexed7daySum != nil {
		toSerialize["logs_indexed_7day_sum"] = o.LogsIndexed7daySum
	}
	if o.LogsIndexed90daySum != nil {
		toSerialize["logs_indexed_90day_sum"] = o.LogsIndexed90daySum
	}
	if o.LogsIndexedCustomRetentionSum != nil {
		toSerialize["logs_indexed_custom_retention_sum"] = o.LogsIndexedCustomRetentionSum
	}
	if o.LogsIndexedSum != nil {
		toSerialize["logs_indexed_sum"] = o.LogsIndexedSum
	}
	if o.LogsIngestedSum != nil {
		toSerialize["logs_ingested_sum"] = o.LogsIngestedSum
	}
	if o.NetworkDeviceTop99p != nil {
		toSerialize["network_device_top99p"] = o.NetworkDeviceTop99p
	}
	if o.NpmFlowSum != nil {
		toSerialize["npm_flow_sum"] = o.NpmFlowSum
	}
	if o.NpmHostSum != nil {
		toSerialize["npm_host_sum"] = o.NpmHostSum
	}
	if o.NpmHostTop99p != nil {
		toSerialize["npm_host_top99p"] = o.NpmHostTop99p
	}
	if o.ProfContainerSum != nil {
		toSerialize["prof_container_sum"] = o.ProfContainerSum
	}
	if o.ProfHostTop99p != nil {
		toSerialize["prof_host_top99p"] = o.ProfHostTop99p
	}
	if o.RumSum != nil {
		toSerialize["rum_sum"] = o.RumSum
	}
	if o.ServerlessInvocationSum != nil {
		toSerialize["serverless_invocation_sum"] = o.ServerlessInvocationSum
	}
	if o.SiemSum != nil {
		toSerialize["siem_sum"] = o.SiemSum
	}
	if o.SyntheticsApiTestsSum != nil {
		toSerialize["synthetics_api_tests_sum"] = o.SyntheticsApiTestsSum
	}
	if o.SyntheticsBrowserChecksSum != nil {
		toSerialize["synthetics_browser_checks_sum"] = o.SyntheticsBrowserChecksSum
	}
	if o.TimeseriesAverage != nil {
		toSerialize["timeseries_average"] = o.TimeseriesAverage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

func (o *UsageBillableSummaryKeys) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		ApmHostSum                    *UsageBillableSummaryBody `json:"apm_host_sum,omitempty"`
		ApmHostTop99p                 *UsageBillableSummaryBody `json:"apm_host_top99p,omitempty"`
		ApmTraceSearchSum             *UsageBillableSummaryBody `json:"apm_trace_search_sum,omitempty"`
		FargateContainerAverage       *UsageBillableSummaryBody `json:"fargate_container_average,omitempty"`
		InfraContainerSum             *UsageBillableSummaryBody `json:"infra_container_sum,omitempty"`
		InfraHostSum                  *UsageBillableSummaryBody `json:"infra_host_sum,omitempty"`
		InfraHostTop99p               *UsageBillableSummaryBody `json:"infra_host_top99p,omitempty"`
		IotTop99p                     *UsageBillableSummaryBody `json:"iot_top99p,omitempty"`
		LambdaFunctionAverage         *UsageBillableSummaryBody `json:"lambda_function_average,omitempty"`
		LogsIndexed15daySum           *UsageBillableSummaryBody `json:"logs_indexed_15day_sum,omitempty"`
		LogsIndexed180daySum          *UsageBillableSummaryBody `json:"logs_indexed_180day_sum,omitempty"`
		LogsIndexed30daySum           *UsageBillableSummaryBody `json:"logs_indexed_30day_sum,omitempty"`
		LogsIndexed3daySum            *UsageBillableSummaryBody `json:"logs_indexed_3day_sum,omitempty"`
		LogsIndexed45daySum           *UsageBillableSummaryBody `json:"logs_indexed_45day_sum,omitempty"`
		LogsIndexed60daySum           *UsageBillableSummaryBody `json:"logs_indexed_60day_sum,omitempty"`
		LogsIndexed7daySum            *UsageBillableSummaryBody `json:"logs_indexed_7day_sum,omitempty"`
		LogsIndexed90daySum           *UsageBillableSummaryBody `json:"logs_indexed_90day_sum,omitempty"`
		LogsIndexedCustomRetentionSum *UsageBillableSummaryBody `json:"logs_indexed_custom_retention_sum,omitempty"`
		LogsIndexedSum                *UsageBillableSummaryBody `json:"logs_indexed_sum,omitempty"`
		LogsIngestedSum               *UsageBillableSummaryBody `json:"logs_ingested_sum,omitempty"`
		NetworkDeviceTop99p           *UsageBillableSummaryBody `json:"network_device_top99p,omitempty"`
		NpmFlowSum                    *UsageBillableSummaryBody `json:"npm_flow_sum,omitempty"`
		NpmHostSum                    *UsageBillableSummaryBody `json:"npm_host_sum,omitempty"`
		NpmHostTop99p                 *UsageBillableSummaryBody `json:"npm_host_top99p,omitempty"`
		ProfContainerSum              *UsageBillableSummaryBody `json:"prof_container_sum,omitempty"`
		ProfHostTop99p                *UsageBillableSummaryBody `json:"prof_host_top99p,omitempty"`
		RumSum                        *UsageBillableSummaryBody `json:"rum_sum,omitempty"`
		ServerlessInvocationSum       *UsageBillableSummaryBody `json:"serverless_invocation_sum,omitempty"`
		SiemSum                       *UsageBillableSummaryBody `json:"siem_sum,omitempty"`
		SyntheticsApiTestsSum         *UsageBillableSummaryBody `json:"synthetics_api_tests_sum,omitempty"`
		SyntheticsBrowserChecksSum    *UsageBillableSummaryBody `json:"synthetics_browser_checks_sum,omitempty"`
		TimeseriesAverage             *UsageBillableSummaryBody `json:"timeseries_average,omitempty"`
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
	if all.ApmHostSum != nil && all.ApmHostSum.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.ApmHostSum = all.ApmHostSum
	if all.ApmHostTop99p != nil && all.ApmHostTop99p.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.ApmHostTop99p = all.ApmHostTop99p
	if all.ApmTraceSearchSum != nil && all.ApmTraceSearchSum.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.ApmTraceSearchSum = all.ApmTraceSearchSum
	if all.FargateContainerAverage != nil && all.FargateContainerAverage.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.FargateContainerAverage = all.FargateContainerAverage
	if all.InfraContainerSum != nil && all.InfraContainerSum.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.InfraContainerSum = all.InfraContainerSum
	if all.InfraHostSum != nil && all.InfraHostSum.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.InfraHostSum = all.InfraHostSum
	if all.InfraHostTop99p != nil && all.InfraHostTop99p.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.InfraHostTop99p = all.InfraHostTop99p
	if all.IotTop99p != nil && all.IotTop99p.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.IotTop99p = all.IotTop99p
	if all.LambdaFunctionAverage != nil && all.LambdaFunctionAverage.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.LambdaFunctionAverage = all.LambdaFunctionAverage
	if all.LogsIndexed15daySum != nil && all.LogsIndexed15daySum.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.LogsIndexed15daySum = all.LogsIndexed15daySum
	if all.LogsIndexed180daySum != nil && all.LogsIndexed180daySum.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.LogsIndexed180daySum = all.LogsIndexed180daySum
	if all.LogsIndexed30daySum != nil && all.LogsIndexed30daySum.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.LogsIndexed30daySum = all.LogsIndexed30daySum
	if all.LogsIndexed3daySum != nil && all.LogsIndexed3daySum.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.LogsIndexed3daySum = all.LogsIndexed3daySum
	if all.LogsIndexed45daySum != nil && all.LogsIndexed45daySum.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.LogsIndexed45daySum = all.LogsIndexed45daySum
	if all.LogsIndexed60daySum != nil && all.LogsIndexed60daySum.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.LogsIndexed60daySum = all.LogsIndexed60daySum
	if all.LogsIndexed7daySum != nil && all.LogsIndexed7daySum.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.LogsIndexed7daySum = all.LogsIndexed7daySum
	if all.LogsIndexed90daySum != nil && all.LogsIndexed90daySum.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.LogsIndexed90daySum = all.LogsIndexed90daySum
	if all.LogsIndexedCustomRetentionSum != nil && all.LogsIndexedCustomRetentionSum.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.LogsIndexedCustomRetentionSum = all.LogsIndexedCustomRetentionSum
	if all.LogsIndexedSum != nil && all.LogsIndexedSum.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.LogsIndexedSum = all.LogsIndexedSum
	if all.LogsIngestedSum != nil && all.LogsIngestedSum.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.LogsIngestedSum = all.LogsIngestedSum
	if all.NetworkDeviceTop99p != nil && all.NetworkDeviceTop99p.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.NetworkDeviceTop99p = all.NetworkDeviceTop99p
	if all.NpmFlowSum != nil && all.NpmFlowSum.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.NpmFlowSum = all.NpmFlowSum
	if all.NpmHostSum != nil && all.NpmHostSum.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.NpmHostSum = all.NpmHostSum
	if all.NpmHostTop99p != nil && all.NpmHostTop99p.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.NpmHostTop99p = all.NpmHostTop99p
	if all.ProfContainerSum != nil && all.ProfContainerSum.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.ProfContainerSum = all.ProfContainerSum
	if all.ProfHostTop99p != nil && all.ProfHostTop99p.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.ProfHostTop99p = all.ProfHostTop99p
	if all.RumSum != nil && all.RumSum.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.RumSum = all.RumSum
	if all.ServerlessInvocationSum != nil && all.ServerlessInvocationSum.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.ServerlessInvocationSum = all.ServerlessInvocationSum
	if all.SiemSum != nil && all.SiemSum.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.SiemSum = all.SiemSum
	if all.SyntheticsApiTestsSum != nil && all.SyntheticsApiTestsSum.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.SyntheticsApiTestsSum = all.SyntheticsApiTestsSum
	if all.SyntheticsBrowserChecksSum != nil && all.SyntheticsBrowserChecksSum.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.SyntheticsBrowserChecksSum = all.SyntheticsBrowserChecksSum
	if all.TimeseriesAverage != nil && all.TimeseriesAverage.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.TimeseriesAverage = all.TimeseriesAverage
	return nil
}
