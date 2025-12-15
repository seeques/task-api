package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	hashString := string(hash)
	return hashString, nil
}

func CheckPassword(password string, passwordHash string) (error) {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))

	if err != nil {
		return err
	}

	// nil on success
	return nil
}