package service

import analytics "google.golang.org/api/analytics/v3"

func (s Service) GetCustomDimSettings(accountID, propertyID, profileID string) ([]*analytics.CustomDimension, error) {
	return nil, nil
}

func (s Service) GetCustomMetricSettings(accountID, propertyID, profileID string) ([]*analytics.CustomMetric, error) {
	return nil, nil
}

func (s Service) GetGoalSettings(accountID, propertyID, profileID string) ([]*analytics.Goal, error) {
	mgmtService := s.getManagementService()
	goalData, err := mgmtService.Goals.List(accountID, propertyID, profileID).Do()
	if err != nil {
		return []*analytics.Goal{}, err
	}
	return goalData.Items, nil
}

func (s Service) GetProfileSettings(accountID, propertyID, profileID string) ([]*analytics.Profile, error) {
	mgmtService := s.getManagementService()
	profileData, err := mgmtService.Profiles.List(accountID, propertyID).Do()
	if err != nil {
		return []*analytics.Profile{}, err
	}
	return profileData.Items, nil
}

func (s Service) GetProfileLinkSettings(accountID, propertyID, profileID string) ([]*analytics.ProfileFilterLink, error) {
	mgmtService := s.getManagementService()
	profileFilterLinkData, err := mgmtService.ProfileFilterLinks.List(accountID, propertyID, profileID).Do()
	if err != nil {
		return []*analytics.ProfileFilterLink{}, err
	}
	return profileFilterLinkData.Items, nil
}
