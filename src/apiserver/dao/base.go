package dao

import (
	"github.com/astaxie/beego/orm"
	"github.com/inspursoft/itpserver/src/apiserver/models"
	_ "github.com/lib/pq"
)

func InitDB() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "postgres://itp:root123@localhost:8882/itpdb?sslmode=disable")
	orm.RegisterModel(new(models.VM), new(models.VMSpec), new(models.Package), new(models.Installation))
}
