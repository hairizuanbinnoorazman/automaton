package googleanalytics_test

import (
	"fmt"
	"testing"

	"gitlab.com/hairizuanbinnoorazman/automaton/audit/googleanalytics"
	analytics "google.golang.org/api/analytics/v3"
)

func TestUnfilteredProfileAvailable(t *testing.T) {
	var x []*analytics.Profile
	var y []*analytics.ProfileFilterLink
	z := googleanalytics.UnfilteredProfileAvailableData{
		Profiles:           x,
		ProfileFilterLinks: y}
	nyaa := googleanalytics.NewUnfilteredProfileAvailable()
	nyaa.Data = z

	nyaa.RunAudit()
	fmt.Printf("%v\n", nyaa.Result)
}
