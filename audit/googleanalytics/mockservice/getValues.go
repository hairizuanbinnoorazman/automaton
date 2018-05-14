package mockservice

import "gitlab.com/hairizuanbinnoorazman/automaton/audit/googleanalytics/models"

func (m MockService) GetCustomDimValues(profileID string) ([]models.CustomDimensionItem, error) {
	return []models.CustomDimensionItem{}, nil
}

func (m MockService) GetCustomMetricValues(profileID string) ([]models.CustomMetricsItem, error) {
	return []models.CustomMetricsItem{}, nil
}

func (m MockService) GetGoalValues(profileID string) ([]models.GoalItem, error) {
	return []models.GoalItem{}, nil
}
