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

// 修改
func (p ArticleTag) UpdateArticleTagById(articleId uint, tagIds string) {
	db.Model(&p).Where("article_id = ?", articleId).Update("tag_id", tagIds)
}
