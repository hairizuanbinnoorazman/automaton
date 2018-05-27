package models

import analytics "google.golang.org/api/analytics/v3"

type ProfileData struct {
	Name               string
	Description        string
	Profiles           []*analytics.Profile
	ProfileFilterLinks []*analytics.ProfileFilterLink
	HasMoreThan1       bool
}

func NewProfileData() ProfileData {
	return ProfileData{Name: "test", Description: "test"}
}

func (p *ProfileData) checkHasMoreThan1() {
	if len(p.Profiles) > 1 {
		p.HasMoreThan1 = true
	}
	p.HasMoreThan1 = false
}
