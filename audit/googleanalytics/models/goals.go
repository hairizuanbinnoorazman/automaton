package models

import analytics "google.golang.org/api/analytics/v3"

type goalsAuditor struct {
	Goals    []*analytics.Goal
	GoalList []customGoalItem
}

type customGoalItem struct {
	Date       string
	GoalID     string
	GoalStarts int
}

func (g goalsAuditor) HasMoreThan0() bool {
	if len(g.Goals) > 0 {
		return true
	}
	return false
}

func (g goalsAuditor) UsedGoals() bool {
	if len(g.GoalList) > 0 {
		return true
	}
	return false
}
