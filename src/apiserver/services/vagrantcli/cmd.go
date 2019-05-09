package vagrantcli

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/astaxie/beego"

	"github.com/inspursoft/itpserver/src/apiserver/models"
)

type vagrantCli struct {
	workpath string
	err      error
}

func NewClient(workpath string) *vagrantCli {
	return &vagrantCli{workpath: workpath, err: nil}
}

func (vc *vagrantCli) init(vmName string) *vagrantCli {
	if _, err := os.Stat(vc.workpath); os.IsNotExist(err) {
		err = os.MkdirAll(vc.workpath, 0755)
		if err != nil {
			beego.Error(fmt.Sprintf("Failed to create Vagrant workpath: %+v", err))
			vc.err = err
			return vc
		}
	}
	return vc
}

func (vc *vagrantCli) generateConfig(vmWithSpec models.VMWithSpec) *vagrantCli {
	t, err := template.ParseFiles("templates/Vagrantfile")
	if err != nil {
		vc.err = err
		return vc
	}
	f, err := os.OpenFile(filepath.Join(vc.workpath, "Vagrantfile"), os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		vc.err = err
		return vc
	}
	defer f.Close()
	t.Execute(f, vmWithSpec)
	return vc
}

func (vc *vagrantCli) CreateVM(vmWithSpec models.VMWithSpec) error {
	vc.init(vmWithSpec.Name).generateConfig(vmWithSpec)
	if vc.err != nil {
		beego.Error(fmt.Sprintf("Failed to create VM with Vagrant: %+v", vc.err))
		return vc.err
	}
	return nil
}
