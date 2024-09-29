package controller

import (
	"net/http"

	usermodel "github.com/AndresBC-Dev/kubera-ms/auth-service/entity/user-model"
	userservice "github.com/AndresBC-Dev/kubera-ms/auth-service/service/user-service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService userservice.UserService // Corrected spelling
}

func New(us userservice.UserService) *UserController {
	return &UserController{
		userService: us, // Corrected spelling
	}
}

func (c *UserController) CreateUserHandler(ctx *gin.Context) {
	var user usermodel.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Proceed with creating user
	err := c.userService.Create(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user created"})
}
