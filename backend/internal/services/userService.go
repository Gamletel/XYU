package services

import (
	"backend/internal/models"
	"backend/internal/repositories"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}
func (s *UserService) GetAllUsers() []models.User {
	return s.userRepo.GetAllUsers()
}
func (s *UserService) CreateUser(user *models.User) (*models.User, error) {
	return s.userRepo.CreateUser(user)
}
