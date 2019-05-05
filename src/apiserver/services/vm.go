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

func (vc *vmConf) Get(vmID string) (vm *models.VM, err error) {
	vms, err := vc.daoHandler.GetVM(vmID)
	if err != nil {
		return
	}
	if len(vms) > 0 {
		vm = vms[0]
	}
	return
}

func (vc *vmConf) GetAll() (vms []*models.VM, err error) {
	vms, err = vc.daoHandler.GetVM("")
	return
}

func (vc *vmConf) Create(vmWithSpec models.VMWithSpec) (status bool, err error) {
	vm := models.VM{VMID: vmWithSpec.VMID, Name: vmWithSpec.Name, OS: vmWithSpec.OS}
	spec := models.VMSpec{
		CPUs:    vmWithSpec.Spec.CPUs,
		Storage: vmWithSpec.Spec.Storage,
		Memory:  vmWithSpec.Spec.Memory,
		Extras:  vmWithSpec.Spec.Extras}
	addedVM, err := vc.daoHandler.AddVM(&vm, &spec)
	status = (err == nil && addedVM.ID != 0)
	return
}

func (vc *vmConf) Delete(vmID string) (status bool, err error) {
	affected, err := vc.daoHandler.DeleteVM(vmID)
	status = (err == nil && affected != 0)
	return
}
