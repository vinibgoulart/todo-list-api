package user

import "github.com/go-chi/chi/v5"

func UserRouter() chi.Router {
	r := chi.NewRouter()

	userRouterHandler := UserRouterHandler{}

	r.Get("/", userRouterHandler.ListUsers)
	r.Post("/", userRouterHandler.CreateUser)
	r.Get("/{id}", userRouterHandler.GetUsers)
	r.Put("/{id}", userRouterHandler.UpdateUser)
	r.Delete("/{id}", userRouterHandler.DeleteUser)

	return r
}
