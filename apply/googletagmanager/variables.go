package googletagmanager

import (
	tagmanager "google.golang.org/api/tagmanager/v1"
)

func CreateGAIDVariable(svc *tagmanager.Service, accountID, containerID, GAID string) (*tagmanager.Variable, error) {
	variableSvc := tagmanager.NewAccountsContainersVariablesService(svc)
	gtmVariable := tagmanager.Variable{
		Name: "UAID",
		Type: "gas",
	}
	createVariableCall := variableSvc.Create(accountID, containerID, &gtmVariable)
	variable, err := createVariableCall.Do()
	return variable, err
}
