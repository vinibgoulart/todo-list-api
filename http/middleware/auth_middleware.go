package server_test

import (
	"database/sql"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/vinibgoulart/todo-list/lib"
	"github.com/vinibgoulart/todo-list/modules/user"
)

func AuthMiddleware(db *sql.DB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenHeader := r.Header.Get("Authorization")

			if tokenHeader == "" {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
				return
			}

			token, err := lib.JWTVerifyToken(tokenHeader)

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
			}

			userId := token.Claims.(jwt.MapClaims)["identifier"].(string)

			userStore := &user.UserStore{}

			user, err := userStore.Get(db, userId)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			if user == nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
