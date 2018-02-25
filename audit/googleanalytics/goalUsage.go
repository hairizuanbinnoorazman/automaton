package googleanalytics

import (
	analytics "google.golang.org/api/analytics/v3"
)

type GoalUsageData struct {
	Goals []*analytics.Goal
}

type GoalUsageResult struct {
	GoalCount int
}

type GoalUsage struct {
	Metadata metadata
	Data     GoalUsageData
	Result   GoalUsageResult
}

func (a *GoalUsage) RunAudit() error {
	a.Result = GoalUsageResult{
		GoalCount: len(a.Data.Goals),
	}
	return nil
}

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
