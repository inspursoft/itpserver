package dao

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/inspursoft/itpserver/src/apiserver/models"
)

type pkgDaoHandler int

func NewPkgDaoHandler() *pkgDaoHandler {
	return new(pkgDaoHandler)
}

func (p *pkgDaoHandler) GetPackage(name ...string) ([]models.Package, error) {
	o := orm.NewOrm()
	q := o.QueryTable("package")
	if name != nil {
		q.Filter("name", name)
	}
	var results []models.Package
	count, err := q.All(&results)
	if err != nil {
		return nil, err
	}
	beego.Info(fmt.Sprintf("Successful got %d package(s) with name: %s", count, name))
	return results, nil

}

func (p *pkgDaoHandler) AddPackages(packages []models.Package) (affected int64, err error) {
	o := orm.NewOrm()
	affected, err = o.InsertMulti(len(packages), packages)
	if err != nil {
		affected = 0
		return
	}
	beego.Info(fmt.Sprintf("Successful inserted %d package(s).", len(packages)))
	return
}

func (p *pkgDaoHandler) DeletePackage(name string) (affected int64, err error) {
	o := orm.NewOrm()
	affected, err = o.QueryTable("package").Filter("name", name).Delete()
	if err != nil {
		affected = 0
		return
	}
	beego.Info(fmt.Sprintf("Successful deleted %d package(s).", affected))
	return
}

func (p *pkgDaoHandler) UpdatePackage(pkg models.Package) (affected int64, err error) {
	o := orm.NewOrm()
	affected, err = o.QueryTable("package").Filter("name", pkg.Name).Update(
		orm.Params{
			"tag": pkg.Tag,
		})
	if err != nil {
		affected = 0
		return
	}
	return
}
