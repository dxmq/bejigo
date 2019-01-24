package controllers

import (
	"bego/models"
	"github.com/astaxie/beego"
)

type IndexController struct {
	BaseController
}

// @router / [get]
func (c *IndexController) Index() {
	// 取出首页文章信息
	var r []models.Result
	models.Article{}.GetArticleData(5, &r) // 取出最新的5条数据
	c.Data["Article"] = r
	c.IndexCommTpl("Index", "index.html", "")
}

// @router /archives [get]
func (c *IndexController) Archives() {
	// 查询出归档信息
	var r []models.Result
	models.Article{}.GetArchives(&r)
	for _, v := range r {
		year := beego.Date(v.CreatedAt, "Y")
		beego.Info(year)
	}
	c.Data["Archives"] = r
	c.IndexCommTpl("Archives", "archives.html", "")
}

// @router /about [get]
func (c *IndexController) About() {
	c.IndexCommTpl("About", "about.html", "about_script.html")
}
