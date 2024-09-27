package userrepository

import (
	"log"

	usermodel "github.com/AndresBC-Dev/kubera-ms/auth-service/entity/user-model"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	Create(u *usermodel.User) error
	FindUser(email string) (*usermodel.User, error)
}

type UserRepository struct {
	DB *gorm.DB
}

func New(db *gorm.DB) UserRepositoryInterface {
	return &UserRepository{
		DB: db,
	}
}

func (UR *UserRepository) Create(u *usermodel.User) error {
	result := UR.DB.Create(u)
	if result.Error != nil {
		log.Printf("Error creating user: %v\n", result.Error)
		return result.Error
	}
	log.Println("User created successfully")
	return nil
}

func (UR *UserRepository) FindUser(email string) (*usermodel.User, error) {
	var user usermodel.User

	err := UR.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		log.Printf("User not found: %v\n", err)
		return nil, err
	}
	return &user, nil
}
