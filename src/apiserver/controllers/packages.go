package controllers

import (
	"fmt"
	"net/http"
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
// @Param Authorization	header	string	true	"Set authorization info."
// @Param source_type	query	string	false	"The source type to retrieve."
// @Param vm_name	query	string	false	"VM name."
// @Param	name	query 	string	false		"The software package name to return"
// @Param tag	query string false "The software package tag to return"
// @Success 200 Successful get all or filter software packages by name.
// @Failure 400 Bad request.
// @Failure 401 Unauthorized.
// @Failure 403 The resouce specified was forbidden to access.
// @Failure 404 The resource specified was not found.
// @Failure 500 Internal error occurred at server side.
// @router / [get]
func (pc *PackagesController) Get() {
	sourceType := pc.GetString("source_type", "ansible")
	if sourceType == "vagrantfile" {
		vmName := pc.requiredParam("vm_name")
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
		resp, err = handler.GetAll()
	} else {
		resp, err = handler.Get(pkgName, pkgTag)
	}
	pc.handleError(err)
	pc.serveJSON(resp)
}

// @Title Upload package.
// @Description Upload software package.
// @Param Authorization	header	string	true	"Set authorization info."
// @Param source_type	query	string	false	"Source type for upload"
// @Param	pkg	formData	file	true		"The package to be uploaded."
// @Param	vm_name	query	string	false	"The target VM to upload packages."
// @Success 200 Successful submitted information about software package.
// @Failure 400 Bad request.
// @Failure 401 Unauthorized.
// @Failure 403 The resouce specified was forbidden to access.
// @Failure 404 The resource specified was not found.
// @Failure 500 Internal error occurred at server side.
// @router / [post]
func (pc *PackagesController) Upload() {
	_, fh, err := pc.GetFile("pkg")
	if err != nil {
		pc.handleError(err)
	}
	sourceName := fh.Filename
	vmName := pc.GetString("vm_name", "")
	sourceType := pc.GetString("source_type", "ansible")
	var uploadPath string
	if sourceType == "vagrantfile" {
		vmName = pc.requiredParam("vm_name")
		pathPrefix := beego.AppConfig.String("pathprefix")
		uploadPath = filepath.Join(pathPrefix, beego.AppConfig.String("vagrant::baseworkpath"))
	} else {
		if !utils.CheckFileExt(sourceName, ".zip") {
			pc.CustomAbort(http.StatusBadRequest, "Only allows file with zip extension.")
		}
		uploadPath = beego.AppConfig.String("ansible::uploadpath")
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
	sshClient, err := utils.NewSecureShell(pc.Ctx.ResponseWriter)

	sshHost := beego.AppConfig.String("ssh-host::host")
	sshPort := beego.AppConfig.String("ssh-host::port")
	sshUsername := beego.AppConfig.String("ssh-host::username")
	scpCommand := fmt.Sprintf("mkdir -p %s && scp -P %s %s@%s:%s %s", targetPath, sshPort, sshUsername, sshHost, targetFullPath, targetFullPath)
	beego.Debug(fmt.Sprintf("SCP command is: %s", scpCommand))
	err = sshClient.ExecuteCommand(scpCommand)
	if err != nil {
		beego.Error(fmt.Sprintf("Failed to SCP with err: %+v", err))
	}
	pkg := models.PackageVO{Name: utils.FileNameWithoutExt(sourceName), SourceName: sourceName}
	handler := services.NewPackageHandler()
	handler.Create(pkg)
}

// @Title Delete
// @Description Delete software package by name and tag.
// @Param Authorization	header	string	true	"Set authorization info."
// @Param	package_name	query 	string	true		"The software package name to be deleted."
// @Param	package_tag		query 	string	false		"The software package tag to be deleted."
// @Success 200 Successful submitted information about software package.
// @Failure 400 Bad request.
// @Failure 401 Unauthorized.
// @Failure 403 The resouce specified was forbidden to access.
// @Failure 404 The resource specified was not found.
// @Failure 500 Internal error occurred at server side.
// @router / [delete]
func (pc *PackagesController) Delete() {
	pkgName := pc.requiredParam("package_name")
	pkgTag := pc.GetString("pkg_tag", "")
	err := services.NewPackageHandler().Delete(pkgName, pkgTag)
	pc.handleError(err)
}
