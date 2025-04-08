package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/adammwaniki/mi-segunda-api-de-golang/service/cart"
	"github.com/adammwaniki/mi-segunda-api-de-golang/service/order"
	"github.com/adammwaniki/mi-segunda-api-de-golang/service/product"
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

	// Register User service on the api
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore) // dependency injection
	userHandler.RegisterRoutes(subrouter)

	// Register product service on the api
	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore, userStore)
	productHandler.RegisterRoutes(subrouter)

	orderStore := order.NewStore(s.db)

	cartHandler := cart.NewHandler(productStore, orderStore, userStore)
	cartHandler.RegisterRoutes(subrouter)

	log.Println("Server running on", s.addr)
	return http.ListenAndServe(s.addr, router)
}

