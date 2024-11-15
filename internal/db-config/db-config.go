package config

import (
	models "backend/pkg/model"
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Config holds the database configuration details
type Config struct {
    Host     string
    Port     int
    User     string
    Password string
    DBName   string
    SSLMode  string
}

var (
    db   *gorm.DB
    once sync.Once
)

// InitializeDB initializes the database connection and performs migrations
func InitializeDB(cfg Config) *gorm.DB {
    once.Do(func() {
        dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
            cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)

        var err error
        db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
            Logger: logger.Default.LogMode(logger.Info), // Adjust log level as needed
        })
        if err != nil {
            log.Fatalf("Failed to connect to database: %v", err)
        }
        log.Println("Database connection established.")

        // Run migrations
        MigrateDB()
    })
    return db
}

// MigrateDB migrates all necessary models to the database
func MigrateDB() {
    if db == nil {
        log.Fatal("Database not initialized. Call InitializeDB first.")
    }

    // Run migrations for all models
    err := db.AutoMigrate(&models.User{}, &models.Note{}, &models.PreviousYearQuestion{}, &models.Book{}, &models.Classroom{})
    if err != nil {
        log.Fatalf("Failed to run migrations: %v", err)
    }
    log.Println("Database migrations completed.")
}

// GetDB returns the instance of the initialized database
func GetDB() *gorm.DB {
    if db == nil {
        log.Fatal("Database not initialized. Call InitializeDB first.")
    }
    return db
}
