package database

import (
	"fmt"
	"github.com/dev-oleksandrv/taskera/internal/config"
	"github.com/dev-oleksandrv/taskera/internal/model/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Init(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.Database.Host,
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if cfg.Server.Env == "development" {
		db.Debug()
	}
	if err := db.AutoMigrate(&domain.User{}); err != nil {
		log.Fatalf("Error migrating database: %s", err)
	}
	return db, nil
}
