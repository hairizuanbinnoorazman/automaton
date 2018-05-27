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
	GetTrafficSourceValues(profileID, startDate, endDate string) ([]models.TrafficSourceItem, error)
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
	EventAuditor         *models.EventsData
	TrafficSourceAuditor *models.TrafficSourceData
}

func (a Auditor) Run(e Extractor) AuditorResults {
	// goalAuditor := GoalAuditor{AccountID: a.AccountID, PropertyID: a.PropertyID, ProfileID: a.ProfileID}
	// goalResults := goalAuditor.Run(e)

	eventAuditor := EventAuditor{ProfileID: a.ProfileID, StartDate: a.StartDate, EndDate: a.EndDate}
	eventResults := eventAuditor.Run(e)

	trafficSourceAuditor := TrafficAuditor{ProfileID: a.ProfileID, StartDate: a.StartDate, EndDate: a.EndDate}
	trafficResults := trafficSourceAuditor.Run(e)
	return AuditorResults{EventAuditor: eventResults, TrafficSourceAuditor: trafficResults}
}

type ProfileAuditor struct {
	AccountID  string
	PropertyID string
	ProfileID  string
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
	StartDate  string
	EndDate    string
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
		fmt.Println(err.Error())
		return &eventsData
	}
	eventsData.Events = temp
	eventsData.RunAudit()
	return &eventsData
}

type TrafficAuditor struct {
	ProfileID string
	StartDate string
	EndDate   string
}

func (t TrafficAuditor) Run(e Extractor) *models.TrafficSourceData {
	trafficSourceData := models.NewTrafficSourceData()
	temp, err := e.GetTrafficSourceValues(t.ProfileID, t.StartDate, t.EndDate)
	if err != nil {
		fmt.Println(err.Error())
		return &trafficSourceData
	}
	trafficSourceData.TrafficSources = temp
	trafficSourceData.RunAudit()
	return &trafficSourceData
}
