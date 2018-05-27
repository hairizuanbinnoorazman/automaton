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
	return ProfileData{
		Name:        "Unfiltered Profile Available",
		Description: "Check to see if there is a Google Analytics Profile that has no filters.",
	}
}

func (p *ProfileData) checkHasMoreThan1() {
	if len(p.Profiles) > 1 {
		p.HasMoreThan1 = true
	}
	p.HasMoreThan1 = false
}

func (p *ProfileData) RunAudit() {
	p.checkHasMoreThan1()
}
