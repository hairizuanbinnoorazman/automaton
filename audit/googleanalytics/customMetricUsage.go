package googleanalytics

import (
	"fmt"
	"net/http"

	analytics "google.golang.org/api/analytics/v3"
	analyticsreporting "google.golang.org/api/analyticsreporting/v4"
)

type CustomMetricUsageData struct {
	CustomMetrics     []*analytics.CustomMetric
	CustomMetricsData map[string][]*analyticsreporting.GetReportsResponse
	auditDetails
}

type CustomMetricUsageResult struct {
	CustomMetricCount      int
	CustomMetricsCollected map[string]bool
}

type CustomMetricUsage struct {
	Metadata metadata
	Data     CustomMetricUsageData
	Result   CustomMetricUsageResult
}

func (a *CustomMetricUsage) Do(mgmtExtractor GaMgmtExtractor, dataExtractor GaDataExtractor) error {
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
	a.Data.CustomMetrics = mgmtData.CustomMetrics

	a.Result = CustomMetricUsageResult{
		CustomMetricCount: len(a.Data.CustomMetrics),
	}

	// Initialize the CustomDimensions ga data map
	a.Data.CustomMetricsData = make(map[string][]*analyticsreporting.GetReportsResponse)

	// Fetching raw GA goals data from Google analytics
	for _, singleCustomMetric := range a.Data.CustomMetrics {
		metricExtractionString := fmt.Sprintf("ga:dimension%s", singleCustomMetric.Id)

		request := map[string][]*analyticsreporting.ReportRequest{
			metricExtractionString: []*analyticsreporting.ReportRequest{
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
							Expression: metricExtractionString,
						},
					},
				},
			},
		}

		value, err := dataExtractor.Extract(a.Data.dataClient, request)
		if err != nil {
			return err
		}
		a.Data.CustomMetricsData[metricExtractionString] = append(a.Data.CustomMetricsData[metricExtractionString], value[metricExtractionString][0])
	}

	// Initialize the goal ga data collected section
	a.Result.CustomMetricsCollected = make(map[string]bool)

	for customDimensionName, singleDimensionGaData := range a.Data.CustomMetricsData {
		dimensionUsed := len(singleDimensionGaData) > 0
		a.Result.CustomMetricsCollected[customDimensionName] = dimensionUsed
	}

	return nil
}

func NewCustomMetricUsage() CustomMetricUsage {
	newMetricUsage := CustomMetricUsage{
		Metadata: metadata{
			Name:        "Custom Metrics Usage",
			Description: "Usage of the custom metrics feature to track metrics that may not be part of web analytics.",
			DataExtractors: dataExtractors{
				GaMgmtProperties: []string{customdimensions, custommetrics},
			},
		},
	}
	return newMetricUsage
}

func NewCustomMetricUsageWithParams(accountID, propertyID, profileID, startDate, endDate string, mgmtClient, dataClient *http.Client) CustomMetricUsage {
	newCustomMetricUsage := NewCustomMetricUsage()
	newAuditDetails := auditDetails{
		AccountID:  accountID,
		PropertyID: propertyID,
		ProfileID:  profileID,
		StartDate:  startDate,
		EndDate:    endDate,
		mgmtClient: mgmtClient,
		dataClient: dataClient,
	}
	newCustomMetricUsage.Data.auditDetails = newAuditDetails
	return newCustomMetricUsage
}
