package vagrantcli_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/astaxie/beego"
	"github.com/inspursoft/itpserver/src/apiserver/models"
	"github.com/inspursoft/itpserver/src/apiserver/services/vagrantcli"
	"github.com/stretchr/testify/assert"
)

const appPath = "../../conf"

func TestMain(m *testing.M) {
	os.Setenv("templatepath", "../../templates")
	beego.LoadAppConfig("ini", filepath.Join(appPath, "app.conf"))
	os.Exit(m.Run())
}
func TestVagrantCli(t *testing.T) {
	vmWithSpec := models.VMWithSpec{
		Name: "ubuntu-vm-1", OS: "ubuntu.box", IP: "172.28.128.16",
		Spec: models.VMSpec{
			CPUs:   1,
			Memory: "1024",
		}}
	t.Run("Create VM", func(t *testing.T) {
		err := vagrantcli.NewClient(vmWithSpec).Create()
		assert := assert.New(t)
		assert.Nil(err)
	})
	// t.Run("Halt VM", func(t *testing.T) {
	// 	err := vagrantcli.NewClient(vmWithSpec).Halt()
	// 	assert := assert.New(t)
	// 	assert.Nil(err)
	// })
	// t.Run("Global status of VM", func(t *testing.T) {
	// 	err := vagrantcli.NewClient(vmWithSpec).GlobalStatus()
	// 	assert := assert.New(t)
	// 	assert.Nil(err)
	// })
}
