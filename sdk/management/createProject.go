package sdk

import (
	"acloud/config"
	"acloud/sdk"
	"fmt"
)

func CreateProject(name string, tags string, defVar bool, description string) error {
	if config.Verbose {
		fmt.Println("Creating new project")
	}
	// Create the payload with the provided parameters
	payload := fmt.Sprintf(`{
		"metadata": {
			"name": "%s",
			"tags": [%s]
		},
		"properties": {
			"description": "%s",
			"default": %t
		}
	}`, name, sdk.FormatList(tags), description, defVar)

	// Use SendPayload from util.go
	url := "projects/"
	method := "POST"

	response, err := sdk.SendPayload(url, method, payload)
	if err != nil {
		return err
	}

	sdk.WriteOutput(response)
	return nil
}
