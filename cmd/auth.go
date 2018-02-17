package cmd

import (
	"context"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2/google"
)

// googleAnalyticsAuth A helper function that help out in
func googleAnalyticsAuth(credFile string) *http.Client {
	cred, _ := ioutil.ReadFile(credFile)
	authConfig, _ := google.JWTConfigFromJSON(cred, "https://www.googleapis.com/auth/analytics", "https://www.googleapis.com/auth/analytics.edit")
	emptyContext := context.Background()
	client := authConfig.Client(emptyContext)
	return client
}
