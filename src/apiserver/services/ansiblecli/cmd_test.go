package ansiblecli_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/inspursoft/itpserver/src/apiserver/models"
	"github.com/inspursoft/itpserver/src/apiserver/services/ansiblecli"
	"github.com/stretchr/testify/assert"
)

const appPath = "../../conf"

func TestMain(m *testing.M) {
	os.Setenv("templatepath", "../../templates")
	beego.LoadAppConfig("ini", filepath.Join(appPath, "app.conf"))
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "postgres://itp:root123@localhost:8882/itpdb?sslmode=disable")
	orm.RegisterModel(new(models.VM), new(models.VMSpec), new(models.Package), new(models.Installation))
	os.Exit(m.Run())
}
func TestAnsibleCli(t *testing.T) {
	vmWithSpec := models.VMWithSpec{IP: "172.28.128.16"}
	pkg := models.PackageVO{Name: "golang", Tag: "1.10"}
	t.Run("Ansible Install", func(t *testing.T) {
		err := ansiblecli.NewClient(vmWithSpec, pkg, os.Stdout).Install()
		assert := assert.New(t)
		itpErr := models.AssertITPError(err)
		assert.True(itpErr.HasNoError())
	})
}
