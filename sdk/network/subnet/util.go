package subnet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"gopkg.in/yaml.v2"
)

// Define the structures for the YAML file and JSON payload.
type Metadata struct {
	Name     string   `yaml:"name" json:"name,omitempty"`
	Location Location `yaml:"location" json:"location,omitempty"`
	Tags     []string `yaml:"tags" json:"tags,omitempty"`
}

type Location struct {
	Value string `yaml:"value" json:"value,omitempty"`
}

type Properties struct {
	Type    string   `yaml:"type" json:"type,omitempty"`
	Default bool     `yaml:"default" json:"default,omitempty"`
	Network Network  `yaml:"network" json:"network,omitempty"`
	DHCP    DHCP     `yaml:"dhcp" json:"dhcp,omitempty"`
	Routes  []Route  `yaml:"routes" json:"routes,omitempty"`
	DNS     []string `yaml:"dns" json:"dns,omitempty"`
}

type Network struct {
	Address string `yaml:"address" json:"address,omitempty"`
}

type DHCP struct {
	Enabled bool   `yaml:"enabled" json:"enabled,omitempty"`
	Range   *Range `yaml:"range,omitempty" json:"range,omitempty"` // Pointer used to omit if nil
}

type Range struct {
	Start string `yaml:"start" json:"start,omitempty"`
	Count int    `yaml:"count" json:"count,omitempty"`
}

type Route struct {
	Address string `yaml:"address" json:"address,omitempty"`
	Gateway string `yaml:"gateway" json:"gateway,omitempty"`
}

type Config struct {
	Metadata   Metadata   `yaml:"metadata" json:"metadata"`
	Properties Properties `yaml:"properties" json:"properties"`
}

// Function to read YAML config file
func readYAMLConfig(filePath string) (Config, error) {
	var config Config

	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return config, fmt.Errorf("error reading YAML file: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return config, fmt.Errorf("error unmarshaling YAML file: %v", err)
	}

	return config, nil
}

// Function to create and send the API request
func createRequestWithPayload(filePath string) error {
	// Read and parse YAML file
	config, err := readYAMLConfig(filePath)
	if err != nil {
		return err
	}

	// Marshal struct into JSON
	payload, err := json.Marshal(config)
	if err != nil {
		return fmt.Errorf("error marshaling config to JSON: %v", err)
	}

	// Prepare the HTTP request
	url := "https://example.com/api/resource" // Replace with your API endpoint
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("error creating HTTP request: %v", err)
	}

	// Set necessary headers
	req.Header.Set("Content-Type", "application/json")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err)
	}

	// Output the response
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", string(body))

	return nil
}
