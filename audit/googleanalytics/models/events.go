package models

type eventsAuditor struct {
	events []eventItem
}

type eventItem struct {
	eventCategory string
	eventAction   string
	eventLabel    string
	eventValue    int
}

func (e eventsAuditor) HasMoreThan0() bool {
	return len(e.events) > 0
}

func (e eventItem) InconsistentCaseEventCategory() bool {
	return true
}

func (e eventItem) InconsistentCaseEventAction() bool {
	return true
}

func (e eventItem) InconsistentCaseEventLabel() bool {
	return true
}
