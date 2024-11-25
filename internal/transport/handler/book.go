package handler

import (
	"CRUDVk/internal/models"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type BookService interface {
	CreateBook(book *models.Book) error
	GetBooks() ([]models.Book, error)
	GetBookID(id int) (*models.Book, error)
	UpdateBook(book *models.Book) error
	DeleteBook(id int) error
}

// Сначала CreateBook

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {

	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Ошибка при декодировании, невалидные данные", http.StatusBadRequest) // ошибка 400
		return
	}
	if err := h.BookService.CreateBook(&book); err != nil {
		http.Error(w, "Ошибка при создании книги", http.StatusInternalServerError) // Ошибка 500
		return
	}
	if err := json.NewEncoder(w).Encode(&book); err != nil {
		return
	}
	log.Printf("Книга успешно добавлена")
}

// GetBooks

func (h *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {

	books, err := h.BookService.GetBooks()
	if err != nil {
		http.Error(w, "Ошибка при получении книг", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(&books); err != nil {
		http.Error(w, "Ошибка при кодировании данных", http.StatusInternalServerError)
	}
	log.Printf("Книги успешно экпортированы")
}

// GetBooksID

func (h *BookHandler) GetBookID(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Path[len("/books_get_id/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный формат ID", http.StatusBadRequest)
		return
	}

	book, err := h.BookService.GetBookID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Книга не найдена", http.StatusNotFound)
			return
		}
		http.Error(w, "Ошибка при получении книги", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(&book); err != nil {
		http.Error(w, "Ошибка при кодировании данных", http.StatusInternalServerError)
	}
	log.Printf("Книга с ID %d успешно экспортирована", id)
}

// UpdateBook

func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/books_update/")
	bookID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный формат ID", http.StatusBadRequest)
		return
	}

	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Ошибка при декодировании данных", http.StatusBadRequest)
		return
	}

	book.ID = bookID

	// Валидация входящих данных

	switch {
	case book.Title == "":
		http.Error(w, "Укажите название книги", http.StatusBadRequest)
		return
	case book.Author == "":
		http.Error(w, "Укажите автора книги", http.StatusBadRequest)
		return
	case book.Publisher == "":
		http.Error(w, "Укажите издателя книги", http.StatusBadRequest)
		return
	}

	if err := h.BookService.UpdateBook(&book); err != nil {
		http.Error(w, "Ошибка при обновлении книги. Возможно эта книга не существует", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&book); err != nil {
		http.Error(w, "Ошибка при кодировании данных", http.StatusInternalServerError)
	}
	log.Printf("Книга с ID %d успешно обновлена", bookID)
}

func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/books_delete/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный формат ID", http.StatusBadRequest)
		return
	}

	if err := h.BookService.DeleteBook(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Книга не найдена", http.StatusNotFound)
			return
		}
		http.Error(w, "Ошибка при удалении книги", http.StatusInternalServerError)
	}
	log.Printf("Книга ID %d успешно удалена", id)
	return
}
