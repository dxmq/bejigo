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

	// 文章添加页面显示
	beego.Router("/admin/article/addindex", &controllers.ArticleController{}, "get:ArticleAddIndex")
	// 文章添加功能
	beego.Router("/admin/article/create", &controllers.ArticleController{}, "post:ArticleCreate")

	// editormd 图片上传
	beego.Router("/admin/article/uploads", &controllers.ArticleController{}, "post:UploadEditorMdPic")

	// 文章列表
	beego.Router("/admin/article/list/:page", &controllers.ArticleController{}, "get:ArticleList")

	// 文章编辑表单显示
	beego.Router("/admin/article/editindex/:id", &controllers.ArticleController{}, "get:ArticleEditIndex")

	// 文章编辑
	beego.Router("/admin/article/edit", &controllers.ArticleController{}, "post:ArticleEdit")
}
