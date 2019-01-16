package controllers

import (
	"bego/models"
	"bego/syserrors"
	"encoding/json"
	"github.com/astaxie/beego"
)

type CategoryController struct {
	BaseController
}
type CategoryDeleteId struct {
	Id string
}

/**
 * 分类列表
 */
func (ctx *CategoryController) CategoryList() {
	ctx.MustLogin()
	// 取出所有的分类信息
	var category []models.Category
	err := models.Category{}.GetAllCategory(&category)
	if err != nil {
		ctx.About500(syserrors.New("暂无分类！", err))
	}
	ctx.Data["Category"] = category
	ctx.AdminCommTpl("category/list.html", "分类列表")
}

/**
 * 分类添加的表单显示
 */
func (ctx *CategoryController) CategoryAddIndex() {
	ctx.MustLogin()
	ctx.AdminCommTpl("category/add.html", "添加分类")
}

/**
 * 分类添加
 */
func (ctx *CategoryController) CategoryCreate() {
	ctx.MustLogin()
	// 接收并验证数据
	categoryName := ctx.GetMustString("category_name", "分类不能为空！")
	categoryName = ctx.GetAllowMaxString("category_name", "长度在0-20个字符之间！", 20)

	// 调用模型的CategoryCreate方法 完成添加
	err := models.Category{}.CategoryCreate(categoryName)
	if err != nil {
		ctx.About500(syserrors.New("添加失败！", err))
	}
	ctx.ReturnJson("添加成功", "/admin/category/list")
}

/**
 * 分类编辑页面的显示
 */
func (ctx *CategoryController) CategoryEditIndex() {
	ctx.MustLogin()
	var c models.Category
	// 获取id
	id := ctx.Ctx.Input.Param(":id")
	// 模型查询出当前id的数据
	category, err := c.GetCategoryById(id)
	if err != nil {
		ctx.About500(syserrors.New("出错了", err))
	}
	ctx.Data["CategoryName"] = category.CategoryName
	ctx.Data["ID"] = category.ID
	ctx.AdminCommTpl("category/edit.html", "编辑分类")
}

/**
 * 分类编辑
 */
func (ctx *CategoryController) CategoryEdit() {
	ctx.MustLogin()
	beego.Info("123")
	// 接收并验证数据
	categoryName := ctx.GetMustString("category_name", "分类不能为空！")
	categoryName = ctx.GetAllowMaxString("category_name", "长度在0-20个字符之间！", 20)
	id := ctx.GetString("ID")
	beego.Info(id)
	beego.Info(categoryName)
	// 调用模型去编辑
	var category models.Category
	err := models.Category{}.CategoryEdit(&category, id, categoryName)
	if err != nil {
		ctx.About500(syserrors.New("编辑失败！", err))
	}
	ctx.ReturnJson("编辑成功！", "/admin/category/list")
}

/**
 * 分类删除
 */
func (ctx *CategoryController) CategoryDelete() {
	ctx.MustLogin()
	var c CategoryDeleteId
	// 接收id
	err := json.Unmarshal(ctx.Ctx.Input.RequestBody, &c)
	if err == nil {
		err1 := models.Category{}.CategoryDeleteById(c.Id)
		if err1 != nil { // 如果有错的话
			ctx.About500(syserrors.New("删除失败！", err))
		} else {
			ctx.ReturnJsonCode("删除成功")
		}
	}
}
