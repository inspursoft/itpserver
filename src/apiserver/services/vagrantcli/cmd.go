package vagrantcli

import (
	"fmt"
	"io"

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
	err        error
}

func NewClient(vmWithSpec models.VMWithSpec, output io.Writer) *vagrantCli {
	sourcePath := beego.AppConfig.String("vagrant::sourcepath")
	workPath := beego.AppConfig.String("vagrant::workpath")
	command := beego.AppConfig.String("vagrant::command")
	return &vagrantCli{sourcePath: sourcePath, workPath: workPath,
		command: command, vmWithSpec: vmWithSpec, output: output}
}

func (vc *vagrantCli) init() *vagrantCli {
	var err error
	vc.sshClient, err = utils.NewSecureShell(vc.output)
	if err != nil {
		vc.err = err
		return vc
	}
	vc.err = vc.sshClient.CheckDir(vc.workPath)
	return vc
}

func (vc *vagrantCli) copySources() *vagrantCli {
	if vc.err != nil {
		return vc
	}
	vc.err = vc.sshClient.ExecuteCommand(fmt.Sprintf("cp -R %s/* %s", vc.sourcePath, vc.workPath))
	return vc
}

func (vc *vagrantCli) generateConfig() *vagrantCli {
	if vc.err != nil {
		return vc
	}
	output, err := utils.ExecuteTemplate(vc.vmWithSpec, "Vagrantfile")
	if err != nil {
		vc.err = err
		return vc
	}
	vc.err = vc.sshClient.SecureCopyData("Vagrantfile", output, vc.workPath)
	return vc
}

func (vc *vagrantCli) executeCommand(action string) *vagrantCli {
	if vc.err != nil {
		return vc
	}
	vc.err = vc.sshClient.ExecuteCommand(fmt.Sprintf("cd %s && %s %s", vc.workPath, vc.command, action))
	return vc
}

func (vc *vagrantCli) record() *vagrantCli {
	if vc.err != nil {
		return vc
	}
	vc.err = services.NewVMHandler().Create(vc.vmWithSpec)
	return vc
}

func (vc *vagrantCli) Create() error {
	vc.init().
		copySources().
		generateConfig().
		executeCommand("up").
		record()
	if vc.err != nil {
		beego.Error(fmt.Sprintf("Failed to create VM with Vagrant: %+v", vc.err))
		return vc.err
	}
	return nil
}

func (vc *vagrantCli) Halt() error {
	vc.init().executeCommand("halt")
	if vc.err != nil {
		beego.Error(fmt.Sprintf("Failed to halt VM with error: %+v", vc.err))
		return vc.err
	}
	return nil
}

func (vc *vagrantCli) GlobalStatus() error {
	vc.init().executeCommand("global-status")
	if vc.err != nil {
		beego.Error(fmt.Sprintf("Failed to get global status of VM with error: %+v", vc.err))
		return vc.err
	}
	return nil
}
