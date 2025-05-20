package utils

import "os"

type UserCtxKey struct{}

var JWTSecret = os.Getenv("JWT_SECRET")
