package entity

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Title string `gorm:"column:title" json:"title"`
	Body  string `gorm:"column:body" json:"body"`
}
