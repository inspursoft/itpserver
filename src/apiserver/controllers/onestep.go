package controllers

import (
	"net/http"

	"github.com/inspursoft/itpserver/src/apiserver/models"
)

type OneStepController struct {
	BaseController
}

// @Title Post
// @Description One step to create VM and install software onto it.
// @Param Authorization	header	string	true	"Set authorization info."
// @Param	one_step	body 	models.OneStepInstallation	true		"The virual machine to submit."
// @Success 200 Successful installed software package onto a virtual machine.
// @Failure 400 Bad request.
// @Failure 401 Unauthorized.
// @Failure 403 The resouce specified was forbidden to access.
// @Failure 404 The resource specified was not found.
// @Failure 500 Internal error occurred at server side.
// @router / [post]
func (ic *OneStepController) Post() {
	var oneStep models.OneStepInstallation
	ic.loadRequestBody(&oneStep)
	ic.proxiedRequest(http.MethodPost, oneStep.VMWithSpec, "VMController.CreateBySpec")
	ic.proxiedRequest(http.MethodPost, oneStep.PackageVO, "InstallationController.Post", ":vm_name", oneStep.VMWithSpec.Name)
	ic.proxiedRequest(http.MethodPost, nil, "VMController.Package", ":vm_name", oneStep.VMWithSpec.Name)
}

// @Title Post
// @Description One step to create VM and install software onto it.
// @Param Authorization	header	string	true	"Set authorization info."
// @Param	vm_name	path	string	true	"VM name."
// @Success 200 Successful installed software package onto a virtual machine.
// @Failure 400 Bad request.
// @Failure 401 Unauthorized.
// @Failure 403 The resouce specified was forbidden to access.
// @Failure 404 The resource specified was not found.
// @Failure 500 Internal error occurred at server side.
// @router /:vm_name [post]
func (ic *OneStepController) PostWithVagrantfile() {
	vmName := ic.requiredParam(":vm_name")
	ic.proxiedRequest(http.MethodPost, nil, "VMController.CreateByVagrantfile", ":vm_name", vmName)
	ic.proxiedRequest(http.MethodPost, nil, "VMController.Package", ":vm_name", vmName)
}
