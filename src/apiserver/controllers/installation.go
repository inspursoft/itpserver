package controllers

import (
	"fmt"
	"net/http"

	"github.com/inspursoft/itpserver/src/apiserver/models"
	"github.com/inspursoft/itpserver/src/apiserver/services"
	"github.com/inspursoft/itpserver/src/apiserver/services/ansiblecli"
)

type InstallationController struct {
	BaseController
}

// @Title Get
// @Description Get virtual machines with software packages installed.
// @Param Authorization	header	string	true	"Set authorization info."
// @Param	vm_name		query 	string	true		"The virtual machine ID which installed software packages."
// @Success 200 {string} 	Successful get virtual machines with software package installed.
// @Failure 400 Bad request.
// @Failure 401 Unauthorized.
// @Failure 403 The resouce specified was forbidden to access.
// @Failure 404 The resource specified was not found.
// @Failure 500 Internal error occurred at server side.
// @router / [get]
func (ic *InstallationController) Get() {
	vmName := ic.requiredParam("vm_name")
	vm, err := services.NewVMHandler().GetByName(vmName)
	ic.handleError(err)
	if vm == nil {
		ic.CustomAbort(http.StatusNotFound, fmt.Sprintf("VM with name: %s does not exist.", vmName))
	}
	installations, err := services.NewInstallationHandler().GetInstalledPackages(vm.ID)
	ic.handleError(err)
	ic.Data["json"] = installations
	ic.ServeJSON()
}

// @Title Post
// @Description Install selected software packages onto a virtual machine.
// @Param Authorization	header	string	true	"Set authorization info."
// @Param	vm_name		path 	string	true		"The virtual machine name which wants to install software packages."
// @Param	pkg		body 	models.PackageVO	true		"The virtual machine ID which wants to install software packages."
// @Success 200 {string} 	Successful installed software package onto a virtual machine.
// @Failure 400 Bad request.
// @Failure 401 Unauthorized.
// @Failure 403 The resouce specified was forbidden to access.
// @Failure 404 The resource specified was not found.
// @Failure 500 Internal error occurred at server side.
// @router /:vm_name [post]
func (ic *InstallationController) Post() {
	vmName := ic.requiredParam(":vm_name")
	vm, err := services.NewVMHandler().GetByName(vmName)
	ic.handleError(err)
	if vm == nil {
		ic.CustomAbort(http.StatusNotFound, fmt.Sprintf("VM with name: %s does not exist.", vmName))
	}
	vmWithSpec := models.VMWithSpec{
		IP: vm.IP, Name: vm.Name, OS: vm.OS,
	}
	var pkg models.PackageVO
	ic.loadRequestBody(&pkg)
	err = ansiblecli.NewClient(vmWithSpec, pkg, ic.Ctx.ResponseWriter).Install()
	ic.handleError(err)
}

// @Title Delete
// @Description Delete selected virtual machine which with software package installed.
// @Param Authorization	header	string	true	"Set authorization info."
// @Param	vm_name	path 	string	true		"The virtual machine ID to be deleted."
// @Param pkg_name	query	string	true	"The package name to be deleted on VM."
// @Param pkg_tag	query	string	false	"The package tag to be deleted on VM."
// @Success 200 {string} 	Successful deleted virtual machine by ID.
// @Failure 400 Bad request.
// @Failure 401 Unauthorized.
// @Failure 403 The resouce specified was forbidden to access.
// @Failure 404 The resource specified was not found.
// @Failure 500 Internal error occurred at server side.
// @router /:vm_name [delete]
func (ic *InstallationController) Delete() {
	vmName := ic.requiredParam(":vm_name")
	vm, err := services.NewVMHandler().GetByName(vmName)
	ic.handleError(err)
	if vm == nil {
		ic.CustomAbort(http.StatusNotFound, fmt.Sprintf("VM with name: %s does not exist.", vmName))
	}
	pkgName := ic.requiredParam("pkg_name")
	pkgTag := ic.GetString("pkg_tag", "")

	handler := services.NewInstallationHandler()
	err = handler.Delete(vm.ID, pkgName, pkgTag)
	ic.handleError(err)
}
