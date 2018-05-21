package googleanalytics

import (
	"fmt"

	analytics "google.golang.org/api/analytics/v3"
	analyticsreporting "google.golang.org/api/analyticsreporting/v4"
)

type CustomDimUsageData struct {
	CustomDimensions     []*analytics.CustomDimension
	CustomDimensionsData map[string][]*analyticsreporting.GetReportsResponse
	auditDetails
}

type CustomDimUsageResult struct {
	CustomDimensionCount     int
	CustomDimensionCollected map[string]bool
}

type CustomDimUsage struct {
	Metadata metadata
	Data     CustomDimUsageData
	Result   CustomDimUsageResult
}

func (a *CustomDimUsage) Do(mgmtExtractor GaMgmtExtractor, dataExtractor GaDataExtractor) error {
	// Validate input data
	err := a.Data.validate()
	if err != nil {
		return err
	}

	// Extract the GA Management Data
	mgmtData, err := mgmtExtractor.Extract(a.Data.auditDetails.mgmtClient, a.Data.auditDetails.AccountID, a.Data.auditDetails.PropertyID,
		a.Data.auditDetails.ProfileID, a.Metadata.DataExtractors.GaMgmtProperties)
	if err != nil {
		return err
	}

	// Assign required data to Custom Dimensions and Custom Metrics
	a.Data.CustomDimensions = mgmtData.CustomDimensions

	a.Result = CustomDimUsageResult{
		CustomDimensionCount: len(a.Data.CustomDimensions),
	}

	// Initialize the CustomDimensions ga data map
	a.Data.CustomDimensionsData = make(map[string][]*analyticsreporting.GetReportsResponse)

	// Fetching raw GA goals data from Google analytics
	for _, singleCustomDimension := range a.Data.CustomDimensions {
		dimensionExtractionString := fmt.Sprintf("ga:dimension%s", singleCustomDimension.Id)

		request := map[string][]*analyticsreporting.ReportRequest{
			dimensionExtractionString: []*analyticsreporting.ReportRequest{
				&analyticsreporting.ReportRequest{
					DateRanges: []*analyticsreporting.DateRange{
						&analyticsreporting.DateRange{
							StartDate: a.Data.StartDate,
							EndDate:   a.Data.EndDate,
						},
					},
					ViewId: a.Data.ProfileID,
					Dimensions: []*analyticsreporting.Dimension{
						&analyticsreporting.Dimension{
							Name: "ga:date",
						},
					},
					Metrics: []*analyticsreporting.Metric{
						&analyticsreporting.Metric{
							Expression: dimensionExtractionString,
						},
					},
				},
			},
		}

		value, err := dataExtractor.Extract(a.Data.dataClient, request)
		if err != nil {
			return err
		}
		a.Data.CustomDimensionsData[dimensionExtractionString] = append(a.Data.CustomDimensionsData[dimensionExtractionString], value[dimensionExtractionString][0])
	}

	// Initialize the goal ga data collected section
	a.Result.CustomDimensionCollected = make(map[string]bool)

	for customDimensionName, singleDimensionGaData := range a.Data.CustomDimensionsData {
		dimensionUsed := len(singleDimensionGaData) > 0
		a.Result.CustomDimensionCollected[customDimensionName] = dimensionUsed
	}

	return nil
}

func NewCustomDimUsage() CustomDimUsage {
	newGoalUsage := CustomDimUsage{
		Metadata: metadata{
			Name:        "Custom Dimensions and Metrics Usage",
			Description: "Usage of the custom dimension and metrics feature to track metrics that may not be part of web analytics.",
			DataExtractors: dataExtractors{
				GaMgmtProperties: []string{customdimensions, custommetrics},
			},
		},
	}
	return newGoalUsage
}
