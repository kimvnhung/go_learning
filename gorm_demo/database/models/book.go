package models

type Book struct {
	ID       int       `gorm:"primaryKey;autoIncrement:true"`
	Name     string    `gorm:"unique"`
	Chapters []Chapter `gorm:""`
}
