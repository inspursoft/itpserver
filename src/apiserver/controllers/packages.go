package controllers

import (
	"fmt"
	"path/filepath"

	"net/http"

	"github.com/inspursoft/itpserver/src/apiserver/services/ansiblecli"
	"github.com/inspursoft/itpserver/src/apiserver/utils"

	"github.com/inspursoft/itpserver/src/apiserver/models"
	"github.com/inspursoft/itpserver/src/apiserver/services"
)

type PackagesController struct {
	BaseController
}

// @Title Get
// @Description Return a list of software packages.
// @Param Authorization	header	string	true	"Set authorization info."
// @Param	name	query 	string	false		"The software package name to return"
// @Param tag	query string false "The software package tag to return"
// @Success 200 {string} 		Successful get all or filter software packages by name.
// @Failure 400 Bad request.
// @Failure 401 Unauthorized.
// @Failure 403 The resouce specified was forbidden to access.
// @Failure 404 The resource specified was not found.
// @Failure 500 Internal error occurred at server side.
// @router / [get]
func (pc *PackagesController) Get() {
	pkgName := pc.GetString("name", "")
	pkgTag := pc.GetString("tag", "")

	handler := services.NewPackageHandler()
	var resp interface{}
	var err error
	if pkgName == "" || pkgTag == "" {
		resp, err = handler.GetAll()
	} else {
		resp, err = handler.Get(pkgName, pkgTag)
	}
	pc.handleError(err)
	pc.serveJSON(resp)
}

// @Title Post
// @Description Upload software package.
// @Param Authorization	header	string	true	"Set authorization info."
// @Param vm_name	formData	string	true	"The VM name to install package."
// @Param	pkg	formData	file	true		"The package to be uploaded."
// @Success 200 {string} 		Successful submitted information about software package.
// @Failure 400 Bad request.
// @Failure 401 Unauthorized.
// @Failure 403 The resouce specified was forbidden to access.
// @Failure 404 The resource specified was not found.
// @Failure 500 Internal error occurred at server side.
// @router / [post]
func (pc *PackagesController) Upload() {
	vmName := pc.requiredParam("vm_name")
	exists, err := services.NewVMHandler().Exists(models.VM{Name: vmName})
	if err != nil {
		pc.handleError(err)
	}
	if !exists {
		pc.CustomAbort(http.StatusNotFound, fmt.Sprintf("VM: %s not found.", vmName))
	}
	_, fh, err := pc.GetFile("pkg")
	if err != nil {
		pc.handleError(err)
	}
	sourceName := fh.Filename
	if !utils.CheckFileExt(sourceName, ".zip") {
		pc.CustomAbort(http.StatusBadRequest, "Only allow file with zip extsion.")
	}
	targetPath, err := utils.CheckDirs("upload", vmName)
	if err != nil {
		pc.handleError(err)
	}
	err = pc.SaveToFile("pkg", filepath.Join(targetPath, sourceName))
	if err != nil {
		pc.handleError(err)
	}
	vmWithSpec := models.VMWithSpec{Name: vmName}
	pkg := models.PackageVO{Name: utils.FileNameWithoutExt(sourceName), SourceName: sourceName}
	err = ansiblecli.NewClient(vmWithSpec, pkg, pc.Ctx.ResponseWriter).Transfer()
	if err != nil {
		pc.handleError(err)
	}
}

// @Title Delete
// @Description Delete software package by name and tag.
// @Param Authorization	header	string	true	"Set authorization info."
// @Param	package_name	query 	string	true		"The software package name to be deleted."
// @Param	package_tag		query 	string	true		"The software package tag to be deleted."
// @Success 200 {string} 		Successful submitted information about software package.
// @Failure 400 Bad request.
// @Failure 401 Unauthorized.
// @Failure 403 The resouce specified was forbidden to access.
// @Failure 404 The resource specified was not found.
// @Failure 500 Internal error occurred at server side.
// @router / [delete]
func (pc *PackagesController) Delete() {
	pkgName := pc.requiredParam("package_name")
	pkgTag := pc.requiredParam("package_tag")
	err := services.NewPackageHandler().Delete(pkgName, pkgTag)
	pc.handleError(err)
}
