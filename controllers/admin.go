package controllers

type AdminController struct {
	BaseController
}

// @router /admin/index [get]
func (p *AdminController) Index() {
	p.MustLogin()
	p.AdminCommTpl("index/index.html", "管理后台")
}
