package controllers

import (
	"github.com/inspursoft/itpserver/src/apiserver/models"
	"github.com/inspursoft/itpserver/src/apiserver/services"
)

type PackagesController struct {
	BaseController
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
	pc.handleError(err)
	pc.serveJSON(resp)
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
	pc.loadRequestBody(&pkg)
	handler := services.NewPackageHandler()
	err := handler.Create(pkg)
	pc.handleError(err)
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
	pkgName := pc.requiredParam("package_name")
	pkgTag := pc.requiredParam("package_tag")
	err := services.NewPackageHandler().Delete(pkgName, pkgTag)
	pc.handleError(err)
}
