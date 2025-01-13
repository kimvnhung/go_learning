package models

type News struct {
	// ID is the primary key for the news model and is serialized as "id" in JSON.
	ID          int    `json:"id" gorm:"primaryKey;autoIncrement=true"`
	Title       string `json:"title" gorm:"unqiueIndex=news"`
	Group       string `json:"group" gorm:"unqiueIndex=news"`
	Content     string `json:"content" gorm:"unqiueIndex=news"`
	Author      string `json:"author" gorm:"unqiueIndex=news"`
	HouseDetail *House `json:"house_detail" default:"null"`
	Date        string `json:"date" gorm:"unqiueIndex=news"`
}
