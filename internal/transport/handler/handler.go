package handler

import (
	"net/http"
)

type BookHandler struct {
	BookService BookService
	UserService UserService
}

// Связь с внешним миром, конструктор

func NewHandler(BookService BookService, UserService UserService) *BookHandler {
	return &BookHandler{
		BookService: BookService,
		UserService: UserService,
	}
}

func (h *BookHandler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/books_add", h.CreateBook)
	mux.HandleFunc("/books_get/", h.GetBooks)
	mux.HandleFunc("/books_get_id/{id}", h.GetBookID)
	mux.HandleFunc("/books_update/{id}", h.UpdateBook)
	mux.HandleFunc("/books_delete/{id}", h.DeleteBook)
	mux.HandleFunc("/user_create", h.UserCreate)
	mux.HandleFunc("/user_login", h.UserLogin)
	return mux
}
