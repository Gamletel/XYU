package models

import "time"

type Todo struct {
	Id        int       `json:"id,omitempty"`
	Title     string    `json:"title,omitempty"`
	Text      string    `json:"text,omitempty"`
	Image     string    `json:"image,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
