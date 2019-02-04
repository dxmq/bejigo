package controllers

import (
	"bego/models"
	"bego/syserrors"
	"encoding/json"
	"errors"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
	"github.com/astaxie/beego/validation"
	"unicode/utf8"
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
	var st models.System
	models.System{}.GetSectionSystem(&st)
	p.Data["SystemData"] = st
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
		p.About500(syserrors.New("用户名或密码错误！", err))
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

// 用户列表
// @router /admin/user/list [get]
func (p *UserController) UserList() {
	p.MustLogin()
	// 查询出所有的用户
	var u []models.User
	models.User{}.GetAllUser(&u)
	p.Data["Users"] = u
	p.AdminCommTpl("user/list.html", "用户列表")
}

// 用户添加页面显示
// @router /admin/user/addindex [get]
func (p *UserController) UserAddIndex() {
	p.MustLogin()
	p.AdminCommTpl("user/add.html", "用户添加")
}

// 用户添加
// @router /admin/user/create [post]
func (p *UserController) UserCreate() {
	p.MustLogin()
	// 接收表单信息
	username := p.GetMustAndInlen("user_name", "用户名不能为空！", "用户名不能超过15个字符！", 15)
	password := p.GetMustString("pass_word", "密码不能为空！")
	if passLen := utf8.RuneCountInString(password); passLen < 6 || passLen > 16 {
		p.About500(errors.New("密码长度6-16位之间！"))
	}
	password = p.GetMd5String(password) // md5 加密
	email := p.GetMustString("email", "邮箱不能为空！")
	valid := validation.Validation{}
	var u models.User
	avatar := p.GetString("avatar", "/static/index/img/face.png")
	role, _ := p.GetInt("role")

	// 赋值给模型
	u.UserName = username
	u.PassWord = password
	u.Email = email
	b, err := valid.Valid(&u)
	if !b {
		p.About500(syserrors.New("邮箱格式不符！", err))
	}
	u.Avatar = avatar
	u.Role = role

	// 调用模型完成添加操作
	err = models.User{}.UserCreate(&u)
	if err == nil {
		p.ReturnJson("添加用户成功！", "/admin/user/list")
	} else {
		p.About500(syserrors.New("用户添加失败！", err))
	}
}

// 用户编辑页面
// @router /admin/user/editindex/:id [get]
func (p *UserController) Profile() {
	p.MustLogin()
	// 接收user_id
	id := p.Ctx.Input.Param(":id")
	// 根据id查询出当前用户的信息
	var user models.User
	models.User{}.GetUserInfoById(id, &user)
	// assign 到页面
	p.Data["User"] = user
	p.AdminCommTpl("user/edit.html", "修改用户资料")
}

// 用户资料修改
// @router /admin/user/edit [post]
func (p *UserController) EditProfile() {
	p.MustLogin()
	// 接收表单信息
	userId, _ := p.GetInt("id")
	username := p.GetMustAndInlen("user_name", "用户名不能为空！", "用户名不能超过15个字符！", 15)
	password := p.GetMustString("pass_word", "密码不能为空！")
	passLen := utf8.RuneCountInString(password)
	if passLen != 32 {
		if passLen < 6 || passLen > 16 {
			p.About500(errors.New("密码长度6-16位之间！"))
		}
		password = p.GetMd5String(password) // md5 加密
	}
	email := p.GetMustString("email", "邮箱不能为空！")
	valid := validation.Validation{}
	var u models.User
	avatar := p.GetString("avatar", "/static/index/img/face.png")
	role, _ := p.GetInt("role")

	// 赋值给模型
	u.ID = uint(userId)
	u.UserName = username
	u.PassWord = password
	u.Email = email
	b, err := valid.Valid(&u)
	if !b {
		p.About500(syserrors.New("邮箱格式不符！", err))
	}
	u.Avatar = avatar
	u.Role = role

	// 调用模型完成添加操作
	err = models.User{}.UserEdit(&u)
	if err == nil {
		p.ReturnJson("修改用户资料成功！", "/admin/user/list")
	} else {
		p.About500(syserrors.New("修改失败！", err))
	}
}

// 用户删除
// @router /admin/user/delete [post]
func (p *UserController) UserDelete() {
	p.MustLogin()
	var uDelId models.UserDeleteId
	if err := json.Unmarshal(p.Ctx.Input.RequestBody, &uDelId); err == nil {
		// 调用模型完成删除操作
		err = models.User{}.UserDelete(uDelId.Id)
		if err == nil {
			p.ReturnJsonCode("删除用户成功！")
		} else {
			p.About500(syserrors.New("删除用户失败！", err))
		}
	}
}
