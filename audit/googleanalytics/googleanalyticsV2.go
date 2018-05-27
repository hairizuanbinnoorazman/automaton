package googleanalytics

import (
	"fmt"

	"gitlab.com/hairizuanbinnoorazman/automaton/audit/googleanalytics/models"
	analytics "google.golang.org/api/analytics/v3"
)

type Extractor interface {
	GetCustomDimSettings(accountID, propertyID, profileID string) ([]*analytics.CustomDimension, error)
	GetCustomMetricSettings(accountID, propertyID, profileID string) ([]*analytics.CustomMetric, error)
	GetGoalSettings(accountID, propertyID, profileID string) ([]*analytics.Goal, error)
	GetProfileSettings(accountID, propertyID, profileID string) ([]*analytics.Profile, error)
	GetProfileLinkSettings(accountID, propertyID, profileID string) ([]*analytics.ProfileFilterLink, error)

	GetCustomDimValues(profileID, startDate, endDate, customDimID string) ([]models.CustomDimensionItem, error)
	GetCustomMetricValues(profileID, startDate, endDate, customMetricID string) ([]models.CustomMetricsItem, error)
	GetGoalValues(profileID, startDate, endDate, goalID string) ([]models.GoalItem, error)
	GetEventValues(profileID, startDate, endDate string) ([]models.EventItem, error)
}

type Auditor struct {
	AccountID  string
	PropertyID string
	ProfileID  string
	StartDate  string
	EndDate    string
}

type AuditorResults struct {
	// GoalAudit    *models.GoalsData
	EventAuditor *models.EventsData
}

func (a Auditor) Run(e Extractor) AuditorResults {
	// goalAuditor := GoalAuditor{AccountID: a.AccountID, PropertyID: a.PropertyID, ProfileID: a.ProfileID}
	// goalResults := goalAuditor.Run(e)

	eventAuditor := EventAuditor{ProfileID: a.ProfileID, StartDate: a.StartDate, EndDate: a.EndDate}
	eventResults := eventAuditor.Run(e)
	return AuditorResults{EventAuditor: eventResults}
}

type GoalAuditor struct {
	AccountID  string
	PropertyID string
	ProfileID  string
	StartDate  string
	EndDate    string
}

func (g GoalAuditor) Run(e Extractor) *models.GoalsData {
	goalData := models.NewGoalsData()
	goalData.Goals, _ = e.GetGoalSettings(g.AccountID, g.PropertyID, g.ProfileID)
	for _, goalSetting := range goalData.Goals {
		values, _ := e.GetGoalValues(g.ProfileID, g.StartDate, g.EndDate, goalSetting.Id)
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

type EventAuditor struct {
	ProfileID string
	StartDate string
	EndDate   string
}

func (a EventAuditor) Run(e Extractor) *models.EventsData {
	eventsData := models.NewEventsData()
	temp, err := e.GetEventValues(a.ProfileID, a.StartDate, a.EndDate)
	if err != nil {
		fmt.Println("Error!!!")
		fmt.Println(err.Error())
		return &eventsData
	}
	eventsData.Events = temp
	eventsData.RunAudit()
	return &eventsData
}
