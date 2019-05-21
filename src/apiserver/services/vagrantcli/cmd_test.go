package vagrantcli_test

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/inspursoft/itpserver/src/apiserver/models"
	"github.com/inspursoft/itpserver/src/apiserver/services/vagrantcli"
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
func TestVagrantCli(t *testing.T) {
	vmWithSpec := models.VMWithSpec{
		Name: "vm-31", OS: "ubuntu.box", IP: "172.28.128.31",
		Spec: models.VMSpec{
			CPUs:   1,
			Memory: "1024",
		}}

	t.Run("Create VM", func(t *testing.T) {
		err := vagrantcli.NewClient(vmWithSpec, os.Stdout).Create()
		assert := assert.New(t)
		itpErr := models.AssertITPError(err)
		assert.True(itpErr.HasNoError())
	})
	t.Run("Halt VM", func(t *testing.T) {
		err := vagrantcli.NewClient(vmWithSpec, os.Stdout).Halt()
		assert := assert.New(t)
		itpErr := models.AssertITPError(err)
		assert.True(itpErr.HasNoError())
	})
	t.Run("Global status of VM", func(t *testing.T) {
		var buf bytes.Buffer
		err := vagrantcli.NewClient(vmWithSpec, &buf).GlobalStatus()
		beego.Debug(models.ResolveGlobalStatus(buf.String()))
		assert := assert.New(t)
		itpErr := models.AssertITPError(err)
		assert.True(itpErr.HasNoError())
	})
	t.Run("Destroy VM", func(t *testing.T) {
		err := vagrantcli.NewClient(vmWithSpec, os.Stdout).Destroy()
		assert := assert.New(t)
		itpErr := models.AssertITPError(err)
		assert.True(itpErr.HasNoError())
	})
}
