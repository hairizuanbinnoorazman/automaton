package cmd

import (
	"context"
	"net/http"

	"golang.org/x/oauth2/google"
)

// googleAnalyticsAuth A helper function that help out in
func googleAnalyticsAuth(cred []byte) *http.Client {
	authConfig, _ := google.JWTConfigFromJSON(cred, "https://www.googleapis.com/auth/analytics", "https://www.googleapis.com/auth/analytics.edit")
	emptyContext := context.Background()
	client := authConfig.Client(emptyContext)
	return client
}


