package service

import "gitlab.com/hairizuanbinnoorazman/automaton/audit/googleanalytics/models"

func (s Service) GetCustomDimValues(profileID string) ([]models.CustomDimensionItem, error) {
	return []models.CustomDimensionItem{}, nil
}

func (s Service) GetCustomMetricValues(profileID string) ([]models.CustomMetricsItem, error) {
	return []models.CustomMetricsItem{}, nil
}

func (s Service) GetGoalValues(profileID string) ([]models.GoalItem, error) {
	return []models.GoalItem{}, nil
}
