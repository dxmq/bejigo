package models

import "github.com/jinzhu/gorm"

//用户表
type User struct {
	gorm.Model
	UserName string `gorm:"unique;size:15;not null;unique_index;"`
	PassWord string `gorm:"size:32;not null;"`
	Email    string `valid:"Email";gorm:"type:varchar(100);unique_index"`
	Avatar   string `gorm:"type:varchar(100)"`
	Role     int    `gorm:default:1` // 1：代表管理员；0代表普通会员
}

// 查询出是否有这个用户
func QueryUser(username, password string) (user User, err error) {
	return user, db.Where("user_name = ? and pass_word = ?", username, password).Take(&user).Error
}

// 用户列表
func (p User) GetAllUser(u *[]User) {
	db.Select("id, user_name, email, role, created_at").Find(&u)
}

// 用户添加
func (p User) UserCreate(u *User) error {
	db.NewRecord(u)
	return db.Create(&u).Error
}

// 根据id查询出用户的信息
func (p User) GetUserInfoById(id string, user *User) {
	db.Where("id = ?", id).Find(&user)
}

// 编辑用户信息
func (p User) UserEdit(user *User) error {
	return db.Save(&user).Error
}

type UserDeleteId struct {
	Id string
}

// 删除用户
func (p User) UserDelete(id string) error {
	return db.Unscoped().Where("id = ?", id).Delete(&User{}).Error
}

// 前台
// 取出超级管理员信息
func (p User) GetSupperInfo(id, role string, u *User) {
	db.Select("user_name, email").Where("id = ?", id).Where("role = ?", 1).Find(&u)
}
