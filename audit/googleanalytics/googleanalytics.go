// Package googleanalytics within the audit package provides the functionality for running Google
// Analytics audits with respect to some of the common best practises in the market. As time goes
// by and if there are new and better ways of looking at the data and if new approaches to data
// appear then, these practises would then slowly be embedded into the package
package googleanalytics

// // Define list of Google Analytics Management Properties
// const profiles = "profiles"
// const filters = "filters"
// const goals = "goals"
// const profileFilterLinks = "profileFilterLinks"
// const customdimensions = "customDimensions"
// const custommetrics = "customMetrics"

// type Audit struct{}

// type metadata struct {
// 	Name           string
// 	Description    string
// 	DataExtractors dataExtractors
// }

// type auditDetails struct {
// 	AccountID  string
// 	PropertyID string
// 	ProfileID  string
// 	StartDate  string
// 	EndDate    string
// 	mgmtClient *http.Client
// 	dataClient *http.Client
// }

// // RenderOutput function will be moved from this package to the cmd package.
// // Current render output here will be depreciated; rendering should not be done on the domain level
// func RenderOutput(w io.Writer, templateFile string, a interface{}) error {
// 	_, templateFileValue := path.Split(templateFile)
// 	t := template.Must(template.New(templateFileValue).ParseFiles(templateFile))

// 	var err error

// 	switch tempStruct := a.(type) {
// 	case *GoalUsage:
// 		err = t.Execute(w, tempStruct)
// 	case *UnfilteredProfileAvailable:
// 		err = t.Execute(w, tempStruct)
// 	default:
// 		err := errors.New("Unable to find the type definition of the audit")
// 		return err
// 	}

// 	if err != nil {
// 		fmt.Println("Unable to render template")
// 		fmt.Println(err.Error())
// 		return err
// 	}
// 	// Add a few new lines
// 	io.WriteString(w, "\n\n")
// 	return nil
// }

// func RenderAllOutput(w io.Writer, config Config, auditOutput Audit) error {
// 	for _, auditItem := range config.AuditItems {
// 		if auditItem.Name == NewUnfilteredProfileAvailable().Metadata.Name {
// 			// err := RenderOutput(w, auditItem.TemplateFile, auditOutput.UnfilteredProfileAvailable)
// 		}
// 		if auditItem.Name == NewGoalUsage().Metadata.Name {
// 			// err := RenderOutput(w, auditItem.TemplateFile, auditOutput.GoalUsage)
// 		}
// 	}

// 	return nil
// }

// func RunAudit(mgmtClient, dataClient *http.Client, config Config) (Audit, error) {
// 	var auditItemNames []string
// 	for _, x := range config.AuditItems {
// 		auditItemNames = append(auditItemNames, x.Name)
// 	}

// 	newAudit := Audit{}

// 	gaMgmtExtractor := NewGaMgmtExtract()
// 	gaDataExtractor := NewGaDataExtract()

// 	for _, item := range config.AuditItems {
// 		if item.Name == NewUnfilteredProfileAvailable().Metadata.Name {
// 			temp := NewUnfilteredProfileAvailableWithParams(
// 				config.AccountID, config.PropertyID, config.ProfileID, mgmtClient)
// 			err := temp.Do(gaMgmtExtractor)
// 			if err != nil {
// 				return Audit{}, err
// 			}
// 			newAudit.UnfilteredProfileAvailable = &temp
// 		}
// 		if item.Name == NewGoalUsage().Metadata.Name {
// 			temp := NewGoalUsageWithParams(
// 				config.AccountID, config.PropertyID, config.ProfileID,
// 				config.StartDate, config.EndDate, mgmtClient, dataClient)
// 			err := temp.Do(gaMgmtExtractor, gaDataExtractor)
// 			if err != nil {
// 				return Audit{}, err
// 			}
// 			newAudit.GoalUsage = &temp
// 		}
// 	}

// 	return newAudit, nil
// }
