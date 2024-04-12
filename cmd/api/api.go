package main

import (
	"net/http"

	server "github.com/vinibgoulart/todo-list/http"
)

func main() {
	r := server.SetupHttpServer()
	http.ListenAndServe(":8080", r)
}
