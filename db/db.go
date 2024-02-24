package db

import (
	"os"
	"voting-system/comman"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() (*gorm.DB, error) {
	err := comman.Getenv()
	if err != nil {
		return nil, err
	}
	dsn := os.Getenv("DSN")
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	return conn, nil
}
