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

func (vc *vmConf) GetByField(value interface{}, field string) (vm *models.VM, err error) {
	query := models.VM{}
	if str, ok := value.(string); ok {
		query.IP = str
	}
	if val, ok := value.(int64); ok {
		query.ID = val
	}
	vm, err = vc.daoHandler.GetVM(query, field)
	if err != nil {
		vc.e.InternalError(err)
		return nil, vc.e
	}
	if vm == nil {
		vc.e.Notfound("VM", fmt.Errorf("No VM was found %s with value:%+v", field, value))
		return nil, vc.e
	}
	return
}

func (vc *vmConf) GetByIP(IP string) (vm *models.VM, err error) {
	return vc.GetByField(IP, "IP")
}

func (vc *vmConf) GetByID(ID int64) (vm *models.VM, err error) {
	return vc.GetByField(ID, "ID")
}

func (vc *vmConf) Exists(query models.VM) (exists bool, err error) {
	vm, err := vc.daoHandler.GetVM(query, "IP")
	if err != nil {
		vc.e.InternalError(err)
		return
	}
	vm, err = vc.daoHandler.GetVM(query, "Name")
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
		Extras:  vmWithSpec.Spec.Extras,
		VID:     vmWithSpec.Spec.VID,
	}
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
	}
	return vc.e
}

func (vc *vmConf) DeleteByID(ID int64) error {
	query := models.VM{ID: ID}
	_, err := vc.daoHandler.DeleteVM(query, "ID")
	if err != nil {
		vc.e.InternalError(err)
		return vc.e
	}
	return nil
}

func (vc *vmConf) DeleteVMByVID(VID string) error {
	_, err := vc.daoHandler.DeleteVMByVID(VID)
	if err != nil {
		vc.e.InternalError(err)
		return vc.e
	}
	return nil
}
