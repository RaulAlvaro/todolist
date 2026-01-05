package db

import (
	"todolist/config"
	"todolist/internal/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ProvideDatabase(cfg *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.DBDSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Añade esto para que no tengas que preocuparte por las tablas:
	db.AutoMigrate(&domain.Todo{})

	return db, nil
}
