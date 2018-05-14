package mockservice

import analytics "google.golang.org/api/analytics/v3"

func (m MockService) GetCustomDimSettings(accountID, viewID, profileID string) ([]*analytics.CustomDimension, error) {
	return nil, nil
}

func (m MockService) GetCustomMetricSettings(accountID, viewID, profileID string) ([]*analytics.CustomMetric, error) {
	return nil, nil
}

func (m MockService) GetGoalSettings(accountID, viewID, profileID string) ([]*analytics.Goal, error) {
	return nil, nil
}

func (m MockService) GetProfileSettings(accountID, viewID, profileID string) ([]*analytics.Profile, error) {
	return nil, nil
}

func (m MockService) GetProfileLinkSettings(accountID, viewID, profileID string) ([]*analytics.ProfileFilterLink, error) {
	return nil, nil
}
