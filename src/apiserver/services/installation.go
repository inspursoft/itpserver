package services

import (
	"github.com/inspursoft/itpserver/src/apiserver/dao"
	"github.com/inspursoft/itpserver/src/apiserver/models"
)

type installationConf struct {
	daoHandler dao.InstallationDaoHandler
}

func NewInstallationHandler() *installationConf {
	return new(installationConf)
}

func (ic *installationConf) GetInstalledPackages(vmID string) ([]models.Package, error) {
	return ic.daoHandler.GetInstallPackages(vmID)
}

func (ic *installationConf) Install(vmID string, pkg models.Package) (status bool, err error) {
	affected, err := ic.daoHandler.InstallPackageToVM(&models.VM{VMID: vmID}, &pkg)
	status = (affected == 1)
	return
}

func (ic *installationConf) Delete(vmID string, pkg models.Package) (status bool, err error) {
	affected, err := ic.daoHandler.RemovePackageFromVM(&models.VM{VMID: vmID}, &pkg)
	status = (affected == 1)
	return
}
