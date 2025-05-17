package handlers

import (
	"backend/internal/models"
	"backend/internal/services"
	"encoding/json"
	"net/http"
	"strconv"
)

type TodoHandler struct {
	s services.TodoService
}

func NewTodoHandler(s *services.TodoService) *TodoHandler {
	return &TodoHandler{s: *s}
}

func (h TodoHandler) GetTodoByTitle(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	userIdStr := r.URL.Query().Get("userId")

	if title == "" || userIdStr == "" {
		http.Error(w, "missing title or userId", http.StatusBadRequest)
		return
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		http.Error(w, "invalid userId", http.StatusBadRequest)
		return
	}

	res, err := h.s.GetTodoByTitle(title, userId)
	if err != nil {
		http.Error(w, "failed to get todo", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (h TodoHandler) GetTodoByUserId(w http.ResponseWriter, r *http.Request) {
	userIdStr := r.URL.Query().Get("userId")

	if userIdStr == "" {
		http.Error(w, "missing userId", http.StatusBadRequest)
		return
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		http.Error(w, "invalid userId", http.StatusBadRequest)
		return
	}

	res, err := h.s.GetTodoByUserId(userId)
	if err != nil {
		http.Error(w, "failed to get todo", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (h TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(todo)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	result, err := h.s.CreateTodo(todo)
	if err != nil {
		http.Error(w, "failed to create todo", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

func (h TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(todo)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	result, err := h.s.UpdateTodo(todo)
	if err != nil {
		http.Error(w, "failed to update todo", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (h TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	result, err := h.s.DeleteTodo(id)
	if err != nil {
		http.Error(w, "failed to delete todo", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
