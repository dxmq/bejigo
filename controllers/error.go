package controllers

import (
	"bejigo/syserrors"
	"github.com/astaxie/beego/logs"
)

type ErrorController struct {
	BaseController
}

// code,msg,reason
func (c *ErrorController) Error404() {
	c.TplName = "error/404.html"
	if c.IsAjax() {
		c.jsonError(syserrors.Error404{})
	}
}

func (c *ErrorController) Error500() {
	c.Data["isIndex"] = false
	c.TplName = "error/500.html"
	err, ok := c.Data["error"].(error)
	if !ok {
		err = syserrors.New("未知错误", nil)
	}

	// 将错误转化为系统错误
	serr, ok := err.(syserrors.Error)
	if !ok {
		serr = syserrors.New(err.Error(), nil)
	}
	if serr.ReasonError() != nil {
		logs.Info(serr.Error(), serr.ReasonError())
	}

	if c.IsAjax() {
		c.jsonError(serr)
	} else {
		c.Data["content"] = serr.Error()
	}
}

func (c *ErrorController) jsonError(serr syserrors.Error) {
	c.Ctx.Output.Status = 200 //状态码改成200
	c.Data["json"] = map[string]interface{}{
		"code": serr.Code(),
		"msg":  serr.Error(),
	}
	c.ServeJSON()
}
