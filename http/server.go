package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/vinibgoulart/todo-list/modules/user"
)

func SetupHttpServer() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(30 * time.Second))

	r.Get("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API Working"))
	})

	r.Mount("/register", user.UserPublicRouter())
	r.Mount("/user", user.UserRouter())

	return r
}
