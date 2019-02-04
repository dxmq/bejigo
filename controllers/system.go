package controllers

import (
	"bego/models"
	"errors"
)

type SystemController struct {
	BaseController
}

func (p *SystemController) SystemIndex() {
	p.MustLogin()
	var st models.System
	models.System{}.GetSystem(&st)
	p.Data["SystemInfo"] = st
	p.AdminCommTpl("system/system.html", "系统设置")
}

func (p *SystemController) SystemSave() {
	p.MustLogin()
	var st models.System

	id, _ := p.GetInt("id", 1)
	webName := p.GetMustAndInlen("web_name", "网站名称不能为空", "网站名称长度在20个字符之间！", 20)
	copyRight := p.GetAllowMaxString("copy_right", "版权信息在50个字符之间！", 50)
	recordNumber := p.GetAllowMaxString("record_number", "备案号在20个字符之间！", 20)
	webTitle := p.GetAllowMaxString("web_title", "网站标题在50个字符之间！", 50)
	webKeywords := p.GetAllowMaxString("web_keywords", "网站描述在100个字符之间！", 100)
	defaultAuthor := p.GetAllowMaxString("default_author", "默认作者长度在15个字符之间！", 15)
	webDescription := p.GetAllowMaxString("web_description", "网站描述在100个字符之间！", 100)
	indexShowNumber, err := p.GetInt("index_show_number")
	if err != nil {
		p.About500(errors.New("首页文章显示数不能为空，且只能是正整数！"))
	}
	webSlogan := p.GetAllowMaxString("web_slogan", "网站标语在100个字符之间！", 100)
	st.ID = uint(id)
	st.WebName = webName
	st.CopyRight = copyRight
	st.RecordNumber = recordNumber
	st.WebTitle = webTitle
	st.WebKeywords = webKeywords
	st.DefaultAuthor = defaultAuthor
	st.WebDescription = webDescription
	st.WebSlogan = webSlogan
	st.IndexShowNumber = uint(indexShowNumber)
	err = models.System{}.SaveSystem(&st)
	if err == nil {
		p.ReturnJson("保存成功！", "/admin/system/index")
	} else {
		p.About500(errors.New("保存失败！"))
	}
}
