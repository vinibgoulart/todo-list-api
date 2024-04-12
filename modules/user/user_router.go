package user

import "github.com/go-chi/chi/v5"

func UserRouter() chi.Router {
	r := chi.NewRouter()

	userRouterHandler := UserRouterHandler{}

	r.Get("/{id}", userRouterHandler.GetUser)
	r.Put("/{id}", userRouterHandler.UpdateUser)
	r.Delete("/{id}", userRouterHandler.DeleteUser)

	return r
}
