package controllers

import (
	"bejigo/models"
	"bejigo/syserrors"
	"encoding/json"
)

type LinkController struct {
	BaseController
}

/*
 * 链接列表
 */
func (p *LinkController) LinkList() {
	p.MustLogin()
	var link []models.Link
	models.Link{}.GetAllLink(&link)
	p.Data["LinkInfo"] = link
	p.AdminCommTpl("link/list.html", "链接列表")
}

/*
 * 链接添加表单显示
 */
func (p *LinkController) LinkAddIndex() {
	p.MustLogin()
	p.AdminCommTpl("link/add.html", "链接添加")
}

/*
 * 链接添加
 */
func (p *LinkController) LinkCreate() {
	linkName := p.GetMustAndInlen("link_name", "链接不能为空！", "链接名称长度在1-20个字符之间！", 20)
	linkUrl := p.GetMustString("link_url", "链接不能为空！")
	var link models.Link
	link.LinkName = linkName
	link.LinkUrl = linkUrl
	err := models.Link{}.LinkCreate(link)
	if err == nil {
		p.ReturnJson("添加成功！", "/admin/link/list")
	} else {
		p.About500(syserrors.New("添加失败", err))
	}
}

/*
* 链接编辑页显示
 */
func (p *LinkController) LinkEditIndex() {
	p.MustLogin()
	// 获取id
	id := p.Ctx.Input.Param(":id")
	// 查询出当前要编辑的链接信息
	var link models.Link
	linkInfo := models.Link{}.GetLinkById(id, link)
	p.Data["LinkInfo"] = linkInfo
	p.AdminCommTpl("link/edit.html", "编辑链接")
}

/*
 * 修改链接
 */
func (p *LinkController) LinkEdit() {
	p.MustLogin()
	// 接收表单数据
	id, _ := p.GetInt("id")
	linkName := p.GetMustAndInlen("link_name", "链接不能为空！", "链接名称长度在1-20个字符之间！", 20)
	linkUrl := p.GetMustString("link_url", "链接不能为空！")

	// 赋值给模型
	var l models.Link
	l.ID = uint(id)
	l.LinkName = linkName
	l.LinkUrl = linkUrl
	err := models.Link{}.EditLink(&l)
	if err == nil {
		p.ReturnJson("修改成功！", "/admin/link/list")
	} else {
		p.About500(syserrors.New("修改失败！", err))
	}
}

/**
 * 删除链接
 */
func (p *LinkController) LinkDelete() {
	p.MustLogin()
	// 接收id
	var id models.LinkDeleteId
	if err1 := json.Unmarshal(p.Ctx.Input.RequestBody, &id); err1 == nil {
		// 根据id删除链接
		err2 := models.Link{}.DeleteLink(id.Id)
		if err2 == nil {
			p.ReturnJsonCode("删除成功！")
		} else {
			p.About500(syserrors.New("删除失败！", err2))
		}
	}
}
