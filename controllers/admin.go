package controllers

type AdminController struct {
	BaseController
}

// @router /admin/index [get]
func (p *AdminController) Index() {
	p.commTpl("index/index.html")
}

func (p *AdminController) commTpl(tpl string) {
	p.Layout = "admin/public/layout.html"
	p.TplName = "admin/" + tpl
}
