package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/inspursoft/itpserver/src/apiserver/controllers:ArchiveController"] = append(beego.GlobalControllerRouter["github.com/inspursoft/itpserver/src/apiserver/controllers:ArchiveController"],
        beego.ControllerComments{
            Method: "Download",
            Router: `/download`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/inspursoft/itpserver/src/apiserver/controllers:ArchiveController"] = append(beego.GlobalControllerRouter["github.com/inspursoft/itpserver/src/apiserver/controllers:ArchiveController"],
        beego.ControllerComments{
            Method: "Upload",
            Router: `/upload`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

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
            Router: `/:vm_name`,
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

    beego.GlobalControllerRouter["github.com/inspursoft/itpserver/src/apiserver/controllers:OneStepController"] = append(beego.GlobalControllerRouter["github.com/inspursoft/itpserver/src/apiserver/controllers:OneStepController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/inspursoft/itpserver/src/apiserver/controllers:OneStepController"] = append(beego.GlobalControllerRouter["github.com/inspursoft/itpserver/src/apiserver/controllers:OneStepController"],
        beego.ControllerComments{
            Method: "PostWithVagrantfile",
            Router: `/:vm_name`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/inspursoft/itpserver/src/apiserver/controllers:PackagesController"] = append(beego.GlobalControllerRouter["github.com/inspursoft/itpserver/src/apiserver/controllers:PackagesController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:vm_name`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/inspursoft/itpserver/src/apiserver/controllers:PackagesController"] = append(beego.GlobalControllerRouter["github.com/inspursoft/itpserver/src/apiserver/controllers:PackagesController"],
        beego.ControllerComments{
            Method: "Upload",
            Router: `/:vm_name`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/inspursoft/itpserver/src/apiserver/controllers:PackagesController"] = append(beego.GlobalControllerRouter["github.com/inspursoft/itpserver/src/apiserver/controllers:PackagesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:vm_name`,
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
            Method: "CreateBySpec",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/inspursoft/itpserver/src/apiserver/controllers:VMController"] = append(beego.GlobalControllerRouter["github.com/inspursoft/itpserver/src/apiserver/controllers:VMController"],
        beego.ControllerComments{
            Method: "CreateByVagrantfile",
            Router: `/:vm_name`,
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

    beego.GlobalControllerRouter["github.com/inspursoft/itpserver/src/apiserver/controllers:VMController"] = append(beego.GlobalControllerRouter["github.com/inspursoft/itpserver/src/apiserver/controllers:VMController"],
        beego.ControllerComments{
            Method: "Package",
            Router: `/:vm_name/package`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
