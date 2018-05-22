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

type Auditor struct {
	AccountID  string
	PropertyID string
	ProfileID  string
}

type AuditorResults struct {
	GoalAudit *models.GoalsAuditResults
}

func (a Auditor) Run(e Extractor) AuditorResults {
	goalAuditor := GoalAuditor{AccountID: a.AccountID, PropertyID: a.PropertyID, ProfileID: a.ProfileID}
	goalResults := goalAuditor.Run(e)
	return AuditorResults{GoalAudit: &goalResults}
}

type GoalAuditor struct {
	AccountID   string
	PropertyID  string
	ProfileID   string
	Name        string
	Description string
}

func (g GoalAuditor) Run(e Extractor) models.GoalsAuditResults {
	goalData := models.GoalsData{}
	goalData.GoalList, _ = e.GetGoalValues(g.ProfileID)
	goalData.Goals, _ = e.GetGoalSettings(g.AccountID, g.PropertyID, g.ProfileID)
	return goalData.RunAudit()
}

func NewGoalAuditor() GoalAuditor {
	return GoalAuditor{Name: "GoalAudit", Description: "Usage of the goals feature to track certain aspects of website metrics that coincide with a conversion on the website."}
}
