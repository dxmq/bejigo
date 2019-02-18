package controllers

import "bejigo/models"

type AdminController struct {
	BaseController
}

// 后台首页
func (p *AdminController) Index() {
	p.MustLogin()
	p.AdminCommTpl("index/index.html", "管理后台")
}

// 进入后台
func (p *AdminController) Admin() {
	if p.GetSession(SESSION_USER_KEY) != nil {
		p.Redirect("/admin/index", 301)
	}

	var st models.System
	models.System{}.GetSectionSystem(&st)
	p.Data["SystemData"] = st
	p.TplName = "admin/login/login.html"
}
