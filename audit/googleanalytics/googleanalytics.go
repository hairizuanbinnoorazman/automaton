// Package googleanalytics
package googleanalytics

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"path"
	"text/template"

	"gitlab.com/hairizuanbinnoorazman/automaton/audit"
	analytics "google.golang.org/api/analytics/v3"
	analyticsreporting "google.golang.org/api/analyticsreporting/v4"
)

// Define list of Google Analytics Management Properties
const profiles = "profiles"
const filters = "filters"
const goals = "goals"
const profileFilterLinks = "profileFilterLinks"
const customdimensions = "customDimensions"
const custommetrics = "customMetrics"

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

func RenderOutput(w io.Writer, templateFile string, a audit.Auditor) error {
	_, templateFileValue := path.Split(templateFile)
	t := template.Must(template.New(templateFileValue).ParseFiles(templateFile))

	var err error

	switch tempStruct := a.(type) {
	case *UnfilteredProfileAvailable:
		err = t.Execute(w, tempStruct)
	case *GoalUsage:
		err = t.Execute(w, tempStruct)
	case *CustomDimMetricUsage:
		err = t.Execute(w, tempStruct)
	default:
		err := errors.New("Unable to find the type definition of the audit")
		return err
	}

	if err != nil {
		fmt.Println("Unable to render template")
		fmt.Println(err.Error())
		return err
	}
	// Add a few new lines
	io.WriteString(w, "\n\n")
	return nil
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
			err = RenderOutput(w, item.TemplateFile, &temp)
			if err != nil {
				return err
			}
		}
		if item.Name == NewGoalUsage().Metadata.Name {
			temp := NewGoalUsage()
			temp.Data = GoalUsageData{Goals: mgmtData.Goals}
			temp.RunAudit()
			err = RenderOutput(w, item.TemplateFile, &temp)
			if err != nil {
				return err
			}
		}
		if item.Name == NewCustomDimMetricUsage().Metadata.Name {
			temp := NewCustomDimMetricUsage()
			temp.Data = CustomDimMetricUsageData{CustomDimensions: mgmtData.CustomDimensions, CustomMetrics: mgmtData.CustomMetrics}
			temp.RunAudit()
			err = RenderOutput(w, item.TemplateFile, &temp)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
