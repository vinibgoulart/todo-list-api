package main

import (
	"net/http"

	"github.com/vinibgoulart/todo-list/user"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API Working"))
	})

	r.Mount("/users", user.UserRouter())

	http.ListenAndServe(":3000", r)
}
