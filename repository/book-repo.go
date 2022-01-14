package repository

import (
	"learn2/models"

	"gorm.io/gorm"
)

type BookRepository interface {
	InsertBook(book models.Book) models.Book
	UpdateBook(book models.Book) models.Book
	DeleteBook(book models.Book)
	Allbook() []models.Book
	FindBookById(ID uint) models.Book
}

type bookrepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *bookrepository {
	return &bookrepository{db}
}

func (r *bookrepository) InsertBook(book models.Book) models.Book {

	r.db.Save(&book)
	r.db.Preload("User").Find(&book)
	return book
}
func (r *bookrepository) UpdateBook(book models.Book) models.Book {
	r.db.Save(&book)
	r.db.Preload("User").Find(&book)
	return book
}
func (r *bookrepository) DeleteBook(book models.Book) {
	r.db.Delete(&book)
}
func (r *bookrepository) Allbook() []models.Book {
	var books []models.Book
	r.db.Preload("User").Find(&books)
	return books
}
func (r *bookrepository) FindBookById(ID uint) models.Book {
	var bookid models.Book
	r.db.Preload("User").Find(&bookid, ID)
	return bookid
}
