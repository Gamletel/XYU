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
func (s *UserService) GetAllUsers() []models.UserResponse {
	return s.userRepo.GetAllUsers()
}
func (s *UserService) GetUserByEmail(email string) (*models.UserResponse, error) {
	return s.userRepo.GetUserByEmail(email)
}
func (s *UserService) CreateUser(user *models.User) (*models.UserResponse, error) {
	return s.userRepo.CreateUser(user)
}
func (s *UserService) UpdateUser(user *models.User) (*models.UserResponse, error) {
	return s.userRepo.UpdateUser(user)
}
func (s *UserService) DeleteUser(id int) (*int, error) {
	return s.userRepo.DeleteUser(id)
}
