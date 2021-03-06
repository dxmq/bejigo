package models

type System struct {
	ID              uint   `gorm:"primary_key"`
	WebName         string `gorm:"size:20;not null;"`
	RecordNumber    string `gorm:"size:20;"`
	CopyRight       string `gorm:"size:50;"`
	WebTitle        string `gorm:"size:50;"`
	WebKeywords     string `gorm:"size:100;"`
	DefaultAuthor   string `gorm:"size:15;"`
	WebDescription  string `gorm:"size:100"`
	WebSlogan       string `gorm:"size:100"`
	IndexShowNumber uint
}

func (p System) GetSystem(st *System) {
	db.Find(&st)
}
func (p System) GetSectionSystem(st *System) {
	db.Select("web_name, default_author").Find(&st)
}

func (p System) SaveSystem(st *System) error {
	return db.Save(&st).Error
}

type IndexShowNumber struct {
	IndexShowNumber uint
}

func (p System) GetIndexShowNumber(isn *IndexShowNumber) {
	db.Model(&p).Select("index_show_number").Scan(&isn)
}
