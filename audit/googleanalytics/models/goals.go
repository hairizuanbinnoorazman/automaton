package models

import analytics "google.golang.org/api/analytics/v3"

type GoalsData struct {
	Goals        []*analytics.Goal
	GoalList     map[string][]GoalItem
	HasMoreThan0 bool
	UsedGoals    map[string]bool
}

type GoalItem struct {
	Date       string
	GoalID     string
	GoalStarts int
}

func (g *GoalsData) checkHasMoreThan0() {
	if len(g.Goals) > 0 {
		g.HasMoreThan0 = true
		return
	}
	g.HasMoreThan0 = false
}

func (g *GoalsData) checkUsedGoals() {
	for id, value := range g.GoalList {
		if len(value) > 0 {
			g.UsedGoals[id] = true
			return
		}
		g.UsedGoals[id] = false
	}
}

func (g *GoalsData) RunAudit() {
	g.checkHasMoreThan0()
	g.checkUsedGoals()
	return
}
