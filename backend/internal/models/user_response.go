package models

type UserResponse struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Login   string `json:"login"`
	Email   string `json:"email"`
	Avatar  string `json:"avatar"`
}
