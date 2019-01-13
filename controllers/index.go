package controllers

type IndexController struct {
	BaseController
}

// @router / [get]
func (c *IndexController) Index() {
	c.IndexCommTpl("Index", "index.html", "")
}

// @router /archives [get]
func (c *IndexController) Archives() {
	c.IndexCommTpl("Archives", "archives.html", "")
}

// @router /about [get]
func (c *IndexController) About() {
	c.IndexCommTpl("About", "index.html", "about_script.tpl")
}
