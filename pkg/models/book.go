package models

import (
	"log"
	"github.com/Jayakumar8055/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB 

type Book struct{
	// gorm.Model 
	// Name string `gorm:""json:"name"`
	// Author string `json:"author"`
	// Publication string `json:"publication"`
	gorm.Model
	Name        string `gorm:"type:varchar(100)" json:"name"`
	Author      string `gorm:"type:varchar(100)" json:"author"`
	Publication string `gorm:"type:varchar(100)" json:"publication"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	// db.AutoMigrate(&Book{})
	if err := db.AutoMigrate(&Book{}).Error; err != nil {
		log.Fatal("Could not migrate database schema:", err)
	}
}

func(b *Book) CreateBook() *Book{
	// db.NewRecord(b)
	// db.Create(&b)
	if err := db.Create(&b).Error; err != nil {
		log.Println("Error creating book:", err)
	}
	return b
}

func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookByID(Id int64) (*Book, *gorm.DB){
	var getBook Book
	//db :=db.Where("ID=?", Id).Find(&getBook)
	db := db.Where("id = ?",Id).Find(&getBook)
	if db.Error != nil {
		log.Println("Error fetching book by ID:", db.Error)
	}
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	//db.Where("ID=?", ID).Delete(book)
	if err := db.Where("id = ?", ID).Delete(&book).Error; err != nil {
		log.Println("Error deleting book:", err)
	}
	return book
}