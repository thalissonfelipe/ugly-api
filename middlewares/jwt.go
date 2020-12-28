package middlewares

import (
	"net/http"
	"strings"

	"github.com/thalissonfelipe/ugly-api/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

func JWTMiddleware(client *mongo.Client) (jwtfunc func(http.HandlerFunc) http.HandlerFunc) {
	jwtfunc = func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("jwt")
			if err == nil {
				ok := utils.VerifyCookie(cookie, w)
				if ok {
					next.ServeHTTP(w, r)
				}
			} else {
				authHeader := r.Header.Get("Authorization")
				if authHeader != "" {
					parts := strings.Split(authHeader, " ")
					if len(parts) != 2 {
						utils.WriteResponse(w, http.StatusUnauthorized, "No token provided.")
						return
					}
					scheme, token := parts[0], parts[1]
					if scheme == "Bearer" {
						ok := utils.VerifyBearerToken(token, w)
						if ok {
							next.ServeHTTP(w, r)
						}
					} else if scheme == "Basic" {
						username, password, ok := r.BasicAuth()
						if ok {
							ok = utils.VerifyBasic(username, password, w, client)
							if ok {
								next.ServeHTTP(w, r)
							}
						}

					} else {
						utils.WriteResponse(w, http.StatusUnauthorized, "Invalid Token.")
					}
				} else {
					utils.WriteResponse(w, http.StatusUnauthorized, "No token provided.")
				}
			}
		})
	}
	return jwtfunc
}
