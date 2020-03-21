package controllers

import (
	"path/filepath"

	"github.com/astaxie/beego"
	"github.com/inspursoft/itpserver/src/apiserver/models"
	"github.com/inspursoft/itpserver/src/apiserver/utils"

	"github.com/inspursoft/itpserver/src/apiserver/services"
)

type PackagesController struct {
	BaseController
}

// @Title Get
// @Description Return a list of software packages.
// @Param	access_token	query	string	false	"Optional access token."
// @Param Authorization	header	string	false	"Set authorization info."
// @Param source_type	query	string	false	"The source type to retrieve."
// @Param vm_name	path	string	true	"VM name for which uploading package."
// @Param	name	query 	string	false		"The software package name to return"
// @Param tag	query string false "The software package tag to return"
// @Success 200 Successful get all or filter software packages by name.
// @Failure 400 Bad request.
// @Failure 401 Unauthorized.
// @Failure 403 The resouce specified was forbidden to access.
// @Failure 404 The resource specified was not found.
// @Failure 500 Internal error occurred at server side.
// @router /:vm_name [get]
func (pc *PackagesController) Get() {
	vmName := pc.requiredParam(":vm_name")
	sourceType := pc.GetString("source_type", "ansible")
	if sourceType == "vagrantfile" {
		fileList, err := services.RetrieveVMFiles(vmName)
		pc.handleError(err)
		pc.serveJSON(fileList)
		return
	}
	pkgName := pc.GetString("name", "")
	pkgTag := pc.GetString("tag", "")
	handler := services.NewPackageHandler()
	var resp interface{}
	var err error
	if pkgName == "" {
		resp, err = handler.GetAllByVMName(vmName)
	} else {
		resp, err = handler.Get(vmName, pkgName, pkgTag)
	}
	pc.handleError(err)
	pc.serveJSON(resp)
}

// @Title Upload package.
// @Description Upload software package.
// @Param	access_token	query	string	false	"Optional access token."
// @Param Authorization	header	string	false	"Set authorization info."
// @Param source_type	query	string	false	"Source type for upload"
// @Param	pkg	formData	file	true		"The package to be uploaded."
// @Param	vm_name	path	string	true	"The target VM to upload packages."
// @Success 200 Successful submitted information about software package.
// @Failure 400 Bad request.
// @Failure 401 Unauthorized.
// @Failure 403 The resouce specified was forbidden to access.
// @Failure 404 The resource specified was not found.
// @Failure 500 Internal error occurred at server side.
// @router /:vm_name [post]
func (pc *PackagesController) Upload() {
	vmName := pc.requiredParam(":vm_name")
	_, fh, err := pc.GetFile("pkg")
	if err != nil {
		pc.handleError(err)
	}
	sourceName := fh.Filename
	sourceType := pc.GetString("source_type", "ansible")
	pathPrefix := beego.AppConfig.String("pathprefix")
	var uploadPath string
	if sourceType == "vagrantfile" {
		uploadPath = filepath.Join(pathPrefix, beego.AppConfig.String("vagrant::baseworkpath"))
	} else {
		uploadPath = filepath.Join(pathPrefix, beego.AppConfig.String("ansible::uploadpath"))
	}
	targetPath, err := utils.CheckDirs(filepath.Join(uploadPath, vmName))
	if err != nil {
		pc.handleError(err)
	}
	targetFullPath := filepath.Join(targetPath, sourceName)
	err = pc.SaveToFile("pkg", targetFullPath)
	if err != nil {
		pc.handleError(err)
	}
	hostMode, _ := beego.AppConfig.Bool("hostmode")
	if !hostMode {
		beego.Debug("Running under Cross Host mode ...")
		sshClient, err := utils.NewSecureShell(pc.Ctx.ResponseWriter)
		if err != nil {
			pc.handleError(err)
		}
		err = sshClient.CheckDir(targetPath)
		if err != nil {
			pc.handleError(err)
		}
		err = sshClient.HostSCP(targetFullPath, targetFullPath, false)
		if err != nil {
			pc.handleError(err)
		}
	}
	pkg := models.PackageVO{VMName: vmName, Name: utils.FileNameWithoutExt(sourceName), SourceName: sourceName}
	handler := services.NewPackageHandler()
	handler.Create(pkg)
}

// @Title Delete
// @Description Delete software package by name and tag.
// @Param	access_token	query	string	false	"Optional access token."
// @Param Authorization	header	string	false	"Set authorization info."
// @Param	package_name	query 	string	true		"The software package name to be deleted."
// @Param	package_tag		query 	string	false		"The software package tag to be deleted."
// @Param	vm_name	path	string	true	"The target VM to upload packages."
// @Success 200 Successful submitted information about software package.
// @Failure 400 Bad request.
// @Failure 401 Unauthorized.
// @Failure 403 The resouce specified was forbidden to access.
// @Failure 404 The resource specified was not found.
// @Failure 500 Internal error occurred at server side.
// @router /:vm_name [delete]
func (pc *PackagesController) Delete() {
	vmName := pc.requiredParam(":vm_name")
	pkgName := pc.requiredParam("package_name")
	pkgTag := pc.GetString("pkg_tag", "")
	err := services.NewPackageHandler().Delete(vmName, pkgName, pkgTag)
	pc.handleError(err)
}
