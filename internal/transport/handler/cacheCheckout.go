package handler

import (
	"net/http"
)

type CacheHandler interface {
	PrintCache()
}

func (h *BookHandler) PrintCache(w http.ResponseWriter, r *http.Request) {
	h.cache.PrintCache()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Содержимое кэша выведено в консоль"))
}
