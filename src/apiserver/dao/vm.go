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

func (v *VMDaoHandler) GetVM(vm models.VM, col ...string) (*models.VM, error) {
	o := orm.NewOrm()
	err := o.Read(&vm, col...)
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	_, err = o.LoadRelated(&vm, "Spec")
	if err != nil {
		return nil, err
	}
	return &vm, nil
}

func (v *VMDaoHandler) GetVMList(vmWithSpec models.VMWithSpec) ([]*models.VM, error) {
	o := orm.NewOrm()
	q := o.QueryTable("vm")
	if vmWithSpec.IP != "" {
		q = q.Filter("IP", vmWithSpec.IP)
	} else if vmWithSpec.Name != "" {
		q = q.Filter("Name", vmWithSpec.Name)
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

func (v *VMDaoHandler) UpdateVM(ID int64, updates map[string]interface{}) (affected int64, err error) {
	o := orm.NewOrm()
	affected, err = o.QueryTable("vm").Filter("ID", ID).Update(orm.Params(updates))
	if err != nil {
		if err == orm.ErrNoRows {
			return 0, nil
		}
		return
	}
	beego.Info(fmt.Sprintf("Successful update VM ID: %d item(s) with updates: %+v", ID, updates))
	return
}

func (v *VMDaoHandler) UpdateVMSpec(ID int64, updates map[string]interface{}) (affected int64, err error) {
	o := orm.NewOrm()
	affected, err = o.QueryTable("vm_spec").RelatedSel().Filter("vm__id", ID).
		Update(orm.Params(updates))
	if err != nil {
		if err == orm.ErrNoRows {
			return 0, nil
		}
		return
	}
	beego.Info(fmt.Sprintf("Successful update VM ID %d item(s) with updates: %+v", ID, updates))
	return
}

func (v *VMDaoHandler) DeleteVM(query models.VM, cols ...string) (affected int64, err error) {
	o := orm.NewOrm()
	affected, err = o.Delete(&query, cols...)
	if err != nil {
		if err == orm.ErrNoRows {
			return 0, nil
		}
		return
	}
	beego.Info(fmt.Sprintf("Successful deleted %d item(s) with VM ID: %d", affected, query.ID))
	return
}
