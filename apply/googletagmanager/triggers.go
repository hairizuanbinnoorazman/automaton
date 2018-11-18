package googletagmanager

import tagmanager "google.golang.org/api/tagmanager/v1"

func CreateEventTrigger(svc *tagmanager.Service, accountID, containerID, GAID string) (*tagmanager.Trigger, error) {
	triggerSvc := tagmanager.NewAccountsContainersTriggersService(svc)
	gtmTrigger := tagmanager.Trigger{
		Name: "Test",
	}
	createTriggerCall := triggerSvc.Create(accountID, containerID, &gtmTrigger)
	createdTrigger, err := createTriggerCall.Do()
	return createdTrigger, err
}
