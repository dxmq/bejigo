package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	gorm.Model
	// Slug string `gorm:"not null;size:15;"`
	CategoryId int `gorm:"index:category"` // 分类id
	// Tag []Tag `gorm:"many2many:article_tags"` // 多对多，使用article_tags表连接
	Title       string `gorm:"size:40;not null;unique;unique_index;"` // 标题
	Author      string `gorm:"size:20;not null;index:author;"`        // 作者
	IsShow      int    `gorm:"type:enum(1,0);default:1"`              // 1是显示，默认不显示
	IsTop       int    `gorm:"type:enum(1,0);default:0"`              // 1是置顶，默认不置顶
	Views       int    // 浏览次数
	Summary     string `gorm:"type:varchar(150)"` // 概要
	Content     string `gorm:"type:text"`         // 内容
	HtmlContent string `gorm:"type:text"`
}

// 文章添加
func (p Article) ArticleAdd(a *Article) (err error) {
	db.NewRecord(a)
	return db.Create(&a).Error
}

// 查询出刚添加后的id
func (p Article) GetIdByTitle(a *Article, title string) {
	db.Where("title = ?", title).Find(&a)
}

type Result struct {
	ID           uint
	CategoryId   int
	TagId        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	CategoryName string
	Title        string
	Author       string
	IsShow       int
	IsTop        int
	Views        int // 浏览次数
	Summary      string
	Content      string
	HtmlContent  string
}

type Search struct {
	ID           uint
	CategoryName string
	Title        string
	CreatedAt    time.Time
}

type Page struct {
	PageNo     int
	PageSize   int
	TotalPage  int
	TotalCount int
	FirstPage  bool
	LastPage   bool
	List       interface{}
}

// 返回分页数据
func (p Article) PageUtil(count int, pageNo int, pageSize int, list interface{}) Page {
	tp := count / pageSize
	if count%pageSize > 0 {
		tp = count/pageSize + 1
	}
	return Page{PageNo: pageNo, PageSize: pageSize, TotalPage: tp, TotalCount: count, FirstPage: pageNo == 1, LastPage: pageNo == tp, List: list}
}

// 根据页码获取文章信息
func (p Article) GetAllArticleByPage(r *[]Result, pageSize, pageId int) {
	db.Table("articles as a").Select("a.id, a.category_id, a.title, a.author, a.is_show, a.is_top, a.views, a.created_at, b.category_name").Joins("left join categories as b on b.id = a.category_id").Limit(pageSize).Offset((pageId - 1) * pageSize).Order("a.created_at desc").Scan(&r)
}

// 获取总的记录数
func (p Article) GetAllCount() int {
	var count int
	db.Model(&Article{}).Count(&count)
	return count
}

// 根据id获取一条文章的信息
func (p Article) GetArticleById(id string, r *Result) {
	db.Table("articles as a").Select("a.id, a.category_id, a.title, a.author, a.is_show, a.is_top, a.content, b.category_name, c.tag_id").Joins("left join categories as b on b.id = a.category_id").Joins("left join article_tags as c on c.article_id = a.id").Where("a.id = ?", id).Scan(&r)
}

// 根据id更新文章信息
func (p Article) UpdateArticleById(a *Article) error {
	return db.Save(&a).Error
}

type ArticleDeleteId struct {
	Id string
}

// 删除文章
func (p Article) DeleteArticleById(id string) error {
	// 强制删除，不是软删除
	return db.Unscoped().Where("id = ?", id).Delete(&Article{}).Error
}

// 统计出拥有当前分类id的文章数目
func (p Article) GetCountByCateId(categoryId string) int {
	var count int
	db.Model(&p).Where("category_id = ?", categoryId).Count(&count)
	return count
}

/******* 前台相关 ******/
// 根据页码获取文章信息
func (p Article) GetArticleData(pageSize int, r *[]Result) {
	db.Table("articles as a").Select("a.id, a.category_id, a.title, a.created_at, a.summary, b.category_name").Where("is_show = ?", 1).Order("a.is_top desc").Order("a.created_at desc").Joins("left join categories as b on b.id = a.category_id").Limit(pageSize).Scan(&r)
}

type NewArticle struct {
	ID    uint
	Title string
}

// 取出最新的文章
func (p Article) GetNewArticle(number int, na *[]NewArticle) {
	db.Model(&p).Order("created_at desc").Select("id, title").Limit(5).Scan(&na)
}

func (p Article) GetArchives(r *[]Result) {
	db.Model(&p).Order("created_at desc").Select("id, title, created_at").Scan(&r)
}

// 根据id取出一条文章信息
func (p Article) GetArticleInfoById(id string, r *Result) {
	db.Table("articles as a").Select("a.id, a.category_id, a.title, a.html_content, a.content, a.created_at, a.author, b.category_name, c.tag_id").Joins("left join categories as b on b.id = a.category_id").Joins("left join article_tags as c on c.article_id = a.id").Where("a.id = ?", id).Scan(&r)
}

type ArticleArch struct {
	ID        uint
	Title     string
	CreatedAt time.Time
}

// 查询出当前分类下的文章
func (p Article) GetArticleByCateName(categoryName string, r *[]Result, limitNum int) {
	db.Table("articles as a").Select("a.id, a.title, a.created_at, b.category_name").Joins("left join categories as b on b.id = a.category_id").Where("b.category_name = ?", categoryName).Where("a.is_show = ?", 1).Order("a.created_at desc").Limit(limitNum).Scan(&r)
}

// 通过article_ids查询出文章
func (p Article) GetArticleByArticleId(id string, ar *[]ArticleArch, limitNum int) {
	db.Raw("SELECT id, title, created_at FROM articles WHERE id in (" + id + ") AND is_show = 1 ORDER BY created_at desc").Limit(limitNum).Scan(&ar)
}

// 取出当前文章的下一篇文章
func (p Article) GetNextArticle(id uint, a *Article) {
	db.Select("id, title").Where("id > ?", id).Where("is_show = ?", 1).Find(&a).Limit(1)
}

// 搜索功能，取出所有文章的id,title,created_at
func (p Article) GetArticleDataForSearch(s *[]Search) {
	db.Table("articles as a").Select("a.id, a.title, a.created_at, b.category_name").Joins("left join categories as b on b.id = a.category_id").Where("a.is_show = ?", 1).Order("a.created_at desc").Scan(&s)
}

type CreatedAt struct {
	ID        uint
	CreatedAt time.Time
}

// 获取所有的文章的添加时间
func (p Article) GetArticleCreateTime(ct *[]CreatedAt) {
	db.Table("articles").Select("id, created_at").Where("is_show = ?", 1).Scan(&ct)
}
