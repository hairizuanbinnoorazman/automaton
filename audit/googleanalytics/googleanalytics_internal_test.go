package googleanalytics

import (
	"errors"
	"testing"
)

func TestAuditDetailsValidator(t *testing.T) {
	type auditDetailsValidator struct {
		Name   string
		Input  auditDetails
		Output error
	}
	testData := []auditDetailsValidator{
		auditDetailsValidator{
			Name:   "Zeroth Case",
			Input:  auditDetails{},
			Output: errors.New("Missing Fields: Account ID, Property ID, View ID, Start Date, End Date, GA Management Client, GA Data Client"),
		},
		auditDetailsValidator{
			Name:   "Account Case",
			Input:  auditDetails{AccountID: "123456789"},
			Output: errors.New("Missing Fields: Property ID, View ID, Start Date, End Date, GA Management Client, GA Data Client"),
		},
	}

	for _, singleTestCase := range testData {
		details := singleTestCase.Input
		err := details.validate()
		if singleTestCase.Output != nil {
			if err.Error() != singleTestCase.Output.Error() {
				t.Errorf("Case: %v Expected: %v Actual: %v", singleTestCase.Name, singleTestCase.Output.Error(), err.Error())
			}
		} else {
			if err != nil {
				t.Errorf("Case: %v Expected: nil Actual: %v", singleTestCase.Name, err.Error())
			}
		}
	}
}
