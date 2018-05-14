package models

import analytics "google.golang.org/api/analytics/v3"

type customMetricsAuditor struct {
	CustomMetrics     []*analytics.CustomMetrics
	CustomMetricsList []CustomMetricsItem
}

type CustomMetricsItem struct {
	Date              string
	CustomMetricID    string
	CustomMetricValue string
	Sessions          int
}

func (c customMetricsAuditor) HasMoreThan0() bool {
	if len(c.CustomMetrics) > 0 {
		return true
	}
	return false
}

func (c customMetricsAuditor) UsedCustomMetrics() bool {
	if len(c.CustomMetricsList) > 0 {
		return true
	}
	return false
}
