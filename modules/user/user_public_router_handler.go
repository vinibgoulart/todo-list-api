package user

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type UserPublicRouterHandler struct {
	storage UserStorage
}

func (u *UserPublicRouterHandler) CreateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		u.storage.Insert(db, user)

		err = json.NewEncoder(w).Encode(user)

		if err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}
	}

}
