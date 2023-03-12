package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title string `json:"title"`
	Price string `json:"price"`
}
