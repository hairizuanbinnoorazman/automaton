package mockservice

import (
	"gitlab.com/hairizuanbinnoorazman/automaton/audit/googleanalytics/models"
	analytics "google.golang.org/api/analytics/v3"
)

type MockService struct {
	ProfilesSettings          []*analytics.Profiles
	ProfileFilterLinkSettings []*analytics.ProfileFilterLink
	GoalSettings              []*analytics.Goal
	CustomDimSettings         []*analytics.CustomDimension
	CustomMetricSettings      []*analytics.CustomMetric

	GoalValues         []models.GoalItem
	CustomDimValues    []models.CustomDimensionItem
	CustomMetricValues []models.CustomMetricsItem
}
