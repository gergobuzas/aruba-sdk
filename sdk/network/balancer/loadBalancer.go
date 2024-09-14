package balancer

import (
	"acloud/sdk"
	"fmt"
	"log"
	"net/url"
)

// ListLoadBalancers retrieves a list of all Load Balancers within a project.
func ListLoadBalancers(projectID, filter, sort, projection string, offset, limit int32) error {
	// Construct the base endpoint
	endpoint := fmt.Sprintf("projects/%s/Aruba.Network/loadBalancers", projectID)

	// Prepare query parameters using url.Values
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
	if limit != 100 { // default is 100, so only add it if it's different
		queryParams.Add("limit", fmt.Sprintf("%d", limit))
	}

	// Append query parameters to the endpoint if present
	if len(queryParams) > 0 {
		endpoint = fmt.Sprintf("%s?%s", endpoint, queryParams.Encode())
	}

	// Send the GET request
	response, err := sdk.SendPayload(endpoint, "GET", "")
	if err != nil {
		log.Printf("Error listing Load Balancers: %v", err)
		return err
	}

	// Print the response
	fmt.Println("Response:")
	fmt.Println(response)

	return nil
}

func GetLoadBalancer(projectID, loadBalancerID string) error {
	endpoint := fmt.Sprintf("projects/%s/providers/Aruba.Network/loadBalancers/%s", projectID, loadBalancerID)
	response, err := sdk.SendPayload(endpoint, "GET", "")
	if err != nil {
		log.Printf("Error fetching Load Balancer details: %v", err)
		return err
	}

	// Print the response
	fmt.Println("Response:")
	fmt.Println(response)

	return nil
}
