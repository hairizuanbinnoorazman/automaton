// Package googleanalytics
package googleanalytics

import (
	"io"
	"net/http"

	analytics "google.golang.org/api/analytics/v3"
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

type gaMgmtProperties struct {
	Profiles           []*analytics.Profile
	Filters            []*analytics.Filter
	ProfileFilterLinks []*analytics.ProfileFilterLink
	Goals              []*analytics.Goal
	CustomDimensions   []*analytics.CustomDimension
	CustomMetrics      []*analytics.CustomMetric
}

func getManagementService(client *http.Client) *analytics.ManagementService {
	analyticsService, _ := analytics.New(client)
	managementService := analytics.NewManagementService(analyticsService)
	return managementService
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

func prepDataExtraction(auditItems []string) (dataExtractors, error) {
	var newDataExtractor dataExtractors
	for _, item := range auditItems {
		if item == NewUnfilteredProfileAvailable().Metadata.Name {
			temp := NewUnfilteredProfileAvailable()
			newDataExtractor = resolveExtractors(newDataExtractor, temp.Metadata.DataExtractors)
		}
	}
	return newDataExtractor, nil
}

func extractGAMgmtData(client *http.Client, mgmtProperties []string, accountID, propertyID, profileID string) (gaMgmtProperties, error) {
	var newGaMgmtProperties gaMgmtProperties
	mgmtService := getManagementService(client)

	for _, item := range mgmtProperties {
		if item == profiles {
			profileData, err := mgmtService.Profiles.List(accountID, propertyID).Do()
			if err != nil {
				var temp gaMgmtProperties
				return temp, err
			}
			newGaMgmtProperties.Profiles = profileData.Items
		}
		if item == goals {
			goalData, err := mgmtService.Goals.List(accountID, propertyID, profileID).Do()
			if err != nil {
				var temp gaMgmtProperties
				return temp, err
			}
			newGaMgmtProperties.Goals = goalData.Items
		}
		if item == profileFilterLinks {
			profileFilterLinksData, err := mgmtService.ProfileFilterLinks.List(accountID, propertyID, profileID).Do()
			if err != nil {
				var temp gaMgmtProperties
				return temp, err
			}
			newGaMgmtProperties.ProfileFilterLinks = profileFilterLinksData.Items
		}
	}
	return newGaMgmtProperties, nil
}

func RunAudit(w io.Writer, client *http.Client, config Config) error {
	var auditItemNames []string
	for _, x := range config.AuditItems {
		auditItemNames = append(auditItemNames, x.Name)
	}
	prep, err := prepDataExtraction(auditItemNames)
	if err != nil {
		return err
	}
	mgmtData, err := extractGAMgmtData(client, prep.GaMgmtProperties, config.AccountID, config.PropertyID, config.ProfileID)
	if err != nil {
		return err
	}

	for _, item := range config.AuditItems {
		if item.Name == NewUnfilteredProfileAvailable().Metadata.Name {
			temp := NewUnfilteredProfileAvailable()
			temp.Data = UnfilteredProfileAvailableData{Profiles: mgmtData.Profiles, ProfileFilterLinks: mgmtData.ProfileFilterLinks}
			temp.RunAudit()
			temp.RenderOutput(w, item.TemplateFile)
		}
	}

	return nil
}
