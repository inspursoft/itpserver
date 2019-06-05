package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/inspursoft/itpserver/src/apiserver/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/inspursoft/itpserver/src/apiserver/controllers:AuthController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/inspursoft/itpserver/src/apiserver/controllers:BaseController"] = append(beego.GlobalControllerRouter["github.com/inspursoft/itpserver/src/apiserver/controllers:BaseController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/inspursoft/itpserver/src/apiserver/controllers:InstallationController"] = append(beego.GlobalControllerRouter["github.com/inspursoft/itpserver/src/apiserver/controllers:InstallationController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/inspursoft/itpserver/src/apiserver/controllers:InstallationController"] = append(beego.GlobalControllerRouter["github.com/inspursoft/itpserver/src/apiserver/controllers:InstallationController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/:vm_name`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/inspursoft/itpserver/src/apiserver/controllers:InstallationController"] = append(beego.GlobalControllerRouter["github.com/inspursoft/itpserver/src/apiserver/controllers:InstallationController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:vm_name`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/inspursoft/itpserver/src/apiserver/controllers:PackagesController"] = append(beego.GlobalControllerRouter["github.com/inspursoft/itpserver/src/apiserver/controllers:PackagesController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/inspursoft/itpserver/src/apiserver/controllers:PackagesController"] = append(beego.GlobalControllerRouter["github.com/inspursoft/itpserver/src/apiserver/controllers:PackagesController"],
        beego.ControllerComments{
            Method: "Upload",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/inspursoft/itpserver/src/apiserver/controllers:PackagesController"] = append(beego.GlobalControllerRouter["github.com/inspursoft/itpserver/src/apiserver/controllers:PackagesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/inspursoft/itpserver/src/apiserver/controllers:VMController"] = append(beego.GlobalControllerRouter["github.com/inspursoft/itpserver/src/apiserver/controllers:VMController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/inspursoft/itpserver/src/apiserver/controllers:VMController"] = append(beego.GlobalControllerRouter["github.com/inspursoft/itpserver/src/apiserver/controllers:VMController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/inspursoft/itpserver/src/apiserver/controllers:VMController"] = append(beego.GlobalControllerRouter["github.com/inspursoft/itpserver/src/apiserver/controllers:VMController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:vm_name`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
