package server

import (
	"github.com/go-chi/chi"
	"github.com/vinibgoulart/todo-list/modules/user"
)

func mountRouter() {
	r := chi.NewRouter()

	r.Mount("/register", user.UserPublicRouter())
	r.Mount("/users", user.UserRouter())
}
