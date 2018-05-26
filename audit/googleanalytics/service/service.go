package service

import (
	"net/http"

	analytics "google.golang.org/api/analytics/v3"
	analyticsreporting "google.golang.org/api/analyticsreporting/v4"
)

type Extractor struct {
	Client *http.Client
}

func (s Extractor) getManagementService() *analytics.ManagementService {
	analyticsService, _ := analytics.New(s.Client)
	managementService := analytics.NewManagementService(analyticsService)
	return managementService
}

func (s Extractor) getGADataService() *analyticsreporting.ReportsService {
	analyticsDataService, _ := analyticsreporting.New(s.Client)
	analyticsReportingService := analyticsreporting.NewReportsService(analyticsDataService)
	return analyticsReportingService
}
