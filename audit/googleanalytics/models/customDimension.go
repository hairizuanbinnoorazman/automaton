package models

import analytics "google.golang.org/api/analytics/v3"

type CustomDimensionData struct {
	Name                string
	Description         string
	CustomDimensions    []*analytics.CustomDimension
	CustomDimensionList map[string][]CustomDimensionItem
	HasMoreThan0        map[string][]bool
	UsedCustomDim       bool
}

type CustomDimensionItem struct {
	Date           string
	DimensionValue string
	Sessions       int
}

func NewCustomDimensionData() CustomDimensionData {
	return CustomDimensionData{Name: "test", Description: "test"}
}

func (c *CustomDimensionData) checkHasMoreThan0() {
	if len(c.CustomDimensions) > 0 {
		return
	}
	return
}

func (c *CustomDimensionData) checkUsedCustomDim() {
	if len(c.CustomDimensionList) > 0 {
		return
	}
	return
}

func (c *CustomDimensionData) RunAudit() {
	c.checkHasMoreThan0()
	c.checkUsedCustomDim()
}
