package mockservice

import "gitlab.com/hairizuanbinnoorazman/automaton/audit/googleanalytics/models"

func (m MockExtractor) GetCustomDimValues(profileID string) ([]models.CustomDimensionItem, error) {
	return m.CustomDimValues, m.Error
}

func (m MockExtractor) GetCustomMetricValues(profileID string) ([]models.CustomMetricsItem, error) {
	return m.CustomMetricValues, m.Error
}

func (m MockExtractor) GetGoalValues(profileID, goalID string) ([]models.GoalItem, error) {
	return m.GoalValues, m.Error
}
