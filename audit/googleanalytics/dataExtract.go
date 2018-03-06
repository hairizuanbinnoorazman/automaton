package googleanalytics

import (
	"errors"
	"net/http"

	analyticsreporting "google.golang.org/api/analyticsreporting/v4"
)

// DataExtractor interface.
// This is the generic interface to ensure that the data being passed around the google analytics methods
type DataExtractor interface {
	Extract(client *http.Client, params interface{}) error
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
	Data GaMgmtProperties
}

// Extract function attached to the GaMgmtExtractor struct.
// This function would extract the Google Analytics data and then store it into the internal dataset
// Data is not returned but instead pulled out of the initialized struct
func (e *GaMgmtExtractor) Extract(client *http.Client, params interface{}) error {
	mgmtParams := params.(GaMgmtParams)
	accountID := mgmtParams.AccountID
	propertyID := mgmtParams.PropertyID
	profileID := mgmtParams.ProfileID

	mgmtService := getManagementService(client)

	for _, item := range mgmtParams.MgmtItems {
		if item == profiles {
			profileData, err := mgmtService.Profiles.List(accountID, propertyID).Do()
			if err != nil {
				return err
			}
			e.Data.Profiles = profileData.Items
		}
		if item == goals {
			goalData, err := mgmtService.Goals.List(accountID, propertyID, profileID).Do()
			if err != nil {
				return err
			}
			e.Data.Goals = goalData.Items
		}
		if item == profileFilterLinks {
			profileFilterLinksData, err := mgmtService.ProfileFilterLinks.List(accountID, propertyID, profileID).Do()
			if err != nil {
				return err
			}
			e.Data.ProfileFilterLinks = profileFilterLinksData.Items
		}
	}
	return nil
}

// GaDataParams is the struct for the parameters of the Google Analytics Data extraction
type GaDataParams struct {
	ReportRequest []*analyticsreporting.ReportRequest
}

type GaDataExtractor struct {
	Data []*analyticsreporting.GetReportsResponse
}

func (e *GaDataExtractor) Extract(client *http.Client, params interface{}) error {
	gaDataParams := params.(GaDataParams)

	dataService := getGADataService(client)

	for _, req := range gaDataParams.ReportRequest {
		reportReq := analyticsreporting.GetReportsRequest{
			ReportRequests: []*analyticsreporting.ReportRequest{req},
		}
		response, err := dataService.Reports.BatchGet(&reportReq).Do()
		if err != nil {
			return err
		}
		if response.HTTPStatusCode != 200 {
			return errors.New("Unable to get values")
		}

		e.Data = append(e.Data, response)
	}

	return nil
}
