package services

import (
	"github.com/inspursoft/itpserver/src/apiserver/dao"
	"github.com/inspursoft/itpserver/src/apiserver/models"
)

type packageConf struct {
	daoHandler dao.PkgDaoHandler
}

func NewPackageHandler() *packageConf {
	return &packageConf{}
}

func (pc *packageConf) Get(name string, tag string) (pkg *models.Package, err error) {
	pkgList, err := pc.daoHandler.GetPackage(name, tag)
	if err != nil {
		return
	}
	if len(pkgList) > 0 {
		pkg = pkgList[0]
	}
	return
}

func (pc *packageConf) GetAll() (pkgList []*models.Package, err error) {
	pkgList, err = pc.daoHandler.GetPackage("", "")
	return
}

func (pc *packageConf) Create(pkg models.Package) (status bool, err error) {
	insertedID, err := pc.daoHandler.AddPackage(&pkg)
	status = (err == nil && insertedID != 0)
	return
}

func (pc *packageConf) Delete(name string, tag string) (status bool, err error) {
	affected, err := pc.daoHandler.DeletePackage(name, tag)
	status = (err == nil && affected == 1)
	return
}
