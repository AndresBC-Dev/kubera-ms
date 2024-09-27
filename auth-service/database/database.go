package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgreSQL struct {
	dsn string
}

func New(dsn string) *PostgreSQL {
	return &PostgreSQL{
		dsn: dsn,
	}
}

func CreateConnection(dsn string) (*gorm.DB, error) {
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Cannot open db connection: %v", err)
		return nil, err
	}
	return connection, nil
}
