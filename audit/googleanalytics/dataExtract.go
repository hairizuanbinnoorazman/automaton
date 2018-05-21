package googleanalytics

import (
	"net/http"

	analyticsreporting "google.golang.org/api/analyticsreporting/v4"
)

// GaMgmtExtractor interface.
// This is the generic interface to ensure that the data being passed around the google analytics methods
// In order to allow differing behaviours during runtime, we would allow the user to provide the client
// struct which would alter the behaviour of the struct - Providing a real client would lead to calling
// APIs but it also becomes possible to provide mocks for testing
type GaMgmtExtractor interface {
	Extract(client *http.Client,
		accountID string,
		propertyID string,
		profileID string,
		mgmtItems []string) (GaMgmtProperties, error)
}

// GaMgmtExtract struct is the main struct that would be used to store information after the extract has been
// extracted
type GaMgmtExtract struct{}

// Extract function attached to the GaMgmtExtractor struct.
// This function would extract the Google Analytics data and then store it into the internal dataset
// Data is not returned but instead pulled out of the initialized struct
func (e *GaMgmtExtract) Extract(client *http.Client, accountID string,
	propertyID string, profileID string, mgmtItems []string) (GaMgmtProperties, error) {

	// Define struct to store all data
	mgmtProperty := GaMgmtProperties{}
	return mgmtProperty, nil
}

// NewGaMgmtExtract returns a pointer to the data extract struct.
// The pointer to the struct is to be given rather than a non-pointer version as
// it is impossible for Interfaces to define pointer functions
// We need to return pointers to struct which implement said functions
func NewGaMgmtExtract() *GaMgmtExtract {
	return &GaMgmtExtract{}
}

// GaDataExtractor interface is an interface which requires a struct to use a slightly altered version of the
// data extractor interface.
type GaDataExtractor interface {
	Extract(client *http.Client,
		reportRequest map[string][]*analyticsreporting.ReportRequest) (map[string][]*analyticsreporting.GetReportsResponse,
		error)
}

// GaDataExtract is the struct for managing the Google Analytics Data extraction method
type GaDataExtract struct{}

// Extract function attached to the GaDataExtractor struct.
// This function would extract the Google Analytics data and then store it into the internal dataset
// Data is not returned but instead should be pulled out of the initialized struct
func (e *GaDataExtract) Extract(client *http.Client,
	reportRequest map[string][]*analyticsreporting.ReportRequest) (map[string][]*analyticsreporting.GetReportsResponse,
	error) {

	// Initialize the results map
	results := make(map[string][]*analyticsreporting.GetReportsResponse)

	return results, nil
}

// NewGaDataExtract returns a pointer to the data extract struct.
// The pointer to the struct is to be given rather than a non-pointer version as
// it is impossible for Interfaces to define pointer functions
// We need to return pointers to struct which implement said functions
func NewGaDataExtract() *GaDataExtract {
	return &GaDataExtract{}
}
