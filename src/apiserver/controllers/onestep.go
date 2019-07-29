package controllers

import (
	"net/http"

	"github.com/inspursoft/itpserver/src/apiserver/models"
	"github.com/inspursoft/itpserver/src/apiserver/services/vagrantcli"
)

type OneStepController struct {
	BaseController
}

// @Title Post
// @Description One step to create VM and install software onto it.
// @Param Authorization	header	string	true	"Set authorization info."
// @Param	one_step	body 	models.OneStepInstallation	true		"The virual machine to submit."
// @Success 200 {string} 	Successful installed software package onto a virtual machine.
// @Failure 400 Bad request.
// @Failure 401 Unauthorized.
// @Failure 403 The resouce specified was forbidden to access.
// @Failure 404 The resource specified was not found.
// @Failure 500 Internal error occurred at server side.
// @router / [post]
func (ic *OneStepController) Post() {
	var oneStep models.OneStepInstallation
	ic.loadRequestBody(&oneStep)
	ic.proxiedRequest(http.MethodPost, oneStep.VMWithSpec, "VMController.Post")
	ic.proxiedRequest(http.MethodPost, oneStep.PackageVO, "PackagesController.Post")
	err := vagrantcli.NewClient(*oneStep.VMWithSpec, ic.Ctx.ResponseWriter).Package()
	ic.handleError(err)
}
