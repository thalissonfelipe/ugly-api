package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword receives a string passwords and returns a hashed password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hashedPassword), err
}

// ComparePassword checks if the password passed is a valid password
func ComparePassword(hashedPassword, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}
