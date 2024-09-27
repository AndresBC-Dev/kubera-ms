package userservice

import (
	"fmt"

	usermodel "github.com/AndresBC-Dev/kubera-ms/auth-service/entity/user-model"
	userrepository "github.com/AndresBC-Dev/kubera-ms/auth-service/repository/user-repository"
)

type UserService struct {
	UR userrepository.UserRepositoryInterface
}

func New(UR userrepository.UserRepositoryInterface) *UserService {
	return &UserService{
		UR: UR,
	}
}
func (US UserService) Create(user *usermodel.User) error {
	email := user.Email
	existingUser, err := US.UR.FindUser(email)
	if err != nil {
		return US.UR.Create(user)
	}

	return fmt.Errorf("user with email %ss already exist", existingUser.Email)
}
