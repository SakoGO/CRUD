package handler

import (
	"CRUDVk/internal/models"
	"encoding/json"
	"net/http"
)

type UserService interface {
	SignUp(username, email, password string) error
	SignIn(username, password string) (string, error)
}

func (h *BookHandler) UserCreate(w http.ResponseWriter, r *http.Request) {

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Неверный формат данных", http.StatusBadRequest) // ошибка 400
		return
	}

	err := h.UserService.SignUp(user.Username, user.Email, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(map[string]string{
		"message": "Пользователь успешно зарегистрирован"})
	if err != nil {
		return
	}
}

func (h *BookHandler) UserLogin(w http.ResponseWriter, r *http.Request) {

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Неверный формат данных", http.StatusBadRequest) // ошибка 400
		return
	}

	token, err := h.UserService.SignIn(user.Username, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(map[string]string{
		"token": token})
	if err != nil {
		return
	}

}
