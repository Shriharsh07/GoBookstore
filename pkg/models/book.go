package models

import (
	"github.com/jinzhu/gorm"
	"github.com/shriharsh07/go-bookstore/pkg/config"
)

var db *gorm.DB

// Creating Book Model
type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// The init func help creates the book table
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// Creating book
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

// Retrieving all books
func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

// Getting Book by Id
func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

// Deleting Book
func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
