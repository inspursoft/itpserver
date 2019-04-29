package dao

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/inspursoft/itpserver/src/apiserver/models"
)

type vmDaoHandler int

func NewVMDaoHandler() *vmDaoHandler {
	return new(vmDaoHandler)
}

func (v *vmDaoHandler) AddVM(vm *models.VM, spec *models.VMSpec) (*models.VM, error) {
	o := orm.NewOrm()
	id, err := o.Insert(vm)
	if err != nil {
		beego.Error(err)
		return nil, err
	}
	beego.Info(fmt.Sprintf("Successful added VM with ID: %d", id))
	spec.VM = vm
	id, err = o.Insert(spec)
	if err != nil {
		beego.Error(err)
		return nil, err
	}
	beego.Info(fmt.Sprintf("Successful added VM Spec with ID: %d", id))
	return vm, nil
}

func (v *vmDaoHandler) GetVM(vmID ...string) ([]models.VM, error) {
	o := orm.NewOrm()
	q := o.QueryTable("vm")
	if vmID != nil {
		q.Filter("vm_id", vmID)
	}
	var results []models.VM
	q.All(&results)
	beego.Info(fmt.Sprintf("Succesful got VM(s) with VM ID: %+v", results))
	return results, nil
}

func (v *vmDaoHandler) GetVMSpec(vmID string) (models.VMSpec, error) {
	var vmSpec models.VMSpec
	o := orm.NewOrm()
	err := o.QueryTable("vm_spec").RelatedSel().Filter("vm__vm_id", vmID).
		One(&vmSpec)
	beego.Info(fmt.Sprintf("Successful got VM Spec %+v with VM ID: %s", vmSpec, vmID))
	return vmSpec, err
}

func (v *vmDaoHandler) UpdateVM(vm models.VM) (affected int64, err error) {
	o := orm.NewOrm()
	affected, err = o.QueryTable("vm").Filter("vm_id", vm.VMID).
		Update(
			orm.Params{
				"vm_name": vm.Name,
				"vm_os":   vm.OS,
			})
	beego.Info(fmt.Sprintf("Successful update VM %d item(s) with VM ID: %s", affected, vm.VMID))
	return
}

func (v *vmDaoHandler) UpdateVMSpec(vm models.VM, spec models.VMSpec) (affected int64, err error) {
	o := orm.NewOrm()
	affected, err = o.QueryTable("vm_spec").RelatedSel().Filter("vm__vm_id", vm.VMID).
		Update(
			orm.Params{
				"cpus":    spec.CPUs,
				"memory":  spec.Memory,
				"storage": spec.Storage,
				"extras":  spec.Extras,
			})
	beego.Info(fmt.Sprintf("Successful update VM Spec %d item(s) with VM ID: %s", affected, vm.VMID))
	return
}

func (v *vmDaoHandler) DeleteVM(vmID string) (affected int64, err error) {
	o := orm.NewOrm()
	affected, err = o.QueryTable("vm").Filter("vm_id", vmID).Delete()
	beego.Info(fmt.Sprintf("Successful deleted %d item(s) with VM ID: %s", affected, vmID))
	return
}
