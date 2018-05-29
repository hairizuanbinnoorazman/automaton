package models

import analytics "google.golang.org/api/analytics/v3"

type CustomMetricsData struct {
	Name              string
	Description       string
	CustomMetrics     []*analytics.CustomMetrics
	CustomMetricsList map[string][]CustomMetricsItem
	HasMoreThan0      map[string][]bool
	UsedCustomMetrics bool
	CustomMetricCount int
}

type CustomMetricsItem struct {
	Date              string
	CustomMetricValue int
	Sessions          int
}

func NewCustomMetricData() CustomMetricsData {
	return CustomMetricsData{
		Name:        "Custom Metrics Usage",
		Description: "Usage of the custom metrics feature to track metrics that may not be part of web analytics.",
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
		c.UsedCustomMetrics = true
	}
	c.UsedCustomMetrics = false
}

func (c *CustomMetricsData) RunAudit() {
	c.checkHasMoreThan0()
	c.checkUsedCustomMetrics()
}
