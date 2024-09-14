package compute

import (
	"acloud/sdk"
	"fmt"
	"log"
)

func GetCloudServer(projectID, cloudServerID string) {
	endpoint := fmt.Sprintf("projects/%s/providers/Aruba.Compute/cloudServers/%s", projectID, cloudServerID)
	response, err := sdk.SendPayload(endpoint, "GET", "")
	if err != nil {
		log.Printf("Error fetching Cloud Server details: %v", err)
		return
	}

	log.Printf("Response: %v", string(response))
}
