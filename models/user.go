package models

import "github.com/jinzhu/gorm"

//用户表
type User struct {
	gorm.Model
	UserName string `gorm:"unique;size:15;not null;unique_index;"`
	PassWord string `gorm:"size:32;not null;"`
	Email    string `gorm:"type:varchar(100);unique_index"`
	Avatar   string `gorm:"type:varchar(100)"`
	Role     int    `gorm:default:1` // 1：代表管理员；0代表普通会员
}

// 查询出是否有这个用户
func QueryUser(username, password string) (user User, err error) {
	return user, db.Where("user_name = ? and pass_word = ?", username, password).Take(&user).Error
}

// 根据id查询出用户的信息
func (p User) GetUserInfoById(id string, user *User) {
	db.Where("id = ?", id).Find(&user)
}
