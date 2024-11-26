package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGormDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Connect database failed: %v", err)
	}
	return db, nil
}
