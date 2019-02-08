package controllers

import (
	"bejigo/models"
	"bejigo/syserrors"
	"encoding/json"
	"strconv"
)

type WhisperController struct {
	BaseController
}

func (p *WhisperController) AddIndex() {
	p.MustLogin()
	p.AdminCommTpl("whisper/add.html", "添加一条轻语")
}

func (p *WhisperController) Create() {
	p.MustLogin()
	userId, _ := p.GetInt("user_id")
	whisper := p.GetMustAndInlen("whisper", "轻语不能为空！", "内容在100个字符之间！", 100)
	var w models.Whisper
	w.UserId = uint(userId)
	w.Whisper = whisper

	err := models.Whisper{}.WhisperAdd(&w)
	if err == nil {
		p.ReturnJson("添加成功！", "/admin/whisper/list/1")
	} else {
		p.About500(syserrors.New("添加失败！", err))
	}
}

func (p *WhisperController) List() {
	p.MustLogin()
	var ws []models.WhisperRes
	var a models.Article
	pageId := p.Ctx.Input.Param(":id")
	pg, _ := strconv.Atoi(pageId)
	count := models.Whisper{}.GetAllCount()
	pageSize := 10
	models.Whisper{}.GetAllWhisper(&ws, pageSize, pg)
	Page := a.PageUtil(count, pg, pageSize, ws)
	p.Data["Page"] = Page
	p.Data["WCount"] = count
	p.Data["WPageSize"] = pageSize
	p.AdminCommTpl("whisper/list.html", "轻语列表")
}

func (p *WhisperController) EditIndex() {
	p.MustLogin()
	id := p.Ctx.Input.Param(":id")
	var w models.Whisper
	wid, _ := strconv.Atoi(id)
	w.ID = uint(wid)
	models.Whisper{}.GetOneWhisper(&w)
	p.Data["Whisper"] = w
	p.AdminCommTpl("whisper/edit.html", "修改轻语")
}

func (p *WhisperController) Edit() {
	p.MustLogin()
	id, _ := p.GetInt("id")
	userId, _ := p.GetInt("user_id")
	whisper := p.GetMustAndInlen("whisper", "轻语不能为空！", "内容在100个字符之间！", 100)
	var w models.Whisper
	w.ID = uint(id)
	w.UserId = uint(userId)
	w.Whisper = whisper

	err := models.Whisper{}.WhisperEdit(&w)
	if err == nil {
		p.ReturnJson("修改成功！", "/admin/whisper/list/1")
	} else {
		p.About500(syserrors.New("修改失败！", err))
	}
}

func (p *WhisperController) Delete() {
	p.MustLogin()
	var id models.WhisperDeleteId
	if err := json.Unmarshal(p.Ctx.Input.RequestBody, &id); err == nil {
		err = models.Whisper{}.WhisperDelete(id.Id)
		if err == nil {
			p.ReturnJsonCode("删除成功！")
		} else {
			p.About500(syserrors.New("删除失败！", err))
		}
	}
}
