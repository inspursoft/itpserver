package dao_test

import (
	"testing"

	"github.com/inspursoft/itpserver/src/apiserver/dao"
	"github.com/inspursoft/itpserver/src/apiserver/models"
	"github.com/stretchr/testify/assert"
)

func TestOperateInstallation(t *testing.T) {
	spec := models.VMSpec{CPUs: 2, Memory: "2Gb", Storage: "512Gb"}
	vm := models.VM{Name: "test-vm-1", VMID: "test-1a2b", OS: "Linux 3.7"}
	pkg := models.Package{Name: "Ruby", Tag: "2.3.1"}

	var insHandler dao.InstallationDaoHandler
	t.Run("Install", func(t *testing.T) {
		var vmDaoHandler dao.VMDaoHandler
		vmDaoHandler.AddVM(&vm, &spec)
		var pkgDaoHandler dao.PkgDaoHandler
		pkgDaoHandler.AddPackage(&pkg)
		affected, err := insHandler.InstallPackageToVM(&vm, &pkg)
		assert := assert.New(t)
		assert.Nil(err)
		assert.Equal(int64(1), affected)
	})
	t.Run("GetInstalledPkgs", func(t *testing.T) {
		pkgList, err := insHandler.GetInstallPackages("test-1a2b")
		assert := assert.New(t)
		assert.Nil(err)
		assert.Len(pkgList, 1)
	})
	t.Run("Remove", func(t *testing.T) {
		affected, err := insHandler.RemovePackageFromVM(&vm, &pkg)
		assert := assert.New(t)
		assert.Nil(err)
		assert.Equal(int64(1), affected)
	})
}
