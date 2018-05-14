package googleanalytics

import (
	"gitlab.com/hairizuanbinnoorazman/automaton/audit/googleanalytics/models"
	analytics "google.golang.org/api/analytics/v3"
)

type Extractor interface {
	GetCustomDimSettings(accountID, viewID, profileID string) ([]*analytics.CustomDimension, error)
	GetCustomMetricSettings(accountID, viewID, profileID string) ([]*analytics.CustomMetric, error)
	GetGoalSettings(accountID, viewID, profileID string) ([]*analytics.Goal, error)
	GetProfileSettings(accountID, viewID, profileID string) ([]*analytics.Profile, error)
	GetProfileLinkSettings(accountID, viewID, profileID string) ([]*analytics.ProfileFilterLink, error)

	GetCustomDimValues(profileID string) ([]models.CustomDimensionItem, error)
	GetCustomMetricValues(profileID string) ([]models.CustomMetricsItem, error)
	GetGoalValues(profileID string) ([]models.GoalItem, error)
}
