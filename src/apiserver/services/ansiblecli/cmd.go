package ansiblecli

import (
	"errors"
	"fmt"

	"github.com/inspursoft/itpserver/src/apiserver/services"

	"github.com/astaxie/beego"

	"github.com/inspursoft/itpserver/src/apiserver/models"
	"github.com/inspursoft/itpserver/src/apiserver/utils"
)

type ansibleCli struct {
	hostIP     string
	command    string
	sourcePath string
	workPath   string
	sshClient  *utils.SecureShell
	install    *models.SimpleInstall
	hosts      *models.Hosts
	err        error
}

func NewClient() *ansibleCli {
	hostIP := beego.AppConfig.String("ansible::hostip")
	command := beego.AppConfig.String("ansible::command")
	sourcePath := beego.AppConfig.String("ansible::sourcepath")
	workPath := beego.AppConfig.String("ansible::workpath")
	return &ansibleCli{hostIP: hostIP, command: command, sourcePath: sourcePath, workPath: workPath}
}

func (ac *ansibleCli) init() *ansibleCli {
	var err error
	ac.sshClient, err = utils.NewSecureShell()
	if err != nil {
		ac.err = err
		return ac
	}
	err = ac.sshClient.CheckDir(ac.workPath)
	if err != nil {
		ac.err = err
	}
	return ac
}

func (ac *ansibleCli) copySources() *ansibleCli {
	if ac.err != nil {
		return ac
	}
	ac.err = ac.sshClient.ExecuteCommand(fmt.Sprintf("cp -R %s/* %s", ac.sourcePath, ac.workPath))
	return ac
}

func (ac *ansibleCli) generateInstall(pkgList []models.PackageVO) *ansibleCli {
	if ac.err != nil {
		return ac
	}
	ac.install = &models.SimpleInstall{PkgName: pkgList[0].Name}
	output, err := utils.ExecuteTemplate(ac.install, "install.yml")
	if err != nil {
		ac.err = err
		return ac
	}
	ac.err = ac.sshClient.SecureCopyData("install.yml", output, ac.workPath)
	return ac
}

func (ac *ansibleCli) generateHosts(name string, ipList ...string) *ansibleCli {
	if ac.err != nil {
		return ac
	}
	ac.hosts = models.YieldHosts(ac.hostIP)
	ac.hosts.AddTarget(name, ipList)
	output, err := utils.ExecuteTemplate(ac.hosts, "hosts")
	if err != nil {
		ac.err = err
		return ac
	}
	ac.err = ac.sshClient.SecureCopyData("hosts", output, ac.workPath)
	return ac
}

func (ac *ansibleCli) executeCommand(action string) *ansibleCli {
	if ac.err != nil {
		return ac
	}
	ac.err = ac.sshClient.ExecuteCommand(fmt.Sprintf("cd %s && %s %s", ac.workPath, ac.command, action))
	return ac
}

func (ac *ansibleCli) record(vmWithSpec models.VMWithSpec, pkgList []models.PackageVO) *ansibleCli {
	if ac.err != nil {
		return ac
	}
	vm, err := services.NewVMHandler().GetByIP(vmWithSpec.IP)
	if err != nil {
		ac.err = err
		return ac
	}
	pkg := pkgList[0]
	ac.err = services.NewInstallationHandler().Install(vm.ID, pkg.Name, pkg.Tag)
	return ac
}

func (ac *ansibleCli) Install(vmWithSpec models.VMWithSpec, pkgList []models.PackageVO) error {
	if len(pkgList) == 0 {
		return errors.New("No package to install")
	}
	ac.init().
		copySources().
		generateInstall(pkgList).
		generateHosts("install", vmWithSpec.IP).
		executeCommand("-i hosts install.yml").
		record(vmWithSpec, pkgList)

	if ac.err != nil {
		beego.Error(fmt.Sprintf("Failed to execute Ansible client: %+v", ac.err))
		return ac.err
	}
	return nil
}
