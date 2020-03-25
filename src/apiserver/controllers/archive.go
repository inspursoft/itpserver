package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/inspursoft/itpserver/src/apiserver/models"
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
	err = services.SCPArtifacts(vmName, ac.Ctx.ResponseWriter)
	if err != nil {
		ac.CustomAbort(http.StatusInternalServerError, fmt.Sprintf("Failed to SCP under Cross Host mode: %+v", err))
	}
	err = services.UploadArtifacts(vmName, repoName, principle, ac.Ctx.ResponseWriter)
	if err != nil {
		ac.CustomAbort(http.StatusInternalServerError, fmt.Sprintf("Failed to upload artifacts: %+v", err))
	}
}

func resetVMPackageStatus() {

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
	vmHandler := services.NewVMHandler()
	vm, err := vmHandler.GetByName(vmName)
	ac.handleError(err)
	if vm == nil {
		ac.CustomAbort(http.StatusNotFound, fmt.Sprintf("VM with name: %s does not exist.", vmName))
	}
	switch vm.PackageStatus {
	case models.Pending:
		ac.serveStatus(http.StatusOK, fmt.Sprintf("VM: %s is packing in progress, please wait.", vmName))
	case models.Initial:
		ac.serveStatus(http.StatusOK, fmt.Sprintf("Start packing VM: %s as box.", vmName))
		go func() {
			vmHandler.UpdateVMPackageStatus(vmName, models.Pending)
			err := ac.proxiedRequest(http.MethodPost, nil, "VMController.Package", ":vm_name", vmName, "access_token", ac.GetString("access_token", ""))
			if err != nil {
				vmHandler.UpdateVMPackageStatus(vmName, models.Initial)
				// ac.handleError(err)
			}
			err = services.SCPArtifacts(vmName, ac.Ctx.ResponseWriter)
			if err != nil {
				vmHandler.UpdateVMPackageStatus(vmName, models.Initial)
				// ac.handleError(err)
			}
			vmHandler.UpdateVMPackageStatus(vmName, models.Finished)
		}()
	case models.Finished:
		ac.Ctx.Output.Download(services.ResolveBoxFilePath(vmName))
		vmHandler.UpdateVMPackageStatus(vmName, models.Initial)
		ac.serveStatus(http.StatusOK, fmt.Sprintf("VM: %s has finished to package and it is ready to download.", vmName))
	}
}

// @Title Reset VM download packaging status manually
// @Description Reset VM download packaging statue manually from ITP service.
// @Param	access_token	query	string	false	"Optional access token."
// @Param Authorization	header	string	false	"Set authorization info."
// @Param	vm_name	query 	string	true		"VM name."
// @Success 200 Successful Reset VM download packaged VM box status from ITP service.
// @Failure 400 Bad request.
// @Failure 401 Unauthorized.
// @Failure 403 The resouce specified was forbidden to access.
// @Failure 404 The resource specified was not found.
// @Failure 500 Internal error occurred at server side.
// @router /download/reset [get]
func (ac *ArchiveController) Reset() {
	vmName := ac.GetString("vm_name")
	vmHandler := services.NewVMHandler()
	vm, err := vmHandler.GetByName(vmName)
	ac.handleError(err)
	if vm == nil {
		ac.CustomAbort(http.StatusNotFound, fmt.Sprintf("VM with name: %s does not exist.", vmName))
	}
	vmHandler.UpdateVMPackageStatus(vmName, models.Initial)
	ac.serveStatus(http.StatusOK, fmt.Sprintf("VM: %s has sucessfully reset to initial status.", vmName))
}
