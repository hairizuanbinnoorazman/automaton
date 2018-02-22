package googleanalytics_test

import (
	"bytes"
	"testing"

	"gitlab.com/hairizuanbinnoorazman/automaton/audit/googleanalytics"
	"gitlab.com/hairizuanbinnoorazman/automaton/helper"
)

func TestCheckUnfilteredProfileAvailable(t *testing.T) {
	client, _ := helper.GetClient("petaccount")
	gaConfig, _ := helper.GetGAConfig("petaccount")

	buf := bytes.NewBufferString("")
	err = googleanalytics.CheckUnfilteredProfileAvailable(buf, client, gaConfig.AccountID, gaConfig.PropertyID, gaConfig.ProfileID)
	if err != nil {
		t.Fatalf("Expected the following account to be able to provide values. %v", err.Error())
	}
}
