package container

import (
	"acloud/sdk"
	"fmt"
	"log"
)

func DeleteKaas(projectId, kaasId string) error {
	endpoint := fmt.Sprintf("projects/%s/providers/Aruba.Container/kaas/%s", projectId, kaasId)
	response, err := sdk.SendPayload(endpoint, "DELETE", "")
	if err != nil {
		log.Printf("Error sending API request: %v", err)
		return fmt.Errorf("error sending API request: %v", err)
	}

	log.Printf("Response: %v", string(response))
	return nil
}
