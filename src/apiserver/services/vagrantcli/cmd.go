package vagrantcli

import (
	"bytes"
	"fmt"
	"io"
	"path/filepath"

	"github.com/inspursoft/itpserver/src/apiserver/services"

	"github.com/astaxie/beego"

	"github.com/inspursoft/itpserver/src/apiserver/models"
	"github.com/inspursoft/itpserver/src/apiserver/utils"
)

type vagrantCli struct {
	sourcePath string
	workPath   string
	command    string
	vmWithSpec models.VMWithSpec
	sshClient  *utils.SecureShell
	output     io.Writer
	err        *models.ITPError
}

func NewClient(vmWithSpec models.VMWithSpec, output io.Writer) *vagrantCli {
	sourcePath := beego.AppConfig.String("vagrant::sourcepath")
	baseWorkPath := beego.AppConfig.String("vagrant::baseworkpath")
	command := beego.AppConfig.String("vagrant::command")
	vc := &vagrantCli{sourcePath: sourcePath, workPath: filepath.Join(baseWorkPath, vmWithSpec.Name),
		command: command, vmWithSpec: vmWithSpec, output: output, err: &models.ITPError{}}
	var err error
	vc.sshClient, err = utils.NewSecureShell(vc.output)
	if err != nil {
		vc.err.InternalError(err)
	}
	return vc
}

func (vc *vagrantCli) init() *vagrantCli {
	vmIP := vc.vmWithSpec.IP
	vmName := vc.vmWithSpec.Name
	exists, err := services.NewVMHandler().Exists(models.VM{IP: vmIP, Name: vmName})
	if err != nil {
		vc.err.InternalError(err)
		return vc
	}
	if exists {
		vc.err.Conflict(fmt.Sprintf("VM name: %s or IP: %s", vmName, vmIP), fmt.Errorf("VM already exists with IP: %s or Name: %s", vmIP, vmName))
		return vc
	}
	err = vc.sshClient.CheckDir(vc.workPath)
	if err != nil {
		vc.err.InternalError(err)
	}
	return vc
}

func (vc *vagrantCli) copySources() *vagrantCli {
	if !vc.err.HasNoError() {
		return vc
	}
	err := vc.sshClient.ExecuteCommand(fmt.Sprintf("cp -R %s/* %s", vc.sourcePath, vc.workPath))
	if err != nil {
		vc.err.InternalError(err)
	}
	return vc
}

func (vc *vagrantCli) generateConfig() *vagrantCli {
	if !vc.err.HasNoError() {
		return vc
	}
	output, err := utils.ExecuteTemplate(vc.vmWithSpec, "Vagrantfile")
	if err != nil {
		vc.err.InternalError(err)
		return vc
	}
	err = vc.sshClient.SecureCopyData("Vagrantfile", output, vc.workPath)
	if err != nil {
		vc.err.InternalError(err)
	}
	return vc
}

func (vc *vagrantCli) executeCommand(action string) *vagrantCli {
	if !vc.err.HasNoError() {
		return vc
	}
	if action == "destroy" {
		action = fmt.Sprintf("destroy -f %s", vc.vmWithSpec.Spec.VID)
	}
	err := vc.sshClient.ExecuteCommand(fmt.Sprintf("cd %s && %s %s", vc.workPath, vc.command, action))
	if err != nil {
		vc.err.InternalError(err)
	}
	return vc
}

func (vc *vagrantCli) updateVID() *vagrantCli {
	if !vc.err.HasNoError() {
		return vc
	}
	var buf bytes.Buffer
	NewClient(vc.vmWithSpec, &buf).GlobalStatus()
	globalStatusList := models.ResolveGlobalStatus(buf.String())
	vc.vmWithSpec.Spec.VID = models.GetVIDByWorkPath(globalStatusList, vc.workPath)
	beego.Debug(fmt.Sprintf("Update VM: %s with VID: %s", vc.vmWithSpec.Name, vc.vmWithSpec.Spec.VID))
	return vc
}

func (vc *vagrantCli) record() *vagrantCli {
	if !vc.err.HasNoError() {
		return vc
	}
	err := services.NewVMHandler().Create(vc.vmWithSpec)
	if err != nil {
		vc.err.InternalError(err)
	}
	return vc
}

func (vc *vagrantCli) remove() *vagrantCli {
	if !vc.err.HasNoError() {
		return vc
	}
	beego.Debug("Removing VM with VID: %s", vc.vmWithSpec.Spec.VID)
	err := services.NewVMHandler().DeleteVMByVID(vc.vmWithSpec.Spec.VID)
	if err != nil {
		vc.err.InternalError(err)
	}
	return vc
}

func (vc *vagrantCli) Create() error {
	vc.init().
		copySources().
		generateConfig().
		executeCommand("up").
		updateVID().
		record()
	if !vc.err.HasNoError() && vc.err != nil {
		beego.Error(fmt.Sprintf("Failed to create VM with Vagrant: %+v", vc.err))
	}
	return vc.err
}

func (vc *vagrantCli) Halt() error {
	vc.executeCommand("halt")
	if !vc.err.HasNoError() && vc.err != nil {
		beego.Error(fmt.Sprintf("Failed to halt VM with error: %+v", vc.err))
	}
	return vc.err
}

func (vc *vagrantCli) Destroy() error {
	vc.updateVID().executeCommand("destroy").remove()
	if !vc.err.HasNoError() && vc.err != nil {
		beego.Error(fmt.Sprintf("Failed to destroy VM with error: %+v", vc.err))
	}
	return vc.err
}

func (vc *vagrantCli) GlobalStatus() error {
	vc.executeCommand("global-status")
	if !vc.err.HasNoError() && vc.err != nil {
		beego.Error(fmt.Sprintf("Failed to get global status of VM with error: %+v", vc.err))
	}
	return vc.err
}
