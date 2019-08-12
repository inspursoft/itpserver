package dao_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/inspursoft/itpserver/src/apiserver/dao"
	"github.com/inspursoft/itpserver/src/apiserver/models"
)

func TestOperateVM(t *testing.T) {
	vm := models.VM{Name: "ubuntu-vm-1", OS: "CentOS"}
	spec := models.VMSpec{CPUs: 2, Memory: "4Gb", Storage: "1T", Extras: "SSD"}
	var h dao.VMDaoHandler
	var v *models.VM
	var err error
	t.Run("CreateVM", func(t *testing.T) {
		v, err = h.AddVM(&vm, &spec)
		assert := assert.New(t)
		assert.NotNil(v, "Got VM object after insertion.")
		assert.Nil(err)
	})
	t.Run("GetVM", func(t *testing.T) {
		query := models.VM{IP: v.IP}
		vm, err := h.GetVM(query, "IP")
		assert := assert.New(t)
		assert.Nil(err)
		assert.NotNil(vm)
	})
	t.Run("GetVMList", func(t *testing.T) {
		query := models.VMWithSpec{IP: v.IP}
		vms, err := h.GetVMList(query)
		assert := assert.New(t)
		assert.Nil(err)
		assert.NotNil(vms)
		assert.NotNil(spec)
	})
	t.Run("UpdateVM", func(t *testing.T) {
		updates := make(map[string]interface{})
		updates["OS"] = "CentOS 7.5"
		affected, err := h.UpdateVM(v.ID, updates)
		assert := assert.New(t)
		assert.Equal(int64(1), affected)
		assert.Nil(err)
	})
	t.Run("UpdateVMSpec", func(t *testing.T) {
		updates := make(map[string]interface{})
		updates["VID"] = "2d4e7f"
		affected, err := h.UpdateVMSpec(v.ID, updates)
		assert := assert.New(t)
		assert.Equal(int64(1), affected)
		assert.Nil(err)
	})
	t.Run("DeleteVMWithVID", func(t *testing.T) {
		affected, err := h.DeleteVMByVID("2d4e7f")
		assert := assert.New(t)
		assert.Equal(int64(1), affected)
		assert.Nil(err)
	})
	// t.Run("DeleteVM", func(t *testing.T) {
	// 	query := models.VM{ID: v.ID}
	// 	affected, err := h.DeleteVM(query, "ID")
	// 	assert := assert.New(t)
	// 	assert.Equal(int64(1), affected)
	// 	assert.Nil(err)
	// })
}
