package service

import (
	"CRUDVk/internal/models"
	"fmt"
	"time"
)

type BookRepository interface {
	CreateBook(book *models.Book) error
	GetBooks() ([]models.Book, error)
	GetBookID(id int) (*models.Book, error)
	UpdateBook(book *models.Book) error
	DeleteBook(id int) error
}

type BookCache interface {
	Set(key string, value interface{}, ttl time.Duration)
	Get(key string) (interface{}, bool)
	Delete(key string)
	PrintCache()
}

type BookService struct {
	repo  BookRepository
	cache BookCache
}

// Дёргать это

func NewBookService(repo BookRepository, cache BookCache) *BookService {
	return &BookService{
		repo:  repo,
		cache: cache,
	}
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

func (s *BookService) GetBooks() ([]models.Book, error) {
	cacheKey := "all_books"

	if cachedBooks, found := s.cache.Get(cacheKey); found {
		return cachedBooks.([]models.Book), nil
	}

	books, err := s.repo.GetBooks()
	if err != nil {
		return nil, err
	}

	s.cache.Set(cacheKey, books, time.Hour*24)
	return books, nil
}

func (s *BookService) GetBookID(id int) (*models.Book, error) {
	cacheKey := cacheKey(id)

	// Чек книги в кэше
	if cachedBook, found := s.cache.Get(cacheKey);
	// Нашлась
		found {
		return cachedBook.(*models.Book), nil
	}
	// Не нашлась, берём из репо
	book, err := s.repo.GetBookID(id)
	// В репо тоже нет, возвращаем ошибку
	if err != nil {
		return nil, err
	}

	s.cache.Set(cacheKey, book, time.Hour*24)
	return book, nil
}

func (s *BookService) DeleteBook(id int) error { return s.repo.DeleteBook(id) }

func cacheKey(id int) string {
	return "book:" + string(id)
}
