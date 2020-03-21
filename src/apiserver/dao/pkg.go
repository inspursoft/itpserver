package dao

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/inspursoft/itpserver/src/apiserver/models"
)

type PkgDaoHandler int

func (p *PkgDaoHandler) GetPackage(vmName string, name string, tag string) (*models.Package, error) {
	o := orm.NewOrm()
	pkg := models.Package{Name: name, Tag: tag, VMName: vmName}
	err := o.Read(&pkg, "Name", "Tag", "VMName")
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &pkg, nil
}

func (p *PkgDaoHandler) GetPackageList(vmName string, name string, tag string) ([]*models.Package, error) {
	o := orm.NewOrm()
	q := o.QueryTable("package")
	q = q.Filter("VMName", vmName)
	if name != "" {
		q = q.Filter("name", name)
	}
	if tag != "" {
		q = q.Filter("tag", tag)
	}
	var results []*models.Package
	count, err := q.All(&results)
	if err != nil {
		return nil, err
	}
	beego.Info(fmt.Sprintf("Successful got %d package(s) with name: %s", count, name))
	return results, nil
}

func (p *PkgDaoHandler) AddPackage(pkg *models.Package) (insertedID int64, err error) {
	o := orm.NewOrm()
	insertedID, err = o.Insert(pkg)
	if err != nil {
		insertedID = 0
		return
	}
	beego.Info(fmt.Sprintf("Successful inserted package with ID: %d.", insertedID))
	return
}

func (p *PkgDaoHandler) DeletePackage(vmName string, name string, tag string) (affected int64, err error) {
	o := orm.NewOrm()
	affected, err = o.QueryTable("package").Filter("vm_name___exact", vmName).Filter("name__exact", name).Filter("tag__exact", tag).Delete()
	if err != nil {
		if err == orm.ErrNoRows {
			return 0, nil
		}
		affected = 0
		return
	}
	beego.Info(fmt.Sprintf("Successful deleted package with ID: %d.", affected))
	return
}

func (p *PkgDaoHandler) UpdatePackage(pkg models.Package) (affected int64, err error) {
	o := orm.NewOrm()
	affected, err = o.QueryTable("package").Filter("name", pkg.Name).Update(
		orm.Params{
			"tag": pkg.Tag,
		})
	if err != nil {
		if err == orm.ErrNoRows {
			return 0, nil
		}
		affected = 0
		return
	}
	return
}
