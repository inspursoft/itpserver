package dao

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/inspursoft/itpserver/src/apiserver/models"
)

type installationDaoHandler int

func NewInstallationDaoHandler() *installationDaoHandler {
	return new(installationDaoHandler)
}

func InstallPackageToVM(vm models.VM, pkg models.Package) (affected int64, err error) {
	o := orm.NewOrm()
	m2m := o.QueryM2M(&vm, "packages")
	affected, err = m2m.Add(&pkg)
	if err != nil {
		affected = 0
		return
	}
	beego.Info(fmt.Sprintf("Successful inserted package %v to VM: %v", pkg, vm))
	return
}

func RemovePackageFromVM(vm models.VM, pkg models.Package) (affected int64, err error) {
	o := orm.NewOrm()
	m2m := o.QueryM2M(&vm, "packages")
	affected, err = m2m.Remove(&pkg)
	if err != nil {
		affected = 0
		return
	}
	beego.Info(fmt.Sprintf("Successful removed package %v to VM: %v", pkg, vm))
	return
}
