package services

import (
	"fyno/server/internal/models"

	"github.com/google/uuid"
)

type userService struct {
	userRepository models.UserRepository
}

func NewUserService(ur models.UserRepository) models.UserService {
	return &userService{
		userRepository: ur,
	}
}

func (us *userService) GetUser(id uuid.UUID) (*models.User, error) {
	return us.userRepository.Get(id)
}

func (us *userService) GetUserByName(name string) (*models.User, error) {
	return us.userRepository.GetByName(name)
}

func (us *userService) CreateUser(u *models.User) (*models.User, error) {
	return us.userRepository.Create(u)
}

func (us *userService) DeleteAllUsers() error {
	return us.userRepository.DeleteAll()
}
