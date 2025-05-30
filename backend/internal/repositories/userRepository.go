package repositories

import (
	"backend/internal/models"
	"database/sql"
	"errors"
	"fmt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAllUsers() []models.User {
	sqlStatement := "SELECT id, email, login, name, surname FROM users"
	rows, err := r.db.Query(sqlStatement)
	if err != nil {
		return nil
	}
	defer rows.Close()
	var users []models.User

	for rows.Next() {
		var user models.User

		err := rows.Scan(&user.Id, &user.Email, &user.Login, &user.Name, &user.Surname)
		if err != nil {
			return nil
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil
	}

	return users
}

func (r *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	sqlStatement := `
        INSERT INTO users(email, login, name, surname, avatar, password)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id`

	var id int
	err := r.db.QueryRow(sqlStatement,
		user.Email, user.Login, user.Name, user.Surname, user.Avatar, user.Password,
	).Scan(&id)

	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	user.Id = id
	return user, nil
}

func (r *UserRepository) UpdateUser(user *models.User) (*models.User, error) {
	sqlStatement := `
		UPDATE users
		SET email = $2, name = $3, surname = $4, avatar = $5
		WHERE id = $1
		RETURNING id, email, name, surname, avatar`

	var userRes models.User
	err := r.db.
		QueryRow(sqlStatement, user.Id, user.Email, user.Name, user.Surname, user.Avatar).
		Scan(&userRes.Id, &userRes.Email, &userRes.Name, &userRes.Surname, &userRes.Avatar)
	if err != nil {
		return nil, err
	}

	return &userRes, nil
}

func (r *UserRepository) DeleteUser(id int) (*int, error) {
	sqlStatement := `DELETE FROM users WHERE id = $1`
	res, err := r.db.Exec(sqlStatement, id)
	if err != nil {
		return nil, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, fmt.Errorf("user with id %v does not exist", id)
	}
	return &id, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	sqlStatement := "SELECT id, email, name, surname FROM users WHERE email = $1"
	row := r.db.QueryRow(sqlStatement, email)

	var user models.User
	if err := row.Scan(&user.Id, &user.Email, &user.Name, &user.Surname); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user not found: %w", err)
		}
		return nil, fmt.Errorf("failed to scan user: %w", err)
	}
	return &user, nil
}
