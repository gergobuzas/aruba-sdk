package container

import (
	"acloud/sdk"
	"fmt"
	"log"
)

func CreateKaasbyId(id, configFile string) error {
	// Prepare the payload
	payload, err := preparePayload(configFile)
	if err != nil {
		return err
	}
	fmt.Println(payload) // TODO REMOVE IF WORKS

	// Send the payload using the SDK (now the payload is a string)
	response, err := sdk.SendPayload("projects/"+id+"/providers/Aruba.Container/kaas", "POST", payload)
	if err != nil {
		log.Printf("Error sending API request: %v", err)
		return fmt.Errorf("error sending API request: %v", err)
	}

	// Log the response (you might want to check the status code, etc.)
	log.Printf("Response: %v", string(response))

	return nil
}

func CreateKaasbyName(projectName, config_file string) {
	// GET ID
	pid := ""
	CreateKaasbyId(pid, config_file)
}
