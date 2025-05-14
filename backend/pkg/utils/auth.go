package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword Хеширование пароля
func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// CheckPasswordHash Проверка пароля
func CheckPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
