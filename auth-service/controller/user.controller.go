package controller

import (
	"encoding/json"
	"log"

	usermodel "github.com/AndresBC-Dev/kubera-ms/auth-service/entity/user-model"
	userservice "github.com/AndresBC-Dev/kubera-ms/auth-service/service/user-service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userSevice userservice.UserService
}

func New(us userservice.UserService) *UserController {
	return &UserController{
		userSevice: us,
	}
}

func (UC *UserController) CreateUserHandler(c *gin.Context) {
	userRequest, err := c.Request.GetBody()
	if err != nil {
		log.Printf("Error getting body: %v", err)
		c.JSON(500, gin.H{"error": "Unable to read request body"})
		return
	}
	defer userRequest.Close()

	var userCleaned usermodel.User

	err = json.NewDecoder(userRequest).Decode(&userCleaned)
	if err != nil {
		log.Printf("Cannot format the user to user: %v", err)
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	err = UC.userSevice.Create(&userCleaned)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		c.JSON(500, gin.H{"error": "Unable to create user"})
		return
	}

	c.JSON(201, gin.H{"message": "User created successfully"}) // Respuesta exitosa
}
