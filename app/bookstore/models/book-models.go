package models

import (
	"github.com/doziestar/gosyn/app/bookstore/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.ConnectDb()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) Create() error {
	return db.Create(b).Error
}

func (b *Book) Update() error {
	return db.Save(b).Error
}

func (b *Book) Delete() error {
	return db.Delete(b).Error
}

func (b *Book) Get() error {
	return db.First(b).Error
}

func GetBooks() ([]Book, error) {
	var books []Book
	err := db.Find(&books).Error
	return books, err
}

func GetBookById(id int64) (Book, error) {
	var book Book
	err := db.First(&book, id).Error
	return book, err
}
