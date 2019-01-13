package models

import (
	"github.com/jinzhu/gorm"
)

type Tag struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null;unique_index;unique"`
}

// 修改标签
func SaveTag(t *Tag, id, name string) error {
	return db.Model(&t).Where("id = ?", id).Update("name", name).Error
}

// 添加标签
func AddTag(t *Tag) error {
	db.NewRecord(t)
	return db.Create(&t).Error
}

func GetAll(a *[]Tag) error {
	return db.Find(&a).Error
}

// 查询当前id的标签信息
func GetTagById(id string) (tag Tag, err error) {
	return tag, db.Where("id = ?", id).Find(&tag).Error
}

// 软删除标签
func DeleteTagById(id string) error {
	return db.Unscoped().Where("id = ?", id).Delete(&Tag{}).Error
}
