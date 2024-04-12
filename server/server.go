package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func SetupServer() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API Working"))
	})

	mountRouter()

	return r
}
