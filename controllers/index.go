package controllers

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	BaseController
}

// @router / [get]
func (c *IndexController) Index() {
	c.commTpl("Index", "index.html", "")
}

// @router /archives [get]
func (c *IndexController) Archives() {
	c.commTpl("Archives", "archives.html", "")
}

// @router /about [get]
func (c *IndexController) About()  {
	c.commTpl("About", "index.html", "about_script.tpl")
}

func (c *IndexController) commTpl(method string, tpl string, sectionTpl string) {
	c.Data["url"] = beego.URLFor("IndexController."+method)
	beego.Info(c.Data["url"])
	c.Layout = "index/public/layout.html"
	c.TplName = "index/"+tpl
	if len(sectionTpl) != 0 {
		c.LayoutSections = make(map[string]string)
		c.LayoutSections["Scripts"] = "index/public/"+sectionTpl
	}
}