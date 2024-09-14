package subnet

import (
	"acloud/sdk"
	"fmt"
)

func ListSubnet(projectId, vpcId string) error {
	endpoint := fmt.Sprintf("projects/%s/providers/Aruba.Network/vpcs/%s/subnets", projectId, vpcId)
	// Call the SendPayload function (since it's a GET request, payload is empty)
	response, err := sdk.SendPayload(endpoint, "GET", "")
	if err != nil {
		fmt.Println("Error listing subnets:", err)
		return err
	}
	fmt.Println("Response:")
	fmt.Println(response)

	return nil
}
