package dao

import (
	"github.com/astaxie/beego/orm"
	"github.com/inspursoft/itpserver/src/apiserver/models"
	_ "github.com/lib/pq"
)

func InitDB() {
	orm.RegisterModel(new(models.VM), new(models.VMSpec), new(models.Package), new(models.Installation))
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "postgres://itp:root123@postgredb:5432/itpdb?sslmode=disable")

}
