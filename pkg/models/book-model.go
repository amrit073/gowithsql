package models

import (
	"github.com/amrit073/gosql/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"json:\"name\""`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Save(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookBYId(Id uint64) *Book {
	var getBook Book
	db.Where("ID=?", Id).Find(&getBook)
	return &getBook
}

func DeleteBook(Id uint64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return book

}