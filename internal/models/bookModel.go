package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ID        int    `json:"id" gorm:"primary_key"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	UserID    int    `json:"user_id"`
}
