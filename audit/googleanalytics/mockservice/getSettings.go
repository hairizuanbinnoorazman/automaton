package mockservice

import analytics "google.golang.org/api/analytics/v3"

func (m MockExtractor) GetCustomDimSettings(accountID, viewID, profileID string) ([]*analytics.CustomDimension, error) {
	return m.CustomDimSettings, m.Error
}

func (m MockExtractor) GetCustomMetricSettings(accountID, viewID, profileID string) ([]*analytics.CustomMetric, error) {
	return m.CustomMetricSettings, m.Error
}

func (m MockExtractor) GetGoalSettings(accountID, viewID, profileID string) ([]*analytics.Goal, error) {
	return m.GoalSettings, m.Error
}

func (m MockExtractor) GetProfileSettings(accountID, viewID, profileID string) ([]*analytics.Profile, error) {
	return m.ProfilesSettings, m.Error
}

func (m MockExtractor) GetProfileLinkSettings(accountID, viewID, profileID string) ([]*analytics.ProfileFilterLink, error) {
	return m.ProfileFilterLinkSettings, m.Error
}
