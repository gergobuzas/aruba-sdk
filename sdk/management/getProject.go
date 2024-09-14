package sdk

import (
	"acloud/sdk"
	"fmt"
)

// GetProject fetches the details of a specific project by ID from the Aruba Cloud API
func GetProjectByID(projectID string) error {
	// Build the endpoint with the project ID
	endpoint := fmt.Sprintf("projects/%s", projectID)

	// Call the SendPayload function (since it's a GET request, payload is empty)
	response, err := sdk.SendPayload(endpoint, "GET", "")
	if err != nil {
		fmt.Println("Error fetching project details:", err)
		return err
	}

	sdk.WriteOutput(response)

	return nil
}

// GetProjectByName fetches the details of a specific project by its name
func GetProjectByName(projectName string) error {
	// Get the project ID from the project name
	projectID := sdk.ProjectNameToProjectID(projectName)
	if projectID == "" {
		return fmt.Errorf("project with name '%s' not found", projectName)
	}

	// Fetch the project details using the project ID
	err := GetProjectByID(projectID)
	if err != nil {
		return fmt.Errorf("error fetching project details for project ID '%s': %v", projectID, err)
	}

	return nil
}
