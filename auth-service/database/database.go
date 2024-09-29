package database

import (
	"fmt"
	"log"
	"time"

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
	var connection *gorm.DB
	var err error
	retries := 5

	for retries > 0 {
		connection, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			return connection, nil // Conexión exitosa, devolver el objeto DB
		}
		log.Printf("Cannot open db connection: %v. Retrying...", err)
		retries--
		time.Sleep(2 * time.Second) // Espera de 2 segundos antes del siguiente intento
	}

	return nil, fmt.Errorf("failed to connect to the database after retries: %v", err)
}
