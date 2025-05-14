package repositories

import (
	"backend/internal/db"
	"backend/internal/models"
	"database/sql"
	"errors"
	"fmt"
)

func CreateTodo(todo *models.Todo, userId int) string {
	sqlStatement := `
		INSERT INTO  todos(title, text, image, user_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id`
	id := 0
	err := db.DB.QueryRow(sqlStatement, todo.Title, todo.Text, todo.Image, userId).Scan(&id)

	switch {
	case err == nil:
		return "Todo created successfully"
	case errors.Is(err, sql.ErrNoRows):
		return "Todo does not exist"
	default:
		return err.Error()
	}
}

func UpdateTodo(todo *models.Todo) string {
	sqlStatement := `
		UPDATE todos
		SET title = $2, text = $3, image = $4
		WHERE id = $1`

	res, err := db.DB.Exec(sqlStatement, todo.Id, todo.Title, todo.Text, todo.Image)
	if err != nil {
		return err.Error()
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err.Error()
	}
	if count == 0 {
		return fmt.Sprintf("Todo with id %v does not exist", todo.Id)
	}
	return fmt.Sprintf("Todo with id %v updated", todo.Id)
}

func DeleteTodo(id int) string {
	sqlStatement := `DELETE FROM todos WHERE id = $1`
	res, err := db.DB.Exec(sqlStatement, id)
	if err != nil {
		return err.Error()
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err.Error()
	}
	if count == 0 {
		return fmt.Sprintf("Todo with id %v does not exist", id)
	}
	return fmt.Sprintf("Todo with id %v deleted", id)
}

func GetTodoByTitle(title string, userId int) (*models.Todo, error) {
	sqlStatement := "SELECT id, title, text, image, created_at FROM todos WHERE title = $1 and user_id = $2"
	row := db.DB.QueryRow(sqlStatement, title, userId)

	var todo models.Todo
	if err := row.Scan(todo.Id, todo.Title, todo.Text, todo.Image); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("todo not found: %w", err)
		}
		return nil, fmt.Errorf("failed to scan todo: %w", err)
	}
	return &todo, nil
}

func GetTodosByUserId(userId int) ([]models.Todo, error) {
	sqlStatement := `SELECT id, title, text, image, created_at FROM todos WHERE user_id = $1`
	rows, err := db.DB.Query(sqlStatement, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var todos []models.Todo

	for rows.Next() {
		var todo models.Todo

		err := rows.Scan(&todo.Id, &todo.Title, &todo.Text, &todo.Image, &todo.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred during iteration: %w", err)
	}

	return todos, nil
}
