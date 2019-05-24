// @APIVersion 1.0.0
// @Title ITP Server API
// @Description Inspur Testing Platform autogenerated API document.
// @Contact wangkun_lc@inspur.com
// @TermsOfServiceUrl http://git.inspur.com/wangkun_lc/itpserver
package routers

import (
	"github.com/inspursoft/itpserver/src/apiserver/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/auth",
			beego.NSInclude(
				&controllers.AuthController{},
			),
		),
		beego.NSNamespace("/vms",
			beego.NSInclude(
				&controllers.VMController{},
			),
		),
		beego.NSNamespace("/packages",
			beego.NSInclude(
				&controllers.PackagesController{},
			),
		),
		beego.NSNamespace("/installations",
			beego.NSInclude(
				&controllers.InstallationController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
