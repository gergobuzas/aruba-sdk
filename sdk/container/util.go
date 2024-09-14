package container

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// Structs for the YAML file structure
type Metadata struct {
	Name     string `yaml:"name" json:"name"`
	Location struct {
		Value string `yaml:"value" json:"value"`
	} `yaml:"location" json:"location"`
	Tags []string `yaml:"tags" json:"tags"`
}

type VPC struct {
	URI string `yaml:"uri" json:"uri"`
}

type Subnet struct {
	URI string `yaml:"uri" json:"uri"`
}

type NodeCidr struct {
	Address string `yaml:"address" json:"address"`
	Name    string `yaml:"name" json:"name"`
}

type SecurityGroup struct {
	Name string `yaml:"name" json:"name"`
}

type KubernetesVersion struct {
	Value string `yaml:"value" json:"value"`
}

type NodePool struct {
	Name       string `yaml:"name" json:"name"`
	Nodes      int    `yaml:"nodes" json:"nodes"`
	Instance   string `yaml:"instance" json:"instance"`
	DataCenter string `yaml:"dataCenter" json:"dataCenter"`
}

type BillingPlan struct {
	BillingPeriod string `yaml:"billingPeriod" json:"billingPeriod"`
}

type Properties struct {
	Preset            bool              `yaml:"preset" json:"preset"`
	VPC               VPC               `yaml:"vpc" json:"vpc"`
	Subnet            Subnet            `yaml:"subnet" json:"subnet"`
	NodeCidr          NodeCidr          `yaml:"nodeCidr" json:"nodeCidr"`
	SecurityGroup     SecurityGroup     `yaml:"securityGroup" json:"securityGroup"`
	KubernetesVersion KubernetesVersion `yaml:"kubernetesVersion" json:"kubernetesVersion"`
	NodePools         []NodePool        `yaml:"nodePools" json:"nodePools"`
	HA                bool              `yaml:"ha" json:"ha"`
	BillingPlan       BillingPlan       `yaml:"billingPlan" json:"billingPlan"`
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
