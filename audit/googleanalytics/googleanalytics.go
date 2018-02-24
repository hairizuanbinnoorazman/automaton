// Package googleanalytics
package googleanalytics

import (
	analyticsreporting "google.golang.org/api/analyticsreporting/v4"
)

// Define list of Google Analytics Management Properties
const profiles = "profiles"
const filters = "filters"
const goals = "goals"
const profileFilterLinks = "profileFilterLinks"

type dataExtractors struct {
	GaMgmtProperties []string
	GaDataProperties []*analyticsreporting.ReportRequest
}

type metadata struct {
	Name           string
	Description    string
	DataExtractors dataExtractors
}

func resolveExtractors(overallDataExtractor, singleDataExtractor dataExtractors) dataExtractors {
	// Resolve properties for managment properties
	for _, singleVal := range singleDataExtractor.GaMgmtProperties {
		found := false
		for _, overallVal := range overallDataExtractor.GaMgmtProperties {
			if singleVal == overallVal {
				found = true
				break
			}
		}
		if found == false {
			overallDataExtractor.GaMgmtProperties = append(overallDataExtractor.GaMgmtProperties, singleVal)
		}
	}

	// Resolve properties for data extractors
	// Not handled properly yet, as no use case available yet. However, a future expected use case would be available soon
	// A few things to handled -> Like how to map the data accordingly to the audit that requested it
	return overallDataExtractor
}
