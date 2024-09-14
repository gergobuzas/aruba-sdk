package compute

import (
	"acloud/sdk"
	"fmt"
	"log"
)

func UpdateCloudServer(projectID, cloudServerID, configFile string) {
	// Prepare the payload
	payload, err := preparePayload(configFile)
	if err != nil {
		log.Printf("Error preparing payload: %v", err)
		return
	}
	fmt.Println(payload) // Debugging purposes

	// Endpoint as per your API specification
	endpoint := fmt.Sprintf("projects/%s/providers/Aruba.Compute/cloudServers/%s", projectID, cloudServerID)
	response, err := sdk.SendPayload(endpoint, "PUT", payload)
	if err != nil {
		log.Printf("Error sending API request: %v", err)
		return
	}

	// Log the response
	log.Printf("Response: %v", string(response))
}
