package user

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
)

func UserPublicRouter(db *sql.DB) chi.Router {
	r := chi.NewRouter()

	userPublicRouterHandler := UserHandler{
		storage: &UserStore{},
	}

	r.Post("/register", userPublicRouterHandler.UserCreate(db))
	r.Post("/login", userPublicRouterHandler.UserAuthenticate(db))

	return r
}
