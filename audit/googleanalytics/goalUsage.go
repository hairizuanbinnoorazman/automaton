package googleanalytics

import (
	analytics "google.golang.org/api/analytics/v3"
)

// GoalUsageData defines the list of data points needed to run this audit.
// It may include both management google analytics or data from the google analytics profile
type GoalUsageData struct {
	Goals []*analytics.Goal
}

// GoalUsageResult defines the list of data points that will be provided after running this audit
type GoalUsageResult struct {
	GoalCount int
}

// GoalUsage is the main object which will interface the data and result structs and provide the RunAudit functionality to run
type GoalUsage struct {
	Metadata metadata
	Data     GoalUsageData
	Result   GoalUsageResult
}

// RunAudit method runs the analysis on the dataset to determine the status of the Goal Usage in the profile
func (a *GoalUsage) RunAudit() error {
	a.Result = GoalUsageResult{
		GoalCount: len(a.Data.Goals),
	}
	return nil
}

// NewGoalUsage is a convenience function to create a new GoalUsage struct with predefined properties such as Name, Description etc
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
