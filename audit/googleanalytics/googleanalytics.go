// Package googleanalytics within the audit package provides the functionality for running Google
// Analytics audits with respect to some of the common best practises in the market. As time goes
// by and if there are new and better ways of looking at the data and if new approaches to data
// appear then, these practises would then slowly be embedded into the package
package googleanalytics

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"path"
	"text/template"

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

type Audit struct {
	UnfilteredProfileAvailable *UnfilteredProfileAvailable `json:"unfiltered_profile_available,omitempty"`
	GoalUsage                  *GoalUsage                  `json:"goal_usage,omitempty"`
}

type metadata struct {
	Name           string
	Description    string
	DataExtractors dataExtractors
}

type auditDetails struct {
	AccountID  string
	PropertyID string
	ProfileID  string
	StartDate  string
	EndDate    string
	mgmtClient *http.Client
	dataClient *http.Client
}

func (a *auditDetails) validate() error {
	var errList []string
	if a.AccountID == "" {
		errList = append(errList, "Account ID")
	}
	if a.PropertyID == "" {
		errList = append(errList, "Property ID")
	}
	if a.ProfileID == "" {
		errList = append(errList, "View ID")
	}
	if a.StartDate == "" {
		errList = append(errList, "Start Date")
	}
	if a.EndDate == "" {
		errList = append(errList, "End Date")
	}
	if a.mgmtClient == nil {
		errList = append(errList, "GA Management Client")
	}
	if a.mgmtClient == nil {
		errList = append(errList, "GA Data Client")
	}
	if len(errList) > 0 {
		errText := "Missing Fields: "
		numMissingFields := len(errList)
		for idx, missingField := range errList {
			if idx < (numMissingFields - 1) {
				errText = errText + missingField + ", "
			} else {
				errText = errText + missingField
			}
		}
		return errors.New(errText)
	}
	return nil
}

type dataExtractors struct {
	GaMgmtProperties []string
	GaDataProperties map[string][]*analyticsreporting.ReportRequest
}

type GaMgmtProperties struct {
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

func getGADataService(client *http.Client) *analyticsreporting.Service {
	analyticsDataService, _ := analyticsreporting.New(client)
	return analyticsDataService
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

func extractGAMgmtData(client *http.Client, mgmtProperties []string, accountID, propertyID, profileID string) (GaMgmtProperties, error) {
	var newGaMgmtProperties GaMgmtProperties
	mgmtService := getManagementService(client)

	for _, item := range mgmtProperties {
		if item == profiles {
			profileData, err := mgmtService.Profiles.List(accountID, propertyID).Do()
			if err != nil {
				var temp GaMgmtProperties
				return temp, err
			}
			newGaMgmtProperties.Profiles = profileData.Items
		}
		if item == goals {
			goalData, err := mgmtService.Goals.List(accountID, propertyID, profileID).Do()
			if err != nil {
				var temp GaMgmtProperties
				return temp, err
			}
			newGaMgmtProperties.Goals = goalData.Items
		}
		if item == profileFilterLinks {
			profileFilterLinksData, err := mgmtService.ProfileFilterLinks.List(accountID, propertyID, profileID).Do()
			if err != nil {
				var temp GaMgmtProperties
				return temp, err
			}
			newGaMgmtProperties.ProfileFilterLinks = profileFilterLinksData.Items
		}
	}
	return newGaMgmtProperties, nil
}

// RenderOutput function will be moved from this package to the cmd package.
// Current render output here will be depreciated; rendering should not be done on the domain level
func RenderOutput(w io.Writer, templateFile string, a interface{}) error {
	_, templateFileValue := path.Split(templateFile)
	t := template.Must(template.New(templateFileValue).ParseFiles(templateFile))

	var err error

	switch tempStruct := a.(type) {
	case *GoalUsage:
		err = t.Execute(w, tempStruct)
	case *UnfilteredProfileAvailable:
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

func RenderAllOutput(w io.Writer, config Config, auditOutput Audit) error {
	for _, auditItem := range config.AuditItems {
		if auditItem.Name == NewUnfilteredProfileAvailable().Metadata.Name {
			err := RenderOutput(w, auditItem.TemplateFile, auditOutput.UnfilteredProfileAvailable)
			if err != nil {
				return err
			}
		}
		if auditItem.Name == NewGoalUsage().Metadata.Name {
			err := RenderOutput(w, auditItem.TemplateFile, auditOutput.GoalUsage)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func RunAudit(mgmtClient, dataClient *http.Client, config Config) (Audit, error) {
	var auditItemNames []string
	for _, x := range config.AuditItems {
		auditItemNames = append(auditItemNames, x.Name)
	}

	newAudit := Audit{}

	gaMgmtExtractor := NewGaMgmtExtract()
	gaDataExtractor := NewGaDataExtract()

	for _, item := range config.AuditItems {
		if item.Name == NewUnfilteredProfileAvailable().Metadata.Name {
			temp := NewUnfilteredProfileAvailableWithParams(
				config.AccountID, config.PropertyID, config.ProfileID, mgmtClient)
			err := temp.Do(gaMgmtExtractor)
			if err != nil {
				return Audit{}, err
			}
			newAudit.UnfilteredProfileAvailable = &temp
		}
		if item.Name == NewGoalUsage().Metadata.Name {
			temp := NewGoalUsageWithParams(
				config.AccountID, config.PropertyID, config.ProfileID,
				config.StartDate, config.EndDate, mgmtClient, dataClient)
			err := temp.Do(gaMgmtExtractor, gaDataExtractor)
			if err != nil {
				return Audit{}, err
			}
			newAudit.GoalUsage = &temp
		}
	}

	return newAudit, nil
}
