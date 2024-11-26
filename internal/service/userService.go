package service

import (
	"CRUDVk/internal/models"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"time"
)

type UserRepository interface {
	UserCreate(user *models.User) error
	FindByUsername(username string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindByID(userID int) (*models.User, error)
}

type UserService struct {
	repo UserRepository
}

func (s *UserService) FindByUsername(username string) (*models.User, error) {
	return s.repo.FindByUsername(username)
}

// Конструктор
func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
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
	keyJWT := GetJWTKey()

	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(keyJWT))
	if err != nil {
		return "", fmt.Errorf("не удалось создать токен: %v", err)
	}
	return tokenString, nil
}

func GetJWTKey() string {
	keyJWT := os.Getenv("JWT_SECRET_KEY")
	fmt.Println(len(keyJWT))
	if keyJWT == "" {
		log.Fatal("JWT секретка не задана")
	}
	return keyJWT
}

func (s *UserService) GetUserByID(userID int) (*models.User, error) {
	return s.repo.FindByID(userID)
}
