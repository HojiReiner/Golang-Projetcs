package models

import (
	"go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() (books []Book) {
	db.Find(&books)
	return books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Find(&getBook, id)
	return &getBook, db
}

func DeleteBook(id int64) {
	db.Unscoped().Delete(&Book{}, id)
}
