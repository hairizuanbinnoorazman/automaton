package googleanalytics_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"gitlab.com/hairizuanbinnoorazman/automaton/audit/googleanalytics"
	analytics "google.golang.org/api/analytics/v3"
)

type TestGoalUsageList struct {
	Name           string
	InputData      googleanalytics.GoalUsageData
	ExpectedOutput googleanalytics.GoalUsageResult
}

func TestGoalUsage(t *testing.T) {
	testLists := []TestGoalUsageList{
		TestGoalUsageList{
			Name:           "Zeroth Test",
			InputData:      googleanalytics.GoalUsageData{Goals: []*analytics.Goal{}},
			ExpectedOutput: googleanalytics.GoalUsageResult{GoalCount: 0},
		},
		TestGoalUsageList{
			Name:           "One goal",
			InputData:      googleanalytics.GoalUsageData{Goals: []*analytics.Goal{&analytics.Goal{}}},
			ExpectedOutput: googleanalytics.GoalUsageResult{GoalCount: 1},
		},
	}

	for _, singleTest := range testLists {
		newGoalUsage := googleanalytics.NewGoalUsage()
		newGoalUsage.Data = singleTest.InputData
		newGoalUsage.RunAudit()
		equalityTest := reflect.DeepEqual(newGoalUsage.Result, singleTest.ExpectedOutput)
		if equalityTest == false {
			expectedValue, _ := json.MarshalIndent(singleTest.ExpectedOutput, "", "\t")
			actualValue, _ := json.MarshalIndent(newGoalUsage.Result, "", "\t")
			t.Errorf("Error in executing the following test: %v. \nExpected Value: %v. \nActual Value: %v",
				singleTest.Name, string(expectedValue), string(actualValue))
		}
	}
}
