package utils

import (
	"context"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/dgrijalva/jwt-go"
	"github.com/thalissonfelipe/ugly-api/config"
	"github.com/thalissonfelipe/ugly-api/models"
	"go.mongodb.org/mongo-driver/bson"
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

func VerifyCookie(cookie *http.Cookie, w http.ResponseWriter) bool {
	claims := &models.Claims{}
	token, err := jwt.ParseWithClaims(cookie.Value, claims, func(token *jwt.Token) (interface{}, error) {
		return config.MyConfig.API.JWTKey, nil
	})
	if err != nil || !token.Valid {
		WriteResponse(w, http.StatusUnauthorized, "Invalid Token.")
		return false
	}
	return true
}

func VerifyBearerToken(tokenString string, w http.ResponseWriter) bool {
	claims := &models.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return config.MyConfig.API.JWTKey, nil
	})
	if err != nil || !token.Valid {
		WriteResponse(w, http.StatusUnauthorized, "Invalid Token.")
		return false
	}
	return true
}

func VerifyBasic(username string, password string, w http.ResponseWriter, client *mongo.Client) bool {
	err := CheckUser(username, password, client)
	if err != nil {
		WriteResponse(w, http.StatusUnauthorized, "Invalid Token.")
		return false
	}
	return true
}

func CheckUser(username string, password string, client *mongo.Client) error {
	collection := client.Database(config.MyConfig.DB.DatabaseName).Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	user := models.User{}
	err := collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return err
	}

	return ComparePassword(user.Password, password)
}
