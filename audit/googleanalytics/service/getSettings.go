package service

import analytics "google.golang.org/api/analytics/v3"

func (s Service) GetCustomDimSettings(accountID, viewID, profileID string) ([]*analytics.CustomDimension, error) {
	return nil, nil
}

func (s Service) GetCustomMetricSettings(accountID, viewID, profileID string) ([]*analytics.CustomMetric, error) {
	return nil, nil
}

func (s Service) GetGoalSettings(accountID, viewID, profileID string) ([]*analytics.Goal, error) {
	return nil, nil
}

func (s Service) GetProfileSettings(accountID, viewID, profileID string) ([]*analytics.Profile, error) {
	return nil, nil
}

func (s Service) GetProfileLinkSettings(accountID, viewID, profileID string) ([]*analytics.ProfileFilterLink, error) {
	return nil, nil
}
