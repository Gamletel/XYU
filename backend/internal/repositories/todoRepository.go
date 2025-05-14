package repositories

import (
	"backend/internal/db"
	"backend/internal/models"
	"database/sql"
	"errors"
	"fmt"
)

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) GetTodoByUserId(userId int) ([]*models.Todo, error) {
	sqlStatement := `SELECT id, title, text, image, created_at FROM todos WHERE user_id = $1`
	rows, err := db.DB.Query(sqlStatement, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var todos []*models.Todo

	for rows.Next() {
		var todo models.Todo

		err := rows.Scan(&todo.Id, &todo.Title, &todo.Text, &todo.Image, &todo.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		todos = append(todos, &todo)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred during iteration: %w", err)
	}

	return todos, nil
}

func (r *TodoRepository) GetTodoByTitle(title string, userId int) ([]*models.Todo, error) {
	sqlStatement := `SELECT id, title, text, image, created_at FROM todos WHERE title = $1 AND user_id = $2`
	rows, err := db.DB.Query(sqlStatement, title, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var todos []*models.Todo

	for rows.Next() {
		var todo models.Todo

		err := rows.Scan(&todo.Id, &todo.Title, &todo.Text, &todo.Image, &todo.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		todos = append(todos, &todo)
	}

	return todos, nil
}

func (r *TodoRepository) CreateTodo(todo *models.Todo) (*models.Todo, error) {
	sqlStatement := `
		INSERT INTO  todos(title, text, image, user_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id, title, text, image, user_id`
	var result models.Todo
	err := db.DB.
		QueryRow(sqlStatement, todo.Title, todo.Text, todo.Image, todo.UserId).
		Scan(&result.Id, &result.Title, &result.Text, &result.Image, &result.UserId)

	switch {
	case err == nil:
		return &result, nil
	case errors.Is(err, sql.ErrNoRows):
		return nil, errors.New(fmt.Sprintf("Todo with id %d does not exist", todo.Id))
	default:
		return nil, err
	}
}

func (r *TodoRepository) UpdateTodo(todo *models.Todo) (*models.Todo, error) {
	sqlStatement := `
		UPDATE todos
		SET title = $2, text = $3, image = $4
		WHERE id = $1
		RETURNING id, title, text, image, user_id`

	var result models.Todo
	err := db.DB.QueryRow(sqlStatement, todo.Id, todo.Title, todo.Text, todo.Image).Scan(&result.Id, &result.Title, &result.Text, &result.Image, &result.UserId)

	switch {
	case err == nil:
		return &result, nil
	case errors.Is(err, sql.ErrNoRows):
		return nil, errors.New(fmt.Sprintf("Todo with id %d does not exist", todo.Id))
	default:
		return nil, err
	}
}

func (r *TodoRepository) DeleteTodo(id int) (*int, error) {
	sqlStatement := `DELETE FROM todos WHERE id = $1 RETURNING id`
	var result int
	err := db.DB.QueryRow(sqlStatement, id).Scan(&result)

	if err != nil {
		return nil, err
	}
	return &result, nil
}
