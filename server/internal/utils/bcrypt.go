package utils

import (
	"golang.org/x/crypto/bcrypt"

	log "github.com/sirupsen/logrus"
)

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Error("Error generating password hash", err)
	}
    return string(bytes), err
}

func CheckPasswordHash(password string, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Error("Error validating password hash", err)
	}
    return err == nil
}
