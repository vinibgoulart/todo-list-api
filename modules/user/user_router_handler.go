package user

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type UserRouterHandler struct {
	storage UserStorage
}

func (u *UserRouterHandler) Get(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 64)

		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		user, err := u.storage.Get(id)

		if err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}

		if user == nil {
			http.Error(w, "User not found", http.StatusNotFound)
		}

		err = json.NewEncoder(w).Encode(user)

		if err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}
	}

}
