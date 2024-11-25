package repository

import (
	models "CRUDVk/internal/models"
	"gorm.io/gorm"
)

type UserRepositoryStr struct {
	db *gorm.DB
}

// Конструктор
func NewUserRepository(db *gorm.DB) (*UserRepositoryStr, error) {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}
	return &UserRepositoryStr{db: db}, nil
}

// TODO: UserCreate

func (r *UserRepositoryStr) UserCreate(user *models.User) error {
	return r.db.Create(user).Error
}

// TODO: Find by username

func (r *UserRepositoryStr) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// TODO: Find by email

func (r *UserRepositoryStr) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// TODO: UserDeleteAccount
