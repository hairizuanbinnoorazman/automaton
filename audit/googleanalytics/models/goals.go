package models

import analytics "google.golang.org/api/analytics/v3"

type GoalsData struct {
	Name        string
	Description string
	Goals       []*analytics.Goal
	GoalList    map[string][]GoalItem
	GoalCount   int
	UsedGoals   map[string]bool
}

type GoalItem struct {
	Date       string
	GoalStarts int
}

func NewGoalsData() GoalsData {
	return GoalsData{
		Name:        "Goal Usage",
		Description: "Usage of the goals feature to track certain aspects of website metrics that coincide with a conversion on the website."}
}

func (g *GoalsData) checkGoalCount() {
	g.GoalCount = len(g.Goals)
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
	g.checkGoalCount()
	g.checkUsedGoals()
	return
}
