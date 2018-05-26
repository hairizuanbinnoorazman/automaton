package models

type EventsData struct {
	Name                          string
	Description                   string
	Events                        []EventItem
	InconsistentCaseEventCategory bool
	InconsistentCaseEventAction   bool
	InconsistentCaseEventLabel    bool
}

type EventItem struct {
	EventCategory string
	EventAction   string
	EventLabel    string
	EventValue    int
}

func NewEventsData() EventsData {
	return EventsData{Name: "test", Description: "test"}
}

func (e *EventsData) checkHasMoreThan0() {
	return
}

func (e *EventsData) checkInconsistentCaseEventCategory() {
	return
}

func (e *EventsData) checkInconsistentCaseEventAction() {
	return
}

func (e *EventsData) checkInconsistentCaseEventLabel() {
	return
}

func (e *EventsData) RunAudit() {
	e.checkHasMoreThan0()
	e.checkInconsistentCaseEventAction()
	e.checkInconsistentCaseEventCategory()
	e.checkInconsistentCaseEventLabel()
}
