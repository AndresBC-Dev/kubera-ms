package main

import (
	"github.com/AndresBC-Dev/kubera-ms/auth-service/controller"
	"github.com/AndresBC-Dev/kubera-ms/auth-service/database"
)

type Config struct {
	userController controller.UserController
	dbConnection   database.PostgreSQL
}

func New(UC controller.UserController, dbConnection database.PostgreSQL) *Config {
	return &Config{
		userController: UC,
		dbConnection:   dbConnection,
	}
}
