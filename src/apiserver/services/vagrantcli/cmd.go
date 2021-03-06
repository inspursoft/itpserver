package vagrantcli

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/astaxie/beego/logs"

	"github.com/inspursoft/itpserver/src/apiserver/services"

	"github.com/astaxie/beego"

	"github.com/inspursoft/itpserver/src/apiserver/models"
	"github.com/inspursoft/itpserver/src/apiserver/utils"
)

type vagrantCli struct {
	sourcePath string
	workPath   string
	outputPath string
	command    string
	vmWithSpec models.VMWithSpec
	sshClient  *utils.SecureShell
	output     io.Writer
	err        *models.ITPError
}

var vagrantCommand = `PATH=/usr/local/bin:$PATH vagrant`
var pathPrefix = beego.AppConfig.String("pathprefix")
var sourcePath = path.Join(pathPrefix, beego.AppConfig.String("vagrant::sourcepath"))
var baseWorkPath = path.Join(pathPrefix, beego.AppConfig.String("vagrant::baseworkpath"))
var outputPath = path.Join(pathPrefix, beego.AppConfig.String("vagrant::outputpath"))

func NewClient(vmWithSpec models.VMWithSpec, output io.Writer) *vagrantCli {
	vc := &vagrantCli{sourcePath: sourcePath, workPath: filepath.Join(baseWorkPath, vmWithSpec.Name),
		outputPath: outputPath, vmWithSpec: vmWithSpec, output: output, err: &models.ITPError{}}
	return vc
}

func NewEaseClient(vmName string, output io.Writer) *vagrantCli {
	var vmWithSpec models.VMWithSpec
	vmWithSpec.Name = vmName
	logs.Debug(fmt.Sprintf("vmName: %s", vmName))
	return NewClient(vmWithSpec, output)
}

func (vc *vagrantCli) newSSHClient() *vagrantCli {
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
	return vc
}

func (vc *vagrantCli) checkDirs() *vagrantCli {
	err := vc.sshClient.CheckDir(vc.sourcePath)
	if err != nil {
		vc.err.InternalError(err)
	}
	err = vc.sshClient.CheckDir(vc.workPath)
	if err != nil {
		vc.err.InternalError(err)
	}
	err = vc.sshClient.CheckDir(vc.outputPath)
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

func (vc *vagrantCli) resolveVagrantfile() *vagrantCli {
	if !vc.err.HasNoError() {
		return vc
	}
	vagrantFilePath := filepath.Join(vc.workPath, "Vagrantfile")
	if _, err := os.Stat(vagrantFilePath); os.IsNotExist(err) {
		vc.err.Notfound("Vagrantfile", err)
		return vc
	}
	f, err := os.Open(vagrantFilePath)
	if err != nil {
		vc.err.InternalError(err)
		return vc
	}
	defer f.Close()
	re, err := regexp.Compile("(ip|box|cpus|memory)\\s*[:=]\\s*[\\\"](.*)[\\\"]$")
	if err != nil {
		vc.err.InternalError(err)
		return vc
	}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if !strings.HasPrefix(line, "#") && re.MatchString(line) {
			groups := re.FindAllStringSubmatch(line, len(line))
			name := groups[0][1]
			value := groups[0][2]
			if name == "ip" {
				vc.vmWithSpec.IP = value
			} else if name == "box" {
				vc.vmWithSpec.OS = value
			} else if name == "cpus" {
				val, _ := strconv.Atoi(value)
				vc.vmWithSpec.Spec.CPUs = int32(val)
			} else if name == "memory" {
				vc.vmWithSpec.Spec.Memory = value
			}
		}
	}
	if vc.vmWithSpec.IP == "" || vc.vmWithSpec.OS == "" ||
		vc.vmWithSpec.Spec.CPUs == 0 || vc.vmWithSpec.Spec.Memory == "" {
		vc.err.Notfound("Vagrantfile", errors.New("Vagrantfile is missing as required value"))
		return vc
	}
	exists, err := services.NewVMHandler().Exists(models.VM{IP: vc.vmWithSpec.IP, Name: vc.vmWithSpec.Name})
	if err != nil {
		vc.err.InternalError(err)
		return vc
	}
	if exists {
		vc.err.Conflict(fmt.Sprintf("VM: %s already exist.", vc.vmWithSpec.Name), fmt.Errorf("VM %s already exist", vc.vmWithSpec.Name))
	}
	return vc
}

func (vc *vagrantCli) executeCommand(command string) *vagrantCli {
	if !vc.err.HasNoError() {
		return vc
	}
	err := vc.sshClient.ExecuteCommand(command)
	if err != nil {
		vc.err.InternalError(err)
	}
	return vc
}
func (vc *vagrantCli) loadSpec() *vagrantCli {
	vm, err := services.NewVMHandler().GetByName(vc.vmWithSpec.Name)
	if err != nil {
		vc.err.InternalError(err)
		return vc
	}
	if vm == nil {
		vc.err.Notfound("VM", fmt.Errorf("VM with name: %s does not exist", vc.vmWithSpec.Name))
		return vc
	}
	vc.vmWithSpec.Spec = *vm.Spec
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
	beego.Debug(fmt.Sprintf("Removing VM with VID: %s", vc.vmWithSpec.Spec.VID))
	err := services.NewVMHandler().DeleteVMByVID(vc.vmWithSpec.Spec.VID)
	if err != nil {
		vc.err.InternalError(err)
	}
	return vc
}

func (vc *vagrantCli) Create() error {
	vc.newSSHClient().
		init().
		checkDirs().
		copySources().
		generateConfig().
		executeCommand(fmt.Sprintf("cd %s && %s up && PATH=/bin:$PATH sh ssh.sh %s", vc.workPath, vagrantCommand, vc.vmWithSpec.IP)).
		updateVID().
		record()
	if !vc.err.HasNoError() && vc.err != nil {
		beego.Error(fmt.Sprintf("Failed to create VM with Vagrant: %+v", vc.err))
		return vc.err
	}
	return nil
}

func (vc *vagrantCli) CreateByVagrantfile() error {
	vc.newSSHClient().
		checkDirs().
		copySources().
		resolveVagrantfile().
		executeCommand(fmt.Sprintf("cd %s && %s up", vc.workPath, vagrantCommand)).
		updateVID().
		record()

	if !vc.err.HasNoError() && vc.err != nil {
		beego.Error(fmt.Sprintf("Failed to create VM with Vagrant: %+v", vc.err))
		return vc.err
	}
	return nil
}

func (vc *vagrantCli) Halt() error {
	vc.newSSHClient().executeCommand(fmt.Sprintf("cd %s", vc.workPath)).executeCommand(fmt.Sprintf("%s halt", vagrantCommand))
	if !vc.err.HasNoError() && vc.err != nil {
		beego.Error(fmt.Sprintf("Failed to halt VM with error: %+v", vc.err))
		return vc.err
	}
	return nil
}

func (vc *vagrantCli) Destroy() error {
	vc.loadSpec().newSSHClient().executeCommand(
		fmt.Sprintf("%s destroy -f %s && rm -rf %s && rm -f %s",
			vagrantCommand, vc.vmWithSpec.Spec.VID, vc.workPath,
			path.Join(vc.outputPath, vc.vmWithSpec.Name+".box"))).remove()
	if !vc.err.HasNoError() && vc.err != nil {
		beego.Error(fmt.Sprintf("Failed to destroy VM with error: %+v", vc.err))
		return vc.err
	}
	return nil
}

func (vc *vagrantCli) GlobalStatus() error {
	vc.newSSHClient().executeCommand(fmt.Sprintf("%s global-status", vagrantCommand))
	if !vc.err.HasNoError() && vc.err != nil {
		beego.Error(fmt.Sprintf("Failed to get global status of VM with error: %+v", vc.err))
		return vc.err
	}
	return nil
}

func (vc *vagrantCli) Package() error {
	vc.loadSpec().newSSHClient().updateVID().executeCommand(fmt.Sprintf("%s package %s --output %s", vagrantCommand, vc.vmWithSpec.Spec.VID, filepath.Join(vc.outputPath, fmt.Sprintf("%s.box", vc.vmWithSpec.Name))))
	if !vc.err.HasNoError() && vc.err != nil {
		beego.Error(fmt.Sprintf("Failed to package VM with error: %+v", vc.err))
		return vc.err
	}
	return nil
}
