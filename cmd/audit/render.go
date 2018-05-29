package audit

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"path"

	"gitlab.com/hairizuanbinnoorazman/automaton/audit/googleanalytics"

	"gitlab.com/hairizuanbinnoorazman/automaton/audit/googleanalytics/models"
)

// RenderOutput function will be moved from this package to the cmd package.
// Current render output here will be depreciated; rendering should not be done on the domain level
func RenderOutput(w io.Writer, templateFile string, a interface{}) error {
	_, templateFileValue := path.Split(templateFile)
	t := template.Must(template.New(templateFileValue).ParseFiles(templateFile))

	var err error

	switch tempStruct := a.(type) {
	case *models.ProfileData:
		err = t.Execute(w, tempStruct)
	case *models.GoalsData:
		err = t.Execute(w, tempStruct)
	case *models.TrafficSourceData:
		enhancedTempStruct := enhanceTrafficSource(tempStruct)
		err = t.Execute(w, enhancedTempStruct)
	case *models.CustomMetricsData:
		err = t.Execute(w, tempStruct)
	default:
		err := errors.New("Unable to find the type definition of the audit")
		return err
	}

	if err != nil {
		fmt.Println("Unable to render template")
		fmt.Println(err.Error())
		return err
	}
	// Add a few new lines
	io.WriteString(w, "\n\n")
	return nil
}

func RenderAllOutput(w io.Writer, output googleanalytics.AuditorResults, auditConfigList ...AuditItem) error {
	var err error
	for _, auditItem := range auditConfigList {
		if auditItem.Name == models.NewProfileData().Name {
			err = RenderOutput(w, auditItem.TemplateFile, output.ProfileAudit)
		}
		if auditItem.Name == models.NewGoalsData().Name {
			err = RenderOutput(w, auditItem.TemplateFile, output.GoalAudit)
		}
	}

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
