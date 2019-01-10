package routers

import "github.com/astaxie/beego"
import "bego/controllers"

func init() {
   beego.ErrorController(&controllers.ErrorController{})

   beego.Include(&controllers.IndexController{})

   beego.Include(&controllers.UserController{})

   beego.Include(&controllers.AdminController{})
}
