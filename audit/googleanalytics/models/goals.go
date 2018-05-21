package models

import analytics "google.golang.org/api/analytics/v3"

type GoalsData struct {
	Goals    []*analytics.Goal
	GoalList []GoalItem
}

type GoalItem struct {
	Date       string
	GoalID     string
	GoalStarts int
}

type GoalsAuditResults struct {
	HasMoreThan0 bool
	UsedGoals    bool
}

func (g GoalsData) HasMoreThan0() bool {
	if len(g.Goals) > 0 {
		return true
	}
	return false
}

func (g GoalsData) UsedGoals() bool {
	if len(g.GoalList) > 0 {
		return true
	}
	return false
}

func (g GoalsData) RunAudit() GoalsAuditResults {
	return GoalsAuditResults{}
}
