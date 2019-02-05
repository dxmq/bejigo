package controllers

import (
	"bego/models"
	"github.com/astaxie/beego"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type IndexController struct {
	BaseController
}

// @router / [get]
func (c *IndexController) Index() {
	// 取出首页文章信息
	var r []models.Result
	var isn models.IndexShowNumber
	models.System{}.GetIndexShowNumber(&isn)
	models.Article{}.GetArticleData(int(isn.IndexShowNumber), &r) // 取出最新的数据
	c.Data["Article"] = r
	c.IndexCommTpl("Index", "index.html", "", "", "")
}

// @router /archives [get]
func (c *IndexController) Archives() {
	// 查询出归档信息
	var ar []models.ArticleArch
	models.Article{}.GetArchives(&ar)
	c.Data["Archives"] = ar
	c.IndexCommTpl("Archives", "archives.html", "", "", "归档")
}

// 按年月归档
// @router /archives/:monthYear [get]
func (c *IndexController) ArchivesByMY() {
	// 接收monthYear
	monthYear := c.Ctx.Input.Param(":monthYear")
	// 取出年份
	myRune := []rune(monthYear)
	year := string(myRune[3:7])
	if len(year) == 0 {
		year = "2019"
	}
	// 取出月份
	month := beego.Substr(monthYear, 0, 2)

	// 查询出所有文章的添加时间
	var ct []models.CreatedAt
	models.Article{}.GetArticleCreateTime(&ct)
	var articleIds = ""
	for _, v := range ct {
		armonth := beego.Date(v.CreatedAt, "m")
		if armonth == month {
			Id := int(v.ID)
			articleIds += strconv.Itoa(Id) + ","
		}
	}
	// 根据article_id查询出文章信息
	articleIds = strings.Trim(articleIds, ",")
	var ar []models.ArticleArch
	models.Article{}.GetArticleByArticleId(articleIds, &ar, 10)
	c.Data["ArticleByMY"] = ar
	c.Data["ArchivesYear"] = year
	c.IndexCommTpl("Archives", "archivesbymonth.html", "", "", "归档")
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
	// 取出下一篇文章
	var a models.Article
	models.Article{}.GetNextArticle(r.ID, &a)
	c.Data["NextArticle"] = a

	// 阅读量
	ct := c.GetSession(id)
	if ct != nil {
		if int(time.Now().Unix())-ct.(int) >= 86400 {
			models.Article{}.UpdateViews(id)
		}
	} else {
		c.SetSession(id, int(time.Now().Unix()))
		models.Article{}.UpdateViews(id)
	}
	c.IndexCommTpl("Index", "article.html", "", "", r.Title)
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
	c.IndexCommTpl("Archives", "category.html", "", "", categoryName)
}

// 某个标签下的文章
// @router /tags/:id [get]
func (c *IndexController) Tags() {
	// 获取标签id
	id := c.Ctx.Input.Param(":id")
	// 根据标签id查询出标签名称
	var t models.Tag
	models.Tag{}.GetTagNameById(id, &t)
	c.Data["TagName"] = t.Name
	// 根据tag_id查询出文章id
	var at []models.ArticleTag
	models.ArticleTag{}.GetAllTagId(&at)
	var ar []models.ArticleArch
	var str = ""
	for _, v := range at {
		if strings.Index(v.TagId, id) != -1 {
			articleId := strconv.Itoa(int(v.ArticleId))
			// 根据articleId查询出文章信息
			str += articleId + ","
		}
	}
	articleIds := strings.Trim(str, ",")
	models.Article{}.GetArticleByArticleId(articleIds, &ar, 10)
	c.Data["ArticleByTag"] = ar
	//models.Article{}.GetArticleByTagId()
	c.IndexCommTpl("Archives", "tag.html", "", "", "标签")
}

//某个页面
// @router /page/:pageAlias [get]
func (c *IndexController) Page() {
	// pageName
	pageAlias := c.Ctx.Input.Param(":pageAlias")
	// 根据分类名称查询出当前分类下的文章
	var s models.SinglePage
	models.SinglePage{}.GetPageByCateName(pageAlias, &s)

	// 查询出所有的pageAlias
	var sg []models.SinglePage
	models.SinglePage{}.GetAllPageAlias(&sg)
	for _, v := range sg {
		if v.PageAlias == pageAlias {
			// 创建文件
			pagePath := "./views/index/page/"
			fileName := pagePath + pageAlias + ".html"
			f, _ := os.Create(fileName)
			io.WriteString(f, s.Content)
			defer f.Close()
		}
	}
	c.IndexCommTpl("Page", pageAlias+".html", "", pageAlias, s.PageName)
}
