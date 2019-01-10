package controllers

import (
	"bego/models"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"regexp"
)

const SESSION_USER_KEY  = "SESSION_USER_KEY" // 定义一个常量，用于保存用户登录session

type BaseController struct {
	beego.Controller
	o orm.Ormer
	controllerName string
	actionName     string
	User models.User
	IsLogin bool // 用户是否登录
}

func (p *BaseController) Prepare() {
	u, ok := p.GetSession("SESSION_USER_KEY").(models.User) // 类型判断
	p.IsLogin = false // 把登录状态设置为false
	if ok {
		p.User = u
		p.IsLogin = true
		p.Data["user"] = p.User // 页面渲染
	}
}
func (p *BaseController)About500(err error) {
	p.Data["error"] = err
	p.Abort("500")
}

// 字符串不能为空判断
func (p *BaseController) GetMustString(key, msg string) string {
	key = p.GetString(key)
	if len(key) == 0 {
		p.About500(errors.New(msg))
	}
	return key
}

//用来做跳转的逻辑展示
func (p *BaseController) History(msg string, url string) {
	if url == ""{
		p.Ctx.WriteString("<script>alert('"+msg+"');window.history.go(-1);</script>")
		p.StopRun()
	}else{
		p.Redirect(url,302)
	}
}

//字串截取
func (p *BaseController) SubString(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}
// 获取文件后缀
func (p *BaseController) GetFileSuffix(s string) string {
	re, _ := regexp.Compile(".(jpg|jpeg|png|gif)")
	suffix := re.ReplaceAllString(s, "")
	return suffix
}
// md5 获取
func (p *BaseController) GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
