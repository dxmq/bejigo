package models

import (
	"github.com/jinzhu/gorm"
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
	return db.Find(&category).Error
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
