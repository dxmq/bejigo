package controllers

import (
	"bego/models"
	"bego/syserrors"
	"encoding/json"
	"errors"
	"os"
	"strconv"
)

type SingleController struct {
	BaseController
}

/**
 * 单页面添加页面显示
 */
func (p *SingleController) AddPageIndex() {
	p.MustLogin()
	p.permission()
	p.AdminCommTpl("single/add.html", "页面添加")
}

// 页面列表
func (p *SingleController) PageList() {
	p.MustLogin()
	var s []models.SinglePage
	models.SinglePage{}.GetAllSinglePage(&s)
	p.Data["SinglePages"] = s
	p.AdminCommTpl("single/list.html", "页面列表")
}

/**
 * 单页面添加
 */
func (p *SingleController) PageAdd() {
	p.MustLogin()
	p.permission()
	// 接收表单数据
	pageName := p.GetMustAndInlen("page_name", "页面名称不能空！", "页面名称的长度在1-15个字符之间", 15)
	pageAlias := p.GetMustAndInlen("page_alias", "页面别名不能空！", "页面别名的长度在1-15个字符之间", 15)
	isShow, _ := p.GetInt("is_show")
	Sort, err1 := p.GetInt("sort")
	if err1 != nil {
		p.About500(errors.New("排序不能为空，且只能是正整数！"))
	}
	content := p.GetMustString("content", "页面内容不能为空！")
	pageIconClass := p.GetMustAndInlen("page_icon_class", "页面图标类名称！", "页面图标类名称长度在15个字符之间", 15)
	// 赋值给模型
	var s models.SinglePage
	s.PageName = pageName
	s.PageAlias = pageAlias
	s.IsShow = isShow
	s.Sort = uint(Sort)
	//s.Content = beego.Htmlquote(content) // 转义html
	s.Content = content
	s.PageIconClass = pageIconClass

	// 调用模型完成添加操作
	err := models.SinglePage{}.SinglePageCreate(&s)
	if err == nil {
		p.ReturnJson("单页面添加成功！", "/admin/single/lst")
	} else {
		p.About500(syserrors.New("添加失败！", err))
	}
}

// 编辑页面显示
func (p *SingleController) EditPageIndex() {
	p.MustLogin()
	p.permission()
	id := p.Ctx.Input.Param(":id")
	var s models.SinglePage
	models.SinglePage{}.GetOneSinglePageById(id, &s)
	p.Data["SinglePage"] = s
	p.AdminCommTpl("single/edit.html", "编辑页面")
}

// 编辑
func (p *SingleController) PageEdit() {
	p.MustLogin()
	p.permission()
	// 接收表单数据
	id, _ := p.GetInt("id")
	pageName := p.GetMustAndInlen("page_name", "页面名称不能空！", "页面名称的长度在1-15个字符之间", 15)
	pageAlias := p.GetMustAndInlen("page_alias", "页面别名不能空！", "页面别名的长度在1-15个字符之间", 15)
	Sort, err1 := p.GetInt("sort")
	if err1 != nil {
		p.About500(errors.New("排序不能为空，且只能是正整数！"))
	}
	isShow, _ := p.GetInt("is_show")
	content := p.GetMustString("content", "页面内容不能为空！")
	pageIconClass := p.GetMustAndInlen("page_icon_class", "页面图标类名称！", "页面图标类名称长度在15个字符之间", 15)
	// 赋值给模型
	var s models.SinglePage
	s.ID = uint(id)
	s.PageName = pageName
	s.PageAlias = pageAlias
	s.IsShow = isShow
	s.Sort = uint(Sort)
	s.Content = content
	s.PageIconClass = pageIconClass

	// 调用模型完成添加操作
	err := models.SinglePage{}.PageEdit(&s)
	if err == nil {
		p.ReturnJson("修改成功！", "/admin/single/lst")
	} else {
		p.About500(syserrors.New("修改失败！", err))
	}
}

// 页面删除
func (p *SingleController) PageDelete() {
	p.MustLogin()
	p.permission()
	var id models.PageDeleteId
	if err := json.Unmarshal(p.Ctx.Input.RequestBody, &id); err == nil {
		// 删除数据表中数据的同时删除创建的文件
		// 根据id查询出PageAlias
		var s models.SinglePage
		models.SinglePage{}.GetPageAliasById(id.Id, &s)
		os.Remove("./views/index/" + s.PageAlias + ".html")
		err = models.SinglePage{}.PageDelete(id.Id)
		if err == nil {
			p.ReturnJsonCode("删除成功")
		} else {
			p.About500(errors.New("删除失败！"))
		}
	}
}

// 页面排序
func (p *SingleController) PageSort() {
	p.MustLogin()
	var ps models.PageSort
	if err := json.Unmarshal(p.Ctx.Input.RequestBody, &ps); err == nil {
		sort, err1 := strconv.Atoi(ps.Sort)
		if err1 == nil && sort >= 0 {
			err = models.SinglePage{}.PageSort(ps.Id, ps.Sort)
			if err == nil {
				p.ReturnJson("排序成功！", "/admin/single/lst")
			} else {
				p.About500(syserrors.New("排序失败！", err))
			}
		} else {
			p.About500(errors.New("排序必须是正整数！"))
		}
	}
}
