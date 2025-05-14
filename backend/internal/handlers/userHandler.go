package handlers

import (
	"backend/internal/models"
	"backend/internal/services"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: *service}
}

var validate = validator.New()

func (h UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := h.service.GetAllUsers()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (h UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := validate.Struct(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.service.CreateUser(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]*models.User{"result": res})
}
