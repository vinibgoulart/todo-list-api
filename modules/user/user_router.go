package user

import "github.com/go-chi/chi/v5"

func UserRouter() chi.Router {
	r := chi.NewRouter()

	userRouterHandler := UserRouterHandler{
		storage: &UserStore{},
	}

	r.Get("/{id}", userRouterHandler.GetUser)

	return r
}
