package handlers

import (
	"backend/internal/db"
	"backend/internal/models"
	"backend/internal/repositories"
	"backend/pkg/utils"
	"encoding/json"
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

	switch {
	case err != nil:
		http.Error(w, "Invalid email or password", http.StatusBadRequest)
		return
	case !utils.CheckPasswordHash(req.Password, user.Password):
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateToken(user.Email)

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
	if err != nil {
		http.Error(w, "Email is already used", http.StatusBadRequest)
		return
	}

	user, err = userRepo.CreateUser(&req.User)
	if err != nil {
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
