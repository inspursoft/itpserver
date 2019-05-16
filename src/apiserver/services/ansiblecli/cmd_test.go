package ansiblecli_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/astaxie/beego"
	"github.com/inspursoft/itpserver/src/apiserver/models"
	"github.com/inspursoft/itpserver/src/apiserver/services/ansiblecli"
	"github.com/stretchr/testify/assert"
)

const appPath = "../../conf"

func TestMain(m *testing.M) {
	os.Setenv("templatepath", "../../templates")
	beego.LoadAppConfig("ini", filepath.Join(appPath, "app.conf"))
	os.Exit(m.Run())
}
func TestAnsibleCli(t *testing.T) {
	vmWithSpec := models.VMWithSpec{IP: "172.28.128.16"}
	pkgList := []models.PackageVO{
		models.PackageVO{Name: "golang", Tag: "1.10"},
	}
	t.Run("Ansible Install", func(t *testing.T) {
		err := ansiblecli.NewClient(os.Stdout).Install(vmWithSpec, pkgList)
		assert := assert.New(t)
		assert.Nil(err)
	})
}
