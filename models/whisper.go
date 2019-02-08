package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Whisper struct {
	gorm.Model
	UserId  uint
	Whisper string `gorm:"index:whisper;not null;size:100"`
}

func (p Whisper) WhisperAdd(w *Whisper) error {
	db.NewRecord(w)
	return db.Create(&w).Error
}

func (p Whisper) GetAllCount() int {
	var count int
	db.Model(&Whisper{}).Count(&count)
	return count
}

type WhisperRes struct {
	ID        uint
	UserName  string
	Whisper   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p Whisper) GetAllWhisper(ws *[]WhisperRes, limit, pageId int) {
	db.Table("whispers as a").Select("a.id,a.whisper,a.created_at, a.updated_at,b.user_name").Joins("left join users as b on a.user_id=b.id").Limit(limit).Offset((pageId - 1) * limit).Order("a.created_at desc").Scan(&ws)
}

func (p Whisper) GetOneWhisper(w *Whisper) {
	db.Select("id, user_id, whisper").Find(&w)
}

func (p Whisper) WhisperEdit(w *Whisper) error {
	return db.Save(&w).Error
}

type WhisperDeleteId struct {
	Id string
}

func (p Whisper) WhisperDelete(id string) error {
	return db.Unscoped().Where("id = ?", id).Delete(&p).Error
}

// 前台
func (p Whisper) GetNewWhisper(number int, w *[]Whisper) {
	db.Select("id, whisper").Order("created_at desc").Limit(number).Find(&w)
}

func (p Whisper) GetWhisperById(id string, ws *WhisperRes) {
	db.Table("whispers as a").Where("a.id = ?", id).Select("a.id,a.whisper,a.created_at,b.user_name").Joins("left join users as b on a.user_id=b.id").Scan(&ws)
}

func (p Whisper) GetNextWhisper(id uint, w *Whisper) {
	db.Select("id, whisper").Where("id > ?", id).Order("created_at desc").Find(&w).Limit(1)
}
