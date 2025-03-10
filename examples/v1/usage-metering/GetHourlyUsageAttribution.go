// Get Hourly Usage Attribution returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("GetHourlyUsageAttribution", true)
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageMeteringApi.GetHourlyUsageAttribution(ctx, time.Now().AddDate(0, 0, -3), datadog.HOURLYUSAGEATTRIBUTIONUSAGETYPE_INFRA_HOST_USAGE, *datadog.NewGetHourlyUsageAttributionOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetHourlyUsageAttribution`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `UsageMeteringApi.GetHourlyUsageAttribution`:\n%s\n", responseContent)
}
