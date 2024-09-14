package subnet

import (
	"acloud/sdk"
	"encoding/json"
	"fmt"
	"log"
)

func UpdateSubnet(projectId, vpcId, subnetId, configFile string) error {
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

	// Send the payload using the SDK
	endpoint := fmt.Sprintf("projects/%s/providers/Aruba.Network/vpcs/%s/subnets/%s", projectId, vpcId, subnetId)
	response, err := sdk.SendPayload(endpoint, "PUT", payload)
	if err != nil {
		log.Printf("Error sending API request: %v", err)
		return fmt.Errorf("error sending API request: %v", err)
	}

	// Log the response
	log.Printf("Response: %v", string(response))

	return nil
}
