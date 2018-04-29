package googleanalytics

import (
	"fmt"
	"strings"
	"time"
)

// Config lists all the configurations that one would need when running in audit tests
type Config struct {
	AuditItems []auditItem `json:"audit_items"`
	AccountID  string      `json:"account_id"`
	PropertyID string      `json:"property_id"`
	ProfileID  string      `json:"profile_id"`
	StartDate  string      `json:"start_date"`
	EndDate    string      `json:"end_date"`
}

type auditItem struct {
	Name         string `json:"name"`
	TemplateFile string `json:"template_file"`
}

func createFileName(auditName string) string {
	auditName = strings.ToLower(auditName)
	auditName = strings.Replace(auditName, " ", "_", -1)
	filename := fmt.Sprintf("./audit/googleanalytics/templates/%v.md", auditName)
	return filename
}

// NewConfig returns a new configuration file that has default values which one can use to modify and append
func NewConfig() Config {
	newConfig := Config{
		AccountID:  "123456789",
		PropertyID: "UA-123456789-1",
		ProfileID:  "1234567890",
		StartDate:  time.Now().Format("2006-01-02"),
		EndDate:    time.Now().AddDate(0, 0, -14).Format("2006-01-02"),
	}

	// Append the new audit items here - we will utilize the name generated from the default struct
	var newAuditItems []auditItem
	newAuditItems = append(newAuditItems, auditItem{Name: NewUnfilteredProfileAvailable().Metadata.Name, TemplateFile: createFileName(NewUnfilteredProfileAvailable().Metadata.Name)})
	newAuditItems = append(newAuditItems, auditItem{Name: NewGoalUsage().Metadata.Name, TemplateFile: createFileName(NewGoalUsage().Metadata.Name)})
	newAuditItems = append(newAuditItems, auditItem{Name: NewCustomDimUsage().Metadata.Name, TemplateFile: createFileName(NewCustomDimUsage().Metadata.Name)})
	newAuditItems = append(newAuditItems, auditItem{Name: NewCustomMetricUsage().Metadata.Name, TemplateFile: createFileName(NewCustomMetricUsage().Metadata.Name)})

	newConfig.AuditItems = newAuditItems
	return newConfig
}
