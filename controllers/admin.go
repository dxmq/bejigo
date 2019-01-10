package controllers

type AdminController struct {
	BaseController
}
// @router /admin/index [get]
func (p *AdminController) Index() {
	p.TplName = "admin/index/index.html"
}
