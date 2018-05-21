package googleanalytics

import (
	"gitlab.com/hairizuanbinnoorazman/automaton/audit/googleanalytics/models"
	analytics "google.golang.org/api/analytics/v3"
)

type Extractor interface {
	GetCustomDimSettings(accountID, propertyID, profileID string) ([]*analytics.CustomDimension, error)
	GetCustomMetricSettings(accountID, propertyID, profileID string) ([]*analytics.CustomMetric, error)
	GetGoalSettings(accountID, propertyID, profileID string) ([]*analytics.Goal, error)
	GetProfileSettings(accountID, propertyID, profileID string) ([]*analytics.Profile, error)
	GetProfileLinkSettings(accountID, propertyID, profileID string) ([]*analytics.ProfileFilterLink, error)

	GetCustomDimValues(profileID string) ([]models.CustomDimensionItem, error)
	GetCustomMetricValues(profileID string) ([]models.CustomMetricsItem, error)
	GetGoalValues(profileID string) ([]models.GoalItem, error)
}

type GoalAuditor struct {
	AccountID  string
	PropertyID string
	ProfileID  string
	Name       string
}

func (g GoalAuditor) Run(e Extractor) models.GoalsAuditResults {
	goalData := models.GoalsData{}
	goalData.GoalList, _ = e.GetGoalValues(g.ProfileID)
	goalData.Goals, _ = e.GetGoalSettings(g.AccountID, g.PropertyID, g.ProfileID)
	return goalData.RunAudit()
}

func NewGoalAuditor() GoalAuditor {
	return GoalAuditor{Name: "GoalAudit"}
}
