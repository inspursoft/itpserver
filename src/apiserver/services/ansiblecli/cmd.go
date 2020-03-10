package ansiblecli

import (
	"fmt"
	"io"
	"path"
	"path/filepath"

	"github.com/astaxie/beego/logs"

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

var ansibleCommand = "PATH=/usr/local/bin:$PATH ansible-playbook"

func NewClient(vmWithSpec models.VMWithSpec, pkg models.PackageVO, output io.Writer) *ansibleCli {
	hostIP := beego.AppConfig.String("ansible::hostip")
	pathPrefix := beego.AppConfig.String("pathprefix")
	uploadPath := path.Join(pathPrefix, beego.AppConfig.String("ansible::uploadpath"))
	sourcePath := path.Join(pathPrefix, beego.AppConfig.String("ansible::sourcepath"))
	baseWorkPath := path.Join(pathPrefix, beego.AppConfig.String("ansible::baseworkpath"))
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

func (ac *ansibleCli) cleanUp() *ansibleCli {
	if !ac.err.HasNoError() {
		return ac
	}
	err := ac.sshClient.RemoveDir(filepath.Join(ac.workPath))
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
	err := ac.sshClient.CheckDir(uploadSourcePath)
	if err != nil {
		ac.err.InternalError(err)
	}
	err = ac.sshClient.SecureCopy(uploadSourcePath, ac.workPath)
	if err != nil {
		ac.err.InternalError(err)
	}
	return ac
}

func (ac *ansibleCli) unzipPackage() *ansibleCli {
	if !ac.err.HasNoError() {
		return ac
	}
	ac.pkg.SourceName = ac.pkg.Name + ac.pkg.Tag + ".zip"
	err := ac.sshClient.ExecuteCommand(fmt.Sprintf("cd %s && unzip %s", ac.workPath, ac.pkg.SourceName))
	if err != nil {
		ac.err.InternalError(err)
	}
	return ac
}

func (ac *ansibleCli) preExecution() *ansibleCli {
	if !ac.err.HasNoError() {
		return ac
	}
	targetPath := filepath.Join(ac.workPath, ac.pkg.Name)
	err := ac.sshClient.ExecuteCommand(fmt.Sprintf("cd %s && sh prepare.sh", targetPath))
	if err != nil {
		logs.Warning("No prepare.sh found while doing pre execution.")
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
		cleanUp().
		transferPackage().
		unzipPackage().
		preExecution().
		generateInstall().
		generateHosts().
		recordPackage()
	if !ac.err.HasNoError() && ac.err != nil {
		beego.Error(fmt.Sprintf("Failed to transfer package to server: %+v", ac.err))
		return ac.err
	}
	return nil
}

func (ac *ansibleCli) TransferWithoutGenerateConfig() error {
	beego.Debug("Start transfering without generating configures...")
	ac.init().
		cleanUp().
		transferPackage().
		unzipPackage().
		preExecution().
		recordPackage()
	if !ac.err.HasNoError() && ac.err != nil {
		beego.Error(fmt.Sprintf("Failed to transfer package without generating configs to server: %+v", ac.err))
		return ac.err
	}
	return nil
}

func (ac *ansibleCli) Install() error {
	targetPath := filepath.Join(ac.workPath, ac.pkg.Name)
	ac.init().
		executeCommand(fmt.Sprintf("cd %s && %s -i hosts install.yml", targetPath, ansibleCommand)).
		recordInstall()
	if !ac.err.HasNoError() && ac.err != nil {
		beego.Error(fmt.Sprintf("Failed to execute Ansible client: %+v", ac.err))
		return ac.err
	}
	return nil
}
