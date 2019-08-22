package controllers

import (
	"net/http"
	"strings"

	"github.com/inspursoft/itpserver/src/apiserver/models"
	"github.com/inspursoft/itpserver/src/apiserver/services"
	"github.com/inspursoft/itpserver/src/apiserver/services/vagrantcli"
)

// Operations about vm
type VMController struct {
	BaseController
}

// @Title Get
// @Description Returns a list of virtual machines or filtered by VM ID.
// @Param Authorization	header	string	true	"Set authorization info."
// @Param	id		query 	int64	false		"The virual machine name to return."
// @Success 200 {string} 	Successful get all or filter virtual machine by name.
// @Failure 400 Bad request.
// @Failure 401 Unauthorized.
// @Failure 403 The resouce specified was forbidden to access.
// @Failure 404 The resource specified was not found.
// @Failure 500 Internal error occurred at server side.
// @router / [get]
func (v *VMController) Get() {
	ID := v.requiredID("id")
	var resp interface{}
	var err error
	handler := services.NewVMHandler()
	if ID == 0 {
		resp, err = handler.GetVMList()
	} else {
		resp, err = handler.GetByID(ID)
	}
	v.handleError(err)
	v.serveJSON(resp)
}

// @Title Post
// @Description Submit to create a virtual machine.
// @Param Authorization	header	string	true	"Set authorization info."
// @Param	build_type	query	string	false	"Specify the build type."
// @Param vm_name	query	string	false	"Specify the VM name."
// @Param	vm_with_spec	body 	models.VMWithSpec	false		"The virual machine to submit."
// @Success 200 {string} 	Successful submitted virtual machine.
// @Failure 400 Bad request.
// @Failure 401 Unauthorized.
// @Failure 403 The resouce specified was forbidden to access.
// @Failure 404 The resource specified was not found.
// @Failure 500 Internal error occurred at server side.
// @router / [post]
func (v *VMController) Post() {
	buildType := v.GetString("build_type", "spec")
	vmName := v.GetString("vm_name", "")
	var err error
	if buildType == "vagrantfile" {
		if len(strings.TrimSpace(vmName)) == 0 {
			v.CustomAbort(http.StatusBadRequest, "VM name is required.")
		}
		err = vagrantcli.NewEaseClient(vmName, v.Ctx.ResponseWriter).CreateByVagrantfile()
		v.handleError(err)
		return
	}
	var vmWithSpec models.VMWithSpec
	v.loadRequestBody(&vmWithSpec)
	err = vagrantcli.NewClient(vmWithSpec, v.Ctx.ResponseWriter).Create()
	v.handleError(err)
}

// @Title Delete
// @Description Delete a virtual machine by ID.
// @Param Authorization	header	string	true	"Set authorization info."
// @Param	vm_name	query 	string	true		"The virtual machine name to be deleted."
// @Success 200 {string} 	Successful deleted virtual machine by name.
// @Failure 400 Bad request.
// @Failure 401 Unauthorized.
// @Failure 403 The resouce specified was forbidden to access.
// @Failure 404 The resource specified was not found.
// @Failure 500 Internal error occurred at server side.
// @router /:vm_name [delete]
func (v *VMController) Delete() {
	vmName := v.requiredParam("vm_name")
	vmWithSpec := models.VMWithSpec{Name: vmName}
	err := vagrantcli.NewClient(vmWithSpec, v.Ctx.ResponseWriter).Destroy()
	v.handleError(err)
}
