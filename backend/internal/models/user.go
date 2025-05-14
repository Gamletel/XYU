package models

type User struct {
	Id       int    `json:"id,omitempty"`
	Email    string `json:"email,omitempty"`
	Login    string `json:"login,omitempty"`
	Password string `json:"-"`
	Name     string `json:"name,omitempty"`
	Surname  string `json:"surname,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
}
