package sdk

import (
	"acloud/config"
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

// Writer is an interface for writing data
type Writer interface {
	Write(v interface{}) error
}

// JSONWriter writes data in JSON format
type JSONWriter struct{}

// Write outputs the data in JSON format
func (w *JSONWriter) Write(v interface{}) error {
	// Handle byte slice input
	if byteData, ok := v.([]byte); ok {
		v = string(byteData)
	}

	// Check if v is a string and attempt to unmarshal it
	if strData, ok := v.(string); ok {
		var unmarshaledData interface{}
		if err := json.Unmarshal([]byte(strData), &unmarshaledData); err != nil {
			// If unmarshaling fails, print the string as is
			fmt.Println(strData)
			return nil
		}
		v = unmarshaledData
	}

	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("JSON marshal error: %w", err)
	}
	fmt.Println(string(data))
	return nil
}

// YAMLWriter writes data in YAML format
type YAMLWriter struct{}

// Write outputs the data in YAML format
func (w *YAMLWriter) Write(v interface{}) error {
	// Handle byte slice input
	if byteData, ok := v.([]byte); ok {
		v = string(byteData)
	}

	// Check if v is a string and attempt to unmarshal it
	if strData, ok := v.(string); ok {
		var unmarshaledData interface{}
		if err := json.Unmarshal([]byte(strData), &unmarshaledData); err != nil {
			// If unmarshaling fails, print the string as is
			fmt.Println(strData)
			return nil
		}
		v = unmarshaledData
	}

	data, err := yaml.Marshal(v)
	if err != nil {
		return fmt.Errorf("YAML marshal error: %w", err)
	}
	fmt.Println(string(data))
	return nil
}

// getWriter returns a Writer object based on the output format specified in config.OutputFormat
func getWriter() Writer {
	switch config.OutputFormat {
	case "yaml":
		return &YAMLWriter{}
	case "json", "": // Default to JSON if not specified
		return &JSONWriter{}
	default:
		return &JSONWriter{} // Default to JSONWriter if format is unrecognized
	}
}

// WriteOutput writes the data using the appropriate writer and handles errors internally
func WriteOutput(data interface{}) {
	// Get the correct writer
	writer := getWriter()

	// Write the data using the selected writer
	if err := writer.Write(data); err != nil {
		fmt.Printf("Error writing data: %v\n", err)
		// Additional error handling can be added here if needed
	}
}
