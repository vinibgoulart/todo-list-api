package user

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
)

func UserRouter(db *sql.DB) chi.Router {
	r := chi.NewRouter()

	userRouterHandler := UserRouterHandler{
		storage: &UserStore{},
	}

	r.Get("/{id}", userRouterHandler.Get(db))

	return r
}
