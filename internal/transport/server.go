package transport

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
	mux        *http.ServeMux
	Addr       string
}

// Дёргать отсюда. Конструктор
func NewServer(mux *http.ServeMux, addr string) *Server {
	return &Server{
		mux:  mux,
		Addr: addr,

		httpServer: &http.Server{
			Addr:    addr,
			Handler: mux,
		},
	}
}

// Сам сервак
func (s *Server) Run(addr string) error {
	logAddress(s.Addr)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return s.httpServer.Shutdown(ctx)
}

func logAddress(addr string) {
	println("Сервер запущен на порте", addr)
}
