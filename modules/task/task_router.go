package task

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
)

func TaskRouter(db *sql.DB) chi.Router {
	r := chi.NewRouter()

	TaskHandler := TaskHandler{
		storage: &TaskStore{},
	}

	r.Post("/", TaskHandler.TaskCreate(db))

	return r
}
