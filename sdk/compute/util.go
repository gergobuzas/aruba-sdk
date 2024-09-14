package compute

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Structs for the YAML file structure
type Metadata struct {
	Name     string   `yaml:"name" json:"name"`
	Location struct {
		Value string `yaml:"value" json:"value"`
	} `yaml:"location" json:"location"`
	Tags []string `yaml:"tags" json:"tags"`
}

type LinkedResource struct {
	URI string `yaml:"uri" json:"uri"`
}

type VPC struct {
	URI string `yaml:"uri" json:"uri"`
}

type Template struct {
	URI string `yaml:"uri" json:"uri"`
}

type ElasticIP struct {
	URI string `yaml:"uri" json:"uri"`
}

type KeyPair struct {
	URI string `yaml:"uri" json:"uri"`
}

type Subnet struct {
	URI string `yaml:"uri" json:"uri"`
}

type SecurityGroup struct {
	URI string `yaml:"uri" json:"uri"`
}

type Volume struct {
	URI string `yaml:"uri" json:"uri"`
}

type Properties struct {
	LinkedResources []LinkedResource `yaml:"linkedResources" json:"linkedResources"`
	DataCenter      string           `yaml:"dataCenter" json:"dataCenter"`
	VPC             VPC              `yaml:"vpc" json:"vpc"`
	VPCPreset       bool             `yaml:"vpcPreset" json:"vpcPreset"`
	FlavorID        string           `yaml:"flavorId" json:"flavorId"`
	Template        Template         `yaml:"template" json:"template"`
	AddElasticIP    bool             `yaml:"addElasticIp" json:"addElasticIp"`
	ElasticIP       ElasticIP        `yaml:"elasticIp" json:"elasticIp"`
	KeyPair         KeyPair          `yaml:"keyPair" json:"keyPair"`
	InitialPassword string           `yaml:"initialPassword" json:"initialPassword"`
	Subnets         []Subnet         `yaml:"subnets" json:"subnets"`
	SecurityGroups  []SecurityGroup  `yaml:"securityGroups" json:"securityGroups"`
	Volumes         []Volume         `yaml:"volumes" json:"volumes"`
}

type Config struct {
	Metadata   Metadata   `yaml:"metadata" json:"metadata"`
	Properties Properties `yaml:"properties" json:"properties"`
}

func readYAMLConfig(filePath string) (Config, error) {
	var config Config

	// Read YAML file
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return config, fmt.Errorf("error reading YAML file: %v", err)
	}

	// Unmarshal YAML into the Config struct
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return config, fmt.Errorf("error unmarshaling YAML file: %v", err)
	}

	return config, nil
}

func preparePayload(configFile string) (string, error) {
	// Read the config from YAML
	config, err := readYAMLConfig(configFile)
	if err != nil {
		log.Printf("Error reading YAML file: %v", err)
		return "", err
	}

	// Marshal the config struct into JSON
	payloadBytes, err := json.Marshal(config)
	if err != nil {
		log.Printf("Error marshaling config to JSON: %v", err)
		return "", fmt.Errorf("error marshaling config to JSON: %v", err)
	}

	// Convert the byte array to a string
	payload := string(payloadBytes)
	return payload, nil
}
