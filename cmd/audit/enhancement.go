package audit

import (
	"bytes"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"gitlab.com/hairizuanbinnoorazman/automaton/audit/googleanalytics/models"
)

type EnhancedTrafficSourceData struct {
	models.TrafficSourceData
	TrafficSourceDataStr string
}

func enhanceTrafficSource(a *models.TrafficSourceData) *EnhancedTrafficSourceData {
	var trafficSourceStr [][]string
	for _, val := range a.TrafficSources {
		sessionValue := strconv.Itoa(val.Sessions)
		trafficSourceStr = append(trafficSourceStr, []string{val.Medium, val.Source, val.Campaign, sessionValue})
	}

	bufferedStr := bytes.NewBufferString("")
	table := tablewriter.NewWriter(bufferedStr)
	table.SetHeader([]string{"Medium", "Source", "Campaign", "Sessions"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(trafficSourceStr)
	table.Render()

	enhancedItem := EnhancedTrafficSourceData{*a, bufferedStr.String()}
	return &enhancedItem
}
