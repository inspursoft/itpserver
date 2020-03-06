package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/inspursoft/itpserver/src/apiserver/services"
)

type ArchiveController struct {
	BaseController
}

// @Title Upload archive
// @Description Upload packaged VM box onto remote Nexus service.
// @Param	access_token	query	string	false	"Optional access token."
// @Param Authorization	header	string	false	"Set authorization info."
// @Param	vm_name	query	string	true	"VM name."
// @Param	repo_name	query	string	true	"Repository name."
// @Param	principle	query	string	true	"Principle name."
// @Success 200 Successful uploaded packaged VM box to remote Nexus service.
// @Failure 400 Bad request.
// @Failure 401 Unauthorized.
// @Failure 403 The resouce specified was forbidden to access.
// @Failure 404 The resource specified was not found.
// @Failure 500 Internal error occurred at server side.
// @router /upload [post]
func (ac *ArchiveController) Upload() {
	vmName := ac.GetString("vm_name")
	vm, err := services.NewVMHandler().GetByName(vmName)
	ac.handleError(err)
	if vm == nil {
		ac.CustomAbort(http.StatusNotFound, fmt.Sprintf("VM with name: %s does not exist.", vmName))
	}
	repoName := ac.GetString("repo_name")
	if len(strings.TrimSpace(repoName)) == 0 {
		ac.CustomAbort(http.StatusBadRequest, "Repo name is required.")
	}
	principle := ac.GetString("principle")
	if len(strings.TrimSpace(principle)) == 0 {
		ac.CustomAbort(http.StatusBadGateway, "Principle is required.")
	}

	err = services.UploadArtifacts(vmName, repoName, principle, ac.Ctx.ResponseWriter)
	if err != nil {
		ac.CustomAbort(http.StatusInternalServerError, fmt.Sprintf("Failed to upload artifacts: %+v", err))
	}
}

// @Title Download archive
// @Description Download packaged VM box from ITP service.
// @Param	access_token	query	string	false	"Optional access token."
// @Param Authorization	header	string	false	"Set authorization info."
// @Param	vm_name	query 	string	true		"VM name."
// @Success 200 Successful download packaged VM box from ITP service.
// @Failure 400 Bad request.
// @Failure 401 Unauthorized.
// @Failure 403 The resouce specified was forbidden to access.
// @Failure 404 The resource specified was not found.
// @Failure 500 Internal error occurred at server side.
// @router /download [get]
func (ac *ArchiveController) Download() {
	vmName := ac.GetString("vm_name")
	vm, err := services.NewVMHandler().GetByName(vmName)
	ac.handleError(err)
	if vm == nil {
		ac.CustomAbort(http.StatusNotFound, fmt.Sprintf("VM with name: %s does not exist.", vmName))
	}
	ac.Ctx.Output.Download(services.ResolveBoxFilePath(vmName))
}
