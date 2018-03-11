// +build extractors
package googleanalytics_test

import (
	"io/ioutil"
	"log"
	"testing"

	"gitlab.com/hairizuanbinnoorazman/automaton/audit/googleanalytics"
	"gitlab.com/hairizuanbinnoorazman/automaton/helper"
	analyticsreporting "google.golang.org/api/analyticsreporting/v4"
)

type TestGaMgmtExtractorInput struct {
	Name   string
	Output googleanalytics.GaMgmtProperties
}

func TestGaMgmtExtractor(t *testing.T) {
	// testData := []TestGaMgmtExtractorInput{
	// 	TestGaMgmtExtractorInput{
	// 		Name: "InitialCase",
	// 		Input: googleanalytics.MgmtParams{
	// 			AccountID:  "",
	// 			PropertyID: "",
	// 			ProfileID:  "",
	// 			MgmtItems:  []string{"goals"},
	// 		},
	// 		Output: googleanalytics.GaMgmtProperties{},
	// 	},
	// },
}

type TestGaDataExtractorInput struct {
	Name      string
	InputData map[string][]*analyticsreporting.ReportRequest
}

func TestGaDataExtractor(t *testing.T) {
	testData := []TestGaDataExtractorInput{
		TestGaDataExtractorInput{
			Name: "Initial Case",
			InputData: map[string][]*analyticsreporting.ReportRequest{
				"Initial Case": []*analyticsreporting.ReportRequest{
					&analyticsreporting.ReportRequest{
						DateRanges: []*analyticsreporting.DateRange{
							&analyticsreporting.DateRange{
								StartDate: "2018-01-01",
								EndDate:   "2018-01-07",
							},
						},
						ViewId: "",
						Dimensions: []*analyticsreporting.Dimension{
							&analyticsreporting.Dimension{
								Name: "ga:source",
							},
						},
						Metrics: []*analyticsreporting.Metric{
							&analyticsreporting.Metric{
								Expression: "ga:users",
							},
						},
					},
				},
			},
		},
	}

	for _, testCase := range testData {
		// Read the cred file
		cred, err := ioutil.ReadFile("cred.json")
		if err != nil {
			log.Printf("Error in loading file. %s", err.Error())
			return
		}

		// Get the client
		client := helper.GoogleAnalyticsReportingAuth(cred)

		// Prep the data extraction
		extractor := googleanalytics.GaDataExtract{}
		output, err := extractor.Extract(client, testCase.InputData)

		// Test the data structures within it
		// Undergoing required changes

		// Check for the following
		// Now only defined very loosely
		if len(output[testCase.Name][0].Reports[0].Data.Rows) <= 0 {
			t.Errorf("Error in data extraction. No data is extracted for the following parameters")
		}

	}
}
