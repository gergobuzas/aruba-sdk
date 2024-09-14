package vpc

import (
	"acloud/sdk"
	"fmt"
)

func DeleteVpc(projectId string, vpcId string) string {
	payload := "" // Create a new POST request with the JSON body
	url := "projects/" + projectId + "/providers/Aruba.Network/vpcs/" + vpcId

	resp, err := sdk.SendPayload(url, "DELETE", payload)
	if err != nil {
		fmt.Println("Something went wrong with ")
	}
	return resp
}
