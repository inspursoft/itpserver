package ansiblecli_test

import (
	"os"
	"testing"

	"github.com/inspursoft/itpserver/src/apiserver/models"
	"github.com/inspursoft/itpserver/src/apiserver/services/ansiblecli"
	"github.com/stretchr/testify/assert"
)

func TestAnsibleCli(t *testing.T) {
	vmWithSpec := models.VMWithSpec{IP: "192.168.1.102"}
	pkgList := []models.PackageVO{
		models.PackageVO{Name: "Ansible", Tag: "15.0"},
		models.PackageVO{Name: "Docker", Tag: "18.09"},
	}
	localIP := "10.0.0.1"
	workpath := "/tmp/installs"
	err := ansiblecli.NewClient(localIP, workpath).Prepare(vmWithSpec, pkgList)
	assert := assert.New(t)
	assert.Nil(err)
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
