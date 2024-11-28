package handler

import (
	"CRUDVk/internal/transport/middleware"
	"net/http"
)

type BookHandler struct {
	BookService BookService
	UserService UserService
	keyJWT      string
	cache       CacheHandler
}

// Связь с внешним миром, конструктор

func NewHandler(BookService BookService, UserService UserService, keyJWT string, cache CacheHandler) *BookHandler {
	return &BookHandler{
		BookService: BookService,
		UserService: UserService,
		keyJWT:      keyJWT,
		cache:       cache,
	}
}

func (h *BookHandler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/books_get/", h.GetBooks)
	mux.HandleFunc("/books_get_id/{id}", h.GetBookID)
	mux.HandleFunc("/cache_print", h.PrintCache)

	mux.HandleFunc("/user_create", h.UserCreate)
	mux.HandleFunc("/user_login", h.UserLogin)

	mux.HandleFunc("/books_add", middleware.JWTMiddleware(h.keyJWT, h.CreateBook))
	mux.HandleFunc("/books_update/{id}", middleware.JWTMiddleware(h.keyJWT, h.UpdateBook))
	mux.HandleFunc("/books_delete/{id}", middleware.JWTMiddleware(h.keyJWT, h.DeleteBook))

	return mux
}
