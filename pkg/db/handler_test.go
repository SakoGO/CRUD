package db

/*
import (
	models "CRUDVk/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestBookRepository_UpdateBookValidating(t *testing.T) {
	dsn := "root:12345@tcp(127.0.0.1:3306)/CRUDDATABASE?utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	assert.Nil(t, err)

	b, err := NewBookRepository(db)
	assert.Nil(t, err)

	var existBook *models.Book
	err = db.First(&existBook).Error

	if len(existBook.Title) == 0 {
		book := &models.Book{
			Title:     "ALFAtest",
			Author:    "CRYPTtest",
			Publisher: "Stest",
		}
		err = db.Create(&book).Error
		assert.Nil(t, err)

		existBook = book
	}

	// Проверка апдейта, если не указано название книги
	InvalidBookTitle := &models.Book{
		ID:        existBook.ID,
		Title:     "",
		Author:    "TestAuthor",
		Publisher: "TestPublisher",
	}
	err = b.UpdateBook(InvalidBookTitle)
	assert.NotNil(t, err)
	assert.Equal(t, "Укажите название книги", err.Error())

	// Проверка апдейта, если не указан автор книги
	InvalidBookAuthor := &models.Book{
		ID:        existBook.ID,
		Title:     "TestTitle",
		Author:    "",
		Publisher: "TestPublisher",
	}
	err = b.UpdateBook(InvalidBookAuthor)
	assert.NotNil(t, err)
	assert.Equal(t, "Укажите автора книги", err.Error())

	// Проверка апдейта, если не указан издатель книги
	InvalidBookPublisher := &models.Book{
		ID:        existBook.ID,
		Title:     "TestTitle",
		Author:    "TestAuthor",
		Publisher: "",
	}
	err = b.UpdateBook(InvalidBookPublisher)
	assert.NotNil(t, err)
	assert.Equal(t, "Укажите издателя книги", err.Error())
}
*/
