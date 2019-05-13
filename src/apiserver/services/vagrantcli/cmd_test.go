package vagrantcli_test

import (
	"testing"

	"github.com/inspursoft/itpserver/src/apiserver/models"
	"github.com/inspursoft/itpserver/src/apiserver/services/vagrantcli"
	"github.com/stretchr/testify/assert"
)

func TestVagrantCli(t *testing.T) {
	vmWithSpec := models.VMWithSpec{Name: "ubuntu", OS: "Ubuntu-18.04", IP: "192.168.1.102"}
	workpath := "/tmp/vagrants"
	err := vagrantcli.NewClient(workpath).CreateVM(vmWithSpec)
	assert := assert.New(t)
	assert.Nil(err)
}
