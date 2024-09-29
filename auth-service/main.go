package main

import (
	"log"

	"github.com/AndresBC-Dev/kubera-ms/auth-service/container"
)

func main() {

	dsn := "host=postgres user=geckode password=s3cr3ts dbname=auth_db port=5432 sslmode=disable TimeZone=America/Bogota"

	cont, err := container.NewContainer(dsn)
	if err != nil {
		log.Fatalf("Error initializating container_config: %v", err)
	}

	cont.Router.POST("/user", cont.UserController.CreateUserHandler)

	err = cont.Router.Run(":8081")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
