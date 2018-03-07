package googleanalytics

import (
	"errors"
	"fmt"
	"net/http"

	analyticsreporting "google.golang.org/api/analyticsreporting/v4"
)

// DataExtractor interface.
// This is the generic interface to ensure that the data being passed around the google analytics methods
// In order to allow differing behaviours during runtime, we would allow the user to provide the client
// struct which would alter the behaviour of the struct - Providing a real client would lead to calling
// APIs but it also becomes possible to provide mocks for testing
type DataExtractor interface {
	Extract(client *http.Client) error
}

// GaMgmtParams struct is the parameters which is used to extract management settings data from
// Google Analytics settings
type GaMgmtParams struct {
	AccountID  string
	PropertyID string
	ProfileID  string
	MgmtItems  []string
}

// GaMgmtExtractor struct is the main struct that would be used to store information after the extract has been
// extracted
type GaMgmtExtractor struct {
	Params  GaMgmtParams
	Results GaMgmtProperties
}

// Extract function attached to the GaMgmtExtractor struct.
// This function would extract the Google Analytics data and then store it into the internal dataset
// Data is not returned but instead pulled out of the initialized struct
func (e *GaMgmtExtractor) Extract(client *http.Client) error {
	accountID := e.Params.AccountID
	propertyID := e.Params.PropertyID
	profileID := e.Params.ProfileID

	mgmtService := getManagementService(client)

	for _, item := range e.Params.MgmtItems {
		if item == profiles {
			profileData, err := mgmtService.Profiles.List(accountID, propertyID).Do()
			if err != nil {
				return err
			}
			e.Results.Profiles = profileData.Items
		}
		if item == goals {
			goalData, err := mgmtService.Goals.List(accountID, propertyID, profileID).Do()
			if err != nil {
				return err
			}
			e.Results.Goals = goalData.Items
		}
		if item == profileFilterLinks {
			profileFilterLinksData, err := mgmtService.ProfileFilterLinks.List(accountID, propertyID, profileID).Do()
			if err != nil {
				return err
			}
			e.Results.ProfileFilterLinks = profileFilterLinksData.Items
		}
	}
	return nil
}

// GaDataParams is the struct for the parameters of the Google Analytics Data extraction
type GaDataParams struct {
	ReportRequest map[string][]*analyticsreporting.ReportRequest
}

// GaDataExtractor is the struct for managing the Google Analytics Data extraction method
type GaDataExtractor struct {
	Params  GaDataParams
	Results map[string][]*analyticsreporting.GetReportsResponse
}

// Extract function attached to the GaDataExtractor struct.
// This function would extract the Google Analytics data and then store it into the internal dataset
// Data is not returned but instead should be pulled out of the initialized struct
func (e *GaDataExtractor) Extract(client *http.Client) error {
	dataService := getGADataService(client)

	for name, req := range e.Params.ReportRequest {
		reportReq := analyticsreporting.GetReportsRequest{
			ReportRequests: []*analyticsreporting.ReportRequest{req[0]},
		}
		response, err := dataService.Reports.BatchGet(&reportReq).Do()
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		if response.HTTPStatusCode != 200 {
			return errors.New("Unable to get values")
		}

		e.Results[name] = append(e.Results[name], response)
	}
	return nil
}
