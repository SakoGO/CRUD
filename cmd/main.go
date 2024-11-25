package main

import (
	"CRUDVk/internal/repository"
	"CRUDVk/internal/service"
	"CRUDVk/internal/transport"
	"CRUDVk/internal/transport/handler"
	"CRUDVk/pkg/db"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// допиши функцию мейн для всего приложения

// исправь ошибки в коде

func main() {
	// инициализируй БДшку

	dsn := "root:12345@tcp(127.0.0.1:3306)/CRUDDATABASE?utf8mb4&parseTime=True&loc=Local"
	db, err := db.NewGormDB(dsn)
	if err != nil {
		log.Fatalf("Ошибка при подключении к базе данных: %v", err)
	}

	bookRepo, err := repository.NewBookRepository(db)
	if err != nil {
		log.Fatalf("Ошибка при создании репозитория для книг %v", err)
	}

	userRepo, err := repository.NewUserRepository(db)
	if err != nil {
		log.Fatalf("Ошибка при создании репозитория для аккаунтов %v", err)
	}

	bookServ := service.NewBookService(bookRepo)
	userServ := service.NewUserService(userRepo, "jwt_key")

	h := handler.NewHandler(bookServ, userServ)

	mux := h.InitRoutes()
	addr := ":8080"

	s := transport.NewServer(mux, addr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		if err := s.Run(s.Addr); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Ошибка при запуске сервера на порте: %v", err)
		}
	}()

	<-quit
	log.Println("Завершение работы...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("Ошибка при завершении работы сервера: %v", err)
	}
	log.Println("Сервер завершил работу")
}
