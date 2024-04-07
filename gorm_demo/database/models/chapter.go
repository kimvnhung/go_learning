package models

type Chapter struct {
	ID      int     `gorm:"primaryKey;autoIncrement:true"`
	Name    string  `gorm:""`
	Entries []Entry `gorm:""`
	Index   int     `gorm:"uniqueIndex:chapteridx"`
	BookID  int     `gorm:"uniqueIndex:chapteridx"`
}
