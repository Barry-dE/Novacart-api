package user

import (
	"net/http"

	"github.com/gorilla/mux"
)


type Handler struct {
}


func NewHandler() *Handler {
	return &Handler{}
}


func (h *Handler) RegisterRoutes(router *mux.Router) {
	
	// Register the "/login" route with the router, mapping it to the handleLogin method
	router.HandleFunc("/login", h.handleLogin).Methods("POST") // Define that this route responds to POST requests

	// Register the "/signup" route with the router, mapping it to the handleSignUp method
	router.HandleFunc("/signup", h.handleSignUp).Methods("POST") 
}

// handleLogin method handles the "/login" route
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	
}

// handleSignUp method handles the "/signup" route
func (h *Handler) handleSignUp(w http.ResponseWriter, r *http.Request) {
	
}
