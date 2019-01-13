package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["bego/controllers:AdminController"] = append(beego.GlobalControllerRouter["bego/controllers:AdminController"],
		beego.ControllerComments{
			Method:           "Index",
			Router:           `/admin/index`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bego/controllers:IndexController"] = append(beego.GlobalControllerRouter["bego/controllers:IndexController"],
		beego.ControllerComments{
			Method:           "Index",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bego/controllers:IndexController"] = append(beego.GlobalControllerRouter["bego/controllers:IndexController"],
		beego.ControllerComments{
			Method:           "About",
			Router:           `/about`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bego/controllers:IndexController"] = append(beego.GlobalControllerRouter["bego/controllers:IndexController"],
		beego.ControllerComments{
			Method:           "Archives",
			Router:           `/archives`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bego/controllers:TagController"] = append(beego.GlobalControllerRouter["bego/controllers:TagController"],
		beego.ControllerComments{
			Method:           "TagStore",
			Router:           `/admin/tag`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bego/controllers:TagController"] = append(beego.GlobalControllerRouter["bego/controllers:TagController"],
		beego.ControllerComments{
			Method:           "TagList",
			Router:           `/admin/tag`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bego/controllers:TagController"] = append(beego.GlobalControllerRouter["bego/controllers:TagController"],
		beego.ControllerComments{
			Method:           "TagShow",
			Router:           `/admin/tag/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bego/controllers:TagController"] = append(beego.GlobalControllerRouter["bego/controllers:TagController"],
		beego.ControllerComments{
			Method:           "TagCreate",
			Router:           `/admin/tag/create`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bego/controllers:TagController"] = append(beego.GlobalControllerRouter["bego/controllers:TagController"],
		beego.ControllerComments{
			Method:           "TagDelete",
			Router:           `/admin/tag/delete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bego/controllers:TagController"] = append(beego.GlobalControllerRouter["bego/controllers:TagController"],
		beego.ControllerComments{
			Method:           "TagEdit",
			Router:           `/admin/tag/edit`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bego/controllers:UserController"] = append(beego.GlobalControllerRouter["bego/controllers:UserController"],
		beego.ControllerComments{
			Method:           "LoginIndex",
			Router:           `/admin/login/index`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bego/controllers:UserController"] = append(beego.GlobalControllerRouter["bego/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Login",
			Router:           `/admin/login/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bego/controllers:UserController"] = append(beego.GlobalControllerRouter["bego/controllers:UserController"],
		beego.ControllerComments{
			Method:           "LoginOut",
			Router:           `admin/login/logout`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
