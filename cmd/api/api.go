package api

import (
	"database/sql"
	"log"
	"net/http"
	"strings"
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

// PrefixHandler strips the prefix before forwarding requests
func PrefixHandler(prefix string, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, prefix) {
			// Remove the prefix before forwarding the request
			r.URL.Path = strings.TrimPrefix(r.URL.Path, prefix)
			if r.URL.Path == "" {
				r.URL.Path = "/"
			}
			handler.ServeHTTP(w, r)
		} else {
			http.NotFound(w, r)
		}
	})
}


// Run starts the HTTP server
func (s *APIServer) Run() error {
	router := http.NewServeMux() // Root router

	// Attach user service routes
	userRoutes := user.NewUserRoutes(s.db) // Get user routes
	router.Handle("/api/v1/users/", PrefixHandler("/api/v1/users", userRoutes))

	log.Println("Server running on", s.addr)
	return http.ListenAndServe(s.addr, router)
}

