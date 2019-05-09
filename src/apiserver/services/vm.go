package services

import (
	"fmt"

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

func (vc *vmConf) GetByID(ID int64) (vm *models.VM, err error) {
	query := models.VM{ID: ID}
	vm, err = vc.daoHandler.GetVM(query, "ID")
	if err != nil {
		vc.e.InternalError(err)
		return nil, vc.e
	}
	if vm == nil {
		vc.e.Notfound("VM", fmt.Errorf("No VM was found with ID: %d", ID))
		return nil, vc.e
	}
	return
}

func (vc *vmConf) Exists(query models.VM) (exists bool, err error) {
	vm, err := vc.daoHandler.GetVM(query, "IP")
	if err != nil {
		vc.e.InternalError(err)
		return
	}
	exists = (vm != nil && vm.ID != 0)
	return
}

func (vc *vmConf) GetVMList(query ...models.VMWithSpec) (vms []*models.VM, err error) {
	var cond models.VMWithSpec
	if len(query) > 0 {
		cond = query[0]
	}
	vms, err = vc.daoHandler.GetVMList(cond)
	if err != nil {
		vc.e.InternalError(err)
		return nil, vc.e
	}
	return
}

func (vc *vmConf) Create(vmWithSpec models.VMWithSpec) error {
	query := models.VM{IP: vmWithSpec.IP}
	exists, err := vc.Exists(query)
	if err != nil {
		vc.e.InternalError(err)
		return vc.e
	}
	if exists {
		vc.e.Conflict(vmWithSpec.IP, err)
		return vc.e
	}
	newVM := models.VM{IP: vmWithSpec.IP, Name: vmWithSpec.Name, OS: vmWithSpec.OS}
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

func (vc *vmConf) UpdateVMID(ID int64, VID string) error {
	updates := map[string]interface{}{"VID": VID}
	_, err := vc.daoHandler.UpdateVMSpec(ID, updates)
	if err != nil {
		vc.e.InternalError(err)
		return vc.e
	}
	return nil
}

func (vc *vmConf) DeleteByID(ID int64) error {
	query := models.VM{ID: ID}
	affected, err := vc.daoHandler.DeleteVM(query, "ID")
	if err != nil {
		vc.e.InternalError(err)
		return vc.e
	}
	if affected == 0 {
		vc.e.Notfound(fmt.Sprintf("VM with ID: %d", ID), err)
		return vc.e
	}
	return nil
}
