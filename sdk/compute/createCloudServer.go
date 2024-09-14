package compute

import (
	"acloud/sdk"
	"fmt"
	"log"
)

func CreateCloudServerByID(projectID, configFile string) error {
	// Prepare the payload
	payload, err := preparePayload(configFile)
	if err != nil {
		return err
	}
	fmt.Println(payload) // Debugging purposes

	// Endpoint as per your API specification
	endpoint := fmt.Sprintf("projects/%s/providers/Aruba.Compute/cloudServers", projectID)

	// Send the payload using the SDK
	response, err := sdk.SendPayload(endpoint, "POST", payload)
	if err != nil {
		log.Printf("Error sending API request: %v", err)
		return fmt.Errorf("error sending API request: %v", err)
	}

	// Log the response
	log.Printf("Response: %v", string(response))

	return nil
}
