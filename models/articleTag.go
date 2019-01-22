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
	// db.Unscoped().Where("tag_id = ?", "").Delete(&ArticleTag{}) // 如果没有标签id就删除
	db.Model(&p).Where("article_id = ?", articleId).Update("tag_id", tagIds)
}

// 删除
func (p ArticleTag) DeleteArticleTagById(id string) {
	db.Unscoped().Where("article_id = ?", id).Delete(&ArticleTag{})
}

// 查询出所有的tag_ids
func (p ArticleTag) GetAllTagId(at *[]ArticleTag) {
	db.Find(&at)
}
