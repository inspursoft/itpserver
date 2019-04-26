package services

import "github.com/inspursoft/itpserver/models"

type packageConf struct{}

func NewPackageHandler() *packageConf {
	return &packageConf{}
}

func (pc *packageConf) Get(name string) []models.Package {
	return nil
}

func (pc *packageConf) Create(p *models.Package) (status bool) {
	status = false
	return
}

func (pc *packageConf) Delete(name string, tag string) (status bool) {
	status = false
	return
}
