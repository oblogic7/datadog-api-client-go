// Create an API test with multi subtype returns "OK - Returns the created test details." response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	body := datadog.SyntheticsAPITest{
		Config: datadog.SyntheticsAPITestConfig{
			ConfigVariables: []datadog.SyntheticsConfigVariable{
				{
					Example: datadog.PtrString("content-type"),
					Name:    "PROPERTY",
					Pattern: datadog.PtrString("content-type"),
					Type:    datadog.SYNTHETICSCONFIGVARIABLETYPE_TEXT,
				},
			},
			Steps: []datadog.SyntheticsAPIStep{
				{
					AllowFailure: datadog.PtrBool(true),
					Assertions: []datadog.SyntheticsAssertion{
						datadog.SyntheticsAssertion{
							SyntheticsAssertionTarget: &datadog.SyntheticsAssertionTarget{
								Operator: datadog.SYNTHETICSASSERTIONOPERATOR_IS,
								Type:     datadog.SYNTHETICSASSERTIONTYPE_STATUS_CODE,
								Target:   200,
							}},
					},
					IsCritical: datadog.PtrBool(true),
					Name:       "request is sent",
					Request: datadog.SyntheticsTestRequest{
						Method:  datadog.HTTPMETHOD_GET.Ptr(),
						Timeout: datadog.PtrFloat64(10),
						Url:     datadog.PtrString("https://datadoghq.com"),
					},
					Retry: &datadog.SyntheticsTestOptionsRetry{
						Count:    datadog.PtrInt64(5),
						Interval: datadog.PtrFloat64(1000),
					},
					Subtype: datadog.SYNTHETICSAPISTEPSUBTYPE_HTTP,
				},
			},
		},
		Locations: []string{
			"aws:us-east-2",
		},
		Message: datadog.PtrString("BDD test payload: synthetics_api_test_multi_step_payload.json"),
		Name:    "Example-Create_an_API_test_with_multi_subtype_returns_OK_Returns_the_created_test_details_response",
		Options: datadog.SyntheticsTestOptions{
			AcceptSelfSigned:   datadog.PtrBool(false),
			AllowInsecure:      datadog.PtrBool(true),
			FollowRedirects:    datadog.PtrBool(true),
			MinFailureDuration: datadog.PtrInt64(10),
			MinLocationFailed:  datadog.PtrInt64(1),
			MonitorName:        datadog.PtrString("Example-Create_an_API_test_with_multi_subtype_returns_OK_Returns_the_created_test_details_response"),
			MonitorPriority:    datadog.PtrInt32(5),
			Retry: &datadog.SyntheticsTestOptionsRetry{
				Count:    datadog.PtrInt64(3),
				Interval: datadog.PtrFloat64(1000),
			},
			TickEvery: datadog.PtrInt64(60),
		},
		Subtype: datadog.SYNTHETICSTESTDETAILSSUBTYPE_MULTI.Ptr(),
		Tags: []string{
			"testing:api",
		},
		Type: datadog.SYNTHETICSAPITESTTYPE_API,
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.SyntheticsApi.CreateSyntheticsAPITest(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.CreateSyntheticsAPITest`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.CreateSyntheticsAPITest`:\n%s\n", responseContent)
}
