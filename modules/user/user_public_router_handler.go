package user

import "net/http"

type UserPublicRouterHandler struct{}

func (u *UserPublicRouterHandler) CreateUser(w http.ResponseWriter, r *http.Request) {}
