package models

import "github.com/golang-jwt/jwt/v5"

// UserClaims данные из jwt
type UserClaims struct {
	Email string `json:"sub"`
	jwt.RegisteredClaims
}
