package controllers

import (
	"bego/models"
	"bego/syserrors"
	"errors"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
)

type UserController struct {
	BaseController
}

var cpt *captcha.Captcha

func init() {
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha", store)
	cpt.ChallengeNums = 4
	cpt.StdWidth = 100
	cpt.StdHeight = 40
}

// 显示登录页面
// @router /admin/login/index [get]
func (p *UserController) LoginIndex() {
	p.TplName = "admin/login/login.html"
}

// 登录
// @router /admin/login/login [post]
func (p *UserController) Login() {
	// username
	username := p.GetMustString("user_name", " 用户名不能为空！")
	// password
	password := p.GetMustString("pass_word", "密码不能为空！")
	password = models.GetMd5String(password) // md5加密
	// captcha
	p.GetMustString("captcha", "验证码不能为空！")
	if cpt.VerifyReq(p.Ctx.Request) == false {
		p.About500(errors.New("验证码错误！"))
	}
	user, err := models.QueryUser(username, password)
	if err != nil {
		p.About500(syserrors.New("登录失败！", err))
	}
	// 登录成功
	p.SetSession(SESSION_USER_KEY, user) // 存 session
	p.ReturnJson("登录成功", "/admin/index") // 提示信息并跳转
}

// 登出
// @router admin/login/logout [post]
func (p *UserController) LoginOut() {
	p.MustLogin()
	// 清空session
	p.DelSession(SESSION_USER_KEY)
	p.ReturnJson("退出成功", "/admin/login/index")
}

// 用户资料
// @router /admin/index/profile/:id [get]
func (p *AdminController) Profile() {
	p.MustLogin()
	id := p.Ctx.Input.Param(":id")
	// 根据id查询出当前用户的信息
	var user models.User
	models.User{}.GetUserInfoById(id, &user)
	// assign 到页面
	p.Data["User"] = user
	p.AdminCommTpl("profile/profile.html", "个人资料")
}
