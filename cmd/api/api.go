package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Barry-dE/Novacart-api/services/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	address string
	database *sql.DB
}


func NewAPIServer(address string, database *sql.DB) *APIServer {

	return &APIServer{
		address: address,
		database: database,
	}
}


func (s *APIServer) Run() error {

	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	
	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)
	
	log.Println("Listening on", s.address)
	return http.ListenAndServe(s.address, router)
}
