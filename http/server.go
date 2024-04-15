package server

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/joho/godotenv/autoload"
	server "github.com/vinibgoulart/todo-list/http/middleware"
	"github.com/vinibgoulart/todo-list/modules/task"
	"github.com/vinibgoulart/todo-list/modules/user"
)

func SetupHttpServer(db *sql.DB) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(server.JSONContentTypeMiddleware)
	r.Use(middleware.Timeout(30 * time.Second))

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Get("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API Working"))
	})
	r.Mount("/", user.UserPublicRouter(db))

	r.With(server.AuthMiddleware(db)).Mount("/user", user.UserRouter(db))
	r.With(server.AuthMiddleware(db)).Mount("/task", task.TaskRouter(db))

	return r
}
