package models

import "time"

type Todo struct {
	Id        int       `json:"id,omitempty"`
	Title     string    `json:"title,omitempty" validate:"required,min=5,max=75"`
	Text      string    `json:"text,omitempty" validate:"min=50,max=2048"`
	Image     string    `json:"image,omitempty"`
	UserId    int       `json:"user_id,omitempty" validate:"required,gte=1"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
