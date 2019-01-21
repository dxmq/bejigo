package main

import (
	"bego/models"
	_ "bego/models" // 引入Model
	_ "bego/routers"
	"encoding/gob"
	"github.com/astaxie/beego"
)

func main() {
	initSession()
	beego.Run()
}

func initSession() {
	gob.Register(models.User{})
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "begoblog"
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "data/session" // session文件目录
}
