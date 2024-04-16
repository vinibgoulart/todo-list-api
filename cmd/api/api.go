package main

import (
	"net/http"

	"github.com/vinibgoulart/todo-list/db"
	server "github.com/vinibgoulart/todo-list/http"
	"github.com/vinibgoulart/todo-list/lib"
)

func main() {
	logger := lib.NewLogger("main")

	db, err := db.ConnectDatabase()

	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("DB connected!")

	r := server.SetupHttpServer(db)

	httpServer := &http.Server{
		Addr:    ":8081",
		Handler: r,
	}

	httpServer.ListenAndServe()
}
