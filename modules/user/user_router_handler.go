package user

import "net/http"

type UserRouterHandler struct{}

func (u UserRouterHandler) ListUsers(w http.ResponseWriter, r *http.Request)  {}
func (u UserRouterHandler) GetUsers(w http.ResponseWriter, r *http.Request)   {}
func (u UserRouterHandler) CreateUser(w http.ResponseWriter, r *http.Request) {}
func (u UserRouterHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {}
func (u UserRouterHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {}
