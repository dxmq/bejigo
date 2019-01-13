package routers

import "github.com/astaxie/beego"
import "bego/controllers"

func init() {
	beego.ErrorController(&controllers.ErrorController{})

	beego.Include(&controllers.IndexController{})

	beego.Include(&controllers.UserController{})

	beego.Include(&controllers.AdminController{})

	beego.Include(&controllers.TagController{})
	/*beego.NSNamespace("/admin",
	   //CRUD Create(创建)、Read(读取)、Update(更新)和Delete(删除)
	   beego.NSNamespace("/tag",
	      beego.NSInclude(&controllers.TagController{}),
	   ),
	)*/
}
