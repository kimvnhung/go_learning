package models

type Entry struct {
	ID        int    `gorm:"primaryKey;autoIncrement:true"`
	Name      string `gorm:""`
	Index     int    `gorm:"uniqueIndex:entryidx"`
	ChapterID int    `gorm:"uniqueIndex:entryidx"`
}
