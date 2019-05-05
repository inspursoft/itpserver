package dao

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/inspursoft/itpserver/src/apiserver/models"
)

type VMDaoHandler int

func (v *VMDaoHandler) AddVM(vm *models.VM, spec *models.VMSpec) (*models.VM, error) {
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

func (v *VMDaoHandler) GetVM(vmID string) ([]*models.VM, error) {
	o := orm.NewOrm()
	q := o.QueryTable("vm")
	if vmID != "" {
		q = q.Filter("vm_id", vmID)
	}
	var results []*models.VM
	count, err := q.All(&results)
	if err != nil {
		return nil, err
	}
	for _, v := range results {
		_, err = o.LoadRelated(v, "Spec")
	}
	if err != nil {
		return nil, err
	}
	beego.Info(fmt.Sprintf("Succesful got %d VM(s) with spec", count))
	return results, nil
}

func (v *VMDaoHandler) UpdateVM(vm models.VM) (affected int64, err error) {
	o := orm.NewOrm()
	affected, err = o.QueryTable("vm").Filter("vm_id", vm.VMID).
		Update(
			orm.Params{
				"vm_name": vm.Name,
				"vm_os":   vm.OS,
			})
	if err != nil {
		if err == orm.ErrNoRows {
			return 0, nil
		}
		affected = 0
		return
	}
	beego.Info(fmt.Sprintf("Successful update VM %d item(s) with VM ID: %s", affected, vm.VMID))
	return
}

func (v *VMDaoHandler) UpdateVMSpec(vm models.VM, spec models.VMSpec) (affected int64, err error) {
	o := orm.NewOrm()
	affected, err = o.QueryTable("vm_spec").RelatedSel().Filter("vm__vm_id", vm.VMID).
		Update(
			orm.Params{
				"cpus":    spec.CPUs,
				"memory":  spec.Memory,
				"storage": spec.Storage,
				"extras":  spec.Extras,
			})
	if err != nil {
		if err == orm.ErrNoRows {
			return 0, nil
		}
		affected = 0
		return
	}
	beego.Info(fmt.Sprintf("Successful update VM Spec %d item(s) with VM ID: %s", affected, vm.VMID))
	return
}

func (v *VMDaoHandler) DeleteVM(vmID string) (affected int64, err error) {
	o := orm.NewOrm()
	affected, err = o.QueryTable("vm").Filter("vm_id__exact", vmID).Delete()
	if err != nil {
		if err == orm.ErrNoRows {
			return 0, nil
		}
		affected = 0
		return
	}
	beego.Info(fmt.Sprintf("Successful deleted %d item(s) with VM ID: %s", affected, vmID))
	return
}
