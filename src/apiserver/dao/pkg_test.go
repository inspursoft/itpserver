package dao_test

import (
	"testing"

	"github.com/inspursoft/itpserver/src/apiserver/dao"
	"github.com/inspursoft/itpserver/src/apiserver/models"
	"github.com/stretchr/testify/assert"
)

func TestOperatePkg(t *testing.T) {
	pkg0 := models.Package{Name: "Docker", Tag: "17.06"}
	pkg1 := models.Package{Name: "Docker", Tag: "18.09.1"}
	var pkgDaoHandler dao.PkgDaoHandler
	t.Run("CreatePkg", func(t *testing.T) {
		affected, err := pkgDaoHandler.AddPackage(&pkg0)
		assert := assert.New(t)
		assert.NotEqual(int64(0), affected)
		assert.Nil(err)
	})
	t.Run("UpdatePkg", func(t *testing.T) {
		affected, err := pkgDaoHandler.UpdatePackage(pkg1)
		assert := assert.New(t)
		assert.Equal(int64(1), affected)
		assert.Nil(err)
	})
	t.Run("GetPkg", func(t *testing.T) {
		pkgList, err := pkgDaoHandler.GetPackage("Docker")
		assert := assert.New(t)
		assert.Nil(err)
		assert.Len(pkgList, 1)
		assert.Equal(pkgList[0].Tag, "18.09.1")
	})
	t.Run("DeletePkg", func(t *testing.T) {
		affected, err := pkgDaoHandler.DeletePackage("Docker", "18.09.1")
		assert := assert.New(t)
		assert.Equal(int64(1), affected)
		assert.Nil(err)
	})
}
