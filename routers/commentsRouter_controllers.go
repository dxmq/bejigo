package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["bejigo/controllers:IndexController"] = append(beego.GlobalControllerRouter["bejigo/controllers:IndexController"],
		beego.ControllerComments{
			Method:           "Index",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bejigo/controllers:IndexController"] = append(beego.GlobalControllerRouter["bejigo/controllers:IndexController"],
		beego.ControllerComments{
			Method:           "Page",
			Router:           `/:pageAlias`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bejigo/controllers:IndexController"] = append(beego.GlobalControllerRouter["bejigo/controllers:IndexController"],
		beego.ControllerComments{
			Method:           "Archives",
			Router:           `/archives`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bejigo/controllers:IndexController"] = append(beego.GlobalControllerRouter["bejigo/controllers:IndexController"],
		beego.ControllerComments{
			Method:           "ArchivesByMY",
			Router:           `/archives/:monthYear`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bejigo/controllers:IndexController"] = append(beego.GlobalControllerRouter["bejigo/controllers:IndexController"],
		beego.ControllerComments{
			Method:           "Categories",
			Router:           `/categories/:cateName`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bejigo/controllers:IndexController"] = append(beego.GlobalControllerRouter["bejigo/controllers:IndexController"],
		beego.ControllerComments{
			Method:           "ArticleDetail",
			Router:           `/detail/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bejigo/controllers:IndexController"] = append(beego.GlobalControllerRouter["bejigo/controllers:IndexController"],
		beego.ControllerComments{
			Method:           "IndexPage",
			Router:           `/page/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bejigo/controllers:IndexController"] = append(beego.GlobalControllerRouter["bejigo/controllers:IndexController"],
		beego.ControllerComments{
			Method:           "Tags",
			Router:           `/tags/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bejigo/controllers:IndexController"] = append(beego.GlobalControllerRouter["bejigo/controllers:IndexController"],
		beego.ControllerComments{
			Method:           "WhisperDetail",
			Router:           `/whisper/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bejigo/controllers:IndexController"] = append(beego.GlobalControllerRouter["bejigo/controllers:IndexController"],
		beego.ControllerComments{
			Method:           "Whisper",
			Router:           `/whisper/page/:pageId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bejigo/controllers:TagController"] = append(beego.GlobalControllerRouter["bejigo/controllers:TagController"],
		beego.ControllerComments{
			Method:           "TagStore",
			Router:           `/admin/tag`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bejigo/controllers:TagController"] = append(beego.GlobalControllerRouter["bejigo/controllers:TagController"],
		beego.ControllerComments{
			Method:           "TagList",
			Router:           `/admin/tag`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bejigo/controllers:TagController"] = append(beego.GlobalControllerRouter["bejigo/controllers:TagController"],
		beego.ControllerComments{
			Method:           "TagShow",
			Router:           `/admin/tag/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bejigo/controllers:TagController"] = append(beego.GlobalControllerRouter["bejigo/controllers:TagController"],
		beego.ControllerComments{
			Method:           "TagCreate",
			Router:           `/admin/tag/create`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bejigo/controllers:TagController"] = append(beego.GlobalControllerRouter["bejigo/controllers:TagController"],
		beego.ControllerComments{
			Method:           "TagDelete",
			Router:           `/admin/tag/delete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bejigo/controllers:TagController"] = append(beego.GlobalControllerRouter["bejigo/controllers:TagController"],
		beego.ControllerComments{
			Method:           "TagEdit",
			Router:           `/admin/tag/edit`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bejigo/controllers:UserController"] = append(beego.GlobalControllerRouter["bejigo/controllers:UserController"],
		beego.ControllerComments{
			Method:           "LoginIndex",
			Router:           `/admin/login/index`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bejigo/controllers:UserController"] = append(beego.GlobalControllerRouter["bejigo/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Login",
			Router:           `/admin/login/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bejigo/controllers:UserController"] = append(beego.GlobalControllerRouter["bejigo/controllers:UserController"],
		beego.ControllerComments{
			Method:           "UserAddIndex",
			Router:           `/admin/user/addindex`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bejigo/controllers:UserController"] = append(beego.GlobalControllerRouter["bejigo/controllers:UserController"],
		beego.ControllerComments{
			Method:           "UserCreate",
			Router:           `/admin/user/create`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bejigo/controllers:UserController"] = append(beego.GlobalControllerRouter["bejigo/controllers:UserController"],
		beego.ControllerComments{
			Method:           "UserDelete",
			Router:           `/admin/user/delete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bejigo/controllers:UserController"] = append(beego.GlobalControllerRouter["bejigo/controllers:UserController"],
		beego.ControllerComments{
			Method:           "EditProfile",
			Router:           `/admin/user/edit`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bejigo/controllers:UserController"] = append(beego.GlobalControllerRouter["bejigo/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Profile",
			Router:           `/admin/user/editindex/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bejigo/controllers:UserController"] = append(beego.GlobalControllerRouter["bejigo/controllers:UserController"],
		beego.ControllerComments{
			Method:           "UserList",
			Router:           `/admin/user/list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bejigo/controllers:UserController"] = append(beego.GlobalControllerRouter["bejigo/controllers:UserController"],
		beego.ControllerComments{
			Method:           "LoginOut",
			Router:           `admin/login/logout`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
