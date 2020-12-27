package middlewares

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/thalissonfelipe/ugly-api/config"
	"github.com/thalissonfelipe/ugly-api/models"
	"github.com/thalissonfelipe/ugly-api/utils"
)

func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("jwt")
		if err == nil {
			tokenString := cookie.Value
			claims := &models.Claims{}
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return config.MyConfig.API.JWTKey, nil
			})
			if err != nil {
				utils.WriteResponse(w, http.StatusUnauthorized, "Invalid Token.")
				return
			}
			if !token.Valid {
				utils.WriteResponse(w, http.StatusUnauthorized, "Invalid Token.")
				return
			}
			next.ServeHTTP(w, r)
		} else {
			authHeader := r.Header.Get("Authorization")
			if authHeader != "" {
				parts := strings.Split(authHeader, " ")
				if len(parts) != 2 {
					utils.WriteResponse(w, http.StatusUnauthorized, "No token provided.")
					return
				}
				scheme, tokenString := parts[0], parts[1]
				if scheme == "Bearer" {
					claims := &models.Claims{}
					token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
						return config.MyConfig.API.JWTKey, nil
					})
					if err != nil {
						utils.WriteResponse(w, http.StatusUnauthorized, "Invalid Token.")
						return
					}
					if !token.Valid {
						utils.WriteResponse(w, http.StatusUnauthorized, "Invalid Token.")
						return
					}
					next.ServeHTTP(w, r)
				} else if scheme == "Basic" {
					// TODO
					utils.WriteResponse(w, http.StatusUnauthorized, "Invalid Token.")
				} else {
					utils.WriteResponse(w, http.StatusUnauthorized, "Invalid Token.")
				}
			} else {
				utils.WriteResponse(w, http.StatusUnauthorized, "No token provided.")
			}
		}
	})
}
