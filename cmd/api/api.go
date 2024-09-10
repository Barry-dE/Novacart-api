// Package api provides the API server and routing setup
package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Barry-dE/Novacart-api/services/user"
	"github.com/gorilla/mux"
)

// APIServer struct holds the server address and database connection
type APIServer struct {
	address  string
	database *sql.DB
}

// NewAPIServer initializes a new APIServer with the given address and database
func NewAPIServer(address string, database *sql.DB) *APIServer {
	return &APIServer{
		address:  address,
		database: database,
	}
}

// Run starts the HTTP server and sets up routing
func (s *APIServer) Run() error {
	// Initialize a new router using Gorilla Mux
	router := mux.NewRouter()

	// Create a subrouter with the prefix /api/v1
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// Initialize a new user handler
	userHandler := user.NewHandler()

	// Register user-related routes with the subrouter
	userHandler.RegisterRoutes(subrouter)

	// Log the server's address
	log.Println("Listening on", s.address)

	// Start the HTTP server on the specified address, using the configured router
	return http.ListenAndServe(s.address, router)
}
