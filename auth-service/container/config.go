package container

import (
	"github.com/AndresBC-Dev/kubera-ms/auth-service/controller"
	"github.com/AndresBC-Dev/kubera-ms/auth-service/database"
	userrepository "github.com/AndresBC-Dev/kubera-ms/auth-service/repository/user-repository"
	userservice "github.com/AndresBC-Dev/kubera-ms/auth-service/service/user-service"
	"github.com/gin-gonic/gin"
)

type Container struct {
	UserController *controller.UserController
	Router         *gin.Engine
}

// Inicializa el contenedor con todas las dependencias
func NewContainer(dsn string) (*Container, error) {
	// Crear la conexión a la base de datos
	db, err := database.CreateConnection(dsn)
	if err != nil {
		return nil, err
	}

	// Crear el repositorio
	userRepo := userrepository.New(db)
	// Crear el servicio
	userServ := userservice.New(userRepo)

	// Crear el controlador
	userCtrl := controller.New(*userServ)

	// Crear el router de Gin
	router := gin.Default()

	// Devolver el contenedor con todas las dependencias
	return &Container{
		UserController: userCtrl,
		Router:         router,
	}, nil
}
