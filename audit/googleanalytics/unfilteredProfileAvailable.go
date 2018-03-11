package googleanalytics

import (
	analytics "google.golang.org/api/analytics/v3"
)

// UnfilteredProfileAvailableData contains data which is needed to audit and check on availablility
// of profiles that contain profiles with no links
type UnfilteredProfileAvailableData struct {
	Profiles           []*analytics.Profile
	ProfileFilterLinks []*analytics.ProfileFilterLink
}

// UnfilteredProfileAvailableResult contains data which is the output of the analysis
type UnfilteredProfileAvailableResult struct {
	ProfileCount               int  `json:"profile_count"`
	UnfilteredProfileAvailable bool `json:"unfiltered_profile_available"`
}

// UnfilteredProfileAvailable is an object which provides data and functionlity to run and obtain
// information on unfiltered profiles available.
type UnfilteredProfileAvailable struct {
	Metadata metadata
	Data     UnfilteredProfileAvailableData
	Result   UnfilteredProfileAvailableResult
}

// Do function runs the data extraction as well as audit.
// There is a reason on why it needs to be paired together; there are certain aspects which require the data to
// be extracted in a 2 way fashion:
//
// 1. Extract the data from management settings state
//
// 2. Extract the data from Google Analytics data which dependes on the management settings state
//
// To make it flexible, we would need to only expose the Do function. But internally the functionality
// would utilize interfaces to switch between the tests and
func (a *UnfilteredProfileAvailable) Do(mgmtExtractor GaMgmtExtractor, dataExtractor GaDataExtractor) error {
	a.Result = UnfilteredProfileAvailableResult{
		ProfileCount:               len(a.Data.Profiles),
		UnfilteredProfileAvailable: false}
	return nil
}

// NewUnfilteredProfileAvailable creates a new unfiltered profile structure that provides functionality
// to audit mentioned property. There is only one exported function here, which is the Do function.
func NewUnfilteredProfileAvailable() UnfilteredProfileAvailable {
	newUnfilteredProfileAvailable := UnfilteredProfileAvailable{
		Metadata: metadata{
			DataExtractors: dataExtractors{
				GaMgmtProperties: []string{profiles, profileFilterLinks},
			},
			Name:        "Unfiltered Profile Available",
			Description: "Check to see if there is a Google Analytics Profile that has no filters.",
		},
	}
	return newUnfilteredProfileAvailable
}
