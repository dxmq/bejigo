package controllers

import (
	"bego/models"
	"io"
	"os"
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
	c.IndexCommTpl("Index", "index.html", "", "")
}

// @router /archives [get]
func (c *IndexController) Archives() {
	// 查询出归档信息
	var r []models.Result
	models.Article{}.GetArchives(&r)
	c.Data["Archives"] = r
	c.IndexCommTpl("Archives", "archives.html", "", "")
}

// 文章详情
// @router /detail/:id [get]
func (c *IndexController) ArticleDetail() {
	// 接收id
	id := c.Ctx.Input.Param(":id")
	// 根据id查询出一条文章数据
	var r models.Result
	models.Article{}.GetArticleInfoById(id, &r)
	var t []models.Tag
	models.Tag{}.GetTags(&t, r.TagId)
	c.Data["Detail"] = r
	c.Data["Tag"] = t
	c.IndexCommTpl("Index", "article.html", "", "")
}

//某个分类下的文章
// @router /categories/:cateName [get]
func (c *IndexController) Categories() {
	// categoryName
	categoryName := c.Ctx.Input.Param(":cateName")
	// 根据分类名称查询出当前分类下的文章
	var r []models.Result
	models.Article{}.GetArticleByCateName(categoryName, &r, 10)
	c.Data["ArticleByCate"] = r
	c.Data["CategoryName"] = categoryName
	c.IndexCommTpl("Archives", "category.html", "", "")
}

// 某个标签下的文章
// @router /tags/:tagName [get]
func (c *IndexController) Tags() {
	// 获取tagName
	/*tagName := c.Ctx.Input.Param(":tagName")
	// 根据tagName查询出文章
	var r []models.Result*/

	c.IndexCommTpl("Archives", "tag.html", "", "")
}

var f *os.File
var err1 error

//某个页面
// @router /:pageAlias [get]
func (c *IndexController) Page() {
	// pageName
	pageAlias := c.Ctx.Input.Param(":pageAlias")
	// 根据分类名称查询出当前分类下的文章
	var s models.SinglePage
	models.SinglePage{}.GetPageByCateName(pageAlias, &s)
	// 创建目录和文件
	pagePath := "./views/index/"
	fileName := pagePath + pageAlias + ".html"
	f, err1 = os.Create(fileName)
	io.WriteString(f, s.Content)
	defer f.Close()
	c.IndexCommTpl("Page", pageAlias+".html", "", pageAlias)
}
