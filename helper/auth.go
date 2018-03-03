package helper

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2/google"
)

type gtmConfig struct {
	Name        string `json:"name"`
	AccountID   string `json:"account_id"`
	ContainerID string `json:"container_id"`
	CredFile    string `json:"cred_file"`
	Scopes      string `json:"scopes"`
}

type GaConfig struct {
	Name       string   `json:"name"`
	AccountID  string   `json:"account_id"`
	PropertyID string   `json:"property_id"`
	ProfileID  string   `json:"profile_id"`
	CredFile   string   `json:"cred_file"`
	Scopes     []string `json:"scopes"`
}

func getGAConfig(name string) (GaConfig, error) {
	configFile := "config.json"
	config, err := ioutil.ReadFile(configFile)
	if err != nil {
		return GaConfig{}, err
	}

	var availableGAConfig []GaConfig
	json.Unmarshal(config, &availableGAConfig)

	currentGAConfig := GaConfig{}
	for _, c := range availableGAConfig {
		if c.Name == name {
			currentGAConfig = c
		}
	}
	return currentGAConfig, nil
}

// GetClient Function that returns a client that can then be used to call Google APIs as needed
func GetClient(name string) (*http.Client, error) {
	currentGAConfig, err := getGAConfig(name)
	if err != nil {
		return nil, err
	}

	cred, err := ioutil.ReadFile(currentGAConfig.CredFile)
	if err != nil {
		return nil, err
	}

	authConfig, _ := google.JWTConfigFromJSON(cred, currentGAConfig.Scopes...)
	emptyContext := context.Background()
	client := authConfig.Client(emptyContext)
	return client, nil
}

// GetGAConfig Function that returns a client that can be used
func GetGAConfig(name string) (GaConfig, error) {
	return getGAConfig(name)
}
