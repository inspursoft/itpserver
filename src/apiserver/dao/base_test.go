package dao_test

import (
	"os"
	"testing"

	"github.com/astaxie/beego/orm"
	"github.com/inspursoft/itpserver/src/apiserver/models"
	_ "github.com/lib/pq"
)

func TestMain(m *testing.M) {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "postgres://itp:root123@localhost:8882/itpdb?sslmode=disable")
	orm.RegisterModel(new(models.VM), new(models.VMSpec), new(models.Package), new(models.Installation))
	os.Exit(func() int {
		defer func() {
			o := orm.NewOrm()
			o.Raw(`delete from vm; delete from vm_spec; delete from package; delete from installation;`).Exec()
		}()
		return m.Run()
	}())
}
