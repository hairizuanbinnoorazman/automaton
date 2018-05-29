package models

import (
	"strings"
)

type EventsData struct {
	Name                          string
	Description                   string
	Events                        []EventItem
	UniqueEventCount              int
	InconsistentCaseEventCategory bool
	InconsistentCaseEventAction   bool
	InconsistentCaseEventLabel    bool
}

type EventItem struct {
	EventCategory string
	EventAction   string
	EventLabel    string
	Sessions      int
}

func NewEventsData() EventsData {
	return EventsData{Name: "Events", Description: "Usage of event tracking allows one to track activities that may co-relate with a business"}
}

func (e *EventsData) checkUniqueEventCount() {
	e.UniqueEventCount = len(e.Events)
}

// Check that its using small case for easier consistency
func (e *EventsData) checkInconsistentCaseEventCategory() {
	for _, val := range e.Events {
		if strings.ToLower(val.EventCategory) != val.EventCategory {
			e.InconsistentCaseEventCategory = false
			return
		}
	}
	e.InconsistentCaseEventCategory = true
}

func (e *EventsData) checkInconsistentCaseEventAction() {
	for _, val := range e.Events {
		if strings.ToLower(val.EventAction) != val.EventAction {
			e.InconsistentCaseEventAction = false
			return
		}
	}
	e.InconsistentCaseEventAction = true
}

func (e *EventsData) checkInconsistentCaseEventLabel() {
	for _, val := range e.Events {
		if strings.ToLower(val.EventLabel) != val.EventLabel {
			e.InconsistentCaseEventLabel = false
			return
		}
	}
	e.InconsistentCaseEventLabel = true
}

func (e *EventsData) RunAudit() {
	e.checkInconsistentCaseEventCategory()
	e.checkInconsistentCaseEventAction()
	e.checkInconsistentCaseEventCategory()
	e.checkInconsistentCaseEventLabel()
}
