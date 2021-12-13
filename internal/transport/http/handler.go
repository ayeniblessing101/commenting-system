package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler - stores the pointer to our comments service
type Handler struct {
	Router *mux.Router
}

// NewHandler - returns a pointer to a handler
func NewHandler() *Handler {
	return &Handler{}
}

// SetupRoutes - set up all routes
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up Routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I am Alive!")
	})
}