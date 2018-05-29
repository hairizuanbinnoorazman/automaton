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
	ProfileAudit       *models.ProfileData
	GoalAudit          *models.GoalsData
	EventAudit         *models.EventsData
	TrafficSourceAudit *models.TrafficSourceData
	CustomDimAudit     *models.CustomDimensionData
	CustomMetricAudit  *models.CustomMetricsData
}

func contains(arr []string, val string) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func (a Auditor) Run(e Extractor, auditList ...string) AuditorResults {
	if len(auditList) == 0 {
		auditList = []string{"profile", "goal", "event", "trafficSource"}
	}

	var profileResults *models.ProfileData
	var goalResults *models.GoalsData
	var eventResults *models.EventsData
	var trafficResults *models.TrafficSourceData

	if contains(auditList, "profile") {
		profileAuditor := ProfileAuditor{AccountID: a.AccountID, PropertyID: a.PropertyID, ProfileID: a.ProfileID}
		profileResults = profileAuditor.Run(e)
	}

	if contains(auditList, "goal") {
		goalAuditor := GoalAuditor{AccountID: a.AccountID, PropertyID: a.PropertyID, ProfileID: a.ProfileID}
		goalResults = goalAuditor.Run(e)
	}

	if contains(auditList, "event") {
		eventAuditor := EventAuditor{ProfileID: a.ProfileID, StartDate: a.StartDate, EndDate: a.EndDate}
		eventResults = eventAuditor.Run(e)
	}

	if contains(auditList, "trafficSource") {
		trafficSourceAuditor := TrafficAuditor{ProfileID: a.ProfileID, StartDate: a.StartDate, EndDate: a.EndDate}
		trafficResults = trafficSourceAuditor.Run(e)
	}

	return AuditorResults{
		ProfileAudit:       profileResults,
		GoalAudit:          goalResults,
		EventAudit:         eventResults,
		TrafficSourceAudit: trafficResults,
	}
}

type ProfileAuditor struct {
	AccountID  string
	PropertyID string
	ProfileID  string
}

func (p ProfileAuditor) Run(e Extractor) *models.ProfileData {
	profileData := models.NewProfileData()
	profileData.Profiles, _ = e.GetProfileSettings(p.AccountID, p.PropertyID, p.ProfileID)
	profileData.ProfileFilterLinks, _ = e.GetProfileLinkSettings(p.AccountID, p.PropertyID, p.ProfileID)
	profileData.RunAudit()
	return &profileData
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

func (c CustomDimAuditor) Run(e Extractor) *models.CustomDimensionData {
	customDimensionsData := models.NewCustomDimensionData()
	customDimensionsData.CustomDimensions, _ = e.GetCustomDimSettings(c.AccountID, c.PropertyID, c.ProfileID)
	for _, customDimSetting := range customDimensionsData.CustomDimensions {
		values, _ := e.GetCustomDimValues(c.ProfileID, c.StartDate, c.EndDate, customDimSetting.Id)
		customDimensionsData.CustomDimensionList[customDimSetting.Id] = values
	}
	customDimensionsData.RunAudit()
	return &customDimensionsData
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
