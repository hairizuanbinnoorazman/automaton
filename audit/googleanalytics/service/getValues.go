package service

import "gitlab.com/hairizuanbinnoorazman/automaton/audit/googleanalytics/models"

func (s Extractor) GetCustomDimValues(profileID string) ([]models.CustomDimensionItem, error) {
	return []models.CustomDimensionItem{}, nil
}

func (s Extractor) GetCustomMetricValues(profileID string) ([]models.CustomMetricsItem, error) {
	return []models.CustomMetricsItem{}, nil
}

func (s Extractor) GetGoalValues(profileID string) ([]models.GoalItem, error) {
	return []models.GoalItem{}, nil
}
