package database

import (
	"fmt"
	"log"

	"github.com/cmrdSurajYadav/auth-service/internal/config"
	"github.com/cmrdSurajYadav/auth-service/internal/modules/auth/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DBHost, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBPort, cfg.DBSSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to DB:", err)
	}

	// Auto migrate User model
	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatal("DB migration failed:", err)
	}

	return db
}
