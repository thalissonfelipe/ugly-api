package utils

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/thalissonfelipe/ugly-api/config"
	"github.com/thalissonfelipe/ugly-api/models"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword receives a string passwords and returns a hashed password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hashedPassword), err
}

// ComparePassword checks if the password passed is a valid password
func ComparePassword(hashedPassword, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}

func CreateToken(username string, w http.ResponseWriter) (string, error) {
	expirationTime := time.Now().Add(time.Duration(config.MyConfig.API.JWTExpireTime) * time.Second)
	claims := &models.Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.MyConfig.API.JWTKey)
	if err != nil {
		return tokenString, err
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "jwt",
		Value:   tokenString,
		Expires: expirationTime,
		Path:    "/",
	})

	return tokenString, err
}

func WriteResponse(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	w.Write([]byte(message))
}
