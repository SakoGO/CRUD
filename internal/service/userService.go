package service

import (
	"CRUDVk/internal/models"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	UserCreate(user *models.User) error
	FindByUsername(username string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	//Register(username, email, password string) (models.User, error)
	//Login(username, password string) (string, error)
}

type UserService struct {
	repo   UserRepository
	keyJWT string
}

// Конструктор
func NewUserService(repo UserRepository, keyJWT string) *UserService {
	return &UserService{repo: repo, keyJWT: keyJWT}
}

func (s *UserService) UserCreate(user *models.User) error { return s.repo.UserCreate(user) }

// USER REGISTRATION FUNCTION
func (s *UserService) SignUp(username, email, password string) error {
	_, err := s.repo.FindByUsername(username)
	if err == nil {
		return fmt.Errorf("пользователь %s уже зарегистрирован ", username)
	}

	_, err = s.repo.FindByEmail(email)
	if err == nil {
		return fmt.Errorf("пользователь %s уже зарегистрирован", email)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	return s.repo.UserCreate(user)
}

// TODO: AUTH USER FUNCTION

func (s *UserService) SignIn(username, password string) (string, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return "", fmt.Errorf("пользователь %s не зарегистрирован", username)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("неправильный пароль")
	}

	token, err := s.generateJWTToken(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *UserService) generateJWTToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(s.keyJWT))
	if err != nil {
		return "", fmt.Errorf("не удалось создать токен: %v", err)
	}
	return tokenString, nil
}

//TODO: Проверить, существует ли пользователем с таким именем. Если его нет, то выдать ошибку fmt.Errorf

//TODO: Раскриптовываем пароль, сравниваем в формате ([]byte(user.Password), []byte(password)). Выдаем ошибку, если пароль неверный

//TODO: Генерация JWT токена отдельной функцией. В ней же подписываем его с использованием keyJWT
