// List permissions for a role returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	// there is a valid "role" in the system
	RoleDataID := os.Getenv("ROLE_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesApi.ListRolePermissions(ctx, RoleDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.ListRolePermissions`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RolesApi.ListRolePermissions`:\n%s\n", responseContent)
}
