package sdk

import (
	"acloud/sdk"
	"fmt"
)

// DeleteProjectByID deletes a project by its ID
func DeleteProjectByID(projectID string) error {
	// Construct the API endpoint, inserting the project ID into the URL
	endpoint := fmt.Sprintf("projects/%s", projectID)

	// Call the SendPayload function to make the DELETE request
	response, err := sdk.SendPayload(endpoint, "DELETE", "")
	if err != nil {
		return fmt.Errorf("error deleting project: %v", err)
	}

	// Check for empty response (since the API returns 204 No Content)
	if response == "" {
		fmt.Println("Project deleted successfully (no content returned).")
	}

	return nil
}
