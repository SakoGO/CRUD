package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func JWTMiddleware(keyJWT string, next http.HandlerFunc) http.HandlerFunc {

	// Проверка наличия хедера авторизации
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Отсутствует хедер авторизации", http.StatusUnauthorized)
			return
		}

		// Разбивка хедера на 2 части (bearer token)
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
			http.Error(w, "Невалидный формат токена", http.StatusUnauthorized)
			return
		}

		// Парсинг и валидация токена
		token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
			return []byte(keyJWT), nil
		})

		// Я ЗАТРАХАЛСЯ СУКА ЛОГИ ПОКАЖИТЕ ЧТО ДЕЛАТЬ НАДО
		if err != nil {
			log.Printf("Ошибка при парсинге ЁБАНОГО токена: %v", err)
			http.Error(w, "Невалидный токен БЛЯДСКОЙ авторизации", http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			// Логируем проблему с валидностью токена
			log.Printf("Токен невалиден, ПОТОМУ ЧТО ОН ГАНДОН: %v", token)
			http.Error(w, "невалидный токен авторизации, ХУЕСОС ЕБЛИВЫЙ", http.StatusUnauthorized)
			return
		}

		// Проверка наличия ошибок при парсинге
		if err != nil || !token.Valid {
			http.Error(w, "невалидный токен авторизации", http.StatusUnauthorized)
			return
		}

		// Извлекаем claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "невалидный пейлод токена", http.StatusUnauthorized)
			return
		}

		// Получение id пользователя из токена
		userID, ok := claims["sub"].(float64)
		if !ok {
			http.Error(w, "невалидный ID пользователя", http.StatusUnauthorized)
			return
		}

		// Добавляем ID пользователя в контекст для получения ID в следующий обработчиках
		ctx := context.WithValue(r.Context(), "userID", int(userID))

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
