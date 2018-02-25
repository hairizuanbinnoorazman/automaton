package googleanalytics

import (
	"fmt"
	"io"
	"path"
	"text/template"

	analytics "google.golang.org/api/analytics/v3"
)

type UnfilteredProfileAvailableData struct {
	Profiles           []*analytics.Profile
	ProfileFilterLinks []*analytics.ProfileFilterLink
}

type UnfilteredProfileAvailableResult struct {
	ProfileCount               int  `json:"profile_count"`
	UnfilteredProfileAvailable bool `json:"unfiltered_profile_available"`
}

type UnfilteredProfileAvailable struct {
	Metadata metadata
	Data     UnfilteredProfileAvailableData
	Result   UnfilteredProfileAvailableResult
}

func (a *UnfilteredProfileAvailable) RunAudit() error {
	a.Result = UnfilteredProfileAvailableResult{
		ProfileCount:               2,
		UnfilteredProfileAvailable: false}
	return nil
}

func (a *UnfilteredProfileAvailable) RenderOutput(w io.Writer, templateFile string) error {
	_, templateFileValue := path.Split(templateFile)
	t := template.Must(template.New(templateFileValue).ParseFiles(templateFile))
	err := t.Execute(w, a)
	if err != nil {
		fmt.Println("Unable to render template")
		fmt.Println(err.Error())
		return err
	}
	// Add a few new lines
	io.WriteString(w, "\n\n")
	return nil
}

func NewUnfilteredProfileAvailable() UnfilteredProfileAvailable {
	unfilteredProfileAvailable := UnfilteredProfileAvailable{
		Metadata: metadata{
			DataExtractors: dataExtractors{
				GaMgmtProperties: []string{profiles, profileFilterLinks},
			},
			Name:        "Unfiltered Profile Available",
			Description: "Check to see if there is a Google Analytics Profile that has no filters.",
		},
	}
	return unfilteredProfileAvailable
}
