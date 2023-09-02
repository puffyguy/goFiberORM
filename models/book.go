package models

import "gorm.io/gorm"

// var DB *mongo.Database
var DB *gorm.DB

// Struct for storing Book details in database
type Book struct {
	Name   string `json:"name" validate:"required"`
	ISBN   string `gorm:"primaryKey" json:"isbn" validate:"required"`
	Author string `json:"author" validate:"required"`
	Price  int    `json:"price" validate:"required"`
}
