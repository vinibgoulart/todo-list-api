package user

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
)

func UserRouter(db *sql.DB) chi.Router {
	r := chi.NewRouter()

	UserHandler := UserHandler{
		storage: &UserStore{},
	}

	r.Get("/{id}", UserHandler.UserGet(db))

	return r
}
