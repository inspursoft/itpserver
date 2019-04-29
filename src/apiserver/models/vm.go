package models

type VM struct {
	ID   int64   `orm:"column(id);auto;pk"`
	VMID string  `json:"vm_id" orm:"column(vm_id)"`
	Name string  `json:"vm_name" orm:"column(vm_name)"`
	OS   string  `json:"vm_os" orm:"column(vm_os)"`
	Spec *VMSpec `json:"vm_spec" orm:"reverse(one);on_delete(cascade)"`
}

func (vm *VM) TableName() string {
	return "vm"
}

type VMSpec struct {
	VM      *VM    `orm:"column(vm_id);rel(one)"`
	ID      int64  `orm:"column(id);auto;pk"`
	CPUs    int32  `json:"cpus" orm:"column(cpus)"`
	Memory  string `json:"memory"`
	Storage string `json:"storage"`
	Extras  string `json:"extras"`
}

func (spec *VMSpec) TableName() string {
	return "vm_spec"
}
