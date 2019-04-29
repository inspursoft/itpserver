package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/inspursoft/itpserver/src/apiserver/models"
	"github.com/inspursoft/itpserver/src/apiserver/services"
)

type InstallationController struct {
	beego.Controller
}

// @Title Get
// @Description Get virtual machines with software packages installed.
// @Param	vm_name		query 	string	false		"The virtual machine name which installed software packages."
// @Success 200 {string} 	Successful get virtual machines with software package installed.
// @Failure 400 Bad request.
// @Failure 401 Unauthorized.
// @Failure 403 The resouce specified was forbidden to access.
// @Failure 404 The resource specified was not found.
// @Failure 500 Internal error occurred at server side.
// @router / [get]
func (ic *InstallationController) Get() {
	vmName := ic.GetString("vm_name", "")
	handler := services.NewInstallationHandler()
	installations := handler.Get(vmName)
	ic.Data["JSON"] = installations
	ic.ServeJSON()
}

// @Title Post
// @Description Install selected software packages onto a virtual machine.
// @Param	vm_id		path 	string	false		"The virtual machine ID which wants to install software packages."
// @Success 200 {string} 	Successful installed software package onto a virtual machine.
// @Failure 400 Bad request.
// @Failure 401 Unauthorized.
// @Failure 403 The resouce specified was forbidden to access.
// @Failure 404 The resource specified was not found.
// @Failure 500 Internal error occurred at server side.
// @router /:vm_id [get]
func (ic *InstallationController) Post() {
	vmID := ic.GetString(":vm_id")
	var pkgs []models.Package
	err := json.Unmarshal(ic.Ctx.Input.RequestBody, &pkgs)
	if err != nil {
		ic.CustomAbort(http.StatusInternalServerError, "Failed to unmarshal request data.")
	}
	handler := services.NewInstallationHandler()
	status := handler.Install(vmID, pkgs)
	if !status {
		ic.CustomAbort(http.StatusExpectationFailed, fmt.Sprintf("Failed to create packages to VM ID: %s", vmID))
	}
}

// @Title Delete
// @Description Delete selected virtual machine which with software package installed.
// @Param	vm_id	path 	string	true		"The virtual machine ID to be deleted."
// @Success 200 {string} 	Successful deleted virtual machine by ID.
// @Failure 400 Bad request.
// @Failure 401 Unauthorized.
// @Failure 403 The resouce specified was forbidden to access.
// @Failure 404 The resource specified was not found.
// @Failure 500 Internal error occurred at server side.
// @router /:vm_id [get]
func (ic *InstallationController) Delete() {
	vmID := ic.GetString(":vm_id")
	handler := services.NewInstallationHandler()
	status := handler.Delete(vmID)
	if !status {
		ic.CustomAbort(http.StatusExpectationFailed, fmt.Sprintf("Failed to delete package to VM ID: %s", vmID))
	}
}
