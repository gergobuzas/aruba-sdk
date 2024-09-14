package vpc

import (
	"acloud/sdk"
	"fmt"
)

func UpdateVpc(projectId string, vpcId string, vpcName string, tags string, location string) string {
	payload := fmt.Sprintf(`{
		"metadata":{
			"name":"%s",
			"tags": [%s],
			"location": {
				"value": "%s"
			}
		}
	}`, vpcName, sdk.FormatList(tags), location)

	// Create a new POST request with the JSON body
	url := "projects/" + projectId + "/providers/Aruba.Network/vpcs/" + vpcId

	resp, err := sdk.SendPayload(url, "PUT", payload)
	if err != nil {
		fmt.Println("Something went wrong during update of vpc...")
	}
	return resp
}
