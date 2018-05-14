package models

import analytics "google.golang.org/api/analytics/v3"

type profileAuditor struct {
	Profiles           []*analytics.Profile
	ProfileFilterLinks []*analytics.ProfileFilterLink
}

func (p profileAuditor) HasMoreThan1() bool {
	if len(p.Profiles) > 1 {
		return true
	}
	return false
}
