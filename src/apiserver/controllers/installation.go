package controllers

import (
	"github.com/inspursoft/itpserver/src/apiserver/models"
	"github.com/inspursoft/itpserver/src/apiserver/services"
	"github.com/inspursoft/itpserver/src/apiserver/services/ansiblecli"
)

type InstallationController struct {
	BaseController
}

// @Title Get
// @Description Get virtual machines with software packages installed.
// @Param	id		query 	int64	true		"The virtual machine ID which installed software packages."
// @Success 200 {string} 	Successful get virtual machines with software package installed.
// @Failure 400 Bad request.
// @Failure 401 Unauthorized.
// @Failure 403 The resouce specified was forbidden to access.
// @Failure 404 The resource specified was not found.
// @Failure 500 Internal error occurred at server side.
// @router / [get]
func (ic *InstallationController) Get() {
	ID := ic.requiredID("id")
	installations, err := services.NewInstallationHandler().GetInstalledPackages(ID)
	ic.handleError(err)
	ic.Data["json"] = installations
	ic.ServeJSON()
}

// @Title Post
// @Description Install selected software packages onto a virtual machine.
// @Param	id		path 	int64	true		"The virtual machine ID which wants to install software packages."
// @Param	pkg		body 	models.PackageVO	true		"The virtual machine ID which wants to install software packages."
// @Success 200 {string} 	Successful installed software package onto a virtual machine.
// @Failure 400 Bad request.
// @Failure 401 Unauthorized.
// @Failure 403 The resouce specified was forbidden to access.
// @Failure 404 The resource specified was not found.
// @Failure 500 Internal error occurred at server side.
// @router /:id [post]
func (ic *InstallationController) Post() {
	ID := ic.requiredID(":id")
	vm, err := services.NewVMHandler().GetByID(ID)
	ic.handleError(err)
	vmWithSpec := models.VMWithSpec{
		IP: vm.IP, Name: vm.Name, OS: vm.OS,
	}
	var pkg models.PackageVO
	ic.loadRequestBody(&pkg)
	err = ansiblecli.NewClient(ic.Ctx.ResponseWriter).Install(vmWithSpec, []models.PackageVO{pkg})
	ic.handleError(err)
}

// @Title Delete
// @Description Delete selected virtual machine which with software package installed.
// @Param	id	path 	int64	true		"The virtual machine ID to be deleted."
// @Param pkg_name	query	string	true	"The package name to be deleted on VM."
// @Param pkg_tag	query	string	true	"The package tag to be deleted on VM."
// @Success 200 {string} 	Successful deleted virtual machine by ID.
// @Failure 400 Bad request.
// @Failure 401 Unauthorized.
// @Failure 403 The resouce specified was forbidden to access.
// @Failure 404 The resource specified was not found.
// @Failure 500 Internal error occurred at server side.
// @router /:id [delete]
func (ic *InstallationController) Delete() {
	ID := ic.requiredID(":id")
	pkgName := ic.requiredParam("pkg_name")
	pkgTag := ic.requiredParam("pkg_tag")

	handler := services.NewInstallationHandler()
	err := handler.Delete(ID, pkgName, pkgTag)
	ic.handleError(err)
}
