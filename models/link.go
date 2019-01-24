package models

import "github.com/jinzhu/gorm"

type Link struct {
	gorm.Model
	LinkName string `gorm:"size:20;not null;unique_index;index:link;"`
	LinkUrl  string `gorm:"not null;"`
}

// 链接添加
func (p Link) LinkCreate(l Link) error {
	db.NewRecord(l)
	return db.Create(&l).Error
}

// 获取所有的链接
func (p Link) GetAllLink(l *[]Link) {
	db.Select("id, link_name, link_url, created_at").Find(&l)
}

// 根据id查询出一条链接
func (p Link) GetLinkById(id string, link Link) Link {
	db.Select("id, link_name, link_url").Where("id = ?", id).Find(&link)
	return link
}

// 修改链接
func (p Link) EditLink(l *Link) error {
	return db.Save(&l).Error
}

// 用于接收删除传过来的id
type LinkDeleteId struct {
	Id string
}

// 根据id删除链接
func (p Link) DeleteLink(id string) error {
	return db.Unscoped().Where("id = ?", id).Delete(&p).Error
}
