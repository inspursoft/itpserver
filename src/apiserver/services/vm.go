package services

import (
	"github.com/inspursoft/itpserver/src/apiserver/dao"
	"github.com/inspursoft/itpserver/src/apiserver/models"
)

type vmConf struct {
	daoHandler dao.VMDaoHandler
}

func NewVMHandler() *vmConf {
	return &vmConf{}
}

func (vc *vmConf) Get(vmID string) (vms []models.VM, err error) {
	return vc.daoHandler.GetVM(vmID)
}

func (vc *vmConf) Create(vm models.VM, spec models.VMSpec) (status bool, err error) {
	addedVM, err := vc.daoHandler.AddVM(&vm, &spec)
	status = (addedVM == nil)
	return
}

func (vc *vmConf) Delete(vmID string) (status bool, err error) {
	affected, err := vc.daoHandler.DeleteVM(vmID)
	status = (affected == 1)
	return
}
