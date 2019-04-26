package services

import "github.com/inspursoft/itpserver/models"

type installationConf struct{}

func NewInstallationHandler() *installationConf {
	return &installationConf{}
}

func (ic *installationConf) Get(vmName string) []models.Installation {
	return nil
}

func (ic *installationConf) Install(vmID string, packages []models.Package) (status bool) {
	status = false
	return
}

func (ic *installationConf) Delete(vmID string) (status bool) {
	status = false
	return
}
