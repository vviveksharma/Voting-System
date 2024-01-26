package db

import (
	"database/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	dbPool *gorm.DB
}


func InitDB() (*gorm.DB, error) {
	dsn := "postgres://postgres:password@db:5432/mydatabase?sslmode=disable"
	var err error
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (db *Database)Begin(opts ...*sql.TxOptions) error {
	tx := db.dbPool.Begin(opts...)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}