package googleanalytics_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"gitlab.com/hairizuanbinnoorazman/automaton/audit/googleanalytics"
	analytics "google.golang.org/api/analytics/v3"
)

type TestUnfilteredProfileAvailableList struct {
	Name           string
	InputData      googleanalytics.UnfilteredProfileAvailableData
	ExpectedOutput googleanalytics.UnfilteredProfileAvailableResult
}

func TestUnfilteredProfileAvailable(t *testing.T) {
	testLists := []TestUnfilteredProfileAvailableList{
		TestUnfilteredProfileAvailableList{
			Name:           "Zeroth Test",
			InputData:      googleanalytics.UnfilteredProfileAvailableData{Profiles: []*analytics.Profile{}, ProfileFilterLinks: []*analytics.ProfileFilterLink{}},
			ExpectedOutput: googleanalytics.UnfilteredProfileAvailableResult{ProfileCount: 0, UnfilteredProfileAvailable: false},
		},
	}

	for _, singleTest := range testLists {
		newUnfilteredProfileAvailable := googleanalytics.NewUnfilteredProfileAvailable()
		newUnfilteredProfileAvailable.Data = singleTest.InputData
		newUnfilteredProfileAvailable.RunAudit()
		equalityTest := reflect.DeepEqual(newUnfilteredProfileAvailable.Result, singleTest.ExpectedOutput)
		if equalityTest == false {
			expectedValue, _ := json.MarshalIndent(singleTest.ExpectedOutput, "", "\t")
			actualValue, _ := json.MarshalIndent(newUnfilteredProfileAvailable.Result, "", "\t")
			t.Errorf("Error in executing the following test: %v. \nExpected Value: %v. \nActual Value: %v",
				singleTest.Name, string(expectedValue), string(actualValue))
		}
	}
}
