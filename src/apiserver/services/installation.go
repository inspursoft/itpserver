package services

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/inspursoft/itpserver/src/apiserver/dao"
	"github.com/inspursoft/itpserver/src/apiserver/models"
)

type installationConf struct {
	daoHandler    dao.InstallationDaoHandler
	daoVMHandler  dao.VMDaoHandler
	daoPkgHandler dao.PkgDaoHandler
}

func NewInstallationHandler() *installationConf {
	return new(installationConf)
}

func (ic *installationConf) GetInstalledPackages(vmID string) ([]models.Package, error) {
	return ic.daoHandler.GetInstallPackages(vmID)
}

func (ic *installationConf) getSpecifiedVMPackage(vmID, pkgName, pkgTag string) (vm *models.VM, pkg *models.Package, err error) {
	vms, err := ic.daoVMHandler.GetVM(vmID)
	if len(vms) == 0 {
		beego.Error(fmt.Sprintf("No VM found with VMID: %s", vmID))
		err = fmt.Errorf(fmt.Sprintf("No VM found with VMID: %s", vmID))
		return
	}
	pkgs, err := ic.daoPkgHandler.GetPackage(pkgName, pkgTag)
	if len(pkgs) == 0 {
		beego.Error(fmt.Sprintf("No package found with name: %s, tag: %s", pkg.Name, pkg.Tag))
		err = fmt.Errorf(fmt.Sprintf("No package found with name: %s, tag: %s", pkg.Name, pkg.Tag))
		return
	}
	vm = vms[0]
	pkg = pkgs[0]
	return
}

func (ic *installationConf) Install(vmID, pkgName, pkgTag string) (status bool, err error) {
	vm, pkg, err := ic.getSpecifiedVMPackage(vmID, pkgName, pkgTag)
	if err != nil {
		return
	}
	affected, err := ic.daoHandler.InstallPackageToVM(vm, pkg)
	status = (err == nil && affected == 1)
	return
}

func (ic *installationConf) Delete(vmID, pkgName, pkgTag string) (status bool, err error) {
	vm, pkg, err := ic.getSpecifiedVMPackage(vmID, pkgName, pkgTag)
	if err != nil {
		return
	}
	affected, err := ic.daoHandler.RemovePackageFromVM(vm, pkg)
	status = (err == nil && affected == 1)
	return
}
