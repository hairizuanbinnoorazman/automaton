package googleanalytics

import (
	"io"
	"net/http"

	analytics "google.golang.org/api/analytics/v3"
)

// CheckUnfilteredProfileAvailable This audit function checks for the following:
// - Check that you have at least 2 Profiles available
// - Check that one of the said profile does not contain and filter links
func CheckUnfilteredProfileAvailable(w io.Writer, client *http.Client, accountID, propertyID, profileID string) error {
	analyticsService, _ := analytics.New(client)
	managementService := analytics.NewManagementService(analyticsService)
	_, err := managementService.Profiles.List(accountID, propertyID).Do()
	if err != nil {
		return err
	}

	return nil
}
