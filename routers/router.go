package routers

import "github.com/astaxie/beego"
import "bego/controllers"

func init() {
	beego.ErrorController(&controllers.ErrorController{})

	beego.Include(&controllers.IndexController{})

	beego.Include(&controllers.UserController{})

	beego.Include(&controllers.AdminController{})

	beego.Include(&controllers.TagController{})

	// 分类列表
	beego.Router("/admin/category/list", &controllers.CategoryController{}, "get:CategoryList")
	// 分类添加表单显示
	beego.Router("/admin/category/addindex", &controllers.CategoryController{}, "get:CategoryAddIndex")
	// 分类添加
	beego.Router("/admin/category/create", &controllers.CategoryController{}, "post:CategoryCreate")
	// 分类页面显示
	beego.Router("/admin/category/editindex/:id", &controllers.CategoryController{}, "get:CategoryEditIndex")
	beego.Router("/admin/category/edit", &controllers.CategoryController{}, "post:CategoryEdit")
	beego.Router("/admin/category/delete", &controllers.CategoryController{}, "post:CategoryDelete")
}
