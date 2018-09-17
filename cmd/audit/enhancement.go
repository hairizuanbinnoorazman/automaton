package audit

import (
	"bytes"
	"strconv"

	"github.com/hairizuanbinnoorazman/automaton/audit/googleanalytics/models"
	"github.com/olekukonko/tablewriter"
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

type EnhancedEventsData struct {
	models.EventsData
	EventsStr string
}

func enhanceEvents(a *models.EventsData) *EnhancedEventsData {
	var eventsStr [][]string
	for _, val := range a.Events {
		sessionValue := strconv.Itoa(val.Sessions)
		eventsStr = append(eventsStr, []string{val.EventCategory, val.EventAction, val.EventLabel, sessionValue})
	}

	bufferedStr := bytes.NewBufferString("")
	table := tablewriter.NewWriter(bufferedStr)
	table.SetHeader([]string{"Event Category", "Event Action", "Event Label", "Sessions"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(eventsStr)
	table.Render()

	enhancedItem := EnhancedEventsData{*a, bufferedStr.String()}
	return &enhancedItem
}
