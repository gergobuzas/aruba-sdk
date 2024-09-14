package sdk

import (
	"acloud/sdk"
	"encoding/json"
	"fmt"
)

// UpdateProjectByID updates the details of a specific project by ID
func UpdateProjectByID(projectID, name, description string, tags []string, isDefault bool) error {
	// Define the API endpoint, inserting the project ID into the path
	endpoint := fmt.Sprintf("projects/%s", projectID)

	// Construct the JSON payload with the provided values
	payloadMap := map[string]interface{}{
		"metadata": map[string]interface{}{
			"name": name,
			"tags": tags,
		},
		"properties": map[string]interface{}{
			"description": description,
			"default":     isDefault,
		},
	}

	// Marshal the payload into JSON
	payloadBytes, err := json.Marshal(payloadMap)
	if err != nil {
		return fmt.Errorf("error marshalling payload: %v", err)
	}

	// Convert the payload to a string for the SendPayload function
	payloadString := string(payloadBytes)

	// Call the SendPayload function to make the PATCH request
	response, err := sdk.SendPayload(endpoint, "PUT", payloadString)
	if err != nil {
		return fmt.Errorf("error sending payload: %v", err)
	}

	sdk.WriteOutput(response)

	return nil
}
