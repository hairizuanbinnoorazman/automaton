package googleanalytics

import (
	"gitlab.com/hairizuanbinnoorazman/automaton/audit/googleanalytics/models"
	analytics "google.golang.org/api/analytics/v3"
)

type Extractor interface {
	GetCustomDimSettings(accountID, viewID, profileID string) ([]*analytics.CustomDimension, error)
	GetCustomDimValues(profileID string) ([]models.CustomDimensionItem, error)
	GetCustomMetricSettings(accountID, viewID, profileID string) ([]*analytics.CustomMetric, error)
	GetCustomMetricValues(profileID string) ([]models.CustomMetricsItem, error)
}
