package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	// Hash the password using bcrypt
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
