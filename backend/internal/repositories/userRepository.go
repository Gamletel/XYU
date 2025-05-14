package repositories

import (
	"backend/internal/db"
	"backend/internal/models"
	"database/sql"
	"errors"
	"fmt"
)

func CreateUser(user *models.User) string {
	sqlStatement := `
		INSERT INTO users(email, login, name, surname, avatar, password)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`
	id := 0
	err := db.DB.QueryRow(sqlStatement, user.Email, user.Login, user.Name, user.Surname, user.Avatar, user.Password).Scan(&id)

	switch {
	case err == nil:
		return "User created successfully"
	case errors.Is(err, sql.ErrNoRows):
		return "User does not exist"
	default:
		return err.Error()
	}
}

func UpdateUser(user *models.User) string {
	sqlStatement := `
		UPDATE users
		SET email = $2, name = $3, surname = $4, avatar = $5
		WHERE id = $1`

	res, err := db.DB.Exec(sqlStatement, user.Id, user.Email, user.Name, user.Surname, user.Avatar)
	if err != nil {
		return err.Error()
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err.Error()
	}
	if count == 0 {
		return fmt.Sprintf("User with id %v does not exist", user.Id)
	}
	return fmt.Sprintf("User with id %v updated", user.Id)
}

func DeleteUser(id int) string {
	sqlStatement := `DELETE FROM users WHERE id = $1`
	res, err := db.DB.Exec(sqlStatement, id)
	if err != nil {
		return err.Error()
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err.Error()
	}
	if count == 0 {
		return fmt.Sprintf("User with id %v does not exist", id)
	}
	return fmt.Sprintf("User with id %v deleted", id)
}

func GetUserByEmail(email string) (*models.User, error) {
	sqlStatement := "SELECT id, email, name, surname FROM users WHERE email = $1"
	row := db.DB.QueryRow(sqlStatement, email)

	var user models.User
	if err := row.Scan(&user.Id, &user.Email, &user.Name, &user.Surname); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user not found: %w", err)
		}
		return nil, fmt.Errorf("failed to scan user: %w", err)
	}
	return &user, nil
}
