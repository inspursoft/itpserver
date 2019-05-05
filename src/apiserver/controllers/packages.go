package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/inspursoft/itpserver/src/apiserver/models"
	"github.com/inspursoft/itpserver/src/apiserver/services"
)

type PackagesController struct {
	beego.Controller
}

// @Title Get
// @Description Return a list of software packages.
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
	if err != nil {
		pc.CustomAbort(http.StatusInternalServerError, "Failed to get package list.")
	}
	pc.Data["json"] = resp
	pc.ServeJSON()
}

// @Title Post
// @Description Submit information about a software package.
// @Param	packages	body 	models.PackageVO	false		"The software package name to submit"
// @Success 200 {string} 		Successful submitted information about software package.
// @Failure 400 Bad request.
// @Failure 401 Unauthorized.
// @Failure 403 The resouce specified was forbidden to access.
// @Failure 404 The resource specified was not found.
// @Failure 500 Internal error occurred at server side.
// @router / [post]
func (pc *PackagesController) Post() {
	var pkg models.Package
	err := json.Unmarshal(pc.Ctx.Input.RequestBody, &pkg)
	if err != nil {
		pc.CustomAbort(http.StatusInternalServerError, "Failed to unmarshal request data.")
	}
	handler := services.NewPackageHandler()
	status, err := handler.Create(pkg)
	if err != nil {
		pc.CustomAbort(http.StatusInternalServerError, "Failed to create package list.")
	}
	if !status {
		pc.CustomAbort(http.StatusExpectationFailed, fmt.Sprintf("Failed to create package: %s", pkg.Name))
	}
}

// @Title Delete
// @Description Delete software package by name and tag.
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
	pkgName := pc.GetString("package_name")
	pkgTag := pc.GetString("package_tag")
	status, err := services.NewPackageHandler().Delete(pkgName, pkgTag)
	if err != nil {
		pc.CustomAbort(http.StatusInternalServerError, "Failed to delete package.")
	}
	if !status {
		pc.CustomAbort(http.StatusExpectationFailed, fmt.Sprintf("Failed to delete package: %s, tag: %s", pkgName, pkgTag))
	}
}
