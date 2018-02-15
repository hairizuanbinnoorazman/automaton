package snapshot

import (
	"fmt"
	"net/http"

	"google.golang.org/api/analytics/v3"
)

func getManagementService(client *http.Client) *analytics.ManagementService {
	analyticsService, _ := analytics.New(client)
	managementService := analytics.NewManagementService(analyticsService)
	return managementService
}

// GetSnapshot exported function
func GetSnapshot(client *http.Client, gaAccountID string, gaPropertyID string, gaViewID string) {
	mgmtService := getManagementService(client)

	filters, err := mgmtService.Filters.List(gaAccountID).Do()
	if err != nil {
		fmt.Println(err.Error())
	}
	data, _ := filters.MarshalJSON()
	fmt.Println(string(data))

	profiles, err := mgmtService.Profiles.List(gaAccountID, gaPropertyID).Do()
	if err != nil {
		fmt.Println(err.Error())
	}
	data, err = profiles.MarshalJSON()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(data))
}
