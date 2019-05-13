package vagrantcli

import (
	"fmt"

	"github.com/astaxie/beego"

	"github.com/inspursoft/itpserver/src/apiserver/models"
	"github.com/inspursoft/itpserver/src/apiserver/utils"
)

type vagrantCli struct {
	workpath string
	err      error
}

func NewClient(workpath string) *vagrantCli {
	return &vagrantCli{workpath: workpath}
}

func (vc *vagrantCli) init() *vagrantCli {
	err := utils.CheckDir(vc.workpath)
	if err != nil {
		vc.err = err
		return vc
	}
	return vc
}

func (vc *vagrantCli) generateConfig(vmWithSpec models.VMWithSpec) *vagrantCli {
	err := utils.ExecuteTemplate(vmWithSpec, "Vagrantfile", vc.workpath)
	if err != nil {
		vc.err = err
		return vc
	}
	return vc
}

func (vc *vagrantCli) CreateVM(vmWithSpec models.VMWithSpec) error {
	vc.init().generateConfig(vmWithSpec)
	if vc.err != nil {
		beego.Error(fmt.Sprintf("Failed to create VM with Vagrant: %+v", vc.err))
		return vc.err
	}
	return nil
}
