package utils

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

func TimeUTC() time.Time {
	return time.Now().UTC()
}

/**
 * create password hash
 */
func HashPassword(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(b), err
}

/**
 * validation password hash
 */
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
