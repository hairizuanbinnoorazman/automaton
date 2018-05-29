package models

import analytics "google.golang.org/api/analytics/v3"

type ProfileData struct {
	Name                       string
	Description                string
	Profiles                   []*analytics.Profile
	ProfileFilterLinks         []*analytics.ProfileFilterLink
	ProfileCount               int
	UnfilteredProfileAvailable bool
}

func NewProfileData() ProfileData {
	return ProfileData{
		Name:        "Unfiltered Profile Available",
		Description: "Check to see if there is a Google Analytics Profile that has no filters.",
	}
}

func (p *ProfileData) checkHasMoreThan1() {
	p.ProfileCount = len(p.Profiles)
}

func (p *ProfileData) RunAudit() {
	p.checkHasMoreThan1()
}
