package user

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UserRouterHandler struct {
}

func (u *UserRouterHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(ListUsers())

	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (u *UserRouterHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	user := GetUser(id)

	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
	}

	err := json.NewEncoder(w).Encode(user)

	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (u *UserRouterHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {}
func (u *UserRouterHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {}
