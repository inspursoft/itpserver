package services

import "github.com/inspursoft/itpserver/src/apiserver/models"

type vmConf struct{}

func NewVMHandler() *vmConf {
	return &vmConf{}
}

func (vc *vmConf) Get(name string) []models.VM {
	return nil
}

func (vc *vmConf) Create(vm *models.VM) (status bool) {
	status = false
	return
}

func (vc *vmConf) Delete(ID string) (status bool) {
	status = false
	return
}
