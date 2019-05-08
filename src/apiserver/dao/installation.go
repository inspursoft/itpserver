package dao

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/inspursoft/itpserver/src/apiserver/models"
)

type InstallationDaoHandler int

func (ins *InstallationDaoHandler) GetInstallPackages(vmID string) (pkgList []models.Package, err error) {
	o := orm.NewOrm()
	count, err := o.QueryTable("package").Filter("VMs__VM__VMID", vmID).All(&pkgList)
	if err != nil {
		return
	}
	beego.Info(fmt.Sprintf("Successful got %d package(s) from VM with ID: %s", count, vmID))
	return
}

func (ins *InstallationDaoHandler) CheckPackagesInstalledToVM(vm *models.VM, pkg *models.Package) (exists bool) {
	o := orm.NewOrm()
	m2m := o.QueryM2M(vm, "packages")
	exists = m2m.Exist(pkg)
	return
}

func (ins *InstallationDaoHandler) InstallPackageToVM(vm *models.VM, pkg *models.Package) (affected int64, err error) {
	o := orm.NewOrm()
	m2m := o.QueryM2M(vm, "packages")
	affected, err = m2m.Add(pkg)
	if err != nil {
		if err == orm.ErrNoRows {
			return 0, nil
		}
		affected = 0
		return
	}
	beego.Info(fmt.Sprintf("Successful inserted package %v to VM: %v", *pkg, *vm))
	return
}

func (ins *InstallationDaoHandler) RemovePackageFromVM(vm *models.VM, pkg *models.Package) (affected int64, err error) {
	o := orm.NewOrm()
	m2m := o.QueryM2M(vm, "packages")
	affected, err = m2m.Remove(pkg)
	if err != nil {
		if err == orm.ErrNoRows {
			return 0, nil
		}
		affected = 0
		return
	}
	beego.Info(fmt.Sprintf("Successful removed package %v to VM: %v", *pkg, *vm))
	return
}
