package service

import (
	"net/http"

	analytics "google.golang.org/api/analytics/v3"
	analyticsreporting "google.golang.org/api/analyticsreporting/v4"
)

type Service struct {
	Client *http.Client
}

func (s Service) getManagementService() *analytics.ManagementService {
	analyticsService, _ := analytics.New(s.Client)
	managementService := analytics.NewManagementService(analyticsService)
	return managementService
}

func (s Service) getGADataService() *analyticsreporting.Service {
	analyticsDataService, _ := analyticsreporting.New(s.Client)
	return analyticsDataService
}
