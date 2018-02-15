package snapshot

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"testing"

	"golang.org/x/oauth2/google"
)

func TestGetSnapshot(t *testing.T) {
	cred, _ := ioutil.ReadFile("cred.json")
	config, _ := ioutil.ReadFile("config.json")

	type gaConfig struct {
		GaAccountID  string
		GaPropertyID string
		GaViewID     string
	}
	var liveGAConfig gaConfig
	json.Unmarshal(config, &liveGAConfig)

	authConfig, _ := google.JWTConfigFromJSON(cred, "https://www.googleapis.com/auth/analytics", "https://www.googleapis.com/auth/analytics.edit")
	emptyContext := context.Background()
	client := authConfig.Client(emptyContext)

	GetSnapshot(client, liveGAConfig.GaAccountID, liveGAConfig.GaPropertyID, liveGAConfig.GaViewID)
}
