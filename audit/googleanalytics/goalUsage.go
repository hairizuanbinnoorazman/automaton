package googleanalytics

import (
	"fmt"
	"net/http"

	analytics "google.golang.org/api/analytics/v3"
	analyticsreporting "google.golang.org/api/analyticsreporting/v4"
)

// GoalUsageData defines the list of data points needed to run this audit.
// It may include both management google analytics or data from the google analytics profile
type GoalUsageData struct {
	Goals       []*analytics.Goal
	GoalsGaData map[string][]*analyticsreporting.GetReportsResponse
	auditDetails
}

// GoalUsageResult defines the list of data points that will be provided after running this audit
type GoalUsageResult struct {
	GoalCount           int
	GoalGaDataCollected map[string]bool
}

// GoalUsage is the main object which will interface the data and result structs and provide the RunAudit functionality to run
type GoalUsage struct {
	Metadata metadata
	Data     GoalUsageData
	Result   GoalUsageResult
}

// Do method runs the analysis on the dataset to determine the status of the Goal Usage in the profile
func (a *GoalUsage) Do(mgmtExtractor GaMgmtExtractor, dataExtractor GaDataExtractor) error {
	// Validate input data
	err := a.Data.validate()
	if err != nil {
		return err
	}

	// Extract the GA Management Data
	mgmtData, err := mgmtExtractor.Extract(a.Data.auditDetails.mgmtClient, a.Data.auditDetails.AccountID, a.Data.auditDetails.PropertyID,
		a.Data.auditDetails.ProfileID, a.Metadata.DataExtractors.GaMgmtProperties)
	if err != nil {
		return err
	}

	// Assign required data to Goals
	a.Data.Goals = mgmtData.Goals

	// Initialize the goals ga data map
	a.Data.GoalsGaData = make(map[string][]*analyticsreporting.GetReportsResponse)

	// Calculate number of goals used
	a.Result.GoalCount = len(a.Data.Goals)

	// Fetching raw GA goals data from Google analytics
	for _, singleGoal := range a.Data.Goals {
		goalExtractionString := fmt.Sprintf("ga:goal%sStarts", singleGoal.Id)

		request := map[string][]*analyticsreporting.ReportRequest{
			goalExtractionString: []*analyticsreporting.ReportRequest{
				&analyticsreporting.ReportRequest{
					DateRanges: []*analyticsreporting.DateRange{
						&analyticsreporting.DateRange{
							StartDate: a.Data.StartDate,
							EndDate:   a.Data.EndDate,
						},
					},
					ViewId: a.Data.ProfileID,
					Dimensions: []*analyticsreporting.Dimension{
						&analyticsreporting.Dimension{
							Name: "ga:date",
						},
					},
					Metrics: []*analyticsreporting.Metric{
						&analyticsreporting.Metric{
							Expression: goalExtractionString,
						},
					},
				},
			},
		}

		value, err := dataExtractor.Extract(a.Data.dataClient, request)
		if err != nil {
			return err
		}
		a.Data.GoalsGaData[goalExtractionString] = append(a.Data.GoalsGaData[goalExtractionString], value[goalExtractionString][0])
	}

	// Initialize the goal ga data collected section
	a.Result.GoalGaDataCollected = make(map[string]bool)

	for goalName, singleGoalGaData := range a.Data.GoalsGaData {
		goalUsed := len(singleGoalGaData) > 0
		a.Result.GoalGaDataCollected[goalName] = goalUsed
	}

	return nil
}

// NewGoalUsage function instantiates the goal with a an almost empty new goal usage struct
func NewGoalUsage() GoalUsage {
	newGoalUsage := GoalUsage{
		Metadata: metadata{
			Name:        "Goal Usage",
			Description: "Usage of the goals feature to track certain aspects of website metrics that coincide with a conversion on the website.",
			DataExtractors: dataExtractors{
				GaMgmtProperties: []string{goals},
			},
		},
	}
	return newGoalUsage
}

// NewGoalUsageWithParams function instantiates the goal audit with all required parameters
func NewGoalUsageWithParams(accountID, propertyID, profileID, startDate, endDate string, mgmtClient, dataClient *http.Client) GoalUsage {
	newGoalUsage := NewGoalUsage()
	newAuditDetails := auditDetails{
		AccountID:  accountID,
		PropertyID: propertyID,
		ProfileID:  profileID,
		StartDate:  startDate,
		EndDate:    endDate,
		mgmtClient: mgmtClient,
		dataClient: dataClient,
	}
	newGoalUsage.Data.auditDetails = newAuditDetails
	return newGoalUsage
}
