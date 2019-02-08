package routers

import "github.com/astaxie/beego"
import "bejigo/controllers"

func init() {
	beego.ErrorController(&controllers.ErrorController{})

	beego.Include(&controllers.IndexController{})

	beego.Include(&controllers.UserController{})

	beego.Include(&controllers.TagController{})

	// 后台首页
	beego.Router("/admin/index", &controllers.AdminController{}, "get:Index")

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

	// 编辑文章
	beego.Router("/admin/article/edit", &controllers.ArticleController{}, "post:ArticleEdit")

	// 文章删除
	beego.Router("/admin/article/delete", &controllers.ArticleController{}, "post:ArticleDelete")

	// 添加表单显示
	beego.Router("/admin/link/addindex", &controllers.LinkController{}, "get:LinkAddIndex")

	// 友链相关
	beego.Router("/admin/link/create", &controllers.LinkController{}, "post:LinkCreate")
	beego.Router("/admin/link/list", &controllers.LinkController{}, "get:LinkList")
	beego.Router("/admin/link/editindex/:id", &controllers.LinkController{}, "get:LinkEditIndex")
	beego.Router("/admin/link/edit", &controllers.LinkController{}, "post:LinkEdit")
	beego.Router("/admin/link/delete", &controllers.LinkController{}, "post:LinkDelete")

	// 单页面相关
	beego.Router("/admin/single/addindex", &controllers.SingleController{}, "get:AddPageIndex")
	beego.Router("/admin/single/lst", &controllers.SingleController{}, "get:PageList")
	beego.Router("/admin/single/add", &controllers.SingleController{}, "post:PageAdd")
	beego.Router("/admin/single/editindex/:id", &controllers.SingleController{}, "get:EditPageIndex")
	beego.Router("/admin/single/edit", &controllers.SingleController{}, "post:PageEdit")
	beego.Router("/admin/single/delete", &controllers.SingleController{}, "post:PageDelete")
	beego.Router("/admin/single/sort", &controllers.SingleController{}, "post:PageSort")

	// 系统设置
	beego.Router("/admin/system/index", &controllers.SystemController{}, "get:SystemIndex")
	beego.Router("/admin/system/save", &controllers.SystemController{}, "post:SystemSave")

	// 轻语
	beego.Router("/admin/whisper/addindex", &controllers.WhisperController{}, "get:AddIndex")
	beego.Router("/admin/whisper/create", &controllers.WhisperController{}, "post:Create")
	beego.Router("/admin/whisper/list/:id", &controllers.WhisperController{}, "get:List")
	beego.Router("/admin/whisper/editindex/:id", &controllers.WhisperController{}, "get:EditIndex")
	beego.Router("/admin/whisper/edit", &controllers.WhisperController{}, "post:Edit")
	beego.Router("/admin/whisper/delete", &controllers.WhisperController{}, "post:Delete")
}
