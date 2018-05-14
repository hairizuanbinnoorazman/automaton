package mockservice

import "gitlab.com/hairizuanbinnoorazman/automaton/audit/googleanalytics/models"

func (m MockService) GetCustomDimValues(profileID string) ([]models.CustomDimensionItem, error) {
	return m.CustomDimValues, m.Error
}

func (m MockService) GetCustomMetricValues(profileID string) ([]models.CustomMetricsItem, error) {
	return m.CustomMetricValues, m.Error
}

func (m MockService) GetGoalValues(profileID string) ([]models.GoalItem, error) {
	return m.GoalValues, m.Error
}
