package models

type ArticleTag struct {
	ArticleId uint   `gorm:"index:article_id"`
	TagId     string `gorm:"index:tag_id"`
}

// 添加
func (p ArticleTag) ArticleTagAdd(at *ArticleTag) {
	db.NewRecord(at)
	db.Create(&at)
}
