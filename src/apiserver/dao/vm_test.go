package dao_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/inspursoft/itpserver/src/apiserver/dao"
	"github.com/inspursoft/itpserver/src/apiserver/models"
)

func TestOperateVM(t *testing.T) {
	vm := models.VM{Name: "ubuntu-vm-1", VMID: "1a2b3c", OS: "CentOS"}
	spec := models.VMSpec{CPUs: 2, Memory: "4Gb", Storage: "1T", Extras: "SSD"}
	var h dao.VMDaoHandler
	t.Run("CreateVM", func(t *testing.T) {
		v, err := h.AddVM(&vm, &spec)
		assert := assert.New(t)
		assert.NotNil(v, "Got VM object after insertion.")
		assert.Nil(err)
	})
	t.Run("GetVM", func(t *testing.T) {
		vm, err := h.GetVM("1a2b3c")
		assert := assert.New(t)
		assert.Nil(err)
		assert.NotNil(vm)
	})
	t.Run("GetVMList", func(t *testing.T) {
		vms, err := h.GetVMList("1a2b3c")
		assert := assert.New(t)
		assert.Nil(err)
		assert.NotNil(vms)
		assert.NotNil(spec)
	})
	t.Run("UpdateVM", func(t *testing.T) {
		affected, err := h.UpdateVM(models.VM{
			Name: "modified-vm", OS: "Ubuntu", VMID: "1a2b3c"})
		assert := assert.New(t)
		assert.Equal(int64(1), affected)
		assert.Nil(err)
	})
	t.Run("UpdateVMSpec", func(t *testing.T) {
		affected, err := h.UpdateVMSpec(models.VM{VMID: "1a2b3c"},
			models.VMSpec{CPUs: 2, Memory: "6Gb", Storage: "3T", Extras: "N/A"})
		assert := assert.New(t)
		assert.Equal(int64(1), affected)
		assert.Nil(err)
	})
	t.Run("DeleteVM", func(t *testing.T) {
		affected, err := h.DeleteVM(vm.VMID)
		assert := assert.New(t)
		assert.Equal(int64(1), affected)
		assert.Nil(err)
	})
}
