package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func NewGormDB(dsn string) (*gorm.DB, error) {
	dsn = "root:12345@tcp(127.0.0.1:3306)/CRUDDATABASE?utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Connect database failed: %v", err)
	}
	return db, nil
}
