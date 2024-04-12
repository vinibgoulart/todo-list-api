package main

import (
	"net/http"

	"github.com/vinibgoulart/todo-list/server"
)

func main() {
	r := server.SetupServer()
	http.ListenAndServe(":3000", r)
}
