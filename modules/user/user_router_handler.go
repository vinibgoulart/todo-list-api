package user

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UserRouterHandler struct {
	storage UserStorage
}

func (u *UserRouterHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(u.storage.ListUsers())

	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (u *UserRouterHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	user := u.storage.GetUser(id)

	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
	}

	err := json.NewEncoder(w).Encode(user)

	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
