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
	CreateBook(book *models.Book, authorID int) error
	GetBooks() ([]models.Book, error)
	GetBookID(id int) (*models.Book, error)
	UpdateBook(book *models.Book) error
	DeleteBook(id int) error
}

// Сначала CreateBook

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		http.Error(w, "Неверный ID пользователя", http.StatusUnauthorized)
		return
	}

	user, err := h.UserService.GetUserByID(userID)
	if err != nil {
		http.Error(w, "Пользователь не найден", http.StatusUnauthorized)
	}

	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Ошибка при декодировании, невалидные данные", http.StatusBadRequest) // ошибка 400
		return
	}

	book.Author = user.Username
	user.UserID = userID

	if err := h.BookService.CreateBook(&book, userID); err != nil {
		http.Error(w, "Ошибка при создании книги", http.StatusInternalServerError) // Ошибка 500
		return
	}
	if err := json.NewEncoder(w).Encode(&book); err != nil {
		return
	}
	log.Printf("Книга успешно добавлена пользователем %s", user.Username)
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

	var updatedBook models.Book
	if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
		http.Error(w, "Ошибка при декодировании данных", http.StatusBadRequest)
		return
	}

	existingBook, err := h.BookService.GetBookID(bookID)
	if err != nil {
		http.Error(w, "Книга не найдена", http.StatusNotFound)
		return
	}
	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		http.Error(w, "Неверный ID пользователя", http.StatusUnauthorized)
		return
	}

	user, err := h.UserService.GetUserByID(userID)
	if err != nil {
		http.Error(w, "Пользователь не найден", http.StatusUnauthorized)
		return
	}

	if existingBook.Author != user.Username {
		http.Error(w, "У вас недостаточно прав для обновления этой книги", http.StatusForbidden)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&updatedBook); err != nil {
		http.Error(w, "Ошибка при кодировании данных", http.StatusInternalServerError)
	}
	log.Printf("Книга № %d успешно обновлена пользователем %s", bookID, user.Username)
}

func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/books_delete/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный формат ID", http.StatusBadRequest)
		return
	}
	existingBook, err := h.BookService.GetBookID(id)
	if err != nil {
		http.Error(w, "Книга не найдена", http.StatusNotFound)
		return
	}
	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		http.Error(w, "Неверный ID пользователя", http.StatusUnauthorized)
		return
	}

	user, err := h.UserService.GetUserByID(userID)
	if err != nil {
		http.Error(w, "Пользователь не найден", http.StatusUnauthorized)
		return
	}

	if existingBook.Author != user.Username {
		http.Error(w, "У вас недостаточно прав для удаления этой книги", http.StatusForbidden)
		return
	}

	if err := h.BookService.DeleteBook(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Книга не найдена", http.StatusNotFound)
			return
		}
		http.Error(w, "Ошибка при удалении книги", http.StatusInternalServerError)
	}
	log.Printf("Книга № %d успешно удалена пользователем %s", id, user.Username)
	return
}
