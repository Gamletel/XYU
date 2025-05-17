package services

import (
	"backend/internal/models"
	"backend/internal/repositories"
)

type TodoService struct {
	todoRepo *repositories.TodoRepository
}

func NewTodoService(todoRepo *repositories.TodoRepository) *TodoService {
	return &TodoService{todoRepo: todoRepo}
}

func (s *TodoService) GetTodoByTitle(title string, userId int) ([]*models.Todo, error) {
	return s.todoRepo.GetTodoByTitle(title, userId)
}

func (s *TodoService) GetTodoByUserId(id int) ([]*models.Todo, error) {
	return s.todoRepo.GetTodoByUserId(id)
}

func (s *TodoService) CreateTodo(todo models.Todo) (*models.Todo, error) {
	return s.todoRepo.CreateTodo(&todo)
}

func (s *TodoService) UpdateTodo(todo models.Todo) (*models.Todo, error) {
	return s.todoRepo.UpdateTodo(&todo)
}

func (s *TodoService) DeleteTodo(id int) (*int, error) {
	return s.todoRepo.DeleteTodo(id)
}
