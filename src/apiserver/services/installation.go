package services

import (
	"fmt"

	"github.com/inspursoft/itpserver/src/apiserver/dao"
	"github.com/inspursoft/itpserver/src/apiserver/models"
)

type installationConf struct {
	daoHandler    dao.InstallationDaoHandler
	daoVMHandler  dao.VMDaoHandler
	svcVMHandler  vmConf
	daoPkgHandler dao.PkgDaoHandler
	e             *models.ITPError
}

func NewInstallationHandler() *installationConf {
	return &installationConf{e: &models.ITPError{}}
}

func (ic *installationConf) GetInstalledPackages(ID int64) ([]models.Package, error) {
	pkgs, err := ic.daoHandler.GetInstallPackages(ID)
	if err != nil {
		ic.e.InternalError(err)
		return nil, ic.e
	}
	return pkgs, nil
}

func (ic *installationConf) getSpecifiedVMPackage(ID int64, pkgName, pkgTag string) (vm *models.VM, pkg *models.Package, err error) {
	vm, err = ic.svcVMHandler.GetByID(ID)
	if err != nil {
		ic.e.InternalError(err)
		return nil, nil, ic.e
	}
	if vm == nil {
		ic.e.Notfound("VM", fmt.Errorf("No VM was found with ID: %d", ID))
		return nil, nil, ic.e
	}
	pkg, err = ic.daoPkgHandler.GetPackage(pkgName, pkgTag)
	if err != nil {
		ic.e.InternalError(err)
		return nil, nil, ic.e
	}
	if pkg == nil {
		ic.e.Notfound("VM", fmt.Errorf("No package was found with name: %s, tag: %s", pkgName, pkgTag))
		return nil, nil, ic.e
	}
	return
}

func (ic *installationConf) Install(ID int64, pkgName, pkgTag string) error {
	vm, pkg, err := ic.getSpecifiedVMPackage(ID, pkgName, pkgTag)
	if err != nil {
		return ic.e
	}
	exists := ic.daoHandler.CheckPackagesInstalledToVM(vm, pkg)
	if exists {
		ic.e.Conflict(fmt.Sprintf("name: %s with tag: %s on ID: %d", pkgName, pkgTag, ID), err)
		return ic.e
	}
	_, err = ic.daoHandler.InstallPackageToVM(vm, pkg)
	if err != nil {
		ic.e.InternalError(err)
		return ic.e
	}
	return nil
}

func (ic *installationConf) Delete(ID int64, pkgName, pkgTag string) error {
	vm, pkg, err := ic.getSpecifiedVMPackage(ID, pkgName, pkgTag)
	if err != nil {
		return ic.e
	}
	_, err = ic.daoHandler.RemovePackageFromVM(vm, pkg)
	if err != nil {
		ic.e.InternalError(err)
		return ic.e
	}
	return nil
}
