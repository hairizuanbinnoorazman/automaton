package googletagmanager

import (
	tagmanager "google.golang.org/api/tagmanager/v1"
)

func CreateGAIDVariable(svc *tagmanager.Service, accountID, containerID, GAID string) error {
	variableSvc := tagmanager.NewAccountsContainersVariablesService(svc)
	gtmVariable := tagmanager.Variable{
		Name: "UAID",
		Type: "gas",
	}
	variableSvc.Create(accountID, containerID, &gtmVariable)
	return nil
}
