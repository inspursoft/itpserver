package ansiblecli

import (
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
	uploadPath string
	sourcePath string
	workPath   string
	vmWithSpec models.VMWithSpec
	pkg        models.PackageVO
	sshClient  *utils.SecureShell
	install    *models.SimpleInstall
	hosts      *models.Hosts
	output     io.Writer
	err        *models.ITPError
}

func NewClient(vmWithSpec models.VMWithSpec, pkg models.PackageVO, output io.Writer) *ansibleCli {
	hostIP := beego.AppConfig.String("ansible::hostip")
	uploadPath := beego.AppConfig.String("ansible::uploadpath")
	sourcePath := beego.AppConfig.String("ansible::sourcepath")
	baseWorkPath := beego.AppConfig.String("ansible::baseworkpath")
	ac := &ansibleCli{hostIP: hostIP,
		vmWithSpec: vmWithSpec, pkg: pkg,
		uploadPath: uploadPath, sourcePath: sourcePath,
		workPath: filepath.Join(baseWorkPath, vmWithSpec.Name),
		output:   output, err: &models.ITPError{}}
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

func (ac *ansibleCli) transferPackage() *ansibleCli {
	if !ac.err.HasNoError() {
		return ac
	}
	vmName := ac.vmWithSpec.Name
	if vmName == "" {
		ac.err.Notfound("VM", fmt.Errorf("VM name is required"))
		return ac
	}
	uploadSourcePath := filepath.Join(ac.uploadPath, vmName)
	err := ac.sshClient.SecureCopy(uploadSourcePath, ac.workPath)
	if err != nil {
		ac.err.InternalError(err)
	}
	return ac
}

func (ac *ansibleCli) unzipPackage() *ansibleCli {
	if !ac.err.HasNoError() {
		return ac
	}
	err := ac.sshClient.ExecuteCommand(fmt.Sprintf("cd %s && unzip %s", ac.workPath, ac.pkg.SourceName))
	if err != nil {
		ac.err.InternalError(err)
	}
	return ac
}

func (ac *ansibleCli) generateInstall() *ansibleCli {
	if !ac.err.HasNoError() {
		return ac
	}
	ac.install = &models.SimpleInstall{PkgName: ac.pkg.Name}
	output, err := utils.ExecuteTemplate(ac.install, "install.yml")
	if err != nil {
		ac.err.InternalError(err)
		return ac
	}
	targetPath := filepath.Join(ac.workPath, ac.pkg.Name)
	err = ac.sshClient.SecureCopyData("install.yml", output, targetPath)
	if err != nil {
		ac.err.InternalError(err)
	}
	return ac
}

func (ac *ansibleCli) generateHosts() *ansibleCli {
	if !ac.err.HasNoError() {
		return ac
	}
	vm, err := services.NewVMHandler().GetByName(ac.vmWithSpec.Name)
	if err != nil {
		ac.err.InternalError(err)
		return ac
	}
	ac.hosts = models.YieldHosts(ac.hostIP)
	ac.hosts.AddTarget("install", vm.IP)
	output, err := utils.ExecuteTemplate(ac.hosts, "hosts")
	if err != nil {
		ac.err.InternalError(err)
		return ac
	}
	targetPath := filepath.Join(ac.workPath, ac.pkg.Name)
	err = ac.sshClient.SecureCopyData("hosts", output, targetPath)
	if err != nil {
		ac.err.InternalError(err)
	}
	return ac
}

func (ac *ansibleCli) executeCommand(command string) *ansibleCli {
	if !ac.err.HasNoError() {
		return ac
	}
	err := ac.sshClient.ExecuteCommand(command)
	if err != nil {
		ac.err.InternalError(err)
	}
	return ac
}

func (ac *ansibleCli) recordPackage() *ansibleCli {
	if !ac.err.HasNoError() {
		return ac
	}
	err := services.NewPackageHandler().Create(ac.pkg)
	if err != nil {
		ac.err.InternalError(err)
	}
	return ac
}

func (ac *ansibleCli) recordInstall() *ansibleCli {
	if !ac.err.HasNoError() {
		return ac
	}
	vm, err := services.NewVMHandler().GetByIP(ac.vmWithSpec.IP)
	if err != nil {
		ac.err.InternalError(err)
		return ac
	}
	err = services.NewInstallationHandler().Install(vm.ID, ac.pkg.Name, ac.pkg.Tag)
	if err != nil {
		ac.err.InternalError(err)
	}
	return ac
}

func (ac *ansibleCli) Transfer() error {
	ac.init().
		transferPackage().
		unzipPackage().
		generateInstall().
		generateHosts().
		recordPackage()
	if !ac.err.HasNoError() && ac.err != nil {
		beego.Error(fmt.Sprintf("Failed to transfer package to server: %+v", ac.err))
		return ac.err
	}
	return nil
}

func (ac *ansibleCli) Install() error {
	targetPath := filepath.Join(ac.workPath, ac.pkg.Name)
	ac.init().
		executeCommand(fmt.Sprintf("cd %s && ansible-playbook -i hosts install.yml", targetPath)).
		recordInstall()
	if !ac.err.HasNoError() && ac.err != nil {
		beego.Error(fmt.Sprintf("Failed to execute Ansible client: %+v", ac.err))
		return ac.err
	}
	return nil
}
