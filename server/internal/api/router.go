package api

import (
	"net/http"
	"strings"

	"github.com/tamakiii/flashcard-app/server/internal/db"
)

// Router handles HTTP routing
type Router struct {
	handler *Handler
}

// NewRouter creates a new Router instance
func NewRouter(db *db.DB) *Router {
	return &Router{
		handler: NewHandler(db),
	}
}

// ServeHTTP implements the http.Handler interface
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// Handle preflight requests
	if req.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Check if the request is to the API endpoint
	path := req.URL.Path
	if !strings.HasPrefix(path, "/api/") {
		http.NotFound(w, req)
		return
	}

	// Extract the API path
	apiPath := strings.TrimPrefix(path, "/api")
	if apiPath == "" {
		apiPath = "/"
	}

	// Route based on path and method
	switch {
	case apiPath == "/flashcards" && req.Method == http.MethodGet:
		r.handler.GetFlashcards(w, req)
	case strings.HasPrefix(apiPath, "/flashcards/") && req.Method == http.MethodGet:
		if strings.Count(apiPath, "/") == 2 {
			r.handler.GetFlashcard(w, req)
		} else {
			http.NotFound(w, req)
		}
	case apiPath == "/flashcards" && req.Method == http.MethodPost:
		r.handler.CreateFlashcard(w, req)
	case strings.HasPrefix(apiPath, "/flashcards/") && req.Method == http.MethodPut:
		if strings.Count(apiPath, "/") == 2 {
			r.handler.UpdateFlashcard(w, req)
		} else {
			http.NotFound(w, req)
		}
	case strings.HasPrefix(apiPath, "/flashcards/") && req.Method == http.MethodDelete:
		if strings.Count(apiPath, "/") == 2 {
			r.handler.DeleteFlashcard(w, req)
		} else {
			http.NotFound(w, req)
		}
	default:
		http.NotFound(w, req)
	}
}
