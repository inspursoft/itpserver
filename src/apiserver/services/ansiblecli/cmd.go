package ansiblecli

import (
	"errors"
	"fmt"
	"io"
	"path/filepath"

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
	vmWithSpec models.VMWithSpec
	sshClient  *utils.SecureShell
	install    *models.SimpleInstall
	hosts      *models.Hosts
	output     io.Writer
	err        *models.ITPError
}

func NewClient(vmWithSpec models.VMWithSpec, output io.Writer) *ansibleCli {
	hostIP := beego.AppConfig.String("ansible::hostip")
	command := beego.AppConfig.String("ansible::command")
	sourcePath := beego.AppConfig.String("ansible::sourcepath")
	baseWorkPath := beego.AppConfig.String("ansible::baseworkpath")
	ac := &ansibleCli{hostIP: hostIP, command: command, vmWithSpec: vmWithSpec,
		sourcePath: sourcePath, workPath: filepath.Join(baseWorkPath, vmWithSpec.Name),
		output: output, err: &models.ITPError{}}
	var err error
	ac.sshClient, err = utils.NewSecureShell(ac.output)
	if err != nil {
		ac.err.InternalError(err)
	}
	return ac
}

func (ac *ansibleCli) init() *ansibleCli {
	vmIP := ac.vmWithSpec.IP
	vmName := ac.vmWithSpec.Name
	exists, err := services.NewVMHandler().Exists(models.VM{IP: vmIP, Name: vmName})
	if err != nil {
		ac.err.InternalError(err)
		return ac
	}
	if !exists {
		ac.err.Notfound(fmt.Sprintf("VM: %s with IP: %s does not exist", vmName, vmIP), err)
		return ac
	}
	err = ac.sshClient.CheckDir(ac.workPath)
	if err != nil {
		ac.err.InternalError(err)
	}
	return ac
}

func (ac *ansibleCli) copySources() *ansibleCli {
	if !ac.err.HasNoError() {
		return ac
	}
	err := ac.sshClient.ExecuteCommand(fmt.Sprintf("cp -R %s/* %s", ac.sourcePath, ac.workPath))
	if err != nil {
		ac.err.InternalError(err)
	}
	return ac
}

func (ac *ansibleCli) generateInstall(pkgList []models.PackageVO) *ansibleCli {
	if !ac.err.HasNoError() {
		return ac
	}
	ac.install = &models.SimpleInstall{PkgName: pkgList[0].Name}
	output, err := utils.ExecuteTemplate(ac.install, "install.yml")
	if err != nil {
		ac.err.InternalError(err)
		return ac
	}
	err = ac.sshClient.SecureCopyData("install.yml", output, ac.workPath)
	if err != nil {
		ac.err.InternalError(err)
	}
	return ac
}

func (ac *ansibleCli) generateHosts(ipList ...string) *ansibleCli {
	if !ac.err.HasNoError() {
		return ac
	}
	ac.hosts = models.YieldHosts(ac.hostIP)
	ac.hosts.AddTarget("install", ipList)
	output, err := utils.ExecuteTemplate(ac.hosts, "hosts")
	if err != nil {
		ac.err.InternalError(err)
		return ac
	}
	err = ac.sshClient.SecureCopyData("hosts", output, ac.workPath)
	if err != nil {
		ac.err.InternalError(err)
	}
	return ac
}

func (ac *ansibleCli) executeCommand(action string) *ansibleCli {
	if !ac.err.HasNoError() {
		return ac
	}
	err := ac.sshClient.ExecuteCommand(fmt.Sprintf("cd %s && sh ssh.sh %s && %s %s", ac.workPath,
		ac.vmWithSpec.IP, ac.command, action))
	if err != nil {
		ac.err.InternalError(err)
	}
	return ac
}

func (ac *ansibleCli) record(pkgList []models.PackageVO) *ansibleCli {
	if !ac.err.HasNoError() {
		return ac
	}
	vm, err := services.NewVMHandler().GetByIP(ac.vmWithSpec.IP)
	if err != nil {
		ac.err.InternalError(err)
		return ac
	}
	pkg := pkgList[0]
	err = services.NewInstallationHandler().Install(vm.ID, pkg.Name, pkg.Tag)
	if err != nil {
		ac.err.InternalError(err)
	}
	return ac
}

func (ac *ansibleCli) Install(pkgList []models.PackageVO) error {
	if len(pkgList) == 0 {
		return errors.New("No package to install")
	}
	ac.init().
		copySources().
		generateInstall(pkgList).
		generateHosts(ac.vmWithSpec.IP).
		executeCommand("-i hosts install.yml").
		record(pkgList)

	if ac.err != nil {
		beego.Error(fmt.Sprintf("Failed to execute Ansible client: %+v", ac.err))
		return ac.err
	}
	return nil
}
