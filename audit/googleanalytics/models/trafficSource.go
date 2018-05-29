package models

type TrafficSourceData struct {
	Name                     string
	Description              string
	TrafficSources           []TrafficSourceItem
	InconsistentCaseMedium   bool
	InconsistentCaseSource   bool
	InconsistentCaseCampaign bool
}

type TrafficSourceItem struct {
	Medium   string
	Source   string
	Campaign string
	Sessions int
}

func NewTrafficSourceData() TrafficSourceData {
	return TrafficSourceData{Name: "Traffic Sources", Description: "This is a check on traffic sources of the website to ensure that it make sense"}
}

func (t *TrafficSourceData) checkInconsistentCaseMedium() {
	return
}

func (t *TrafficSourceData) checkInconsistentCaseSource() {
	return
}

func (t *TrafficSourceData) checkInconsistentCaseCampaign() {
	return
}

func (t *TrafficSourceData) RunAudit() {
	t.checkInconsistentCaseMedium()
	t.checkInconsistentCaseSource()
	t.checkInconsistentCaseCampaign()
}
