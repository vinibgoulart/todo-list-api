package user

import "github.com/go-chi/chi/v5"

func UserPublicRouter() chi.Router {
	r := chi.NewRouter()

	userPublicRouterHandler := UserPublicRouterHandler{}

	r.Post("/", userPublicRouterHandler.CreateUser)

	return r
}
