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

func (pc *packageConf) Get(name string) (pkgList []models.Package, err error) {
	pkgList, err = pc.daoHandler.GetPackage(name)
	return
}

func (pc *packageConf) Create(pkg models.Package) (status bool, err error) {
	insertedID, err := pc.daoHandler.AddPackage(&pkg)
	status = (insertedID != 0)
	return
}

func (pc *packageConf) Delete(name string, tag string) (status bool, err error) {
	affected, err := pc.daoHandler.DeletePackage(name, tag)
	status = (affected == 1)
	return
}
