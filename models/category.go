package models

import (
	"github.com/jinzhu/gorm"
	"strconv"
)

type Category struct {
	gorm.Model
	CategoryName string `gorm:"type:varchar(20);not null;unique;unique_index"`
}

// 分类添加
func (c Category) CategoryCreate(CategoryName string) (err error) {
	c.CategoryName = CategoryName
	db.NewRecord(c)
	return db.Create(&c).Error
}

// 获取所有的分类信息
func (c Category) GetAllCategory(category *[]Category) error {
	return db.Select("id, created_at, category_name").Find(&category).Error
}

// 获取当前id的那条分类信息
func (c Category) GetCategoryById(id string) (category Category, err error) {
	return category, db.Where("id = ?", id).Find(&category).Error
}

// 编辑分类
func (c Category) CategoryEdit(category *Category, id, CategoryName string) error {
	return db.Model(&category).Where("id = ?", id).Update("category_name", CategoryName).Error
}

// 删除分类
func (c Category) CategoryDeleteById(id string) error {
	return db.Unscoped().Where("id = ?", id).Delete(&Category{}).Error
}

type CateAndArticleNum struct {
	Id           uint
	CategoryName string
	Number       int
}

// 前台
// 取出所有分类及当前分类下的文章数目
func (c Category) GetAllCateAndNum() (cate []CateAndArticleNum) {
	db.Model(&c).Select("id, category_name").Order("created_at desc").Scan(&cate)
	for inx, v := range cate {
		categoryId := strconv.Itoa(int(v.Id))
		var count int
		db.Model(&Article{}).Where("category_id = ?", categoryId).Where("is_show = ?", 1).Count(&count)
		cate[inx].Number = count
	}
	return cate
}
