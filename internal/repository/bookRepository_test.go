package repository

import (
	models "CRUDVk/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestBookRepository_GetBooks(t *testing.T) {

	dsn := "root:12345@tcp(127.0.0.1:3306)/CRUDDATABASE?utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	assert.Nil(t, err)

	b, err := NewBookRepository(db)
	assert.Nil(t, err)

	var existBooks []*models.Book
	err = db.Find(&existBooks).Error
	assert.Nil(t, err)
	if len(existBooks) == 0 {
		books := []*models.Book{
			{
				Title:     "ALFA",
				Author:    "CRYPT",
				Publisher: "S",
			},
			{
				Title:     "BORYAN",
				Author:    "CRYPT1",
				Publisher: "ASS",
			},
		}

		err = db.Create(&books).Error
		assert.Nil(t, err)

		existBooks = books
	}

	repoBooks, err := b.GetBooks()
	assert.Nil(t, err)
	assert.Equal(t, len(existBooks), len(repoBooks))

}

func TestBookRepository_GetBookID(t *testing.T) {

	dsn := "root:12345@tcp(127.0.0.1:3306)/CRUDDATABASE?utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	assert.Nil(t, err)

	b, err := NewBookRepository(db)
	assert.Nil(t, err)

	var existBook *models.Book
	err = db.First(&existBook).Error
	assert.Nil(t, err)

	// Проверка несуещствующей в бд книги
	ifIDInvalid := 99999999999
	repoBook, err := b.GetBookID(ifIDInvalid)
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
	assert.Nil(t, repoBook)

	// Проверка что при id 0 выдаст ошибку
	ifIDZero := 0
	repoBookId, err := b.GetBookID(ifIDZero)
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
	assert.Nil(t, repoBookId)

	// Проверка что при отрицательном id выдаст ошибку
	ifIDNegative := -1
	repoBookIdNegative, err := b.GetBookID(ifIDNegative)
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
	assert.Nil(t, repoBookIdNegative)

}

func TestBookRepository_UpdateBook(t *testing.T) {

	dsn := "root:12345@tcp(127.0.0.1:3306)/CRUDDATABASE?utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	assert.Nil(t, err)

	b, err := NewBookRepository(db)
	assert.Nil(t, err)

	var existBook *models.Book
	err = db.First(&existBook).Error
	assert.Nil(t, err)

	// Проверка несуществующей в бд книги

	ifIDInvalid := 99999999999
	bookToUpdate := &models.Book{
		ID:        ifIDInvalid,
		Title:     "TestTitle",
		Author:    "TestAuthor",
		Publisher: "TestPublisher",
	}
	err = b.UpdateBook(bookToUpdate)
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)

	// Проверка что при id 0 выдаст ошибку
	ifIDZero := 0
	bookToUpdateZero := &models.Book{
		ID:        ifIDZero,
		Title:     "TestTitle",
		Author:    "TestAuthor",
		Publisher: "TestPublisher",
	}
	err = b.UpdateBook(bookToUpdateZero)
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)

	// Проверка что при отрицательном id выдаст ошибку
	IfIDNegative := -1

	bookToUpdateNegative := &models.Book{
		ID:        IfIDNegative,
		Title:     "TestTitle",
		Author:    "TestAuthor",
		Publisher: "TestPublisher",
	}
	err = b.UpdateBook(bookToUpdateNegative)
	assert.NotNil(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}
