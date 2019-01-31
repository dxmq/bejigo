package controllers

import (
	"bego/models"
	"bego/syserrors"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"os"
	"strconv"
	"strings"
	"time"
)

type ArticleController struct {
	BaseController
}

/**
 * 文章列表
 */
func (p *ArticleController) ArticleList() {
	p.MustLogin()
	var r []models.Result
	var a = models.Article{}
	page := p.Ctx.Input.Param(":page")
	pg, _ := strconv.Atoi(page)
	count := a.GetAllCount()
	a.GetAllArticleByPage(&r, 2, pg)
	Page := a.PageUtil(count, pg, 2, r)
	// assign到页面
	p.Data["Page"] = Page
	// 显示文章列表页
	p.AdminCommTpl("article/list.html", "文章列表")
}

/**
 * 显示文章添加表单
 */
func (p *ArticleController) ArticleAddIndex() {
	p.MustLogin()
	p.Data["Category"] = p.GetAllCategory() // 取出所有的分类并assign到页面
	p.Data["Tag"] = p.GetAllTag()           // 获取所有的标签并assign到页面
	p.AdminCommTpl("article/add.html", "添加文章")

}

/**
 * 文章添加
 */
func (p *ArticleController) ArticleCreate() {
	p.MustLogin()
	// 验证并接收表单数据
	title := p.GetMustAndInlen("title", "标题不能为空！", "标题在1-40个字符之间！", 40)
	author := p.GetMustAndInlen("author", "作者不能为空！", "作者在1-20个字符之间！", 20)

	// 获取分类
	categoryId := p.GetMustString("category_id", "分类不能为空！")
	// categoryId, _ := p.GetInt("category_id")
	// 获取标签id数组
	tagId := p.GetStrings("tag_id[]")
	// 获取是否展示
	isShow, _ := p.GetInt("is_show")
	// 获取是否是否置顶
	isTop, _ := p.GetInt("is_top")
	// markdown 内容
	htmlContent := p.GetString("editormd-html-code")
	htmlContent = beego.Htmlquote(htmlContent) // 转义内容
	// content 内容
	content := p.GetMustString("content", "内容不能为空")
	// summary 内容
	summary := beego.Substr(content, 0, 200)

	// 赋值给模型
	var a models.Article
	a.Title = title
	a.Author = author
	a.CategoryId, _ = strconv.Atoi(categoryId)
	a.IsShow = isShow
	a.IsTop = isTop
	a.Summary = summary
	a.Content = content
	a.HtmlContent = htmlContent

	// 调用模型完成添加操作
	err := models.Article{}.ArticleAdd(&a)
	if err != nil {
		p.About500(syserrors.New("添加失败！", err))
	} else {
		if len(tagId) != 0 { // 如果有选择标签才维护
			/** 维护article_tag表 **/
			// 获取刚添加的文章id
			var article models.Article
			models.Article{}.GetIdByTitle(&article, a.Title)
			// 将tagId数组转化为用','分隔的字符串
			tagIds := strings.Replace(strings.Trim(fmt.Sprint(tagId), "[]"), " ", ",", -1)
			// 添加到article_tag表
			var at models.ArticleTag
			at.TagId = tagIds
			at.ArticleId = a.ID
			models.ArticleTag{}.ArticleTagAdd(&at)
			// 提示并跳转
			p.ReturnJson("添加成功！", "/admin/article/list/1")
		} else {
			p.ReturnJson("添加成功！", "/admin/article/list/1")
		}
	}
}

// editorMd 图片上传
func (p *ArticleController) UploadEditorMdPic() {
	p.MustLogin()
	// 获取文件信息
	file, information, err := p.GetFile("editormd-image-file")
	if err != nil {
		beego.Info(err)
	}
	file.Close()
	// 文件名
	fileName := information.Filename
	// 日期字符串
	dateStr := beego.Date(time.Now(), "Ymd")
	// 创建文件夹
	uploadPath := beego.AppConfig.String("upload_article_pic_path")
	filePath := uploadPath + dateStr
	if err1 := os.MkdirAll(filePath, 0777); err1 != nil { // 创建上传目录
		panic("failed" + err1.Error())
	}
	// folderPath := p.CreateDateDir("/static/uploads/")
	// 移动文件到创建好的文件夹内
	err2 := p.SaveToFile("editormd-image-file", filePath+"/"+fileName)
	if err2 == nil { // 如果没错，返回url,success,message
		p.Data["json"] = map[string]interface{}{
			"url":     "/static/uploads/" + dateStr + "/" + fileName,
			"success": 1,
			"message": "upload success!",
		}
		p.ServeJSON()
	}
}

/**
 * 显示编辑文章的页面
 */
func (p *ArticleController) ArticleEditIndex() {
	p.MustLogin()
	// 查询出所有的分类并分配到模板
	category := p.GetAllCategory()
	p.Data["Category"] = category

	// 查询出所有的标签并分配到模板
	tags := p.GetAllTag()
	p.Data["Tag"] = tags
	// 获取当前的文章的id
	id := p.Ctx.Input.Param(":id")
	// 根据id查询文章表
	var r models.Result
	models.Article{}.GetArticleById(id, &r)
	r.Content = beego.Htmlunquote(r.Content)
	/*r.Content = beego.HTML2str(r.Content)*/
	// 分配到模板
	p.Data["Article"] = r

	// 根据tag_ids 查询出当前的标签信息
	//beego.Info(r.TagId)
	var tag []models.Tag
	models.Tag{}.GetTags(&tag, r.TagId)
	// 将当前文章的标签信息分配到模板
	p.Data["TagCurrent"] = tag

	// 将当前的文章的categoryId转化为uint类型
	categoryId := uint(r.CategoryId)
	// 分配到模板
	p.Data["CategoryId"] = categoryId

	// 显示页面
	p.AdminCommTpl("article/edit.html", "编辑文章")
}

/**
 * 文章编辑
 */
func (p *ArticleController) ArticleEdit() {
	p.MustLogin()
	// 接收id
	id, _ := p.GetInt("id")
	// 接收其它表单数据
	title := p.GetMustAndInlen("title", "标题不能为空！", "标题在1-40个字符之间！", 40)
	author := p.GetMustAndInlen("author", "作者不能为空！", "作者在1-20个字符之间！", 20)
	categoryId := p.GetMustString("category_id", "分类不能为空！")
	cateId, _ := strconv.Atoi(categoryId) // 字符串转整型
	tagId := p.GetStrings("tag_id[]")
	isShow, _ := p.GetInt("is_show")
	isTop, _ := p.GetInt("is_top")
	htmlContent := p.GetString("editormd-html-code")
	htmlContent = beego.Htmlquote(htmlContent) // html转义
	content := p.GetMustString("content", "内容不能为空！")
	summary := beego.Substr(content, 0, 200)

	// 赋值给结构体
	var a models.Article
	a.ID = uint(id)
	a.Title = title
	a.Author = author
	a.CategoryId = cateId
	a.IsShow = isShow
	a.IsTop = isTop
	a.HtmlContent = htmlContent
	a.Content = content
	a.Summary = summary

	// 调用模型完成更新操作
	err := models.Article{}.UpdateArticleById(&a)
	if err != nil {
		p.About500(syserrors.New("编辑失败！", err))
	} else {
		if len(tagId) != 0 { // 如果有选择标签才维护
			// 维护article_tag表
			// 将tagId数组转化为用','分隔的字符串
			tagIds := strings.Replace(strings.Trim(fmt.Sprint(tagId), "[]"), " ", ",", -1)
			// 添加到article_tag表
			models.ArticleTag{}.UpdateArticleTagById(a.ID, tagIds)
			// 提示并跳转
			p.ReturnJson("编辑成功！", "/admin/article/list/1")
		} else {
			p.ReturnJson("编辑成功！", "/admin/article/list/1")
		}
	}
}

/**
 * 删除文章
 */
func (p *ArticleController) ArticleDelete() {
	p.MustLogin()
	var aid models.ArticleDeleteId // 用结构体接收ajax过来的数据
	err1 := json.Unmarshal(p.Ctx.Input.RequestBody, &aid)
	if err1 == nil {
		err2 := models.Article{}.DeleteArticleById(aid.Id)
		if err2 == nil {
			// 维护article_tag表，删除文章时删除对应的标签数据
			models.ArticleTag{}.DeleteArticleTagById(aid.Id)
			p.ReturnJsonCode("删除成功！")
		} else {
			p.About500(syserrors.New("删除失败！", err2))
		}
	}
}
