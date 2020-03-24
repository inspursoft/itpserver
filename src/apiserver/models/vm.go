package models

type VMPackageStatus int

const (
	Initial VMPackageStatus = iota
	Pending
	Finished
)

type VM struct {
	ID            int64           `json:"id" orm:"column(id);auto;pk"`
	IP            string          `json:"vm_ip" orm:"column(vm_ip)"`
	Name          string          `json:"vm_name" orm:"column(vm_name)"`
	OS            string          `json:"vm_os" orm:"column(vm_os)"`
	Spec          *VMSpec         `json:"vm_spec" orm:"reverse(one);on_delete(cascade)"`
	PackageStatus VMPackageStatus `json:"vm_package_status" orm:"column(vm_package_status)"`
	Packages      []*Package      `json:"-" orm:"rel(m2m);rel_through(github.com/inspursoft/itpserver/src/apiserver/models.Installation)"`
}

func (vm *VM) TableName() string {
	return "vm"
}

type VMSpec struct {
	VM      *VM    `json:"-" orm:"column(vm_id);rel(one)"`
	ID      int64  `json:"-" orm:"column(id);auto;pk"`
	VID     string `json:"vid" orm:"column(vid)"`
	CPUs    int32  `json:"cpus" orm:"column(cpus)"`
	Memory  string `json:"memory" orm:"column(memory)"`
	Storage string `json:"storage" orm:"column(storage)"`
	Extras  string `json:"extras" orm:"column(extras)"`
}

func (spec *VMSpec) TableName() string {
	return "vm_spec"
}

type VMWithSpec struct {
	IP   string `json:"vm_ip"`
	Name string `json:"vm_name"`
	OS   string `json:"vm_os"`
	Spec VMSpec `json:"vm_spec"`
}
