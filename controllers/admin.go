package controllers

type AdminController struct {
	BaseController
}

// 后台首页
func (p *AdminController) Index() {
	p.MustLogin()
	p.AdminCommTpl("index/index.html", "管理后台")
}
