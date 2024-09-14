package container

import (
	"acloud/sdk"
	"fmt"
	"net/url"
)

func ListKaas(projectId, filter, sort, projection string, offset, limit int32) error {
	// API endpoint for listing projects
	endpoint := fmt.Sprintf("projects/%s/providers/Aruba.Container/kaas", projectId)

	// Prepare the query parameters
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
	if limit != 0 {
		queryParams.Add("limit", fmt.Sprintf("%d", limit))
	}

	// Construct the full endpoint with query parameters
	if len(queryParams) > 0 {
		endpoint = endpoint + "?" + queryParams.Encode()
	}

	// Call the SendPayload function (since it's a GET request, payload is empty)
	response, err := sdk.SendPayload(endpoint, "GET", "")
	if err != nil {
		fmt.Println("Error listing containers:", err)
		return err
	}

	fmt.Println("Response:")
	fmt.Println(response)

	return nil
}
