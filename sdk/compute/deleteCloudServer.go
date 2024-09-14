package compute

import (
	"acloud/sdk"
	"fmt"
	"log"
)

func DeleteCloudServer(projectID, cloudServerID string) error {
	endpoint := fmt.Sprintf("projects/%s/providers/Aruba.Compute/cloudServers/%s", projectID, cloudServerID)
	response, err := sdk.SendPayload(endpoint, "DELETE", "")
	if err != nil {
		log.Printf("Error sending API request: %v", err)
		return fmt.Errorf("error sending API request: %v", err)
	}

	log.Printf("Response: %v", string(response))
	return nil
}
