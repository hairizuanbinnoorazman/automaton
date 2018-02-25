package googleanalytics

import analytics "google.golang.org/api/analytics/v3"

type CustomDimMetricUsageData struct {
	CustomDimensions []*analytics.CustomDimension
	CustomMetrics    []*analytics.CustomMetric
}

type CustomDimMetricUsageResult struct {
	CustomDimensionCount int
	CustomMetricCount    int
}

type CustomDimMetricUsage struct {
	Metadata metadata
	Data     CustomDimMetricUsageData
	Result   CustomDimMetricUsageResult
}

func (a *CustomDimMetricUsage) RunAudit() error {
	a.Result = CustomDimMetricUsageResult{
		CustomDimensionCount: len(a.Data.CustomDimensions),
		CustomMetricCount:    len(a.Data.CustomMetrics),
	}
	return nil
}

func NewCustomDimMetricUsage() CustomDimMetricUsage {
	newGoalUsage := CustomDimMetricUsage{
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
