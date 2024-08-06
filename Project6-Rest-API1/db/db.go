package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	return gorm.Open(postgres.Open("postgres://postgres:123456@localhost:5432/rest_api"), &gorm.Config{})
}
