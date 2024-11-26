package service

import (
	"CRUDVk/internal/models"
	"fmt"
)

type BookRepository interface {
	CreateBook(book *models.Book) error
	GetBooks() ([]models.Book, error)
	GetBookID(id int) (*models.Book, error)
	UpdateBook(book *models.Book) error
	DeleteBook(id int) error
}

type BookService struct {
	repo BookRepository
}

// Дёргать это

func NewBookService(repo BookRepository) *BookService {
	return &BookService{repo}
}

func (s *BookService) CreateBook(book *models.Book, userID int) error {
	book.UserID = userID
	return s.repo.CreateBook(book)
}
func (s *BookService) UpdateBook(book *models.Book) error {

	if book.ID == 0 {
		return fmt.Errorf("невозможно обновить книгу без указания ID")
	}

	existingBook, err := s.repo.GetBookID(book.ID)
	if err != nil {
		return err
	}

	if book.Title != "" {
		existingBook.Title = book.Title
	}
	if book.Publisher != "" {
		existingBook.Publisher = book.Publisher
	}

	return s.repo.UpdateBook(existingBook)
}

func (s *BookService) GetBooks() ([]models.Book, error) { return s.repo.GetBooks() }

func (s *BookService) GetBookID(id int) (*models.Book, error) { return s.repo.GetBookID(id) }

func (s *BookService) DeleteBook(id int) error { return s.repo.DeleteBook(id) }
