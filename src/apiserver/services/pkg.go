package services

import (
	"fmt"

	"github.com/inspursoft/itpserver/src/apiserver/dao"
	"github.com/inspursoft/itpserver/src/apiserver/models"
)

type packageConf struct {
	daoHandler dao.PkgDaoHandler
	e          *models.ITPError
}

func NewPackageHandler() *packageConf {
	return &packageConf{e: &models.ITPError{}}
}

func (pc *packageConf) Get(vmName string, name string, tag string) (pkg *models.Package, err error) {
	pkg, err = pc.daoHandler.GetPackage(vmName, name, tag)
	if err != nil {
		pc.e.InternalError(err)
		return nil, pc.e
	}
	return
}

func (pc *packageConf) Exists(vmName string, name string, tag string) (exists bool, err error) {
	pkg, err := pc.Get(vmName, name, tag)
	return (pkg != nil), nil
}

func (pc *packageConf) GetAllByVMName(vmName string) (pkgList []*models.Package, err error) {
	pkgList, err = pc.daoHandler.GetPackageList(vmName, "", "")
	if err != nil {
		pc.e.InternalError(err)
		return nil, pc.e
	}
	return
}

func (pc *packageConf) Create(pkg models.PackageVO) error {
	query, err := pc.daoHandler.GetPackage(pkg.Name, pkg.Tag, pkg.VMName)
	if err != nil {
		pc.e.InternalError(err)
		return pc.e
	}
	if query != nil {
		pc.e.Conflict("Package", fmt.Errorf("name: %s with tag: %s already exists in VM: %s", pkg.Name, pkg.Tag, pkg.VMName))
		return pc.e
	}
	_, err = pc.daoHandler.AddPackage(&models.Package{VMName: pkg.VMName, Name: pkg.Name, Tag: pkg.Tag})
	if err != nil {
		pc.e.InternalError(err)
		return pc.e
	}
	return nil
}

func (pc *packageConf) Delete(vmName string, name string, tag string) error {
	_, err := pc.daoHandler.DeletePackage(vmName, name, tag)
	if err != nil {
		pc.e.InternalError(err)
		return pc.e
	}
	return nil
}
