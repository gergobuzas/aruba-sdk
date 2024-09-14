// src/sdk/container/updateKaas.go
package container

import (
	"acloud/sdk"
	"fmt"
	"log"
)

func UpdateKaas(projectId, kaasId, configFile string) {
	// Prepare the payload
	payload, err := preparePayload(configFile)
	if err != nil {
		return
	}
	fmt.Println(payload) // TODO REMOVE IF WORKS

	// Send the payload using the SDK (now the payload is a string)
	endpoint := fmt.Sprintf("projects/%s/providers/Aruba.Container/kaas/%s", projectId, kaasId)
	response, err := sdk.SendPayload(endpoint, "PUT", payload)
	if err != nil {
		log.Printf("Error sending API request: %v", err)
		return
	}

	// Log the response (you might want to check the status code, etc.)
	log.Printf("Response: %v", string(response))
}
