package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/adammwaniki/mi-segunda-api-de-golang/service/user"
	"github.com/gorilla/mux"
)

type APIServer struct{
	addr string
	db *sql.DB
}

// NewAPIServer initializes a new API server
func NewAPIServer (addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db: db,
	}
}

// Run starts the HTTP server
func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)

	log.Println("Server running on", s.addr)
	return http.ListenAndServe(s.addr, router)
}

