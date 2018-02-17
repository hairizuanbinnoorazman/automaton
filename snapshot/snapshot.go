package snapshot

import (
	"encoding/json"
	"fmt"
	"net/http"

	"google.golang.org/api/analytics/v3"
)

type snapshot struct {
	Profiles         []*analytics.Profile
	Filters          []*analytics.Filter
	FilterLinks      []*analytics.ProfileFilterLink
	Goals            []*analytics.Goal
	CustomDimensions []*analytics.CustomDimension
	CustomMetrics    []*analytics.CustomMetric
}

func getManagementService(client *http.Client) *analytics.ManagementService {
	analyticsService, _ := analytics.New(client)
	managementService := analytics.NewManagementService(analyticsService)
	return managementService
}

// GetSnapshot exported function
func GetSnapshot(client *http.Client, gaAccountID string, gaPropertyID string, gaViewID string) []byte {
	mgmtService := getManagementService(client)

	profiles, err := mgmtService.Profiles.List(gaAccountID, gaPropertyID).Do()
	if err != nil {
		fmt.Println(err.Error())
	}

	filters, err := mgmtService.Filters.List(gaAccountID).Do()
	if err != nil {
		fmt.Println(err.Error())
	}

	filterLinks, err := mgmtService.ProfileFilterLinks.List(gaAccountID, gaPropertyID, gaViewID).Do()
	if err != nil {
		fmt.Println("Error during profile filter link extraction")
		fmt.Println(err.Error())
	}

	goals, err := mgmtService.Goals.List(gaAccountID, gaPropertyID, gaViewID).Do()
	if err != nil {
		fmt.Println("Error during goals extraction")
		fmt.Println(err.Error())
	}

	customDimensions, err := mgmtService.CustomDimensions.List(gaAccountID, gaPropertyID).Do()
	if err != nil {
		fmt.Println("Error during Custom Dimensions extraction")
		fmt.Println(err.Error())
	}

	customMetrics, err := mgmtService.CustomMetrics.List(gaAccountID, gaPropertyID).Do()
	if err != nil {
		fmt.Println("Error during Custom Metrics extraction")
		fmt.Println(err.Error())
	}

	currentSnapshot := snapshot{
		profiles.Items,
		filters.Items,
		filterLinks.Items,
		goals.Items,
		customDimensions.Items,
		customMetrics.Items,
	}

	output, err := json.MarshalIndent(currentSnapshot, "", "\t")
	return output
}
