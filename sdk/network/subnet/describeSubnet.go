package subnet

import (
	"acloud/sdk"
	"fmt"
	"log"
)

func DescribeSubnet(projectId, vpcId, subnetId string) (string, error) {
	// Construct the endpoint URL
	endpoint := fmt.Sprintf("projects/%s/providers/Aruba.Network/vpcs/%s/subnets/%s", projectId, vpcId, subnetId)

	// Send the GET request using the SDK
	response, err := sdk.SendPayload(endpoint, "GET", "")
	if err != nil {
		log.Printf("Error sending API request: %v", err)
		return "", fmt.Errorf("error sending API request: %v", err)
	}

	// Log the response
	log.Printf("Response: %v", string(response))

	return string(response), nil
}
