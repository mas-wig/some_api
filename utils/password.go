package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashedPassword(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("could not hashed password %w", err)
	}
	return string(hashedPass), nil
}

func VerifyPassword(hashedPass string, candidatePass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(candidatePass))
}
