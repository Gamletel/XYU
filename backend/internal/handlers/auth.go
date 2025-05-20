package handlers

import (
	"backend/internal/db"
	"backend/internal/models"
	"backend/internal/repositories"
	"backend/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDTO struct {
	Login string `json:"login"`
	Email string `json:"email"`
}

type loginResponse struct {
	Token string  `json:"token"`
	User  UserDTO `json:"user"`
}

type registerRequest struct {
	models.User
}

// Login Авторизация
func Login(w http.ResponseWriter, r *http.Request) {
	var req loginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userRepo := repositories.NewUserRepository(db.DB)
	user, err := userRepo.GetUserByEmail(req.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	password, err := userRepo.GetUserPassword(user.Id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !utils.CheckPasswordHash(*password, req.Password) {
		http.Error(w, fmt.Sprintf("wrong email/password %v, %v", req.Password, *password), http.StatusBadRequest)
		return
	}

	token, err := utils.GenerateToken(user.Email)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	res := loginResponse{
		User: UserDTO{
			Login: user.Login,
			Email: user.Email,
		},
		Token: token,
	}
	json.NewEncoder(w).Encode(res)
}

// Register Регистрация пользователя
func Register(w http.ResponseWriter, r *http.Request) {
	var req registerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userRepo := repositories.NewUserRepository(db.DB)
	user, err := userRepo.GetUserByEmail(req.Email)
	if user != nil {
		http.Error(w, "User is already exist", http.StatusBadRequest)
		return
	}

	passwordHashed, _ := utils.HashPassword(req.Password)

	userReq := models.User{
		Email:    req.Email,
		Login:    req.Login,
		Password: passwordHashed,
		Name:     req.Name,
		Surname:  req.Surname,
		Avatar:   req.Avatar,
	}

	user, err = userRepo.CreateUser(&userReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
