package vpc

import (
	"acloud/sdk"
	"fmt"
)

func ListVpcs(projectId string) string {
	payload := ""

	// Create a new POST request with the JSON body
	url := "projects/" + projectId + "/providers/Aruba.Network/vpcs"

	resp, err := sdk.SendPayload(url, "GET", payload)
	if err != nil {
		fmt.Println("Something went wrong...")
	}
	return resp
}
