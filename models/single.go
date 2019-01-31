package models

import "time"

type SinglePage struct {
	ID            uint   `gorm:"primary_key"`
	PageName      string `gorm:"size:15;unique;not null;"`
	PageAlias     string `gorm:"size:15;unique;not null;"`
	Sort          uint   // 排序
	IsShow        int    `gorm:"type:enum(1,0);default:1"` // 默认显示
	PageIconClass string `gorm:"size:15;not null;"`
	Content       string `gorm:"type:text"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// 单页面添加
func (p SinglePage) SinglePageCreate(s *SinglePage) error {
	db.NewRecord(s)
	return db.Create(&s).Error
}

// 取出所有的单页面，单页面列表
func (p SinglePage) GetAllSinglePage(s *[]SinglePage) {
	db.Select("id, page_name, page_icon_class, is_show, sort, created_at, page_alias").Find(&s)
}

// 取出一条单页面信息
func (p SinglePage) GetOneSinglePageById(id string, s *SinglePage) {
	db.Select("id, page_name, page_alias, content, is_show, sort, page_icon_class").Where("id = ?", id).Find(&s)
}

// 编辑单页面
func (p SinglePage) PageEdit(s *SinglePage) error {
	return db.Save(&s).Error
}

type PageDeleteId struct {
	Id string
}

// 删除单页面
func (p SinglePage) PageDelete(id string) error {
	return db.Unscoped().Where("id = ?", id).Delete(&SinglePage{}).Error
}

type PageSort struct {
	Id   string
	Sort string
}

// 页面排序
func (p SinglePage) PageSort(id, sort string) error {
	return db.Model(SinglePage{}).Where("id = ?", id).Update("sort", sort).Error
}

// 前台

// 取出页面名称和图标类
func (p SinglePage) GetAllPage(s *[]SinglePage) {
	db.Where("is_show = ?", 1).Order("sort asc").Select("id, page_name, page_alias, page_icon_class").Find(&s)
}

// 根据别名取出某个页面
func (p SinglePage) GetPageByCateName(pageAlias string, s *SinglePage) {
	db.Select("id, content").Where("page_alias = ?", pageAlias).Find(&s)
}
