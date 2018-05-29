package audit

import (
	"time"

	"gitlab.com/hairizuanbinnoorazman/automaton/audit/googleanalytics/models"
)

// Config lists all the configurations that one would need when running in audit tests
type Config struct {
	AuditItems []AuditItem `json:"audit_items"`
	AccountID  string      `json:"account_id"`
	PropertyID string      `json:"property_id"`
	ProfileID  string      `json:"profile_id"`
	StartDate  string      `json:"start_date"`
	EndDate    string      `json:"end_date"`
}

type AuditItem struct {
	Name         string `json:"name"`
	TemplateFile string `json:"template_file"`
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
	var newAuditItems []AuditItem
	newAuditItems = append(newAuditItems, AuditItem{
		Name:         models.NewProfileData().Name,
		TemplateFile: "./templates/googleAnalyticsAudit/unfiltered_profile_available.md",
	})
	newAuditItems = append(newAuditItems, AuditItem{
		Name:         models.NewGoalsData().Name,
		TemplateFile: "./templates/googleAnalyticsAudit/goal_usage.md",
	})
	newAuditItems = append(newAuditItems, AuditItem{Name: models.NewEventsData().Name})
	newAuditItems = append(newAuditItems, AuditItem{Name: models.NewTrafficSourceData().Name})

	newConfig.AuditItems = newAuditItems
	return newConfig
}
