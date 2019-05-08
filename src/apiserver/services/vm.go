package services

import (
	"github.com/inspursoft/itpserver/src/apiserver/dao"
	"github.com/inspursoft/itpserver/src/apiserver/models"
)

type vmConf struct {
	daoHandler dao.VMDaoHandler
	e          *models.ITPError
}

func NewVMHandler() *vmConf {
	return &vmConf{e: &models.ITPError{}}
}

func (vc *vmConf) Get(vmID string) (vm *models.VM, err error) {
	vm, err = vc.daoHandler.GetVM(vmID)
	if err != nil {
		vc.e.InternalError(err)
		return nil, vc.e
	}
	return
}

func (vc *vmConf) GetAll() (vms []*models.VM, err error) {
	vms, err = vc.daoHandler.GetVMList("")
	if err != nil {
		vc.e.InternalError(err)
		return nil, vc.e
	}
	return
}

func (vc *vmConf) Create(vmWithSpec models.VMWithSpec) error {
	newVM := models.VM{VMID: vmWithSpec.VMID, Name: vmWithSpec.Name, OS: vmWithSpec.OS}
	vm, err := vc.daoHandler.GetVM(newVM.VMID)
	if err != nil {
		vc.e.InternalError(err)
		return vc.e
	}
	if vm != nil {
		vc.e.Conflict(vm.VMID, err)
		return vc.e
	}
	spec := models.VMSpec{
		CPUs:    vmWithSpec.Spec.CPUs,
		Storage: vmWithSpec.Spec.Storage,
		Memory:  vmWithSpec.Spec.Memory,
		Extras:  vmWithSpec.Spec.Extras}
	_, err = vc.daoHandler.AddVM(&newVM, &spec)
	if err != nil {
		vc.e.InternalError(err)
		return vc.e
	}
	return nil
}

func (vc *vmConf) Delete(vmID string) error {
	_, err := vc.daoHandler.DeleteVM(vmID)
	if err != nil {
		vc.e.InternalError(err)
		return vc.e
	}
	return nil
}
