package controllers

import (
	"bejigo/models"
	"bejigo/syserrors"
	"encoding/json"
	"errors"
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
	ctx.Data["Category"] = ctx.GetAllCategory()
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
	categoryName := ctx.GetMustAndInlen("category_name", "分类名称不能为空！", "分类名称长度在1-20个字符之间！", 20)
	// 调用模型的CategoryCreate方法 完成添加
	err := models.Category{}.CategoryCreate(categoryName)
	if err != nil {
		ctx.About500(syserrors.New("添加失败！", err))
	}
	// 初始化搜索数据，用于前台搜索
	ctx.initSearchData()

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
	// 接收并验证数据
	categoryName := ctx.GetMustAndInlen("category_name", "分类不能为空！", "分类长度在1-20个字符之间！", 20)
	id := ctx.GetString("ID")
	// 调用模型去编辑
	var category models.Category
	err := models.Category{}.CategoryEdit(&category, id, categoryName)
	if err != nil {
		ctx.About500(syserrors.New("编辑失败！", err))
	}

	// 初始化搜索数据，用于前台搜索
	ctx.initSearchData()

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
		// 要先查询当前分类下还有没有文章
		count := models.Article{}.GetCountByCateId(c.Id)
		if count != 0 {
			ctx.About500(errors.New("当前分类下还有文章，请先删除文章！"))
		}
		// 如果当前分类下没有文章
		err1 := models.Category{}.CategoryDeleteById(c.Id)
		if err1 == nil { // 如果没有错的话
			// 初始化搜索数据，用于前台搜索
			ctx.initSearchData()

			ctx.ReturnJsonCode("删除成功")
		} else {
			ctx.About500(syserrors.New("删除失败！", err))
		}
	}
}
