package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/spf13/viper"
)

func Authenticate(clientId string, clientSecret string) error {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", clientId)
	data.Set("client_secret", clientSecret)
	// Create a new POST request with the form data
	req, err := http.NewRequest("POST", "https://login.aruba.it/auth/realms/cmp-new-apikey/protocol/openid-connect/token", bytes.NewBufferString(data.Encode()))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}

	// Set the appropriate headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return err
	}

	// Extract access_token from the map
	if accessToken, ok := result["access_token"].(string); ok {
		viper.Set("api.bearer_token", accessToken)
		viper.WriteConfig()
	} else {
		return err
	}

	return nil
}

func SetToken(token string) error {
	viper.Set("api.bearer_token", token)
	err := viper.WriteConfig()
	return err
}

func Logout() error {
	viper.Set("api.bearer_token", "NONE")
	err := viper.WriteConfig()
	return err
}
