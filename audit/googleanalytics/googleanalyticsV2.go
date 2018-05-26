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
	GetGoalValues(profileID string, goalID string) ([]models.GoalItem, error)
}

type Auditor struct {
	AccountID  string
	PropertyID string
	ProfileID  string
}

type AuditorResults struct {
	GoalAudit *models.GoalsData
}

func (a Auditor) Run(e Extractor) AuditorResults {
	goalAuditor := GoalAuditor{AccountID: a.AccountID, PropertyID: a.PropertyID, ProfileID: a.ProfileID}
	goalResults := goalAuditor.Run(e)
	return AuditorResults{GoalAudit: goalResults}
}

type GoalAuditor struct {
	AccountID  string
	PropertyID string
	ProfileID  string
}

func (g GoalAuditor) Run(e Extractor) *models.GoalsData {
	goalData := models.NewGoalsData()
	goalData.Goals, _ = e.GetGoalSettings(g.AccountID, g.PropertyID, g.ProfileID)
	for _, goalSetting := range goalData.Goals {
		values, _ := e.GetGoalValues(g.ProfileID, goalSetting.Id)
		goalData.GoalList[goalSetting.Id] = values
	}
	goalData.RunAudit()
	return &goalData
}

type CustomDimAuditor struct {
	AccountID  string
	PropertyID string
	ProfileID  string
}
