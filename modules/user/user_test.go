package user

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var fakeUsers = []*User{{
	ID:       "1",
	Name:     "Vinicius Goulart",
	Email:    "viblaziusgoulart@gmail.com",
	Password: "123456",
}}

type fakeStorage struct {
}

func (s fakeStorage) GetUser(_ string) *User {
	return fakeUsers[0]
}

func (s fakeStorage) StoreUser(_ User) {
}

func (s fakeStorage) ListUsers() []*User {
	return fakeUsers
}

func TestGetUsersHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/users/1", nil)

	w := httptest.NewRecorder()

	userHandler := UserRouterHandler{
		storage: fakeStorage{},
	}

	userHandler.GetUser(w, req)

	res := w.Result()

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	user := User{}
	json.Unmarshal(data, &user)

	if user.Name != "Vinicius Goulart" {
		t.Errorf("expected name got %v", string(data))
	}
}
