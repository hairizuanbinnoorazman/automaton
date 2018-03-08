// +build extractors
package googleanalytics_test

import (
	"io/ioutil"
	"log"
	"testing"
	"time"

	"gitlab.com/hairizuanbinnoorazman/automaton/audit/googleanalytics"
	"gitlab.com/hairizuanbinnoorazman/automaton/helper"
	analyticsreporting "google.golang.org/api/analyticsreporting/v4"
)

type TestGaMgmtExtractorInput struct {
	Name   string
	Input  googleanalytics.GaMgmtParams
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
	Name  string
	Input googleanalytics.GaDataParams
}

func TestGaDataExtractor(t *testing.T) {
	testData := []TestGaDataExtractorInput{
		TestGaDataExtractorInput{
			Name: "Initial Case",
			Input: googleanalytics.GaDataParams{
				ReportRequest: map[string][]*analyticsreporting.ReportRequest{
					"test": []*analyticsreporting.ReportRequest{
						&analyticsreporting.ReportRequest{
							DateRanges: []*analyticsreporting.DateRange{
								&analyticsreporting.DateRange{
									StartDate: time.Now().AddDate(0, 0, -14).Format("2006-01-02"),
									EndDate:   time.Now().Format("2006-01-02"),
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
		extractor := googleanalytics.GaDataExtractor{}
		extractor.Params = googleanalytics.GaDataParams{ReportRequest: testCase.Input.ReportRequest}
		extractor.Extract(client)

		// Test the data structures within it
		// Undergoing required changes
	}
}
