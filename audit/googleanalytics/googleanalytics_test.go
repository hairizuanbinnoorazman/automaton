// +build extractors
package googleanalytics_test

import (
	"io/ioutil"
	"testing"
	"time"

	"gitlab.com/hairizuanbinnoorazman/automaton/audit/googleanalytics"
	"gitlab.com/hairizuanbinnoorazman/automaton/cmd"
	analyticsreporting "google.golang.org/api/analyticsreporting/v4"
)

type TestGaMgmtExtractorInput struct {
	Name   string
	Input  googleanalytics.MgmtParams
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
				ReportRequest: []*analyticsreporting.ReportRequest{
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
	}

	for _, testCase := range testData {
		cred, _ := ioutil.ReadFile("cred.json")
		client := cmd.GoogleAnalyticsReportingAuth(cred)
		input := googleanalytics.GaDataParams{ReportRequest: testCase.Input.ReportRequest}
		lol := googleanalytics.GaDataExtractor{}
		lol.Extract(client, input)
		t.Error(lol.Data[0].HTTPStatusCode)
		t.Error(len(lol.Data[0].Reports))
		t.Error(lol.Data[0].Reports[0].Data.RowCount)
	}
}
