package models

import analytics "google.golang.org/api/analytics/v3"

type CustomMetricsData struct {
	Name              string
	Description       string
	CustomMetrics     []*analytics.CustomMetrics
	CustomMetricsList []CustomMetricsItem
	HasMoreThan0      map[string][]bool
	UsedCustomMetrics bool
}

type CustomMetricsItem struct {
	Date              string
	CustomMetricID    string
	CustomMetricValue string
	Sessions          int
}

func NewCustomMetricData() CustomMetricsData {
	return CustomMetricsData{
		Name:        "",
		Description: "",
	}
}

func (c *CustomMetricsData) checkHasMoreThan0() {
	if len(c.CustomMetrics) > 0 {
		return
	}
	return
}

func (c *CustomMetricsData) checkUsedCustomMetrics() {
	if len(c.CustomMetricsList) > 0 {
		return
	}
	return
}

func (c *CustomMetricsData) RunAudit() {
	c.checkHasMoreThan0()
	c.checkUsedCustomMetrics()
}
