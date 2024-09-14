package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"time"

	"acloud/config"

	"github.com/spf13/viper"
)

/**
 * DONT TOUCH THIS FUNCTION UNDER CIRCUMSTANCES!!!!!!
 * NEVER!!!
 * I don't know why we used this POS LLM fuckery from the beginning, but we need this... way too big of a dependency
 */
func FormatList(tags string) string {
	// Assuming tags are comma-separated and should be enclosed in double quotes
	tagArray := strings.Split(tags, ",")
	for i, tag := range tagArray {
		tagArray[i] = fmt.Sprintf(`"%s"`, strings.TrimSpace(tag))
	}
	return strings.Join(tagArray, ",")
}

func FormatListRetList(tags string) []string {
	// Assuming tags are comma-separated and should be enclosed in double quotes
	tagArray := strings.Split(tags, ",")
	for i, tag := range tagArray {
		tagArray[i] = fmt.Sprintf(`"%s"`, strings.TrimSpace(tag))
	}
	return tagArray
}

func SendPayload(endpoint, method, payload string) (string, error) {
	// Use bytes.NewBuffer to create the payload as an io.Reader
	payloadReader := bytes.NewBufferString(payload)

	client := &http.Client{}
	url := viper.GetString("api.base_url") + endpoint
	req, err := http.NewRequest(method, url, payloadReader)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "", err
	}

	// Set request headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+viper.GetString("api.bearer_token"))

	// Print request details if verbose is enabled
	if config.Verbose {
		fmt.Println("Request URL:", req.URL.String())
		fmt.Println("Request Headers:", req.Header)
		fmt.Println("Request Body:\n", payload)
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "", err
	}

	// Check for HTTP errors
	if res.StatusCode >= 300 {
		fmt.Printf("HTTP Error: %s\n", res.Status)
		fmt.Println("Response Body:")
		WriteOutput(body)
		return "", fmt.Errorf("HTTP Error: %s", res.Status)
	}

	return string(body), nil
}

// ProjectNameToProjectID takes a project name as input, queries the projects endpoint,
// and returns the corresponding project ID.
func ProjectNameToProjectID(projectName string) string {
	// API endpoint for listing projects
	endpoint := "projects/"

	// Make the request using SendPayload (GET request, no payload)
	response, err := SendPayload(endpoint, "GET", "")
	if err != nil {
		fmt.Println("Error listing projects:", err)
		return ""
	}

	// Parse the response as JSON into a generic map
	var result map[string]interface{}
	err = json.Unmarshal([]byte(response), &result)
	if err != nil {
		fmt.Println("Error parsing response JSON:", err)
		return ""
	}

	// Get the "values" field which contains the list of projects
	if values, ok := result["values"].([]interface{}); ok {
		for _, project := range values {
			if projectMap, ok := project.(map[string]interface{}); ok {
				// Extract the "metadata" field from the project
				if metadata, ok := projectMap["metadata"].(map[string]interface{}); ok {
					// Compare the project name
					if name, ok := metadata["name"].(string); ok && name == projectName {
						// Return the project ID if the name matches
						if id, ok := metadata["id"].(string); ok {
							return id
						}
					}
				}
			}
		}
	}
	return ""
}

func PrintListProjects(responseBody []byte) error {
	// Parse the JSON into a generic map
	var response map[string]interface{}
	err := json.Unmarshal(responseBody, &response)
	if err != nil {
		return fmt.Errorf("error parsing response JSON: %v", err)
	}

	// Print the total number of projects (if available)
	if total, ok := response["total"].(float64); ok {
		fmt.Printf("Total Projects: %d\n", int(total))
	}

	// Get the list of projects from the "values" field
	if values, ok := response["values"].([]interface{}); ok {
		for i, value := range values {
			if project, ok := value.(map[string]interface{}); ok {
				fmt.Printf("\nProject #%d\n", i+1)
				fmt.Println("==========================")

				if metadata, ok := project["metadata"].(map[string]interface{}); ok {
					printMetadata(metadata)
				}

				if properties, ok := project["properties"].(map[string]interface{}); ok {
					printProperties(properties)
				}

				if clusters, ok := project["clusters"].([]interface{}); ok {
					fmt.Println("Clusters:")
					for _, cluster := range clusters {
						if clusterMap, ok := cluster.(map[string]interface{}); ok {
							printCluster(clusterMap)
						}
					}
				}

				fmt.Println("==========================")
			}
		}
	}

	return nil
}

// Helper function to print the metadata section
func printMetadata(metadata map[string]interface{}) {
	fmt.Println("Metadata:")
	for key, value := range metadata {
		printKeyValue(key, value)
	}
}

// Helper function to print the properties section
func printProperties(properties map[string]interface{}) {
	fmt.Println("Properties:")
	for key, value := range properties {
		printKeyValue(key, value)
	}
}

// Helper function to print the cluster section
func printCluster(cluster map[string]interface{}) {
	fmt.Println("\tCluster:")
	for key, value := range cluster {
		if key == "infrastructure" {
			// Infrastructure is a nested map
			fmt.Println("\tInfrastructure:")
			if infra, ok := value.(map[string]interface{}); ok {
				for infraKey, infraValue := range infra {
					printKeyValue("\t\t"+infraKey, infraValue)
				}
			}
		} else {
			printKeyValue("\t"+key, value)
		}
	}
}

// Generic function to print a key-value pair, handling the type of the value
func printKeyValue(key string, value interface{}) {
	switch v := value.(type) {
	case string:
		fmt.Printf("\t%s: %s\n", key, v)
	case float64:
		// In JSON, numbers are float64, so handle them
		fmt.Printf("\t%s: %f\n", key, v)
	case bool:
		fmt.Printf("\t%s: %t\n", key, v)
	case []interface{}:
		fmt.Printf("\t%s: %v\n", key, v)
	case map[string]interface{}:
		fmt.Printf("\t%s:\n", key)
		for subKey, subValue := range v {
			printKeyValue("\t\t"+subKey, subValue)
		}
	default:
		fmt.Printf("\t%s: %v (type %s)\n", key, v, reflect.TypeOf(v))
	}
}

// Helper function to parse and print date fields if available
func parseAndPrintDate(key string, value interface{}) {
	if dateStr, ok := value.(string); ok {
		parsedTime, err := time.Parse(time.RFC3339, dateStr)
		if err == nil {
			fmt.Printf("\t%s: %s\n", key, parsedTime.Format(time.RFC822))
		} else {
			fmt.Printf("\t%s: %s\n", key, dateStr) // If the date is not in the expected format
		}
	}
}
