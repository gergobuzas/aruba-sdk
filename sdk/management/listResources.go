package sdk

import (
	"acloud/sdk"
	"fmt"
	"net/url"
)

// ListResources fetches resources across all projects with optional query parameters
func ListResources(filter, sort, projection string, offset, limit int32) error {
	// Construct the base endpoint
	endpoint := "projects/resources"

	// Prepare query parameters
	queryParams := url.Values{}
	if filter != "" {
		queryParams.Add("filter", filter)
	}
	if sort != "" {
		queryParams.Add("sort", sort)
	}
	if projection != "" {
		queryParams.Add("projection", projection)
	}
	if offset != 0 {
		queryParams.Add("offset", fmt.Sprintf("%d", offset))
	}
	if limit != 100 { // default is 100, only add it if it's different
		queryParams.Add("limit", fmt.Sprintf("%d", limit))
	}

	// Append query parameters to the endpoint if present
	if len(queryParams) > 0 {
		endpoint = fmt.Sprintf("%s?%s", endpoint, queryParams.Encode())
	}

	// Call the SendPayload function to make the GET request (no payload for GET)
	response, err := sdk.SendPayload(endpoint, "GET", "")
	if err != nil {
		return fmt.Errorf("error listing resources: %v", err)
	}

	sdk.WriteOutput(response)

	return nil
}
