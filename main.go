package main

import (
	"bejigo/models"
	_ "bejigo/models"
	_ "bejigo/routers"
	"encoding/gob"
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

func main() {
	initSession()
	StrReplace()
	beego.Run()
}

func initSession() {
	gob.Register(models.User{})
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "bejigoblog"
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "data/session" // session文件目录
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 3600 * 24       // 过期秒数
}

func StrReplace() {
	err := beego.AddFuncMap("replace", func(str, old, new string, n int) string {
		return strings.Replace(str, old, new, -1)
	})
	if err != nil {
		fmt.Println(err)
	}
}
