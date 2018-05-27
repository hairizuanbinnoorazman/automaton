package service

import (
	"fmt"
	"strconv"

	"gitlab.com/hairizuanbinnoorazman/automaton/audit/googleanalytics/models"
	analyticsreporting "google.golang.org/api/analyticsreporting/v4"
)

func (s Extractor) GetCustomDimValues(profileID, startDate, endDate, customDimID string) ([]models.CustomDimensionItem, error) {
	return []models.CustomDimensionItem{}, nil
}

func (s Extractor) GetCustomMetricValues(profileID, startDate, endDate, customMetricID string) ([]models.CustomMetricsItem, error) {
	return []models.CustomMetricsItem{}, nil
}

func (s Extractor) GetTrafficSourceValues(profileID, startDate, endDate string) ([]models.TrafficSourceItem, error) {
	request := analyticsreporting.GetReportsRequest{
		ReportRequests: []*analyticsreporting.ReportRequest{
			&analyticsreporting.ReportRequest{
				DateRanges: []*analyticsreporting.DateRange{
					&analyticsreporting.DateRange{
						StartDate: startDate,
						EndDate:   endDate,
					},
				},
				ViewId: profileID,
				Dimensions: []*analyticsreporting.Dimension{
					&analyticsreporting.Dimension{
						Name: "ga:medium",
					},
					&analyticsreporting.Dimension{
						Name: "ga:source",
					},
					&analyticsreporting.Dimension{
						Name: "ga:campaign",
					},
				},
				Metrics: []*analyticsreporting.Metric{
					&analyticsreporting.Metric{
						Expression: "ga:sessions",
					},
				},
			},
		},
	}

	gaDataService := s.getGADataService()
	response, err := gaDataService.BatchGet(&request).Do()
	if err != nil {
		fmt.Println("ERRROR")
		fmt.Println(err.Error())
		return []models.TrafficSourceItem{}, err
	}

	trafficSourceItems := []models.TrafficSourceItem{}
	rows := response.Reports[0].Data.Rows

	for _, val := range rows {
		sessionValue := val.Metrics[0].Values[0]
		sessionValueInt, _ := strconv.Atoi(sessionValue)
		singleTrafficSourceItem := models.TrafficSourceItem{
			Medium:   val.Dimensions[0],
			Source:   val.Dimensions[1],
			Campaign: val.Dimensions[2],
			Sessions: sessionValueInt}
		trafficSourceItems = append(trafficSourceItems, singleTrafficSourceItem)
	}

	return trafficSourceItems, nil
}

func (s Extractor) GetGoalValues(profileID, startDate, endDate, goalID string) ([]models.GoalItem, error) {
	goalExtractionString := fmt.Sprintf("ga:goal%sStarts", goalID)

	request := analyticsreporting.GetReportsRequest{
		ReportRequests: []*analyticsreporting.ReportRequest{
			&analyticsreporting.ReportRequest{
				DateRanges: []*analyticsreporting.DateRange{
					&analyticsreporting.DateRange{
						StartDate: startDate,
						EndDate:   endDate,
					},
				},
				ViewId: profileID,
				Dimensions: []*analyticsreporting.Dimension{
					&analyticsreporting.Dimension{
						Name: "ga:date",
					},
				},
				Metrics: []*analyticsreporting.Metric{
					&analyticsreporting.Metric{
						Expression: goalExtractionString,
					},
				},
			},
		},
	}

	gaDataService := s.getGADataService()
	response, _ := gaDataService.BatchGet(&request).Do()

	goalItems := []models.GoalItem{}
	rows := response.Reports[0].Data.Rows

	for _, val := range rows {
		goalStartValue := val.Metrics[0].Values[0]
		goalStartInt, _ := strconv.Atoi(goalStartValue)
		singleGoalItem := models.GoalItem{Date: val.Dimensions[0], GoalStarts: goalStartInt}
		goalItems = append(goalItems, singleGoalItem)
	}

	return goalItems, nil
}

func (s Extractor) GetEventValues(profileID, startDate, endDate string) ([]models.EventItem, error) {
	request := analyticsreporting.GetReportsRequest{
		ReportRequests: []*analyticsreporting.ReportRequest{
			&analyticsreporting.ReportRequest{
				DateRanges: []*analyticsreporting.DateRange{
					&analyticsreporting.DateRange{
						StartDate: startDate,
						EndDate:   endDate,
					},
				},
				ViewId: profileID,
				Dimensions: []*analyticsreporting.Dimension{
					&analyticsreporting.Dimension{
						Name: "ga:eventCategory",
					},
					&analyticsreporting.Dimension{
						Name: "ga:eventAction",
					},
					&analyticsreporting.Dimension{
						Name: "ga:eventLabel",
					},
				},
				Metrics: []*analyticsreporting.Metric{
					&analyticsreporting.Metric{
						Expression: "ga:sessions",
					},
				},
			},
		},
	}

	gaDataService := s.getGADataService()
	response, err := gaDataService.BatchGet(&request).Do()
	if err != nil {
		fmt.Println("ERRROR")
		fmt.Println(err.Error())
		return []models.EventItem{}, err
	}

	eventItems := []models.EventItem{}
	rows := response.Reports[0].Data.Rows

	for _, val := range rows {
		sessionValue := val.Metrics[0].Values[0]
		sessionValueInt, _ := strconv.Atoi(sessionValue)
		singleEventItem := models.EventItem{
			EventCategory: val.Dimensions[0],
			EventAction:   val.Dimensions[1],
			EventLabel:    val.Dimensions[2],
			Sessions:      sessionValueInt}
		eventItems = append(eventItems, singleEventItem)
	}

	return eventItems, nil
}
