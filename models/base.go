package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
)

var (
	db *gorm.DB
)

func init() {
	var err error
	if err = os.MkdirAll("data", 0777); err != nil { // 创建数据库目录
		panic("failed" + err.Error())
	}
	db, err = gorm.Open("sqlite3", "data/data.db")
	if err != nil {
		panic("连接数据库失败")
	}
	db.AutoMigrate(&User{}, &Tag{})
	var count int
	// 如果数据里边没有数据，新增一条记录
	if err := db.Model(&User{}).Count(&count).Error; err == nil && count == 0 {
		db.Create(&User{
			UserName: "admin",
			PassWord: "123",
			Email:    "admin888@qq.com",
			Avatar:   "/static/index/img/face.png",
			Role:     1,
		})
	}
}
