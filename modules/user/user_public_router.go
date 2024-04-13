package user

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
)

func UserPublicRouter(db *sql.DB) chi.Router {
	r := chi.NewRouter()

	userPublicRouterHandler := UserPublicRouterHandler{
		storage: &UserStore{},
	}

	r.Post("/", userPublicRouterHandler.CreateUser(db))

	return r
}
