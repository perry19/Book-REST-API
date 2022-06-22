package models

import "gorm.io/gorm"

// Book structure
type Book struct {
	gorm.Model
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}
