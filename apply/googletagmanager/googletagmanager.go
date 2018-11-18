// Package googletagmanager This package handles the logic to set the setting on the
// google tag manager
package googletagmanager

import (
	"net/http"

	tagmanager "google.golang.org/api/tagmanager/v1"
)

func NewTagManagerService(oauthHttpClient *http.Client) (*tagmanager.Service, error) {
	tagmanagerService, err := tagmanager.New(oauthHttpClient)
	if err != nil {
		return nil, err
	}
	return tagmanagerService, nil
}
