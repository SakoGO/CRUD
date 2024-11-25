package repository

import (
	"CRUDVk/internal/models"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

// Конструктор
func NewBookRepository(db *gorm.DB) (*BookRepository, error) {

	err := db.AutoMigrate(&models.Book{})
	if err != nil {
		return nil, err
	}
	return &BookRepository{db: db}, nil
}

func (r *BookRepository) CreateBook(book *models.Book) error {
	return r.db.Create(book).Error
}

func (r *BookRepository) GetBooks() ([]models.Book, error) {
	var books []models.Book
	err := r.db.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (r *BookRepository) GetBookID(id int) (*models.Book, error) {
	var book models.Book
	err := r.db.First(&book, id).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *BookRepository) UpdateBook(book *models.Book) error {
	var update models.Book
	err := r.db.First(&update, book.ID).Error
	if err != nil {
		return err
	}

	update.Title = book.Title
	update.Author = book.Author
	update.Publisher = book.Publisher

	return r.db.Save(&update).Error
}

func (r *BookRepository) DeleteBook(id int) error {
	var book models.Book

	errF := r.db.First(&book, id).Error
	if errF != nil {
		return errF
	}

	errD := r.db.Delete(&book).Error
	if errD != nil {
		return errD
	}
	return nil

}
