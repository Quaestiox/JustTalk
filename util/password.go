package util

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pd string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pd), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
