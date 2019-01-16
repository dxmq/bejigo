package controllers

import (
	"bego/models"
	"bego/syserrors"
	"encoding/json"
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
	name := p.GetMustString("name", "标签不能为空！")
	name = p.GetAllowMaxString("name", "长度在0-20个字符之间", 20)
	var t models.Tag
	t.Name = name
	if err := models.AddTag(&t); err != nil {
		p.About500(syserrors.New("添加失败", err))
	}
	p.ReturnJson("添加成功", "/admin/tag/create")
}

/**
 * 标签列表页面
 */
// @router /admin/tag [get]
func (p *TagController) TagList() {
	p.MustLogin()
	var a []models.Tag
	err := models.GetAll(&a)
	if err != nil {
		p.About500(syserrors.New("暂无标签！", err))
	}
	p.Data["tag"] = a
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
	name := p.GetMustString("name", "标签不能为空！")
	name = p.GetAllowMaxString("name", "长度在0-20个字符之间！", 20)
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
	json.Unmarshal(p.Ctx.Input.RequestBody, &t)
	// 删除
	if err := models.DeleteTagById(t.Id); err != nil {
		p.About500(syserrors.New("删除失败", err))
	}
	p.ReturnJsonCode("删除成功")
}
