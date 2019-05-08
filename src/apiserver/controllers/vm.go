package controllers

import (
	"net/http"

	"github.com/inspursoft/itpserver/src/apiserver/models"
	"github.com/inspursoft/itpserver/src/apiserver/services"
)

// Operations about vm
type VMController struct {
	BaseController
}

// @Title Get
// @Description Returns a list of virtual machines or filtered by VM ID.
// @Param	vm_id		query 	string	false		"The virual machine name to return."
// @Success 200 {string} 	Successful get all or filter virtual machine by name.
// @Failure 400 Bad request.
// @Failure 401 Unauthorized.
// @Failure 403 The resouce specified was forbidden to access.
// @Failure 404 The resource specified was not found.
// @Failure 500 Internal error occurred at server side.
// @router / [get]
func (v *VMController) Get() {
	vmID := v.GetString("vm_id", "")
	var resp interface{}
	var err error
	if vmID == "" {
		resp, err = services.NewVMHandler().GetAll()
	} else {
		resp, err = services.NewVMHandler().Get(vmID)
	}
	v.handleError(err)
	if resp == nil {
		v.CustomAbort(http.StatusNotFound, "No VM(s) was found.")
	}
	v.Data["json"] = resp
	v.ServeJSON()
}

// @Title Post
// @Description Submit to create a virtual machine.
// @Param	vm_with_spec	body 	models.VMWithSpec	false		"The virual machine to submit."
// @Success 200 {string} 	Successful submitted virtual machine.
// @Failure 400 Bad request.
// @Failure 401 Unauthorized.
// @Failure 403 The resouce specified was forbidden to access.
// @Failure 404 The resource specified was not found.
// @Failure 500 Internal error occurred at server side.
// @router / [post]
func (v *VMController) Post() {
	var vmWithSpec models.VMWithSpec
	v.loadRequestBody(&vmWithSpec)
	err := services.NewVMHandler().Create(vmWithSpec)
	v.handleError(err)
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
// @router /:vm_id [delete]
func (v *VMController) Delete() {
	vmID := v.requiredParam(":vm_id")
	err := services.NewVMHandler().Delete(vmID)
	v.handleError(err)
}
