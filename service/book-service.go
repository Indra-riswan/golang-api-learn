package service

import (
	"fmt"
	"learn2/dto"
	"learn2/models"
	"learn2/repository"

	"github.com/mashingan/smapping"
)

type BookService interface {
	CreateBook(book dto.DTOBookregister) models.Book
	UpdateBook(book dto.DTOBookUpdate) models.Book
	FindBookById(ID uint) models.Book
	DeleteBook(models.Book)
	AllBook() []models.Book
	AlllowEdited(userID string, bookID uint) bool
}

type bookservice struct {
	repository repository.BookRepository
}

func NewBookService(repo repository.BookRepository) *bookservice {
	return &bookservice{repo}
}

func (s *bookservice) CreateBook(book dto.DTOBookregister) models.Book {
	var books = models.Book{}
	err := smapping.FillStruct(&books, smapping.MapFields(&book))
	if err != nil {
		fmt.Println("Failed Create Book")
	}
	s.repository.InsertBook(books)
	return books
}

func (s *bookservice) UpdateBook(book dto.DTOBookUpdate) models.Book {
	var books = models.Book{}
	err := smapping.FillStruct(&books, smapping.MapFields(&book))
	if err != nil {
		fmt.Println("Failed Update Book")
	}
	s.repository.UpdateBook(books)
	return books
}

func (s *bookservice) FindBookById(ID uint) models.Book {
	return s.repository.FindBookById(ID)
}

func (s *bookservice) DeleteBook(book models.Book) {
	s.repository.DeleteBook(book)
}

func (s *bookservice) AllBook() []models.Book {
	return s.repository.Allbook()
}

func (s *bookservice) AlllowEdited(userID string, bookID uint) bool {
	book := s.repository.FindBookById(bookID)
	id := fmt.Sprintf("%v", book.ID)

	return userID == id

}
