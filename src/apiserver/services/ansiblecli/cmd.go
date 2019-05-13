package ansiblecli

import (
	"fmt"
	"path/filepath"

	"github.com/astaxie/beego"

	"github.com/inspursoft/itpserver/src/apiserver/models"
	"github.com/inspursoft/itpserver/src/apiserver/utils"
)

type ansibleCli struct {
	localIP  string
	workpath string
	install  *models.Install
	hosts    *models.Hosts
	err      error
}

func NewClient(localIP, workpath string) *ansibleCli {
	return &ansibleCli{localIP: localIP, workpath: workpath}
}

func (ac *ansibleCli) init() *ansibleCli {
	err := utils.CheckDir(ac.workpath)
	if err != nil {
		ac.err = err
		return ac
	}
	return ac
}

func (ac *ansibleCli) generateInstall(pkgList []models.PackageVO) *ansibleCli {
	ac.install = models.YieldInstall("install")
	ac.install.AddRoles(pkgList)
	err := utils.MarshalToYAML(ac.install, filepath.Join(ac.workpath, "install.yml"))
	if err != nil {
		ac.err = err
	}
	return ac
}

func (ac *ansibleCli) generateHosts(name string, ipList []string) *ansibleCli {
	ac.hosts = models.YieldHosts(ac.localIP)
	ac.hosts.AddTarget(name, ipList)
	err := utils.ExecuteTemplate(ac.hosts, "hosts", ac.workpath)
	if err != nil {
		ac.err = err
	}
	return ac
}

func (ac *ansibleCli) Prepare(vmWithSpec models.VMWithSpec, pkgList []models.PackageVO) error {
	ac.init().generateInstall(pkgList).generateHosts("install", []string{vmWithSpec.IP})
	if ac.err != nil {
		beego.Error(fmt.Sprintf("Failed to prepare Ansible client: %+v", ac.err))
		return ac.err
	}
	return nil
}
