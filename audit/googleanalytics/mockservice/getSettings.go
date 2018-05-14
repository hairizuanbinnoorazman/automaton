package mockservice

import analytics "google.golang.org/api/analytics/v3"

func (m MockService) GetCustomDimSettings(accountID, viewID, profileID string) ([]*analytics.CustomDimension, error) {
	return m.CustomDimSettings, m.Error
}

func (m MockService) GetCustomMetricSettings(accountID, viewID, profileID string) ([]*analytics.CustomMetric, error) {
	return m.CustomMetricSettings, m.Error
}

func (m MockService) GetGoalSettings(accountID, viewID, profileID string) ([]*analytics.Goal, error) {
	return m.GoalSettings, m.Error
}

func (m MockService) GetProfileSettings(accountID, viewID, profileID string) ([]*analytics.Profile, error) {
	return m.ProfilesSettings, m.Error
}

func (m MockService) GetProfileLinkSettings(accountID, viewID, profileID string) ([]*analytics.ProfileFilterLink, error) {
	return m.ProfileFilterLinkSettings, m.Error
}
