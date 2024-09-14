package subnet

import (
	"acloud/sdk"
	"encoding/json"
	"fmt"
	"log"
)

func CreateSubnet(projectId, vpcId, configFile string) error {
	// Read the config from YAML
	config, err := readYAMLConfig(configFile)
	if err != nil {
		log.Printf("Error reading YAML file: %v", err)
		return err
	}

	// Marshal the config struct into JSON
	payloadBytes, err := json.Marshal(config)
	if err != nil {
		log.Printf("Error marshaling config to JSON: %v", err)
		return fmt.Errorf("error marshaling config to JSON: %v", err)
	}

	// Convert the byte array to a string
	payload := string(payloadBytes)

	// Log the payload for debugging
	fmt.Println("Payload being sent:", payload)

	// Build the endpoint URL
	endpoint := fmt.Sprintf("projects/%s/providers/Aruba.Network/vpcs/%s/subnets", projectId, vpcId)

	// Send the payload using the already written SDK function
	response, err := sdk.SendPayload(endpoint, "POST", payload)
	if err != nil {
		log.Printf("Error sending API request: %v", err)
		return fmt.Errorf("error sending API request: %v", err)
	}

	// Log the response
	log.Printf("Response: %v", string(response))

	return nil
}
