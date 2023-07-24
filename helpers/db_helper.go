package helpers

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDb() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(os.Getenv("POSTGRES_DSN")), &gorm.Config{})
	if err != nil {
		return &gorm.DB{}, err
	}
	return db, nil
}
