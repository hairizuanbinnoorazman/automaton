package googleanalytics

import (
	analytics "google.golang.org/api/analytics/v3"
)

type UnfilteredProfileAvailableData struct {
	Profiles           []*analytics.Profile
	ProfileFilterLinks []*analytics.ProfileFilterLink
}

type UnfilteredProfileAvailableResult struct {
	ProfileCount               int  `json:"profile_count"`
	UnfilteredProfileAvailable bool `json:"unfiltered_profile_available"`
}

type UnfilteredProfileAvailable struct {
	Metadata metadata
	Data     UnfilteredProfileAvailableData
	Result   UnfilteredProfileAvailableResult
}

func (a *UnfilteredProfileAvailable) RunAudit() error {
	a.Result = UnfilteredProfileAvailableResult{
		ProfileCount:               len(a.Data.Profiles),
		UnfilteredProfileAvailable: false}
	return nil
}

func NewUnfilteredProfileAvailable() UnfilteredProfileAvailable {
	unfilteredProfileAvailable := UnfilteredProfileAvailable{
		Metadata: metadata{
			DataExtractors: dataExtractors{
				GaMgmtProperties: []string{profiles, profileFilterLinks},
			},
			Name:        "Unfiltered Profile Available",
			Description: "Check to see if there is a Google Analytics Profile that has no filters.",
		},
	}
	return unfilteredProfileAvailable
}
