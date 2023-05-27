package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
)

type MyServer struct {
	port int
	db   *sql.DB
}

func NewServer(p int, db *sql.DB) *MyServer {
	return &MyServer{
		port: p,
		db:   db,
	}
}

func (s *MyServer) StartServer() error {
	r := chi.NewRouter()
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})
	r.Post("/api/v1/user", s.CreateUser)
	r.Get("/api/v1/user/{id}", s.GetUser)
	r.Post("/api/v1/task", s.CreateTask)
	r.Get("/api/v1/task/{id}", s.GetTask)
	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), r)
}
