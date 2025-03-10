// Delete an SLO correction returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("DeleteSLOCorrection", true)
	apiClient := datadog.NewAPIClient(configuration)
	r, err := apiClient.ServiceLevelObjectiveCorrectionsApi.DeleteSLOCorrection(ctx, "slo_correction_id")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectiveCorrectionsApi.DeleteSLOCorrection`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
