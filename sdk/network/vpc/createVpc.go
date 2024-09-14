package vpc

import (
	"acloud/sdk"
	"fmt"
)

func CreateVpc(projectId string, vpcName string, tags string, location string) error {
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
	url := "projects/" + projectId + "/providers/Aruba.Network/vpcs"

	sdk.SendPayload(url, "POST", payload)
	return nil
}
