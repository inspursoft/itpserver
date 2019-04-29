package models

type Installation struct {
	ID  int64    `orm:"column(id);auto;pk"`
	VM  *VM      `orm:"column(vm_id);rel(fk)"`
	Pkg *Package `orm:"column(package_id);rel(fk)"`
}

func (i *Installation) TableName() string {
	return "installation"
}
