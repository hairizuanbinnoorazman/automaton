package models

import analytics "google.golang.org/api/analytics/v3"

type customDimensionAuditor struct {
	CustomDimensions    []*analytics.CustomDimension
	CustomDimensionList []CustomDimensionItem
}

type CustomDimensionItem struct {
	Date           string
	DimensionID    string
	DimensionValue string
	Sessions       int
}

func (c customDimensionAuditor) HasMoreThan0() bool {
	if len(c.CustomDimensions) > 0 {
		return true
	}
	return false
}

func (c customDimensionAuditor) UsedCustomDim() bool {
	if len(c.CustomDimensionList) > 0 {
		return true
	}
	return false
}
