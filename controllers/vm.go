package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/inspursoft/itpserver/models"
	"github.com/inspursoft/itpserver/services"
)

// Operations about vm
type VMController struct {
	beego.Controller
}

// @Title Get
// @Description Returns a list of virtual machines or filtered by name.
// @Param	vm_name		query 	string	false		"The virual machine name to return."
// @Success 200 {string} 	Successful get all or filter virtual machine by name.
// @Failure 400 Bad request.
// @Failure 401 Unauthorized.
// @Failure 403 The resouce specified was forbidden to access.
// @Failure 404 The resource specified was not found.
// @Failure 500 Internal error occurred at server side.
// @router / [get]
func (v *VMController) Get() {
	vmName := v.GetString("name", "")
	vmHandler := services.NewVMHandler()
	vms := vmHandler.Get(vmName)
	v.Data["JSON"] = vms
	v.ServeJSON()
}

// @Title Post
// @Description Submit to create a virtual machine.
// @Param	vm	body 	string	false		"The virual machine to submit."
// @Success 200 {string} 	Successful submitted virtual machine.
// @Failure 400 Bad request.
// @Failure 401 Unauthorized.
// @Failure 403 The resouce specified was forbidden to access.
// @Failure 404 The resource specified was not found.
// @Failure 500 Internal error occurred at server side.
// @router / [post]
func (v *VMController) Post() {
	var vm models.VM
	err := json.Unmarshal(v.Ctx.Input.RequestBody, &vm)
	if err != nil {
		v.CustomAbort(http.StatusInternalServerError, "Failed to unmarshal request data.")
	}
	status := services.NewVMHandler().Create(&vm)
	if !status {
		v.CustomAbort(http.StatusExpectationFailed, fmt.Sprintf("Failed to create VM: %s", vm.Name))
	}
}

// @Title Delete
// @Description Delete a virtual machine by ID.
// @Param	vm_id	path 	string	true		"The virtual machine ID to be deleted."
// @Success 200 {string} 	Successful deleted virtual machine by ID.
// @Failure 400 Bad request.
// @Failure 401 Unauthorized.
// @Failure 403 The resouce specified was forbidden to access.
// @Failure 404 The resource specified was not found.
// @Failure 500 Internal error occurred at server side.
// @router /:vm_id [get]
func (v *VMController) Delete() {
	vmID := v.GetString("vm_id", "")
	status := services.NewVMHandler().Delete(vmID)
	if !status {
		v.CustomAbort(http.StatusExpectationFailed, fmt.Sprintf("Failed to delete VM by ID: %s", vmID))
	}
}
