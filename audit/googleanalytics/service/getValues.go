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
						Expression: "ga:eventValue",
					},
				},
			},
		},
	}

	gaDataService := s.getGADataService()
	response, _ := gaDataService.BatchGet(&request).Do()

	eventItems := []models.EventItem{}
	rows := response.Reports[0].Data.Rows

	for _, val := range rows {
		eventValue := val.Metrics[0].Values[0]
		eventValueInt, _ := strconv.Atoi(eventValue)
		singleEventItem := models.EventItem{
			EventCategory: val.Dimensions[0],
			EventAction:   val.Dimensions[1],
			EventLabel:    val.Dimensions[2],
			EventValue:    eventValueInt}
		eventItems = append(eventItems, singleEventItem)
	}

	return eventItems, nil
}
