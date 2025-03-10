// Delete a global variable returns "OK" response

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
	apiClient := datadog.NewAPIClient(configuration)
	r, err := apiClient.SyntheticsApi.DeleteGlobalVariable(ctx, "variable_id")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.DeleteGlobalVariable`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
