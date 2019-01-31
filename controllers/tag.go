package controllers

import (
	"bego/models"
	"bego/syserrors"
	"encoding/json"
	"strings"
)

type TagController struct {
	BaseController
}

// 用于接收删除时ajax传过来的 id
type Object struct {
	Id string
}

/**
 * 标签添加表单显示
 */
// @router /admin/tag/create [get]
func (p *TagController) TagCreate() {
	p.MustLogin() // 必须登录
	p.AdminCommTpl("tag/add.html", "标签添加")
}

/**
 * 标签添加
 */
// @router /admin/tag [post]
func (p *TagController) TagStore() {
	p.MustLogin()
	name := p.GetMustAndInlen("name", "标签不能为空！", "标签长度在1-20个字符之间！", 20)
	var t models.Tag
	t.Name = name
	if err := models.AddTag(&t); err != nil {
		p.About500(syserrors.New("添加失败", err))
	}
	p.ReturnJson("添加成功", "/admin/tag")
}

/**
 * 标签列表页面
 */
// @router /admin/tag [get]
func (p *TagController) TagList() {
	p.MustLogin()
	p.Data["tag"] = p.GetAllTag()
	//fmt.Printf("%v", a)
	p.AdminCommTpl("tag/list.html", "标签列表")
}

/**
 * 修改页面显示
 */
// @router /admin/tag/:id [get]
func (p *TagController) TagShow() {
	p.MustLogin()
	// 查询出当前标签的信息
	var id = p.Ctx.Input.Param(":id")
	tag, err := models.GetTagById(id)
	if err != nil {
		p.About500(syserrors.New("出错了", err))
	}
	p.Data["Name"] = tag.Name
	p.Data["ID"] = tag.ID
	p.AdminCommTpl("tag/edit.html", "标签修改")
}

/**
 * 修改标签
 */
// @router /admin/tag/edit [post]
func (p *TagController) TagEdit() {
	p.MustLogin()
	name := p.GetMustAndInlen("name", "标签不能为空！", "标签长度在1-20个字符之间！", 20)
	id := p.GetString("ID")
	var t models.Tag
	if err := models.SaveTag(&t, id, name); err != nil {
		p.About500(syserrors.New("修改失败", err))
	}
	p.ReturnJson("修改成功", "/admin/tag")
}

/**
 * 强删除标签
 */
// @router /admin/tag/delete [post]
func (p *TagController) TagDelete() {
	p.MustLogin()
	var t Object
	// 接收 id
	err := json.Unmarshal(p.Ctx.Input.RequestBody, &t)
	if err == nil {
		// 删除
		if err1 := models.DeleteTagById(t.Id); err1 != nil {
			p.About500(syserrors.New("删除失败", err1))
		}
		// 维护article_tags表，删除tag时，删除文章对应的tag
		// 查询出所有的tag_ids
		var at []models.ArticleTag
		models.ArticleTag{}.GetAllTagId(&at)
		for _, v := range at {
			option := strings.Index(v.TagId, t.Id+",")
			option1 := strings.Index(v.TagId, t.Id)
			var newTagIds string
			if option != -1 { // 1,2 如果"1,"存在，将"1,"替换为""
				newTagIds = strings.Replace(v.TagId, t.Id+",", "", -1)
			} else {
				if option1 == -1 { // 如果没有
					newTagIds = strings.Replace(v.TagId, ","+t.Id, "", -1)
				}
			}
			models.ArticleTag{}.UpdateArticleTagById(v.ArticleId, newTagIds)
		}
		p.ReturnJsonCode("删除成功")
	}
}
