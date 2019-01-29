package models

import "time"

type Profile struct {
	UserId    uint   `gorm:"primary_key"`
	About     string `gorm:"type:text;"`
	CreatedAt time.Time
}

// 根据id获取关于我
func (p Profile) GetAboutInfoById(id string, profile *Profile) {
	db.Where("user_id = ?", id).Find(&profile)
}

// 个人介绍修改
func (p Profile) ProfileEdit(profile *Profile) {
	db.Save(&profile)
}
